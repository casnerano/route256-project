package cart

//go:generate mockgen -destination=mock/cart.go -source=cart.go

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "sync"

    "go.uber.org/zap"
    "route256/cart/internal/model"
    "route256/cart/internal/repository"
    "route256/cart/internal/service/cart/worker_pool"
    "route256/cart/internal/service/client/pim"
    "route256/cart/pkg/cache"
)

var ErrNotFound = errors.New("not found")
var ErrEmptyCart = errors.New("empty cart")
var ErrPIMProductNotFound = errors.New("pim product not found")
var ErrPIMLowAvailability = errors.New("pim product low availability in stock")

type PIMClient interface {
    GetProductInfo(ctx context.Context, sku model.SKU) (*model.ProductInfo, error)
}

type LOMSClient interface {
    GetStockInfo(ctx context.Context, sku model.SKU) (uint32, error)
    CreateOrder(ctx context.Context, userID model.UserID, items []*model.Item) (model.OrderID, error)
}

type Cart struct {
    rep    repository.Cart
    pim    PIMClient
    loms   LOMSClient
    cache  cache.Cache
    logger *zap.Logger
}

func New(rep repository.Cart, pim PIMClient, loms LOMSClient, cache cache.Cache, logger *zap.Logger) *Cart {
    return &Cart{
        rep:    rep,
        pim:    pim,
        loms:   loms,
        cache:  cache,
        logger: logger,
    }
}

func (c *Cart) Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint32) (*model.Item, error) {
    _, err := c.getProductInfo(ctx, sku)
    if err != nil {
        if errors.Is(err, pim.ErrProductNotFound) {
            return nil, ErrPIMProductNotFound
        }

        return nil, err
    }

    available, err := c.loms.GetStockInfo(ctx, sku)
    if err != nil {
        return nil, err
    }

    if count > available {
        return nil, ErrPIMLowAvailability
    }

    return c.rep.Add(ctx, userID, sku, count)
}

func (c *Cart) Delete(ctx context.Context, userID model.UserID, sku model.SKU) error {
    err := c.rep.DeleteBySKU(ctx, userID, sku)
    if err != nil {
        if errors.Is(err, repository.ErrNotFound) {
            return ErrNotFound
        }

        return err
    }

    return nil
}

// List returns a detailed list of the cart items.
// The items is a combination of cart values and product info from ProductService.
func (c *Cart) List(ctx context.Context, wp worker_pool.WorkerPool, userID model.UserID) ([]*model.ItemDetail, error) {
    list, err := c.rep.FindByUser(ctx, userID)
    if err != nil {
        if errors.Is(err, repository.ErrNotFound) {
            return nil, ErrNotFound
        }

        return nil, err
    }

    // Function for processing task in worker.
    processor := func(ctx context.Context, task worker_pool.Task) *worker_pool.Result {
        productInfo, err := c.getProductInfo(ctx, task.SKU)
        if err != nil {
            log.Println("Failed processing task in worker pool: ", err)
            return nil
        }

        return &worker_pool.Result{
            SKU: task.SKU,
            ProductInfo: model.ProductInfo{
                Name:  productInfo.Name,
                Price: productInfo.Price,
            },
        }
    }

    tasksCh := make(chan worker_pool.Task)
    resultsCh := wp.Run(ctx, tasksCh, processor)

    productList := make(map[model.SKU]*model.ProductInfo)

    wg := &sync.WaitGroup{}
    wg.Add(1)

    // Goroutine for listen workers result channel.
    // Read len(list) results.
    go func() {
        defer wg.Done()
        for i := 0; i < len(list); i++ {
            select {
            case result, ok := <-resultsCh:
                if ok {
                    if result != nil {
                        productList[result.SKU] = &result.ProductInfo
                    }

                }
            case <-ctx.Done():
                return
            }
        }
    }()

    // Send tasks to the tasks channel.
    for _, item := range list {
        tasksCh <- worker_pool.Task{SKU: item.SKU}
    }

    // Wait getting the results of all tasks
    wg.Wait()

    close(tasksCh)

    detailList := make([]*model.ItemDetail, 0, len(list))
    for k := range list {
        productInfo, ok := productList[list[k].SKU]
        if !ok {
            return nil, fmt.Errorf("failed to get product info from product.service with sku %d", list[k].SKU)
        }

        detailList = append(detailList, &model.ItemDetail{
            Item: model.Item{
                SKU:   list[k].SKU,
                Count: list[k].Count,
            },
            Price: productInfo.Price,
            Name:  productInfo.Name,
        })
    }

    return detailList, nil
}

func (c *Cart) Clear(ctx context.Context, userID model.UserID) error {
    err := c.rep.DeleteByUser(ctx, userID)
    if err != nil {
        if errors.Is(err, repository.ErrNotFound) {
            return ErrNotFound
        }

        return err
    }

    return nil
}

func (c *Cart) Checkout(ctx context.Context, userID model.UserID) (model.OrderID, error) {
    items, err := c.rep.FindByUser(ctx, userID)
    if err != nil {
        if errors.Is(err, repository.ErrNotFound) {
            return 0, ErrEmptyCart
        }

        return 0, err
    }

    if len(items) == 0 {
        return 0, ErrEmptyCart
    }

    orderID, err := c.loms.CreateOrder(ctx, userID, items)
    if err != nil {
        return 0, err
    }

    err = c.rep.DeleteByUser(ctx, userID)
    if err != nil {
        return 0, err
    }

    return orderID, nil
}

func (c *Cart) getProductInfo(ctx context.Context, sku model.SKU) (*model.ProductInfo, error) {
    key := fmt.Sprint(sku)

    cacheData, err := c.cache.Get(ctx, key)
    if err != nil {
        c.logger.Error("Failed get data from cache.", zap.Error(err))
    }

    var data *model.ProductInfo

    if cacheData == nil {
        data, err = c.pim.GetProductInfo(ctx, sku)
        if err != nil {
            return nil, err
        }

        var bData []byte
        bData, err = json.Marshal(data)
        if err != nil {
            c.logger.Error("Failed marshal data.", zap.Error(err))
        } else {
            if err = c.cache.Set(ctx, key, string(bData)); err != nil {
                c.logger.Error("Failed set data to cache.", zap.Error(err))
            }
        }

        return data, nil
    }

    if err = json.Unmarshal([]byte(*cacheData), &data); err != nil {
        c.logger.Error("Failed unmarshal data.", zap.Error(err))

        data, err = c.pim.GetProductInfo(ctx, sku)
        if err != nil {
            return nil, err
        }
    }

    return data, nil
}

package cart

import (
	"context"
	"errors"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
	mock_repository "route256/cart/internal/repository/mock"
	mock_cart "route256/cart/internal/service/cart/mock"
	"route256/cart/internal/service/cart/worker_pool"
	mock_worker_pool "route256/cart/internal/service/cart/worker_pool/mock"
	"route256/cart/internal/service/client/pim"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

var fakeError = errors.New("fake error")

type fakeItem struct {
	UserID model.UserID
	SKU    model.SKU
}

type CartTestSuite struct {
	suite.Suite
	cart *Cart
	rep  *mock_repository.MockCart
	pim  *mock_cart.MockPIMClient
	loms *mock_cart.MockLOMSClient
	wp   *mock_worker_pool.MockWorkerPool
}

func (s *CartTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.rep = mock_repository.NewMockCart(ctrl)
	s.pim = mock_cart.NewMockPIMClient(ctrl)
	s.loms = mock_cart.NewMockLOMSClient(ctrl)
	s.wp = mock_worker_pool.NewMockWorkerPool(ctrl)

	s.cart = New(s.rep, s.pim, s.loms)
}

func (s *CartTestSuite) TestCartAdd() {
	fItem := fakeItem{
		UserID: 1,
		SKU:    2,
	}

	ctx := context.Background()

	s.Run("Product not found in pim", func() {
		s.pim.EXPECT().GetProductInfo(gomock.Any(), fItem.SKU).Return(nil, pim.ErrProductNotFound)

		_, err := s.cart.Add(ctx, fItem.UserID, fItem.SKU, 1)
		s.ErrorIs(err, ErrPIMProductNotFound)
	})

	s.Run("Unknown error from pim", func() {
		s.pim.EXPECT().GetProductInfo(gomock.Any(), fItem.SKU).Return(nil, fakeError)

		_, err := s.cart.Add(ctx, fItem.UserID, fItem.SKU, 1)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Stock other error", func() {
		s.pim.EXPECT().GetProductInfo(gomock.Any(), fItem.SKU).Return(nil, nil)
		s.loms.EXPECT().GetStockInfo(gomock.Any(), fItem.SKU).Return(uint32(0), fakeError)

		_, err := s.cart.Add(ctx, fItem.UserID, fItem.SKU, 1)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Product low availability in stock", func() {
		s.pim.EXPECT().GetProductInfo(gomock.Any(), fItem.SKU).Return(nil, nil)
		s.loms.EXPECT().GetStockInfo(gomock.Any(), fItem.SKU).Return(uint32(10), nil)

		_, err := s.cart.Add(ctx, fItem.UserID, fItem.SKU, 100)
		s.ErrorIs(err, ErrPIMLowAvailability)
	})

	s.Run("Unknown Add error from repo", func() {
		var availStock uint32 = 100
		var addCount uint32 = 10

		s.pim.EXPECT().GetProductInfo(gomock.Any(), fItem.SKU).Return(nil, nil)
		s.loms.EXPECT().GetStockInfo(gomock.Any(), fItem.SKU).Return(availStock, nil)
		s.rep.EXPECT().Add(gomock.Any(), fItem.UserID, fItem.SKU, addCount).Return(nil, fakeError)

		_, err := s.cart.Add(ctx, fItem.UserID, fItem.SKU, addCount)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Successful Add", func() {
		var availStock uint32 = 100
		var addCount uint32 = 10

		wantModel := &model.Item{
			ID:     1,
			UserID: fItem.UserID,
			SKU:    fItem.SKU,
			Count:  addCount,
		}

		s.pim.EXPECT().GetProductInfo(gomock.Any(), fItem.SKU).Return(nil, nil)
		s.loms.EXPECT().GetStockInfo(gomock.Any(), fItem.SKU).Return(availStock, nil)
		s.rep.EXPECT().Add(gomock.Any(), fItem.UserID, fItem.SKU, addCount).Return(wantModel, nil)

		gotModel, err := s.cart.Add(ctx, fItem.UserID, fItem.SKU, addCount)
		s.Require().NoError(err)
		s.Equal(wantModel, gotModel)
	})
}

func (s *CartTestSuite) TestCartDelete() {
	fItem := fakeItem{
		UserID: 1,
		SKU:    2,
	}

	ctx := context.Background()

	s.Run("Not found", func() {
		s.rep.EXPECT().DeleteBySKU(gomock.Any(), fItem.UserID, fItem.SKU).Return(repository.ErrNotFound)

		err := s.cart.Delete(ctx, fItem.UserID, fItem.SKU)
		s.ErrorIs(err, ErrNotFound)
	})

	s.Run("Unknown error from repo", func() {
		s.rep.EXPECT().DeleteBySKU(gomock.Any(), fItem.UserID, fItem.SKU).Return(fakeError)

		err := s.cart.Delete(ctx, fItem.UserID, fItem.SKU)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Successful delete", func() {
		s.rep.EXPECT().DeleteBySKU(gomock.Any(), fItem.UserID, fItem.SKU).Return(nil)

		err := s.cart.Delete(ctx, fItem.UserID, fItem.SKU)
		s.NoError(err)
	})
}

func (s *CartTestSuite) TestCartList() {
	fItem := fakeItem{
		UserID: 1,
	}

	ctx := context.Background()
	emptyList := make([]*model.Item, 0)

	s.Run("Cart not found", func() {
		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(emptyList, repository.ErrNotFound)

		_, err := s.cart.List(ctx, s.wp, fItem.UserID)
		s.ErrorIs(err, ErrNotFound)
	})

	s.Run("Unknown error from repo", func() {
		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(emptyList, fakeError)

		_, err := s.cart.List(ctx, s.wp, fItem.UserID)
		s.ErrorIs(err, fakeError)
	})

	fillList := []*model.Item{
		{
			ID:     1,
			UserID: 1,
			SKU:    1,
			Count:  1,
		},
		{
			ID:     2,
			UserID: 2,
			SKU:    2,
			Count:  2,
		},
		{
			ID:     3,
			UserID: 3,
			SKU:    3,
			Count:  3,
		},
		{
			ID:     4,
			UserID: 4,
			SKU:    4,
			Count:  4,
		},
	}

	s.Run("Successful result cart list with worker pool", func() {
		fakeProductInfo := model.ProductInfo{
			Name:  "",
			Price: 0,
		}

		fakeResults := make(chan *worker_pool.Result)

		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(fillList, nil)
		s.pim.EXPECT().GetProductInfo(gomock.Any(), gomock.Any()).Return(&fakeProductInfo, nil).MaxTimes(len(fillList))
		s.wp.EXPECT().Run(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, tasks <-chan worker_pool.Task, proc worker_pool.Processor) <-chan *worker_pool.Result {
			go func() {
				for {
					select {
					case <-ctx.Done():
						return
					case task, ok := <-tasks:
						if !ok {
							return
						}
						result := proc(ctx, task)
						select {
						case <-ctx.Done():
							return
						case fakeResults <- result:
						}
					}
				}
			}()

			return fakeResults
		}).MaxTimes(len(fillList))

		_, err := s.cart.List(ctx, s.wp, fItem.UserID)
		s.Require().NoError(err)
	})
}

func (s *CartTestSuite) TestCartClear() {
	fItem := fakeItem{
		UserID: 1,
	}

	ctx := context.Background()

	s.Run("Not found", func() {
		s.rep.EXPECT().DeleteByUser(gomock.Any(), fItem.UserID).Return(repository.ErrNotFound)

		err := s.cart.Clear(ctx, fItem.UserID)
		s.ErrorIs(err, ErrNotFound)
	})

	s.Run("Unknown error from repo", func() {
		s.rep.EXPECT().DeleteByUser(gomock.Any(), fItem.UserID).Return(fakeError)

		err := s.cart.Clear(ctx, fItem.UserID)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Successful clear", func() {
		s.rep.EXPECT().DeleteByUser(gomock.Any(), fItem.UserID).Return(nil)

		err := s.cart.Clear(ctx, fItem.UserID)
		s.NoError(err)
	})
}

func (s *CartTestSuite) TestCartCheckout() {
	fItem := fakeItem{
		UserID: 1,
	}

	ctx := context.Background()
	emptyList := make([]*model.Item, 0)

	s.Run("Empty cart (with not found)", func() {
		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(emptyList, repository.ErrNotFound)

		_, err := s.cart.Checkout(ctx, fItem.UserID)
		s.ErrorIs(err, ErrEmptyCart)
	})

	s.Run("Unknown error from repo", func() {
		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(emptyList, fakeError)

		_, err := s.cart.Checkout(ctx, fItem.UserID)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Empty cart (with zero items)", func() {
		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(emptyList, nil)

		_, err := s.cart.Checkout(ctx, fItem.UserID)
		s.ErrorIs(err, ErrEmptyCart)
	})

	s.Run("Unknown error from loms", func() {
		var notEmptyList = []*model.Item{
			{
				ID:     1,
				UserID: 2,
				SKU:    3,
				Count:  4,
			},
		}

		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(notEmptyList, nil)
		s.loms.EXPECT().CreateOrder(gomock.Any(), fItem.UserID, notEmptyList).Return(model.OrderID(0), fakeError)

		_, err := s.cart.Checkout(ctx, fItem.UserID)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Clear basket after checkout", func() {
		var notEmptyList = []*model.Item{
			{
				ID:     1,
				UserID: 2,
				SKU:    3,
				Count:  4,
			},
		}

		var wantOrderID model.OrderID = 1

		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(notEmptyList, nil)
		s.loms.EXPECT().CreateOrder(gomock.Any(), fItem.UserID, notEmptyList).Return(wantOrderID, nil)
		s.rep.EXPECT().DeleteByUser(gomock.Any(), fItem.UserID).Return(fakeError)

		_, err := s.cart.Checkout(ctx, fItem.UserID)
		s.ErrorIs(err, fakeError)
	})

	s.Run("Successful checkout", func() {
		var notEmptyList = []*model.Item{
			{
				ID:     1,
				UserID: 2,
				SKU:    3,
				Count:  4,
			},
		}

		var wantOrderID model.OrderID = 1

		s.rep.EXPECT().FindByUser(gomock.Any(), fItem.UserID).Return(notEmptyList, nil)
		s.loms.EXPECT().CreateOrder(gomock.Any(), fItem.UserID, notEmptyList).Return(wantOrderID, nil)
		s.rep.EXPECT().DeleteByUser(gomock.Any(), fItem.UserID).Return(nil)

		gotOrderID, err := s.cart.Checkout(ctx, fItem.UserID)
		s.Require().NoError(err)
		s.Equal(wantOrderID, gotOrderID)
	})
}

func TestCartTestSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}

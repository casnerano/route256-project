package cart

import (
    "context"
    "errors"
    "testing"

    "github.com/stretchr/testify/suite"
    "go.uber.org/mock/gomock"

    "route256/cart/internal/model"
    mock_repository "route256/cart/internal/repository/mock"
    mock_cart "route256/cart/internal/service/cart/mock"
    "route256/cart/internal/service/client/pim"
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
}

func (s *CartTestSuite) SetupSuite() {
    ctrl := gomock.NewController(s.T())
    defer ctrl.Finish()

    s.rep = mock_repository.NewMockCart(ctrl)
    s.pim = mock_cart.NewMockPIMClient(ctrl)
    s.loms = mock_cart.NewMockLOMSClient(ctrl)

    s.cart = New(s.rep, s.pim, s.loms)
}

func (s *CartTestSuite) TestCartAdd() {
    fItem := fakeItem{
        UserID: 1,
        SKU:    2,
    }

    ctx := context.Background()

    s.Run("Product not found in pim", func() {
        s.pim.EXPECT().GetProductInfo(gomock.Any(), gomock.Any()).Return(nil, pim.ErrProductNotFound)
        _, err := s.cart.Add(ctx, fItem.UserID, fItem.SKU, 1)
        s.ErrorIs(err, ErrPIMProductNotFound)
    })

    s.Run("Unknown error from pim", func() {
        s.pim.EXPECT().GetProductInfo(gomock.Any(), gomock.Any()).Return(nil, fakeError)
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
}

func TestCartTestSuite(t *testing.T) {
    suite.Run(t, new(CartTestSuite))
}

package cart

import (
	"context"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"reflect"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
	mock_repository "route256/cart/internal/repository/mock"
	mock_cart "route256/cart/internal/service/cart/mock"
	"testing"
)

type fixture struct {
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
	//ctx := context.Background()

	s.Run("", func() {
		//item, err := s.cart.Add(ctx, 1, 1, 1)
		//s.Require().NoError(err)
	})

}

func TestCart_Checkout(t *testing.T) {
	type fields struct {
		rep  repository.Cart
		pim  PIMClient
		loms LOMSClient
	}
	type args struct {
		ctx    context.Context
		userID model.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.OrderID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				rep:  tt.fields.rep,
				pim:  tt.fields.pim,
				loms: tt.fields.loms,
			}
			got, err := c.Checkout(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Checkout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_Clear(t *testing.T) {
	type fields struct {
		rep  repository.Cart
		pim  PIMClient
		loms LOMSClient
	}
	type args struct {
		ctx    context.Context
		userID model.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				rep:  tt.fields.rep,
				pim:  tt.fields.pim,
				loms: tt.fields.loms,
			}
			if err := c.Clear(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Clear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCart_Delete(t *testing.T) {
	type fields struct {
		rep  repository.Cart
		pim  PIMClient
		loms LOMSClient
	}
	type args struct {
		ctx    context.Context
		userID model.UserID
		sku    model.SKU
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				rep:  tt.fields.rep,
				pim:  tt.fields.pim,
				loms: tt.fields.loms,
			}
			if err := c.Delete(tt.args.ctx, tt.args.userID, tt.args.sku); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCart_List(t *testing.T) {
	type fields struct {
		rep  repository.Cart
		pim  PIMClient
		loms LOMSClient
	}
	type args struct {
		ctx    context.Context
		wp     *worker_pool.WorkerPool
		userID model.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ItemDetail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				rep:  tt.fields.rep,
				pim:  tt.fields.pim,
				loms: tt.fields.loms,
			}
			got, err := c.List(tt.args.ctx, tt.args.wp, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		rep  repository.Cart
		pim  PIMClient
		loms LOMSClient
	}
	tests := []struct {
		name string
		args args
		want *Cart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rep, tt.args.pim, tt.args.loms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartTestSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}

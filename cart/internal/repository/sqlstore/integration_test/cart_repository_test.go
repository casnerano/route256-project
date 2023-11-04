//go:build integration

package integration_test

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
	"route256/cart/internal/repository/sqlstore"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

var defaultDatabaseDSN = "postgres://user:password@pgbouncer:5432/cart?sslmode=disable"

type fakeItem struct {
	UserID model.UserID
	SKU    model.SKU
	Count  uint32
}

type CartRepositoryTestSuite struct {
	suite.Suite
	pgxpool *pgxpool.Pool
	repo    repository.Cart
}

func (s *CartRepositoryTestSuite) SetupSuite() {
	var err error

	s.pgxpool, err = pgxpool.New(context.Background(), defaultDatabaseDSN)
	s.Require().NoError(err)

	s.repo = sqlstore.NewCartRepository(s.pgxpool)
}

func (s *CartRepositoryTestSuite) SetupTest() {
	s.clearItems()
}

func (s *CartRepositoryTestSuite) TearDownSuite() {
	s.clearItems()
	s.pgxpool.Close()
}

func (s *CartRepositoryTestSuite) TestAdd() {
	fItem := fakeItem{
		UserID: 1,
		SKU:    2,
		Count:  3,
	}

	_, err := s.repo.Add(context.Background(), fItem.UserID, fItem.SKU, fItem.Count)
	s.Require().NoError(err)

	s.Equal(1, s.getItemsCount())
}

func (s *CartRepositoryTestSuite) TestFindByUser() {
	fUserID := model.UserID(1)

	fItems := []fakeItem{
		{
			UserID: fUserID,
			SKU:    2,
			Count:  3,
		},
		{
			UserID: fUserID,
			SKU:    5,
			Count:  6,
		},
		{
			UserID: 7,
			SKU:    8,
			Count:  9,
		},
	}

	for _, fItem := range fItems {
		_, err := s.repo.Add(context.Background(), fItem.UserID, fItem.SKU, fItem.Count)
		s.Require().NoError(err)
	}

	result, err := s.repo.FindByUser(context.Background(), fUserID)
	s.Require().NoError(err)

	s.Equal(2, len(result))
}

func (s *CartRepositoryTestSuite) TestDeleteBySKU() {
	fUserID := model.UserID(1)
	fSKU := model.SKU(2)

	fItems := []fakeItem{
		{
			UserID: fUserID,
			SKU:    fSKU,
			Count:  3,
		},
		{
			UserID: fUserID,
			SKU:    fSKU,
			Count:  6,
		},
		{
			UserID: 7,
			SKU:    8,
			Count:  9,
		},
	}

	for _, fItem := range fItems {
		_, err := s.repo.Add(context.Background(), fItem.UserID, fItem.SKU, fItem.Count)
		s.Require().NoError(err)
	}

	err := s.repo.DeleteBySKU(context.Background(), fUserID, fSKU)
	s.Require().NoError(err)

	s.Equal(1, s.getItemsCount())
}

func (s *CartRepositoryTestSuite) TestDeleteByUser() {
	fUserID := model.UserID(1)

	fItems := []fakeItem{
		{
			UserID: fUserID,
			SKU:    2,
			Count:  3,
		},
		{
			UserID: fUserID,
			SKU:    5,
			Count:  6,
		},
		{
			UserID: 7,
			SKU:    8,
			Count:  9,
		},
	}

	for _, fItem := range fItems {
		_, err := s.repo.Add(context.Background(), fItem.UserID, fItem.SKU, fItem.Count)
		s.Require().NoError(err)
	}

	err := s.repo.DeleteByUser(context.Background(), fUserID)
	s.Require().NoError(err)

	s.Equal(1, s.getItemsCount())
}

func (s *CartRepositoryTestSuite) getItemsCount() int {
	row := s.pgxpool.QueryRow(context.Background(), "SELECT count(*) FROM cart")
	count := 0
	err := row.Scan(&count)
	s.Require().NoError(err)
	return count
}

func (s *CartRepositoryTestSuite) clearItems() {
	_, err := s.pgxpool.Exec(context.Background(), "truncate table cart")
	s.Require().NoError(err)
}

func TestCartRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CartRepositoryTestSuite))
}

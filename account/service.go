package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}
type accountService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &accountService{r}
}

// PostAccount implements Service.
func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	a := &Account{
		Name: name,
		ID:   ksuid.New().String(),
	}
	if err := s.repository.PutAccount(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil

}

// GetAccount implements Service.
func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetAccountByID(ctx, id)
}

// GetAccounts implements Service.
func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 || take == 0) {
		take = 100
	}
	return s.repository.ListAccounts(ctx, skip, take)
}

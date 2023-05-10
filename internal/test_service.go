package internal

import (
	"context"
	"from_scratch_wep_api/config"
	"from_scratch_wep_api/internal/db"
	"from_scratch_wep_api/types"
)

type TestService struct {
	store db.Storer
}

func NewTestService(cfg *config.Config) *TestService {
	svc := &TestService{store: db.NewPostgresProvider(cfg)}
	return svc
}

func (t *TestService) GetTests(ctx context.Context) ([]*types.TestRequest, error) {
	return t.store.GetTests()
}

func (t *TestService) CreateTest(ctx context.Context, tr types.TestRequest) (types.TestRequest, error) {
	return t.store.CreateTest(tr)
}

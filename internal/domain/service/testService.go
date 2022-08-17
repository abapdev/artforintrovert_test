package service

import (
	"artforintrovert_test/internal/adapters/mongo"
	"artforintrovert_test/internal/domain/models"
	"context"
)

type TestService struct {
	ctx        context.Context
	db         *mongo.Storage
	cachedData []models.Data
}

func (t TestService) List(ctx context.Context) ([]models.Data, error) {
	list, err := t.db.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (t TestService) Update(ctx context.Context, data *models.Data) error {
	//TODO implement me
	panic("implement me")
}

func (t TestService) Delete(ctx context.Context, data *models.Data) error {
	//TODO implement me
	panic("implement me")
}

func New(ctx context.Context, db *mongo.Storage) (*TestService, error) {
	list, err := db.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &TestService{
		ctx:        ctx,
		db:         db,
		cachedData: list,
	}, nil
}

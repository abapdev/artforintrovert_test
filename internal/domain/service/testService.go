package service

import (
	"artforintrovert_test/internal/adapters/mongo"
	"context"
)

type TestService struct {
	ctx context.Context
	db  *mongo.Storage
}

func New(ctx context.Context, db *mongo.Storage) *TestService {
	return &TestService{
		ctx: ctx,
		db:  db,
	}
}

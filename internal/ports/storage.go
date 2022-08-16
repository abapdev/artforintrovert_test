package ports

import (
	"artforintrovert_test/internal/domain/models"
	"context"
)

type Storage interface {
	GetList(ctx context.Context) ([]models.Data, error)
	UpdateData(ctx context.Context, data *models.Data) error
	DeleteData(ctx context.Context, data *models.Data) error
}

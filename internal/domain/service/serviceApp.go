package service

import (
	"artforintrovert_test/internal/domain/models"
	"context"
)

type ServiceApp interface {
	//Бизнес логика основного сервиса:
	//Предоставление данных
	List(ctx context.Context, name string) (models.CachedData, error)
	//Обновление данных
	Update(ctx context.Context, data *models.Data) error
	//Удаление данных
	Delete(ctx context.Context, data *models.Data) error
}

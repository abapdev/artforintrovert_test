package service

import (
	"artforintrovert_test/internal/adapters/mongo"
	"artforintrovert_test/internal/domain/models"
	"context"
	"sync"
	"time"
)

type TestService struct {
	ctx        context.Context
	db         *mongo.Storage
	cachedData models.CachedData
}

func (t *TestService) List(ctx context.Context, name string) (models.CachedData, error) {
	//list, err := t.db.GetList(ctx, name)
	//if err != nil {
	//	return nil, err
	//}
	//return list, nil
	if name != "" {
		singleData := make(models.CachedData)
		value, ok := t.cachedData[name]
		if ok != true {
			return nil, nil
		}
		singleData[name] = value
		return singleData, nil
	}
	return t.cachedData, nil
}

func (t *TestService) Update(ctx context.Context, data *models.DataJSON) error {
	if err := t.db.UpdateData(ctx, data); err != nil {
		return err
	}
	return nil
}

func (t *TestService) Delete(ctx context.Context, data *models.DataJSON) error {
	if err := t.db.DeleteData(ctx, data); err != nil {
		return err
	}
	return nil
}

func New(ctx context.Context, db *mongo.Storage) (*TestService, error) {

	list, err := db.GetList(ctx, "")
	if err != nil {
		return &TestService{
			ctx: ctx,
			db:  db,
			//cachedData: cachMap,
		}, err
	}
	cachMap := make(models.CachedData)
	for _, value := range list {
		cachMap[value.Name] = value.Phone
	}
	return &TestService{
		ctx:        ctx,
		db:         db,
		cachedData: cachMap,
	}, nil
}
func (t *TestService) Refresh(ctx context.Context) {
	var mutex sync.Mutex
	go func(mutex *sync.Mutex) {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				time.Sleep(60 * time.Second)
				list, err := t.db.GetList(ctx, "")
				if err != nil {
					return
				}
				cachMap := make(models.CachedData)
				for _, value := range list {
					cachMap[value.Name] = value.Phone
				}
				mutex.Lock()
				t.cachedData = cachMap
				mutex.Unlock()
			}
		}
	}(&mutex)
}

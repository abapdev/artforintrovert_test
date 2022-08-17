package mongo

import (
	"artforintrovert_test/internal/config"
	"artforintrovert_test/internal/domain/models"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	ConnectionString string
}

func New(cfg *config.Config) (*Storage, error) {
	connectString := cfg.Listen.Mongo.Connect
	logrus.Info(connectString)
	cOpts := options.Client().ApplyURI(connectString)
	mClient, err := mongo.Connect(context.Background(), cOpts)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := mClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	err = mClient.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	//обойдемся без миграций, просто наполним БД хардкодом для теста API
	//mCollection := mClient.Database("artforintrovert").Collection("test")
	//data := []models.Data{
	//	{
	//		ID:    primitive.NewObjectID(),
	//		Name:  "Jhon",
	//		Phone: "11-22-33",
	//	},
	//	{
	//		ID:    primitive.NewObjectID(),
	//		Name:  "Boris",
	//		Phone: "44-55-66",
	//	},
	//	{
	//		ID:    primitive.NewObjectID(),
	//		Name:  "Serena",
	//		Phone: "77-88-99",
	//	},
	//	{
	//		ID:    primitive.NewObjectID(),
	//		Name:  "Jack",
	//		Phone: "00-11-22",
	//	},
	//}
	//items := make([]interface{}, 0, len(data))
	//for _, b := range data {
	//	items = append(items, b)
	//}
	//
	//res, err := mCollection.InsertMany(context.Background(), items)
	//if err != nil {
	//	panic(err)
	//}
	//logrus.Info("inserted ids: %v", res.InsertedIDs)

	return &Storage{
		ConnectionString: connectString,
	}, nil
}
func (st *Storage) GetList(ctx context.Context, name string) ([]models.Data, error) {
	logrus.Info("Refresh data from Mongo Storage")
	sd, err := st.readFromMongoDB(ctx, &models.DataJSON{
		Name:  name,
		Phone: "",
	})
	if err != nil {
		logrus.WithError(err).Error("Can't read from MongoDB")
		return nil, err
	}
	return sd, nil
}
func (st *Storage) UpdateData(ctx context.Context, dataJson *models.DataJSON) error {
	logrus.Info("UpdateData in Mongo Storage")
	cOpts := options.Client().ApplyURI(st.ConnectionString)
	mClient, err := mongo.Connect(ctx, cOpts)
	if err != nil {
		return err
	}

	defer func() {
		if err := mClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mCollection := mClient.Database("artforintrovert").Collection("test")
	data := models.Data{
		ID:    primitive.NewObjectID(),
		Name:  dataJson.Name,
		Phone: dataJson.Phone,
	}
	uOpts := &options.UpdateOptions{}
	filter := bson.D{{"Name", data.Name}}
	update := bson.D{{"$set", bson.D{{"Phone", data.Phone}}}}
	res, err := mCollection.UpdateOne(ctx, filter, update, uOpts)
	if err != nil {
		return err
	}
	logrus.Info("Updated ids: ", res.ModifiedCount)
	return nil
}
func (st *Storage) DeleteData(ctx context.Context, data *models.DataJSON) error {
	logrus.Info("UpdateData in Mongo Storage")
	cOpts := options.Client().ApplyURI(st.ConnectionString)
	mClient, err := mongo.Connect(ctx, cOpts)
	if err != nil {
		return err
	}

	defer func() {
		if err := mClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mCollection := mClient.Database("artforintrovert").Collection("test")

	dOpts := &options.DeleteOptions{}
	filter := bson.D{{"Name", data.Name}}
	res, err := mCollection.DeleteOne(ctx, filter, dOpts)
	if err != nil {
		return err
	}
	logrus.Info("Delete ids: ", res.DeletedCount)
	return nil
}
func (st *Storage) readFromMongoDB(ctx context.Context, data *models.DataJSON) ([]models.Data, error) {
	cOpts := options.Client().ApplyURI(st.ConnectionString)
	mClient, err := mongo.Connect(ctx, cOpts)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := mClient.Disconnect(ctx); err != nil {
			logrus.WithError(err).Error("Cannot close connection with MongoDB")
		}
	}()

	mCollection := mClient.Database("artforintrovert").Collection("test")

	filter := bson.D{}
	if data.Name != "" {
		filter = bson.D{
			{"Name", bson.D{{"$eq", data.Name}}},
		}
	}

	cursor, err := mCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var dataSlice []models.Data
	for cursor.Next(ctx) {
		var dt models.Data
		if err := cursor.Decode(&dt); err != nil {
			break
		}
		//logrus.Info("find data item: ", dt.Name)

		dataSlice = append(dataSlice, dt)
	}
	if len(dataSlice) == 0 {
		return nil, errors.New("Data not found")
	}
	return dataSlice, nil
}

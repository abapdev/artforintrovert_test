package mongo

import (
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

func New(ConnectionString string) (*Storage, error) {
	cOpts := options.Client().ApplyURI(ConnectionString)
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
	return &Storage{
		ConnectionString: ConnectionString,
	}, nil
}
func (st *Storage) GetList(ctx context.Context) ([]models.Data, error) {
	logrus.Info("GetList in Mongo Storage")
	sd, err := st.readFromMongoDB(ctx, &models.Data{
		ID:    primitive.ObjectID{},
		Name:  "",
		Phone: "",
	})
	if err != nil {
		logrus.WithError(err).Error("Can't read user from MongoDB")
		return nil, err
	}
	return sd, nil
}
func (st *Storage) UpdateData(ctx context.Context, data *models.Data) error {
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

	uOpts := &options.UpdateOptions{}
	res, err := mCollection.UpdateByID(ctx, data.ID, data, uOpts)
	if err != nil {
		return err
	}
	logrus.Info("Updated ids: ", res.UpsertedID)
	return nil
}
func (st *Storage) DeleteData(ctx context.Context, data *models.Data) error {
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
	filter := bson.D{
		{"Name", bson.D{{"$eq", data.Name}}},
	}
	res, err := mCollection.DeleteOne(ctx, filter, dOpts)
	if err != nil {
		return err
	}
	logrus.Info("Delete ids: ", res.DeletedCount)
	return nil
}
func (st *Storage) readFromMongoDB(ctx context.Context, data *models.Data) ([]models.Data, error) {
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
		logrus.Info("find data item: ", dt.Name)
		dataSlice = append(dataSlice, dt)
	}
	if len(dataSlice) == 0 {
		return nil, errors.New("Data not found")
	}
	return dataSlice, nil
}

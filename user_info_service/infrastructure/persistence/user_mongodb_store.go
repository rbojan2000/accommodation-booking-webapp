package persistence

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func (u UserMongoDBStore) GetUserByID(id primitive.ObjectID) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*model.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (products []*model.User, err error) {
	for cursor.Next(context.TODO()) {
		var product model.User
		err = cursor.Decode(&product)
		if err != nil {
			return
		}
		products = append(products, &product)
	}
	err = cursor.Err()
	return
}

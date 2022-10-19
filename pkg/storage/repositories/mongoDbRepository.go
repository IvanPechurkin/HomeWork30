package repositories

import (
	"30/cmd/utils"
	"30/pkg/model"
	"30/pkg/storage/interfaces"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoDbRepository struct {
	db     *mongo.Client
	users  *mongo.Collection
	config *utils.Configuration
}

func NewMongoDbRepository(db *mongo.Client, cfg *utils.Configuration) interfaces.Repository {
	return &MongoDbRepository{
		db:     db,
		users:  db.Database(cfg.Database.DbName).Collection(cfg.Database.Collection),
		config: cfg,
	}
}

// Добавление пользователя в хранилище
func (r MongoDbRepository) AddUser(ctx context.Context, user *model.User) {
	_, err := r.users.InsertOne(ctx, user)
	if err != nil {
		return
	}
}

// Получение пользователя по идентификатору
func (r MongoDbRepository) Get(ctx context.Context, userId int) (*model.User, error) {
	filter := bson.D{{"_id", userId}}

	var user *model.User
	err := r.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Получение всех пользователей в хранилище
func (r MongoDbRepository) GetAll(ctx context.Context) []*model.User {
	users := make([]*model.User, 0)
	filter := bson.D{{}}

	result, err := r.users.Find(ctx, filter)
	if err != nil {
		return nil
	}

	for result.Next(ctx) {
		var user *model.User
		err := result.Decode(&user)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	if err := result.Err(); err != nil {
		return nil
	}

	result.Close(ctx)

	return users
}

// Получение друзей пользователя
func (r MongoDbRepository) GetFriends(ctx context.Context, userId int) ([]*model.User, error) {
	friends := make([]*model.User, 0)
	filter := bson.D{{"_id", userId}}

	var user *model.User
	err := r.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	for _, val := range user.Friends {
		friend, err := r.Get(ctx, val)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}

// Удаление пользователя из хранилища
func (r MongoDbRepository) DeleteUser(ctx context.Context, userId int) error {
	filter := bson.M{"_id": userId}

	_, err := r.users.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// Обновление возраста пользователя
func (r MongoDbRepository) UpdateUserAge(ctx context.Context, userId int, age int) error {
	_, err := r.users.UpdateOne(
		ctx,
		bson.M{"_id": userId},
		bson.D{
			{"$set", bson.D{{"age", age}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// Добавление пользователя в друзья
func (r MongoDbRepository) LinkUsers(ctx context.Context, userLinkFrom int, userLinkTo int) error {
	return nil
}

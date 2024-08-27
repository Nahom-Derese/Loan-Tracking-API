package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	"github.com/Nahom-Derese/Loan-Tracking-API/domain/forms"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database       mongo.Database
	collectionName string
}

func NewUserRepository(db mongo.Database, collection string) entities.UserRepository {
	return &userRepository{
		database:       db,
		collectionName: collection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *entities.User) (*entities.User, error) {

	collection := ur.database.Collection(ur.collectionName)

	res, err := collection.InsertOne(c, user)

	if err != nil {
		return nil, err
	}
	// Find the inserted user by ID
	insertedID, _ := res.InsertedID.(primitive.ObjectID)
	var insertedUser entities.User
	err = collection.FindOne(c, bson.M{"_id": insertedID}).Decode(&insertedUser)

	if err != nil {
		return nil, custom_error.ErrErrorCreatingUser
	}

	return &insertedUser, err
}
func (ur *userRepository) IsOwner(c context.Context) (bool, error) {
	collection := ur.database.Collection(ur.collectionName)
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return false, custom_error.ErrErrorCreatingUser
	}
	return count == 0, nil
}
func (ur *userRepository) GetAllUsers(c context.Context) ([]entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	var users []entities.User
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return []entities.User{}, err
	}
	err = cursor.All(c, &users)
	if err != nil {
		return []entities.User{}, err
	}

	return users, nil

}
func (ur *userRepository) UpdateRefreshToken(c context.Context, userID string, refreshToken string) error {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	_, err = collection.UpdateOne(c, filter, bson.M{"$push": bson.M{"tokens": refreshToken}})
	return err
}

func (ur *userRepository) UpdateLastLogin(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	_, err = collection.UpdateOne(c, filter, bson.M{"$set": bson.M{"last_login": primitive.NewDateTimeFromTime(time.Now())}})

	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}

	return err
}

func (ur *userRepository) GetUserById(c context.Context, userId string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	var user entities.User

	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, custom_error.ErrInvalidID
	}

	err = collection.FindOne(c, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, custom_error.ErrUserNotFound
	}

	return &user, err

}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	var user entities.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, custom_error.ErrUserNotFound
	}
	return &user, err
}

func (ur *userRepository) GetUsers(c context.Context, filter bson.M, userFilter entities.UserFilter) (*[]entities.User, mongopagination.PaginationData, error) {
	collection := ur.database.Collection(ur.collectionName)

	projectQuery := bson.M{"$project": bson.M{
		"password":     0,
		"tokens":       0,
		"verfiy_token": 0,
		"is_owner":     0,
	}}

	var aggUserList []entities.User = make([]entities.User, 0)

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(userFilter.Limit).Page(userFilter.Pages).Aggregate(filter, projectQuery)

	if err != nil {
		return &[]entities.User{}, mongopagination.PaginationData{}, custom_error.ErrFilteringUsers
	}

	for _, raw := range paginatedData.Data {
		var user *entities.User
		if marshallErr := bson.Unmarshal(raw, &user); marshallErr == nil {
			aggUserList = append(aggUserList, *user)
		}

	}

	return &aggUserList, paginatedData.Pagination, nil

}

func (ur *userRepository) RevokeRefreshToken(c context.Context, userID, refreshToken string) error {
	collection := ur.database.Collection(ur.collectionName)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": objID}, bson.M{"$pull": bson.M{"tokens": refreshToken}})

	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return custom_error.ErrTokenNotFound
	}
	return nil
}

func (ur *userRepository) UpdateUser(c context.Context, userID string, updatedUser *forms.RegisterUserForm) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedUser}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser entities.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, custom_error.ErrErrorUpdatingUser
	}

	return &ResultUser, nil
}

func (ur *userRepository) ActivateUser(c context.Context, userID string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_active": true}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser entities.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, custom_error.ErrErrorUpdatingUser
	}

	return &ResultUser, nil
}

func (ur *userRepository) DeleteUser(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(c, filter)

	if res.DeletedCount == 0 {
		return custom_error.ErrUserNotFound
	}

	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}

	return nil
}
func (ur *userRepository) IsUserActive(c context.Context, userID string) (bool, error) {
	collection := ur.database.Collection(ur.collectionName)
	var user entities.User
	err := collection.FindOne(c, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return false, custom_error.ErrUserNotFound
	}
	return user.Active, err

}
func (ur *userRepository) ResetUserPassword(c context.Context, userID string, resetPassword *forms.ResetPasswordForm) error {
	collection := ur.database.Collection(ur.collectionName)
	ObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": resetPassword.NewPassword}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}
	return nil
}
func (ur *userRepository) UpdateUserPassword(c context.Context, userID string, updatePassword *forms.UpdatePasswordForm) error {
	collection := ur.database.Collection(ur.collectionName)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": updatePassword.NewPassword}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}
	return nil
}

// GetRefreshToken implements entities.UserRepository.
func (ur *userRepository) RefreshTokenExist(c context.Context, userID, refreshToken string) (bool, error) {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, err
	}
	filter := bson.M{
		"_id":    id,
		"tokens": refreshToken, // Check if the refreshToken exists tokens[]
	}
	err = collection.FindOne(c, filter).Decode(&entities.User{})
	if err != nil {
		return false, nil
	}

	return true, nil

}

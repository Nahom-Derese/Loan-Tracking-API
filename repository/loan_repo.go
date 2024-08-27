package repository

import (
	"context"

	"github.com/Nahom-Derese/Loan-Tracking-API/domain/entities"
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type loanRepository struct {
	database       mongo.Database
	collectionName string
}

// NewLoanRepository creates a new instance of loanRepository.
func NewLoanRepository(db mongo.Database, collection string) entities.LoanRepository {
	return &loanRepository{
		database:       db,
		collectionName: collection,
	}
}

// CreateLoan inserts a new loan into the database.
func (lr *loanRepository) CreateLoan(ctx context.Context, loan *entities.Loan) (*entities.Loan, error) {
	collection := lr.database.Collection(lr.collectionName)

	res, err := collection.InsertOne(ctx, loan)
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted loan by its ID
	insertedID, _ := res.InsertedID.(primitive.ObjectID)
	var insertedLoan entities.Loan
	err = collection.FindOne(ctx, bson.M{"_id": insertedID}).Decode(&insertedLoan)
	if err != nil {
		return nil, custom_error.ErrErrorCreatingLoan
	}

	return &insertedLoan, nil
}

// GetLoanByID retrieves a loan by its ID from the database.
func (lr *loanRepository) GetLoanByID(ctx context.Context, loanID string) (*entities.Loan, error) {
	collection := lr.database.Collection(lr.collectionName)
	var loan entities.Loan

	id, err := primitive.ObjectIDFromHex(loanID)
	if err != nil {
		return nil, custom_error.ErrInvalidID
	}

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&loan)
	if err != nil {
		return nil, custom_error.ErrLoanNotFound
	}

	return &loan, nil
}

func (lr *loanRepository) GetLoans(ctx context.Context, limit int64, page int64) (*[]entities.Loan, mongopagination.PaginationData, error) {
	collection := lr.database.Collection(lr.collectionName)

	var aggLoanList []entities.Loan = make([]entities.Loan, 0)

	paginatedData, err := mongopagination.New(collection).Context(ctx).Limit(limit).Page(page).Aggregate()

	if err != nil {
		return &[]entities.Loan{}, mongopagination.PaginationData{}, custom_error.ErrFilteringUsers
	}

	for _, raw := range paginatedData.Data {
		var loan *entities.Loan
		if marshallErr := bson.Unmarshal(raw, &loan); marshallErr == nil {
			aggLoanList = append(aggLoanList, *loan)
		}

	}

	return &aggLoanList, paginatedData.Pagination, nil

}

// DeleteLoan removes a loan by its ID from the database.
func (lr *loanRepository) DeleteLoan(ctx context.Context, loanID string) error {
	collection := lr.database.Collection(lr.collectionName)

	id, err := primitive.ObjectIDFromHex(loanID)
	if err != nil {
		return custom_error.ErrInvalidID
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return custom_error.ErrLoanNotFound
	}

	return nil
}

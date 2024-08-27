package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName       string             `bson:"first_name" json:"firstName" validate:"required,min=2,max=100"`
	LastName        string             `bson:"last_name" json:"lastName" validate:"required,min=2,max=100"`
	Email           string             `bson:"email" json:"email" validate:"required,email"`
	Password        string             `bson:"password" json:"-" validate:"required,min=6"`
	Phone           string             `bson:"phone" json:"phone" validate:"required,e164"`
	Address         string             `bson:"address" json:"address" validate:"required,min=5,max=200"`
	Role            string             `bson:"role" json:"role" validate:"required,oneof=user admin"`
	TotalLoanAmount float64            `bson:"total_loan_amount" json:"totalLoanAmount" validate:"gte=0"`
	OutstandingDebt float64            `bson:"outstanding_debt" json:"outstandingDebt" validate:"gte=0"`
	Loans           []Loan             `bson:"loans" json:"loans"`
	CreatedAt       int64              `bson:"created_at" json:"createdAt"`
	UpdatedAt       int64              `bson:"updated_at" json:"updatedAt"`
}

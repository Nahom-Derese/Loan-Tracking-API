package forms

type RegisterUserForm struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=100"`
	LastName  string `json:"lastName" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Phone     string `json:"phone" validate:"required,e164"`
	Address   string `json:"address" validate:"required,min=5,max=200"`
}

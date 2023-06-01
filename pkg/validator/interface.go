package validator

type validator interface {
	Password(password string) error
	Email(email string) error
}

type Validator struct {
}

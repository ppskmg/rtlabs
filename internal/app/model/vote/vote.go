package vote

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Vote struct {
	WorkID    int64  `json:"work_id"`
	UserEmail string `json:"user_email"`
	Contest   string `json:"contest"`
}

type Count struct {
	WorkID int64 `json:"work_id"`
	Count  int64 `json:"count"`
}

type Winner struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

func (s *Vote) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.WorkID, validation.Required),
		validation.Field(&s.UserEmail, validation.Required),
		validation.Field(&s.Contest, validation.Required),
	)
}
func (s *Winner) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Email, validation.Required),
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
		validation.Field(&s.MiddleName, validation.Required),
	)
}

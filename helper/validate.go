package helper

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GoInn/internal/dto"
	"regexp"
)

const (
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

func Validate(user *dto.User) error {
	if len(user.FirstName) < minFirstNameLen {
		return fmt.Errorf("first name must be at least %d characters", minFirstNameLen)
	}
	if len(user.LastName) < minLastNameLen {
		return fmt.Errorf("last name must be at least %d characters", minLastNameLen)
	}
	if len(user.Password) < minPasswordLen {
		return fmt.Errorf("password must be at least %d characters", minPasswordLen)
	}
	if !isEmailValid(user.Email) {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func isEmailValid(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

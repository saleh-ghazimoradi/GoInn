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

func Validate(user *dto.User) map[string]string {
	errors := map[string]string{}
	if len(user.FirstName) < minFirstNameLen {
		errors["firstname"] = fmt.Sprintf("first name must be at least %d characters", minFirstNameLen)
	}
	if len(user.LastName) < minLastNameLen {
		errors["lastname"] = fmt.Sprintf("last name must be at least %d characters", minLastNameLen)
	}
	if len(user.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password must be at least %d characters", minPasswordLen)
	}
	if !isEmailValid(user.Email) {
		errors["email"] = fmt.Sprintf("invalid email")
	}
	return errors
}

func isEmailValid(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

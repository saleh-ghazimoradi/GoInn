package helper

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GoInn/logger"
	"net/http"
)

type Error struct {
	logger *logger.Logger
}

func (e *Error) LogError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	e.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (e *Error) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	env := Envelope{"error": message}
	if err := WriteJSON(w, status, env, nil); err != nil {
		e.LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (e *Error) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.LogError(r, err)
	message := "the server encountered a problem and could not process your request"
	e.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func (e *Error) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (e *Error) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	e.ErrorResponse(w, r, http.StatusNotFound, message)
}

func (e *Error) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	env := Envelope{"errors": errors}
	if err := WriteJSON(w, http.StatusUnprocessableEntity, env, nil); err != nil {
		e.LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (e *Error) InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	e.ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func (e *Error) RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	e.ErrorResponse(w, r, http.StatusTooManyRequests, message)
}

func (e *Error) EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	e.ErrorResponse(w, r, http.StatusConflict, message)
}

func (e *Error) InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	e.ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func (e *Error) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	e.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (e *Error) AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	e.ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func (e *Error) NotPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account does not have the necessary permissions to access this resource"
	e.ErrorResponse(w, r, http.StatusForbidden, message)
}

func (e *Error) InactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	e.ErrorResponse(w, r, http.StatusForbidden, message)
}

func NewError(logger *logger.Logger) *Error {
	return &Error{logger: logger}
}

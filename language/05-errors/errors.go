/*
Package errorkata goes deep on Go's error handling: wrapping with %w,
sentinel errors and errors.Is, custom error types and errors.As,
errors.Join, and turning a panic back into an error with defer/recover.
*/
package errorkata

import "errors"

var (
	ErrAgeOutOfRange = errors.New("age out of range")
	ErrNotFound      = errors.New("not found")
	ErrPanicked      = errors.New("recovered from panic")
)

func ParseAge(s string) (int, error) {
	return 0, nil
}

type User struct {
	ID    int
	Name  string
	Email string
}

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return ""
}

func ValidateUser(u User) error {
	return nil
}

type UserStore struct {
	users map[int]User
}

func NewUserStore(users ...User) *UserStore {
	return &UserStore{}
}

func (s *UserStore) FindUser(id int) (User, error) {
	return User{}, nil
}

func (s *UserStore) EmailFor(id int) (string, error) {
	return "", nil
}

func SafeCall(fn func()) error {
	return nil
}

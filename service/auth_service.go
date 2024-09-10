package service

import (
	"context"
	"errors"
	"github/luqxus/spxce/database"
	"github/luqxus/spxce/tokens"
	"github/luqxus/spxce/types"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	datastore database.Database
}

// create a new user
// prepares user for insertion in data store
func (s *AuthService) CreateUser(ctx context.Context, data *types.CreateUserRequest) (string, error) {
	// check if provided email exists in the data store
	count, err := s.datastore.CountEmail(ctx, data.Email)
	if err != nil {
		return "", err
	}

	// check if count is
	// greater than 0 => email already in use
	// else => email available for use
	if count > 0 {
		// return error if email already in use
		return "", errors.New("email already in use")
	}

	user := new(types.User)

	// generate new user uid
	uid := uuid.NewString()

	user.UID = uid
	user.Email = data.Email
	user.Username = data.Username
	user.Password, _ = hashPassword(data.Password)
	// hash password

	// call CreateUser in datastore to create new user
	err = s.datastore.CreateUser(ctx, user)
	if err != nil {
		// return empty string and error on failure
		return "", err
	}

	// generate new jwt token from uid and email
	token, err := tokens.GenerateJWT(uid, data.Email)

	// return jwt token and nil on success
	return token, err
}

// gets user from datastore with matching email
// compare found user password and given password
// return jwt token and nil or empty string and error on failure
func (s *AuthService) Login(ctx context.Context, data *types.LoginRequest) (string, error) {
	// get user from datastor matching email
	user, err := s.datastore.GetUser(ctx, data.Email)
	if err != nil {
		return "", err
	}

	// check if found password matches given password
	err = validatePassword(user.Password, data.Password)
	if err != nil {
		// return empty string and error if no match
		return "", errors.New("wrong email or password")
	}

	// generate a new jwt token
	token, _ := tokens.GenerateJWT(user.UID, user.Email)

	return token, nil
}

// hash plain password
func hashPassword(password string) (string, error) {

	// generate new hash
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(b), err
}

// compare given password and found password
// return nil on match otherwise error
func validatePassword(foundPassword, givenPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(foundPassword), []byte(givenPassword))
}

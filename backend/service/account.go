package service

import (
	"main/database"
	"main/errors"
	"main/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(request *models.SignupRequest) (*models.User, error) {
	if _, err := database.GetUserByUsername(request.Username); err != nil {
		if !errors.Is(err, errors.ErrUserNotFound) {
			return nil, err
		}
	} else {
		return nil, errors.ErrUsernameAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.NewString(),
		Username: request.Username,
		Password: hash,
	}

	if err := database.InsertUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func Signin(request *models.SigninRequest) (*models.Session, error) {
	user, err := database.GetUserByUsername(request.Username)
	if err != nil {
		if errors.Is(err, errors.ErrUserNotFound) {
			return nil, errors.ErrInvalidUsernameOrPassword
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
		// Return generic error to prevent password timing attacks
		return nil, errors.ErrInvalidUsernameOrPassword
	}

	session := models.NewSession(user.ID)

	if err := database.InsertSession(session); err != nil {
		return nil, err
	}

	return session, nil
}

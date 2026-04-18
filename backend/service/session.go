package service

import (
	"main/database"
	"main/errors"
	"main/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(request *models.SignupRequest) (*models.Session, error) {
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

	session, err := CreateSession(user)
	if err != nil {
		return nil, err
	}

	return session, nil
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
		return nil, errors.ErrInvalidUsernameOrPassword
	}

	session, err := CreateSession(user)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func GetSessionStatus(id string) (*models.SessionStatusResponse, error) {
	session, err := database.GetSession(id)
	if err != nil {
		if errors.Is(err, errors.ErrSessionNotFound) {
			return &models.SessionStatusResponse{}, nil
		}
	}

	return &models.SessionStatusResponse{
		Exists:  true,
		Expired: time.Now().After(session.ExpiresAt),
		UserID:  session.UserID,
	}, nil
}

func CreateSession(user *models.User) (*models.Session, error) {
	session := &models.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour),
	}

	if err := database.InsertSession(session); err != nil {
		return nil, err
	}

	return session, nil
}

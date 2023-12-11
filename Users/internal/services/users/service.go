package services

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/users"
	"net/http"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	Users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving Users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return Users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	User, err := repository.GetUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving Users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return User, err
}

func PostUser(User models.User) (*models.User, error) {

	User, err := repository.PostUser(User)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving Users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return User, err
}

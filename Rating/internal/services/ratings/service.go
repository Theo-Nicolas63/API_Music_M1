package ratings

import (
	"database/sql"
	"errors"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllratings() ([]models.Rating, error) {
	var err error
	// calling repository
	ratings, err := repository.GetAllratings()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving ratings : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return ratings, nil
}

func GetratingById(id uuid.UUID) (*models.Rating, error) {
	rating, err := repository.GetratingById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving ratings : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}

func PostRating(rating *models.Rating) (*models.Rating, error) {

	rating, err := repository.Postrating(rating)

	if err != nil {
		logrus.Errorf("error retrieving ratings : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}

func PutRating(rating *models.Rating) (*models.Rating, error) {

	rating, err := repository.Putrating(rating)

	if err != nil {
		logrus.Errorf("error retrieving ratings : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}

func DeleteRating(id uuid.UUID) (*models.Rating, error) {

	rating, err := repository.Deleterating(id)

	if err != nil {
		logrus.Errorf("error retrieving ratings : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rating, err
}

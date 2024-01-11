package ratings

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"
)

func PutRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	ratingIDstr := chi.URLParam(r, "id")
	ratingID, err2 := uuid.FromString(ratingIDstr)

	if err2 != nil {
		logrus.Errorf("error parsing rating ID : %s", err2.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&rating)

	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rating.Id = &ratingID
	_, err = ratings.PutRating(&rating)

	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)

		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(rating)
	_, _ = w.Write(body)
	return
}

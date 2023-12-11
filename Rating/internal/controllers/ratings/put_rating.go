package ratings

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"
)

func PutRating(w http.ResponseWriter, r *http.Request) {
	var rating models.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)

	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

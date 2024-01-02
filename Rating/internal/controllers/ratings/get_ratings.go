package ratings

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Getratings
// @Tags         ratings
// @Summary      Get ratings.
// @Description  Get ratings.
// @Success      200            {array}  models.rating
// @Failure      500             "Something went wrong"
// @Router       /ratings [get]
func GetRatings(w http.ResponseWriter, _ *http.Request) {
	// calling service
	ratings, err := ratings.GetAllratings()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(ratings)
	_, _ = w.Write(body)
	return
}

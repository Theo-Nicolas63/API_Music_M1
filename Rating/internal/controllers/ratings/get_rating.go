package ratings

import (
	"encoding/json"
	"fmt"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// Getrating
// @Tags         ratings
// @Summary      Get a rating.
// @Description  Get a rating.
// @Param        id           	path      string  true  "rating UUID formatted ID"
// @Success      200            {object}  models.rating
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [get]
func Getrating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("ratingId").(uuid.UUID)
	fmt.Println(ratingId)
	rating, err := ratings.GetratingById(ratingId)
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

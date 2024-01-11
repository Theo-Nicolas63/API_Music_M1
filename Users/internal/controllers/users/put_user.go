package users

import (
	"encoding/json"
	"middleware/example/internal/models"
	"net/http"
	"middleware/example/internal/services/users"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func PutUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	UserIDstr := chi.URLParam(r, "id")
	UserID, err2 := uuid.FromString(UserIDstr)

	if err2 != nil {
		logrus.Errorf("error parsing user ID : %s", err2.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&User)

	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	User.Id = &UserID

	_, err = services.PutUser(&User)

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
	body, _ := json.Marshal(User)
	_, _ = w.Write(body)
	return
}

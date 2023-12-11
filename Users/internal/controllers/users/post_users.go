package users

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	services "middleware/example/internal/services/users"
	"net/http"
)

// GetCollections
// @Tags         users
// @Summary      Get users.
// @Description  Get users.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /users [get]

func PostUsers(w http.ResponseWriter, _ *http.Request) {
	// calling service

	//Scan JSON + recup variable user a metre dans fonction ci dessous
	Users, err := services.PostUser()

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
	body, _ := json.Marshal(Users)
	_, _ = w.Write(body)
	return
}

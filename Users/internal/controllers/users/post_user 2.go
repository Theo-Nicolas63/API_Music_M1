package users

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
	"net/http"
)

// GetCollections
// @Tags         users
// @Summary      Get users.
// @Description  Get users.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /users [get]

func PostUser(w http.ResponseWriter, r *http.Request) {
	// calling service

	//Scan JSON + recup variable user a metre dans fonction ci dessous

	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	fmt.Print(newUser)
	user, err := services.PostUser(&newUser)

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

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}

package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PostCollection(w http.ResponseWriter, r *http.Request) {
    var collection models.Collection
    err := json.NewDecoder(r.Body).Decode(&collection)
    if err != nil {
        logrus.Errorf("error : %s", err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    collection, err = PostCollection(&collection)
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
    body, _ := json.Marshal(collection)
    _, _ = w.Write(body)
    return
}

package songs

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/songs"
	"net/http"
)

//func PutSong(w http.ResponseWriter, r *http.Request) {
//	var song models.Song
//	err := json.NewDecoder(r.Body).Decode(&song)
//	print("Debug 1 \n ")
//	if err != nil {
//		logrus.Errorf("error : %s", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	print("Debug 2 \n ")
//	_, err = songs.PutSong(&song)
//	print("Debug 3 \n ")
//	if err != nil {
//		logrus.Errorf("error : %s", err.Error())
//		customError, isCustom := err.(*models.CustomError)
//		print("Debug 4 \n ")
//		if isCustom {
//			w.WriteHeader(customError.Code)
//			print("Debug 5 \n ")
//			body, _ := json.Marshal(customError)
//			_, _ = w.Write(body)
//		} else {
//			w.WriteHeader(http.StatusInternalServerError)
//			print("Debug 6 \n ")
//		}
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	body, _ := json.Marshal(song)
//	_, _ = w.Write(body)
//	return
//}

func PutSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	songIDstr := chi.URLParam(r, "id")
	songID, err := uuid.FromString(songIDstr)
	if err != nil {
		logrus.Errorf("error parsing song ID : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = songs.PutSong(songID, song.Name, song.Singer)
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
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}

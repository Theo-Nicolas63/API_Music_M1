package songs

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"
)

// GetSong
// @Tags         songs
// @Summary      Get a Song by Name.
// @Description  Get a Song by Name.
// @Param        id           	path      string  true  "Song Name"
// @Success      200            {object}  models.Song
// @Failure      422            "Invalid name"
// @Failure      500            "Something went wrong"
// @Router       /songs/name/{name} [get]
func GetSongByName(w http.ResponseWriter, r *http.Request) {

	songName := chi.URLParam(r, "name")
	//ctx := r.Context()
	//var songName string = r.URL.Query().Get("name")
	//songName, _ = ctx.Value("songName").(string)

	song, err := songs.GetSongByName(songName)
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

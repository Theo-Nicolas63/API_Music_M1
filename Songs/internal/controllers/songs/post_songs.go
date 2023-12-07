package songs

import (
	"encoding/json"
	"middleware/example/internal/models"

	//	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	//"middleware/example/internal/models"
	"middleware/example/internal/repositories/songs"
	"net/http"
)

// PostSong
// @Tags         songs
// @Summary      Add a new song.
// @Description  Add a new song to the collection.
// @Accept       json
// @Produce      json
// @Param        song body     models.Song  true  "Song to add"
// @Success      201    {object}  string     "Successfully added"
// @Failure      400    "Invalid input"
// @Failure      500    "Something went wrong"
// @Router       /songs [post]

func PostSong(w http.ResponseWriter, r *http.Request) {
	var newSong models.Song
	err := json.NewDecoder(r.Body).Decode(&newSong)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Décodage du payload JSON entrant
	/*if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		logrus.Errorf("error decoding song: %s", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}*/

	song, err := songs.AddSong(newSong)
	if err != nil {
		logrus.Errorf("error adding song: %s", err.Error())
		http.Error(w, "Failed to add the song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
}

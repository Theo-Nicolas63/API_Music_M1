package songs

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Content)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Content)
	if err != nil {
		return nil, err
	}
	return &song, err
}

// AddSong ajoute une nouvelle chanson à la base de données
func AddSong(song models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	randomUUID, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}
	_, err = db.Exec("INSERT INTO songs (id, content) VALUES (?, ?)", randomUUID.String(), song.Content)
	row := db.QueryRow("SELECT * FROM songs WHERE id = ?", randomUUID)
	helpers.CloseDB(db)
	err = row.Scan(&song.Id, &song.Content)
	if err != nil {
		return &song, err
	}

	//song, _ = GetSongById(randomUUID)

	return &song, err
}

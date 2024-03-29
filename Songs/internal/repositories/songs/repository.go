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
		err = rows.Scan(&data.Id, &data.Name, &data.Singer)
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
	err = row.Scan(&song.Id, &song.Name, &song.Singer)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func GetSongByName(name string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE name=?", name)
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Name, &song.Singer)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func PostSong(song *models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()

	randomUUID, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}
	_, err = db.Exec("INSERT INTO songs (id, name,singer ) VALUES (?, ?, ? )", randomUUID.String(), song.Name, song.Singer)

	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return song, err
}
func DeleteSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("DELETE FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func PutSong(id uuid.UUID, newName string, newSinger string) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("UPDATE songs SET name=?, singer=? WHERE id=?", newName, newSinger, id.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return nil, err
}

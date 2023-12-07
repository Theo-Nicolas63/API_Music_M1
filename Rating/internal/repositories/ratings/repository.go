package ratings

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllratings() ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.User_id, &data.Song_id, &data.Content)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return ratings, err
}

func GetratingById(id uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM RATINGS WHERE id=?", id.String())
	helpers.CloseDB(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.User_id, &rating.Song_id, &rating.Content)
	if err != nil {
		return nil, err
	}
	return &rating, err
}

func Postrating(rating *models.Rating) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("INSERT INTO ratings (id, id_user, id_song, content) VALUES (?,?, ?, ?)", rating.Id.String(), rating.User_id.String(), rating.Song_id.String(), rating.Content)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return rating, err
}

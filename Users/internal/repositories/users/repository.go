package repositories

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()

	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM Users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	Users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Name)
		if err != nil {
			return nil, err
		}
		Users = append(Users, data)
	}

	// don't forget to close rows
	_ = rows.Close()

	return Users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var User models.User
	err = row.Scan(&User.Id, &User.Name)
	if err != nil {
		return nil, err
	}
	return &User, err
}

func PostUser(User models.User) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("INSERT INTO Users (Id,Name,MusicLiked) VALUES (?,?,?)", User.Id.String(), User.Name, User.MusicLiked)
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}
	return &User, err
}

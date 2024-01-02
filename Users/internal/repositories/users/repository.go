package repositories

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()

	if err != nil {
		helpers.CloseDB(db)
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
		err = rows.Scan(&data.Id, &data.Name, &data.MusicLiked)
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
		helpers.CloseDB(db)
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var Users models.User
	err = row.Scan(&Users.Id, &Users.Name, &Users.MusicLiked)
	if err != nil {
		return nil, err
	}
	return &Users, err
}

func PostUser(Users *models.User) (*models.User, error) {
	db, err := helpers.OpenDB()

	randomUUID, err := uuid.NewV4()

	if err != nil {
		helpers.CloseDB(db)
		return nil, err
	}

	_, err = db.Exec("INSERT INTO Users (Id,Name,MusicLiked) VALUES (?,?,?)", randomUUID.String(), Users.Name, Users.MusicLiked)
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}
	return Users, err
}

func DeleteUser(Id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()

	if err != nil {
		helpers.CloseDB(db)
		return nil, err
	}

	_, err = db.Exec("DELETE FROM Users WHERE id=?", Id.String())
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}
	return nil, err
}

func PutUser(Users *models.User) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("UPDATE Users SET name=? WHERE id=?", Users.Name, Users.Id.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return Users, err
}

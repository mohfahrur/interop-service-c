package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohfahrur/interop-service-c/entity"
)

var db *sql.DB

type UserAgent interface {
	GetUserByID(userID int) (entity.User, error)
}

type UserDomain struct {
	db *sql.DB
}

func NewUserDomain(db *sql.DB) *UserDomain {

	return &UserDomain{
		db: db,
	}
}

func (d *UserDomain) GetUserInfo(userID string) (userData *entity.User, err error) {

	stmt, err := db.Prepare("SELECT ID, username, password, role FROM users WHERE ID = ?")
	if err != nil {
		log.Println("Failed to prepare the SQL statement:", err)
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)
	user := entity.User{}
	err = row.Scan(&user.ID, &user.Password, &user.Role)
	if err != nil {
		log.Println("Failed to retrieve user information:", err)
		return
	}
	userData = &user

	return
}

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
	GetUserAuth(pass string) (entity.User, error)
	GetUsers() (usersData []*entity.User, err error)
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
	query := "SELECT id, username, role FROM users WHERE ID = '" + userID + "'"
	log.Println(query)
	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Println("Failed to prepare the SQL statement:", err)
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	user := entity.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		log.Println("Failed to retrieve user information:", err)
		return
	}
	userData = &user

	return
}

func (d *UserDomain) GetUserAuth(pass string) (userData *entity.User, err error) {
	stmt, err := d.db.Prepare("SELECT id, username, role FROM users WHERE password = ?")
	if err != nil {
		log.Println("Failed to prepare the SQL statement:", err)
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(pass)
	user := entity.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		log.Println("Failed to retrieve user information:", err)
		return
	}
	userData = &user

	return
}

func (d *UserDomain) GetUsers() (usersData []*entity.User, err error) {

	stmt, err := d.db.Prepare("SELECT id, username, role FROM users")
	if err != nil {
		log.Println("Failed to prepare the SQL statement:", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Failed to prepare the SQL statement:", err)
		return nil, err
	}
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Role)
		if err != nil {
			log.Println("Failed to retrieve user information:", err)
			return
		}
		usersData = append(usersData, &user)
	}

	return
}

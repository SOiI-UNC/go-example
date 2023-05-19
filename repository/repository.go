package repository

import (
	"database/sql"
	"example-auth/model"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Connect() error {
	db, err := sql.Open("sqlite3", "./users.db")

	if err != nil {
		return err
	}

	database = db

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS USERS(id INTEGER PRIMARY KEY,USERNAME TEXT,PASSWORD TEXT)")
	if err != nil {
		fmt.Println(err)
	}

	statement.Exec()

	return nil
}

func Close() {
	database.Close()
}

func Save(u model.User) {

	stm, err := database.Prepare("INSERT INTO USERS (USERNAME,PASSWORD) VALUES (?,?);")

	if err != nil {
		log.Println(err)
	}

	defer stm.Close()

	safePass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}
	_, err = stm.Exec(u.Username, safePass)

	if err != nil {
		log.Println(err)
	}

}

func Get(u model.User) model.User {
	stm, err := database.Prepare("SELECT * FROM USERS WHERE USERNAME=?;")

	if err != nil {
		log.Println("Error in prepare query, get function")
		log.Println(err)
	}

	defer stm.Close()

	var result model.User
	var id int

	err = stm.QueryRow(u.Username).Scan(&id, &result.Username, &result.Password)

	if err != nil {
		log.Println("Error in query row, get function")
		log.Println(err)
	}

	return result

}

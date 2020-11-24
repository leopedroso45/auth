package dto

import (
	"auth/db"
	"auth/model"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
)

func Auth(username, password string) (model.User, error){
	conn := db.Connect()
	user := model.User{}

	encPassword := encrypting(password)
	sqlstate := fmt.Sprintf(`SELECT user.id, user.email, user.age, user.CreatedAt FROM user WHERE user.username= '%s' AND user.password= '%s'`, username, encPassword)
	row := conn.QueryRow(sqlstate)
	err := row.Scan(&user.ID, &user.Email, &user.Age, &user.CreatedAt)
	user.Username = username
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned")
		return user, err
	case nil:
		return user, nil
	default:
		panic(err)
	}

}

func encrypting(pass string) string {
	sha1Instance := sha1.New()
	sha1Instance.Write([]byte(pass))
	passwordCrypt := sha1Instance.Sum(nil)[:20]
	stringPasswordCrypt := hex.EncodeToString(passwordCrypt)
	return stringPasswordCrypt
}

func CreateUser(user model.User) (result bool) {
	//conn := db.Connect()

	newPass := encrypting(user.Password)
	user.Password = newPass

	//sqlstate := fmt.Sprintf("INSERT INTO user (username, email, password, age, CreatedAt) VALUES ('%s', '%s', '%s', '%s', '%s');", user.Username, user.Email, user.Password, user.Age, user.CreatedAt)

	//result, err := conn.Query("INSERT INTO user (username, email, password, age, CreatedAt) VALUES ('%s', '%s', '%s', '%s', '%s');", user.Username, user.Email, user.Password, user.Age, user.CreatedAt))
	//if err != nil {
	///	log.Fatal(err)
		return false
	//}
	//defer result.Close()
	//return true
}
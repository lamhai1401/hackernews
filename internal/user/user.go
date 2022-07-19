package user

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/lamhai1401/hackernews/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	statement, err := db.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := statement.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func (user *User) Authenticate() bool {
	statement, err := db.Db.Prepare("select Password from Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword) // return all row in struct line by line
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	statement, err := db.Db.Prepare("select ID from Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	row := statement.QueryRow(username)

	var id int

	err = row.Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return id, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

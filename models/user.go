package models

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

var (
	UserList map[string]*User
)

/* func init() {
	UserList = make(map[string]*User)
	u := User{1, "Sonom", "Sonom", "srnoor95@gmail.com", "01677734142", "1234", "29-04-1998"}
	UserList["user_11111"] = &u
} */

func init() {
	Objects = make(map[string]*Object)

}

type User struct {
	Id          int64
	Firstname   string
	Lastname    string
	Email       string
	Phone       string
	Password    string
	DateOfBirth string
}

func AddUser(u User) string {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "newPassword"
		dbname   = "user_db"
	)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()
	//validationofEmail
	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	email := emailRegexp.MatchString(u.Email)
	fmt.Println(email)

	//validationOfPhone
	var phoneRegx = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	phone := phoneRegx.MatchString(u.Phone)
	fmt.Println(phone)
	//validationOfdateOfBirth
	var dobRegx = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	dob := dobRegx.MatchString(u.DateOfBirth)
	fmt.Println(dob)

	sql := `INSERT INTO "user_table"("id","first_name", "last_name", "email", "phone", "password", "dob") VALUES ($1, $2, $3, $4, $5, $6,$7)`
	_, e := db.Exec(sql, u.Id, u.Firstname, u.Lastname, u.Email, u.Phone, u.Password, u.DateOfBirth)
	CheckError(e)
	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	fmt.Println(u.Email)

	return u.Email
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

/* func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
*/

package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	Student = iota
	Staff
	Faculty
)

type User struct {
	Username string
	Password string
	Property int
	Address  string
	Phone    string
}

func SearchUser(user User) (*User, error) {
	fmt.Println("Searhing: username:", user.Username, " password:", user.Password)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	var result = &User{}
	err = db.QueryRow("SELECT * FROM \"Users\" WHERE username = $1 AND password = $2", user.Username, user.Password).Scan(&result.Username, &result.Password, &result.Property, &result.Address, &result.Phone)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

func RegisterUser(user User) error {
	fmt.Println("Adding: username:", user.Username, " password:", user.Password, " address:", user.Address, " phone:", user.Phone)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO \"Users\" (username,password,property,address,phone) VALUES ($1,$2,$3,$4,$5)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, user.Property, user.Address, user.Phone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Insert success!")

	return nil
}

func GetUserProperty(user User) (int, error) {
	fmt.Println("Searching property: username:", user.Username, " password:", user.Password)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	defer db.Close()
	var result = &User{}
	err = db.QueryRow("SELECT * FROM \"Users\" WHERE username = $1 AND password = $2", user.Username, user.Password).Scan(&result.Username, &result.Password, &result.Property, &result.Address, &result.Phone)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return result.Property, nil
}

func GetAllUsers() ([][5]string, error) {
	fmt.Println("geting all users")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM \"Users\"")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result [][5]string

	for rows.Next() {
		var str1, str2, str3, str4, str5 string
		err = rows.Scan(&str1, &str2, &str3, &str4, &str5)
		if err != nil {
			return result, err

		}
		result = append(result, [5]string{str1, str2, str3, str4, str5})
	}

	return result, nil

}

func DeleteUser(user User) error {
	fmt.Println("Deletinging: username:", user.Username, " password:", user.Password)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM \"Users\" WHERE username=$1 AND password=$2")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Delete success")
	return nil
}

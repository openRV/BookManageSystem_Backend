package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Username string
	Password string
	Property *int
}

func SearchUser(user User) (*User, error) {
	fmt.Println("Searching: username:", user.Username, " password:", user.Password)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	var result = &User{}
	err = db.QueryRow("SELECT * FROM \"Users\" WHERE username = $1 AND password = $2", user.Username, user.Password).Scan(&result.Username, &result.Password, &result.Property)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

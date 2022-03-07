package Database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

func Borrow(book Book, user User) error {
	fmt.Println("Borrowing book:", book.Title, " to:", user.Username)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		return err
	}
	defer db.Close()

	BookInfo := &Book{}
	row := db.QueryRow("SELECT * FROM Book WHERE bookid=$1 ", book.ID)
	err = row.Scan(&BookInfo.ID, &BookInfo.Title, &BookInfo.CopyNum, &BookInfo.Author, &BookInfo.PublicationDate, &BookInfo.PublisherName, &BookInfo.PublisherAddress)
	if err != nil {
		fmt.Println(err)
		return err
	}
	copies, err := GetAllCopy(*BookInfo)
	if err != nil {
		return err
	}
	borrowed, err := GetBorrowed(*BookInfo)
	if err != nil {
		return err
	}

	if len(borrowed) == BookInfo.CopyNum {
		return errors.New("no book copy available, book copy:" + strconv.Itoa(BookInfo.CopyNum) + " borrowed num:" + strconv.Itoa(len(borrowed)))
	}

	valid := false
	borrowingId := 0

	for _, value := range copies {
		valid = true
		for _, a := range borrowed {
			if value[1] == strconv.Itoa(a.CopyID) {
				valid = false
			}
		}
		if valid == true {
			borrowingId, _ = strconv.Atoi(value[1])
			break
		}
	}

	stmt, err := db.Prepare("INSERT INTO Borrow (bookid,copyid,username,userpassword,borrowdate,returndate) VALUES ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(BookInfo.ID, borrowingId, user.Username, user.Password, time.Now(), time.Now().Add(time.Hour*24*30))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// TODO: insert information into table BorrowHistory
	fmt.Println("Borrow success")

	return nil
}

func Return(book Book, user User) error {
	fmt.Println("Returning book:", book.Title, " for:", user.Username)

	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Borrow WHERE bookid=$1 AND username=$2 AND password=$3")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// TODO: insert information into table BorrowHistory
	fmt.Println("Return success")
	return nil
}

func GetBorrowed(book Book) ([]Copy, error) {
	fmt.Println("searching borrowed copy of book:", book.Title)

	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Borrow WHERE bookid = $1", book.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []Copy

	for rows.Next() {
		var str1, str2, str3, str4, str5, str6 string
		err = rows.Scan(&str1, &str2, &str3, &str4, &str5, &str6)
		if err != nil {
			return result, err

		}
		copyid, _ := strconv.Atoi(str2)
		bookid, _ := strconv.Atoi(str1)
		result = append(result, Copy{CopyID: copyid, BookID: bookid})
	}

	return nil, nil
}

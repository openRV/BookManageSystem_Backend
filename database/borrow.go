package Database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

type BorrowInfo struct {
	BookId     int
	Title      string
	Author     string
	Publisher  string
	BorrowDate string
	ReturnDate string
}

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

	_, err = stmt.Exec(BookInfo.ID, borrowingId, user.Username, user.Password, time.Now().String()[:10], time.Now().Add(time.Hour * 24 * 30).String()[:10])
	if err != nil {
		fmt.Println(err)
		return err
	}

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

	//borrowing, err := GetBorrowingBy(user)
	//if err != nil {
	//	return err
	//}

	//if len(borrowing) != 0 {
	//	return errors.New("have unreturned books, please return books first")
	//}

	var bid, cid, uname, upass, bdate, rdate string
	row := db.QueryRow("SELECT * FROM Borrow WHERE bookid=$1 AND username=$2 AND userpassword=$3", book.ID, user.Username, user.Password)
	err = row.Scan(&bid, &cid, &uname, &upass, &bdate, &rdate)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM Borrow WHERE bookid=$1 AND username=$2 AND userpassword=$3")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.ID, user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	stmt, err = db.Prepare("INSERT INTO Borrowhistory (bookid,copyid,username,userpassword,borrowdate,returndate) VALUES ($1,$2,$3,$4,$5,$6)")
	_, err = stmt.Exec(bid, cid, uname, upass, bdate, time.Now().String()[:10])
	if err != nil {
		return err
	}

	fmt.Println("Return success")
	return nil
}

func GetBorrowed(book Book) ([]Copy, error) {
	fmt.Println("searching borrowed copy of book:", book.ID)

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
			return nil, err
		}
		copyid, _ := strconv.Atoi(str2)
		bookid, _ := strconv.Atoi(str1)
		result = append(result, Copy{CopyID: copyid, BookID: bookid})
	}

	return result, nil
}

func GetBorrowingBy(user User) ([]BorrowInfo, error) {
	fmt.Println("Searching books borrowing by :", user.Username)

	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Borrow WHERE username= $1 AND userpassword=$2", user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []BorrowInfo

	for rows.Next() {
		var str1, str2, str3, str4, str5, str6 string
		err = rows.Scan(&str1, &str2, &str3, &str4, &str5, &str6)
		if err != nil {
			return nil, err
		}
		bookid, _ := strconv.Atoi(str1)
		err := db.QueryRow("SELECT * FROM Book WHERE bookid=$1", bookid).Scan(&str1, &str2, &str1, &str3, &str1, &str4, &str1)
		if err != nil {
			return nil, err
		}
		result = append(result, BorrowInfo{
			BookId:     bookid,
			Title:      str2,
			Author:     str3,
			Publisher:  str4,
			BorrowDate: str5,
			ReturnDate: str6,
		})
	}

	return result, nil
}

func GetBorrowedBy(user User) ([]BorrowInfo, error) {
	fmt.Println("Searching books borrowed by :", user.Username)

	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Borrowhistory WHERE username= $1 AND userpassword=$2", user.Username, user.Password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []BorrowInfo

	for rows.Next() {
		var str1, str2, str3, str4, str5, str6 string
		err = rows.Scan(&str1, &str2, &str3, &str4, &str5, &str6)
		if err != nil {
			return nil, err
		}
		bookid, _ := strconv.Atoi(str1)
		err := db.QueryRow("SELECT * FROM Book WHERE bookid=$1", bookid).Scan(&str1, &str2, &str1, &str3, &str1, &str4, &str1)
		if err != nil {
			return nil, err
		}
		result = append(result, BorrowInfo{
			BookId:     bookid,
			Title:      str2,
			Author:     str3,
			Publisher:  str4,
			BorrowDate: str5,
			ReturnDate: str6,
		})
	}

	return result, nil
}

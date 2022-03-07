package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Book struct {
	ID               int
	Title            string
	CopyNum          int
	Author           string
	PublicationDate  string
	PublisherName    string
	PublisherAddress string
}

type Copy struct {
	BookID          int
	CopyID          int
	LibraryName     string
	LibraryLocation string
}

func GetAllBook() ([][7]string, error) {
	fmt.Println("geting all books")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Book")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result [][7]string

	for rows.Next() {
		var str1, str2, str3, str4, str5, str6, str7 string
		err = rows.Scan(&str1, &str2, &str3, &str4, &str5, &str6, &str7)
		if err != nil {
			return result, err

		}
		result = append(result, [7]string{str1, str2, str3, str4, str5, str6, str7})
	}

	return result, nil
}

func GetAllCopy(book Book) ([][4]string, error) {

	fmt.Println("geting all copies with book id:", book.ID)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM BookCopy WHERE bookid = $1", book.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result [][4]string

	for rows.Next() {
		var str1, str2, str3, str4 string
		err = rows.Scan(&str1, &str2, &str3, &str4)
		if err != nil {
			return result, err

		}
		result = append(result, [4]string{str1, str2, str3, str4})
	}

	return result, nil
}

func DelBook(book Book) error {
	fmt.Println("deleting all Books and Copies with id:", book.ID)
	err := DelCopy(book)
	if err != nil {
		return err
	}

	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Book WHERE bookid=$1")
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
	fmt.Println("Delete Book success")
	return nil
}

func DelCopy(book Book) error {
	fmt.Println("deleting all Copies with id:", book.ID)

	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM BookCopy WHERE bookid=$1")
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
	fmt.Println("Delete Copy success")
	return nil
}

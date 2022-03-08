package Database

import (
	"database/sql"
	"fmt"
	"strings"
)

type JournalData struct {
	JournalId    string
	JournalTitle string `json:"JournalTitle"`
	Author       string `json:"JournalEditor"`
	Scope        string `json:"scope"`
}

func SearchJournal(JournalTitle string, Scope string) ([][5]string, error) {
	fmt.Println("Getting Journals...")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Journal")
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
		if strings.Contains(str2, JournalTitle) && strings.Contains(str4, Scope) {
			result = append(result, [5]string{str1, str2, str3, str4, str5})
		}
	}

	fmt.Println("Successfully get!")
	return result, nil

}

func InsertJournal(data JournalData) error {
	fmt.Println("Insert a Journal...")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Journal (journalid,journalititle,author,scope) VALUES ($1,$2,$3,$4)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		data.JournalId,
		data.JournalTitle,
		data.Author,
		data.Scope)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Insert success!")

	return nil

}

func DeleteJournal(JournalId string) error {
	fmt.Println("Deletinging a Journal")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Journal WHERE journalid = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(JournalId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Delete success")
	return nil
}
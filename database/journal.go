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

func SearchJournal(JournalTitle string, Scope string) ([][4]string, error) {
	fmt.Println("Getting Journals...")
	fmt.Println("JournalTitle: " + JournalTitle)
	fmt.Println("JournalCope: " + Scope)
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

	var result [][4]string

	for rows.Next() {
		var str1, str2, str3, str4 string
		err = rows.Scan(&str1, &str2, &str3, &str4)
		if err != nil {
			return result, err
		}
		if strings.Contains(str2, JournalTitle) && strings.Contains(str4, Scope) {
			result = append(result, [4]string{str1, str2, str3, str4})
		}
	}

	fmt.Println("Successfully get journals!")
	return result, nil

}

func InsertJournal(data JournalData) error {
	fmt.Println("Insert a Journal...")
	fmt.Println("JournalId: ", data.JournalId)
	fmt.Println("journalTitle: ", data.JournalTitle)
	fmt.Println(("Author: " + data.Author))
	fmt.Println(("Scope: " + data.Scope))
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Journal (journalid,journaltitle,author,scope) VALUES ($1,$2,$3,$4)")
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

	fmt.Println("Insert Journal success!")

	return nil

}

func DeleteJournal(JournalId string) error {
	fmt.Println("Deletinging a Journal...")
	fmt.Println("JournalId: " + JournalId)
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
	fmt.Println("Delete journal success")
	return nil
}

package Database

import (
	"database/sql"
	"fmt"
	"strings"
)

type VolumnData struct {
	JournalId       string
	VolumnId        string `json:"VolumeNum"`
	Volumneditior   string `json:"VolumnEditior"`
	PublicationDate string `json:"publicationDate"`
}

func SearchVolumn(volumnId string, volumnEditor string) ([][4]string, error) {
	//fmt.Println("geting all conference")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Volumn")
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
		if strings.Contains(str1, volumnId) && strings.Contains(str3, volumnEditor) {
			result = append(result, [4]string{str1, str2, str3, str4})
		}
	}

	return result, nil

}

func InsertVolumn(data VolumnData) error {
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Volumn (journalid,wolumnid,wolumneditor,publicationdate) VALUES ($1,$2,$3,$4)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		data.JournalId,
		data.VolumnId,
		data.Volumneditior,
		data.PublicationDate)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Insert success!")

	return nil
}

func DeleteVolumn(journalid string, volumnid string) error {
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Journal WHERE journalid = $1 AND volumnid = $2")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(journalid, volumnid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Delete success")
	return nil

}

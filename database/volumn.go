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

func SearchVolumn1(journalId string, volumnId string) (string, error) {
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT publicationdate FROM Volumn WHERE journalid = $1 AND volumnid =$2", journalId, volumnId).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return result, nil
}
func SearchVolumn(volumnId string, volumnEditor string) ([][4]string, error) {
	fmt.Println("getting Volumn...")
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
	fmt.Println("Successfully get Volumn")
	return result, nil

}

func InsertVolumn(data VolumnData) error {
	fmt.Println("Inserting a volumn...")
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
	fmt.Println("Deleting a volumn...")
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

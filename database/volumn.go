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
	fmt.Println("Get journal paper's publication date...")
	fmt.Println("JournalId: " + journalId)
	fmt.Println(("VolumnId: " + volumnId))
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

	fmt.Println("Get journal paper's publication date successfully!")

	return result, nil
}
func SearchVolumn(volumnId string, volumnEditor string) ([][4]string, error) {
	fmt.Println("getting Volumn...")
	fmt.Println("VolumnId: " + volumnId)
	fmt.Println("volumnEditor+ " + volumnEditor)
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
	fmt.Println("JournalId: " + data.JournalId + "\n" + "VolumnId: " + data.VolumnId + "\n" + "VolumnEditor: " + data.Volumneditior + "\n" + "PublicationDate: " + data.PublicationDate + "\n")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Volumn (journalid,volumnid,volumneditor,publicationdate) VALUES ($1,$2,$3,$4)")
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

	fmt.Println("Insert volumn success!")

	return nil
}

func DeleteVolumn(journalid string, volumnid string) error {
	fmt.Println("Deleting a volumn...")
	fmt.Println("JournalId: " + journalid + "\n" + "VolumnId: " + volumnid + "\n")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Volumn WHERE journalid = $1 AND volumnid = $2")
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
	fmt.Println("Delete volumn success")
	return nil

}

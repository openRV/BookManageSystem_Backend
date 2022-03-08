package Database

import (
	"database/sql"
	"fmt"
	"strings"
)

type ConferenceData struct {
	ConferenceId     string
	ConferenceTitle  string `json:"ConferenceTitle"`
	ProceedingEditor string `json:"ProceedingEditor"`
	PublishDate      string `json:"publicDate"`
	PublishAddress   string `json:"publicAddress"`
}

func SearchConference1(conferenceId string) (string, error) {
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT publiahsdate FROM Conference WHERE conferenceid = $1", conferenceId).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return result, nil
}

func SearchConference(conferenceTitle string, proceedingEditor string) ([][5]string, error) {
	fmt.Println("geting all conference")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Conference")
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
		if strings.Contains(str1, conferenceTitle) && strings.Contains(str2, proceedingEditor) {
			result = append(result, [5]string{str1, str2, str3, str4, str5})
		}
	}

	return result, nil
}

func InsertConference(data ConferenceData) error {
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Conference (conferenceid,conferencetitle,proceedingeditor,publishdate,publishaddress) VALUES ($1,$2,$3,$4,$5)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		data.ConferenceId,
		data.ConferenceTitle,
		data.ProceedingEditor,
		data.PublishDate,
		data.PublishAddress)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Insert success!")

	return nil
}
func DeleteConference(conferenceId string) error {
	//fmt.Println("Deletinging: username:", user.Username, " password:", user.Password)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Conference WHERE conferenceid = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(conferenceId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Delete success")
	return nil
}

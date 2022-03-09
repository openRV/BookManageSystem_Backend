package Database

import (
	"database/sql"
	"fmt"
)

type DownloadInfo struct {
	Username     string
	Userpassword string
	Paperid      string
	Papertitle   string
	Paperauthor  string
	Downloaddate string
}

func SearchDownload(Username string, Userpassword string) ([][4]string, error) {

	fmt.Println("Getting download history...")
	fmt.Println("Username: " + Username + "\n" + "Userpassword: " + Userpassword + "\n")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT paperid,papertitle,paperauthor,downloaddate FROM Downloadhistory WHERE username = $1 AND userpassword = $2", Username, Userpassword)
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

	fmt.Println("Successfully get download history!")
	return result, nil
}

func InsertDownload(info DownloadInfo) error {
	fmt.Println("Insert a download history...")
	fmt.Println("Username: " + info.Username + "\n" + "UserPassword: " + info.Userpassword)
	fmt.Println("PaperId: " + info.Paperid)
	fmt.Println("PaperTtiel: " + info.Papertitle)
	fmt.Println("PaperAuthor:" + info.Paperauthor)
	fmt.Println("Downloaddate: " + info.Downloaddate)
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Downloadhistory (username,userpassword,paperid,papertitle,paperauthor,downloaddate) VALUES ($1,$2,$3,$4,$5,$6)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		info.Username,
		info.Userpassword,
		info.Paperid,
		info.Papertitle,
		info.Paperauthor,
		info.Downloaddate)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Insert download history successfully!")

	return nil
}

func DeleteHistory(Username string, Userpassword string, Paperid string, Downloaddate string) error {
	//TODO: To Complete
	return nil
}

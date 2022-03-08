package Database

import (
	"database/sql"
	"fmt"
	"strings"
)

type PaperData struct {
	PaperId      string `json:"paperId"`
	Author       string `json:"author"`
	PaperTitle   string `json:"title"`
	JournalId    string
	VolumnId     string `json:"VolumeNum"`
	ConferenceId string
	Link         string `json:"Link"`
	IsOpen       string
}

func GetOpenPaper() ([][5]string, error) {

	var result [][5]string
	temp, err := SearchPaper("", "")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, value := range temp {
		if value[7] == "True" {
			if value[5] != "" {
				conferenceResult, err := SearchConference1(value[5])
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				result = append(result, [5]string{value[0], value[2], value[1], conferenceResult, value[6]})
			}
			if value[3] != "" && value[4] != "" {
				volumnresult, err := SearchVolumn1(value[3], value[4])
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				result = append(result, [5]string{value[0], value[2], value[1], volumnresult, value[6]})
			}
		}
	}
	return result, nil

}

func SearchPaper1(Conferenceid string) (bool, error) {
	fmt.Println("Verify whether are more papers in same Conference...")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT papertitle FROM Conference WHERE conferenceid = $1", Conferenceid).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil

}

func SearchPaper2(Journalid string, Volumnid string) (bool, error) {
	fmt.Println("Examing whether there are more papers in same JournalColumn...")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT papertitle FROM Conference WHERE journalid = $1 AND volumnid = $2", Journalid, Volumnid).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil

}

func SearchPaper3(Journalid string) (bool, error) {
	fmt.Println("Examing whether there are more papers in same Journal...")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer db.Close()

	var result string
	err = db.QueryRow("SELECT papertitle FROM Conference WHERE journalid = $1", Journalid).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil

}

func SearchPaper(PaperTitle string, PaperAuthor string) ([][8]string, error) {
	fmt.Println("getting papers...")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	//paper id, author, paper title, journal id, column id, conference id, link, is open
	rows, err := db.Query("SELECT * FROM Paper")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result [][8]string

	for rows.Next() {
		var str1, str2, str3, str4, str5, str6, str7, str8 string
		err = rows.Scan(&str1, &str2, &str3, &str4, &str5, &str6, &str7, &str8)
		if err != nil {
			return result, err
		}
		if strings.Contains(str3, PaperTitle) && strings.Contains(str2, PaperAuthor) {
			result = append(result, [8]string{str1, str2, str3, str4, str5, str6, str7, str8})
		}
	}

	fmt.Println("Successfully get!")
	return result, nil
}

func SearchConferencePaper(PaperTitle string, PaperAuthor string, ConferenceTitle string, ProceedingEditor string) ([][7]string, error) {
	fmt.Println("Getting papers from certain Conference...")
	var result [][7]string
	conference, err := SearchConference(ConferenceTitle, ProceedingEditor)
	if err != nil {
		return result, err
	}
	paper, err := SearchPaper(PaperTitle, PaperAuthor)
	if err != nil {
		return result, err
	}
	for _, paperArr := range paper {
		for _, conferenceArr := range conference {
			if paperArr[5] == conferenceArr[0] {
				//paperid,papertitle,author,publishdate,publishaddress,conferencetitle,proceedingeditor
				result = append(result, [7]string{paperArr[0], paperArr[2], paperArr[1], conferenceArr[3], conferenceArr[4], conferenceArr[1], conferenceArr[2]})
			}
		}
	}

	fmt.Println("Successfull get!")
	return result, nil
}

func SearchJournalPaper(PaperTitle string, PaperAuthor string, JournalTitle string, JournalScope string, VolumnNum string, VolumnEditor string) ([][9]string, error) {
	fmt.Println("Getting papers from sertain Journal Volumn...")
	var result [][9]string
	paper, err := SearchPaper(PaperTitle, PaperAuthor)
	if err != nil {
		return result, err
	}
	journal, err := SearchJournal(JournalTitle, JournalScope)
	if err != nil {
		return result, err
	}
	volumn, err := SearchVolumn(VolumnNum, VolumnEditor)
	for _, paperArr := range paper {
		for _, journalArr := range journal {
			for _, volumnArr := range volumn {
				if paperArr[3] == journalArr[0] && paperArr[3] == volumnArr[0] && paperArr[4] == volumnArr[1] {
					//paperid,papertitlr,author,publicationdate,scope,journaltitle,journaleditor,volumnid,volumneditor
					result = append(result, [9]string{
						paperArr[0],
						paperArr[2],
						paperArr[1],
						volumnArr[3],
						journalArr[3],
						journalArr[1],
						journalArr[2],
						volumnArr[1],
						volumnArr[2]})
				}
			}
		}
	}
	fmt.Println("Successfully get!")
	return result, nil
}

//TODO: To complete
func InsertPaper(data PaperData) error {
	/*
		db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer db.Close()

		//针对相关Journal Volumn/Conference做处理

		//检测某会议是否空
		if data.ConferenceId != "" {
			var test1 bool
			test1, err = SearchPaper1(data.ConferenceId)
			if !test1 {
				InsertConference(result.ConferenceId)
			}
		}
		//检测期刊的某册是否空
		if result.JournalId != "" && result.VolumnId != "" {
			var test2 bool
			test2, err = SearchPaper2(result.JournalId, result.VolumnId)
			if !test2 {
				DeleteVolumn(result.JournalId, result.VolumnId)
				//某期刊是否空
				var test3 bool
				test3, err = SearchPaper3(result.JournalId)
				if !test3 {
					DeleteJournal(result.JournalId)
				}
			}
		}*/
	var err error
	return err

}

func DeletePaper(PaperId string) error {
	fmt.Println("Deleting a paper...")
	db, err := sql.Open(DBTYPE, DBTYPE+"://"+USERNAME+":"+PASSWORD+"@"+HOST+":"+PORT+"/"+DBNAME+"?sslmode="+SSLMODE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Paper WHERE paperid = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(PaperId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//针对相关Journal Volumn/Conference做处理

	var result = &PaperData{}
	err = db.QueryRow("SELECT * FROM Paper WHERE paperid = $1", PaperId).Scan(&result.PaperId, &result.PaperTitle, &result.JournalId, &result.VolumnId, &result.ConferenceId, &result.Link)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//检测某会议是否空
	if result.ConferenceId != "" {
		var test1 bool
		test1, err = SearchPaper1(result.ConferenceId)
		if !test1 {
			DeleteConference(result.ConferenceId)
		}
	}
	//检测期刊的某册是否空
	if result.JournalId != "" && result.VolumnId != "" {
		var test2 bool
		test2, err = SearchPaper2(result.JournalId, result.VolumnId)
		if !test2 {
			DeleteVolumn(result.JournalId, result.VolumnId)
			//某期刊是否空
			var test3 bool
			test3, err = SearchPaper3(result.JournalId)
			if !test3 {
				DeleteJournal(result.JournalId)
			}
		}
	}

	fmt.Println("Delete success")
	return nil
}

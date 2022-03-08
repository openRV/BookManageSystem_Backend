package Interface

import (
	"github.com/gin-gonic/gin"
)

type ConferencePaperData struct {
	PaperId          string `json:"paperId"`
	PaperTitle       string `json:"title"`
	PaperAuthor      string `json:"author"`
	PublishDate      string `json:"publicDate"`
	PublishAddress   string `json:"publicAddress"`
	ConferenceTitle  string `json:"ConferenceTitle"`
	ProceedingEditor string `json:"proceedingEditor"`
}

type ConferencePaperRet struct {
	Success bool                  `json:"success"`
	Data    []ConferencePaperData `json:"data"`
	Total   int                   `json:"total"`
}

func SearchConferencebook(c *gin.Context) {

}

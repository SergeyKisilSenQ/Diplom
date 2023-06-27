package repo_support

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type StorageSupport []*SupportData

func NewStorageSupport() *StorageSupport {
	return &StorageSupport{}
}

func (*StorageSupport) GetSupportData() []int {
	res, err := http.Get(os.Getenv("SUPPORT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var s []*SupportData
	if res.StatusCode == http.StatusOK {
		if err := json.Unmarshal(body, &s); err != nil {
			log.Fatal(err)
		}
	}
	return SupportDataStatus(s)
}

func SupportDataStatus(s []*SupportData) []int {
	var supportDataStatus []int
	var sumActiveTickets int
	for _, val := range s {
		sumActiveTickets += val.ActiveTickets
	}
	if sumActiveTickets < 9 {
		supportDataStatus = append(supportDataStatus, 1)
	} else if sumActiveTickets > 9 && sumActiveTickets < 17 {
		supportDataStatus = append(supportDataStatus, 2)
	} else {
		supportDataStatus = append(supportDataStatus, 3)
	}
	supportDataStatus = append(supportDataStatus, sumActiveTickets*3)
	return supportDataStatus
}

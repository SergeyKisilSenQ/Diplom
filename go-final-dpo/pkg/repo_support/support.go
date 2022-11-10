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

type StorageSupport map[int]*SupportData

func NewStorageSupport() StorageSupport {
	return make(map[int]*SupportData)
}

func (SSD StorageSupport) Put(Country *SupportData) {
	SSD[len(SSD)] = Country
}
func (SSD StorageSupport) GetSupport() (SS []*SupportData) {
	res, err := http.Get(os.Getenv("SUPPORT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	var s []*SupportData
	if res.StatusCode == http.StatusOK {
		if e := json.Unmarshal(body, &s); e != nil {
			return SS
		}
		for i, _ := range s {
			SSD.Put(s[i])
		}
	}
	if res.StatusCode == http.StatusInternalServerError {
		return SS
	}
	//for i, _ := range SSD {
	//	fmt.Println(SSD[i])
	//}
	return SS
}

func SupportDataStatus(SSD StorageSupport) []int {
	var supportDataStatus []int
	var sumActiveTickets int
	for i, _ := range SSD {
		sumActiveTickets += SSD[i].ActiveTickets
	}
	if sumActiveTickets < 9 {
		supportDataStatus = append(supportDataStatus, 1)
	} else if sumActiveTickets > 9 && sumActiveTickets < 17 {
		supportDataStatus = append(supportDataStatus, 2)
	} else {
		supportDataStatus = append(supportDataStatus, 3)
	}
	supportDataStatus = append(supportDataStatus, sumActiveTickets*3)
	//fmt.Println(supportDataStatus)
	return supportDataStatus
}

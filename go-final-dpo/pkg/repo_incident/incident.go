package repo_incident

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

type StorageIncident map[int]*IncidentData

func NewStorageIncident() StorageIncident {
	return make(map[int]*IncidentData)
}

func (ID StorageIncident) Put(Country *IncidentData) {
	ID[len(ID)] = Country
}

func (ID StorageIncident) GetIncident() (IS []*IncidentData) {
	res, err := http.Get(os.Getenv("INCIDENT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	var s []*IncidentData
	if res.StatusCode == http.StatusOK {
		if e := json.Unmarshal(body, &s); e != nil {
			return IS
		}
		for i, _ := range s {
			ID.Put(s[i])
		}
	}
	if res.StatusCode == http.StatusInternalServerError {
		return IS
	}
	//for i, _ := range ID {
	//	fmt.Println(ID[i])
	//}
	return IS
}

func AddIncidentData(ID StorageIncident) []*IncidentData {
	addIncidentData := make([]*IncidentData, len(ID))
	for i := 1; i < len(ID); i++ {
		j := i
		for j > 0 {
			if ID[j-1].Status > ID[j].Status {
				ID[j-1], ID[j] = ID[j], ID[j-1]
			}
			j = j - 1
		}
		for k, _ := range ID {
			addIncidentData[k] = ID[k]
		}
	}
	//fmt.Println(addIncidentData[0].Status)
	return addIncidentData
}

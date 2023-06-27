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

type StorageIncident []*IncidentData

func NewStorageIncident() *StorageIncident {
	return &StorageIncident{}
}

func (*StorageIncident) GetIncidentData() []*IncidentData {
	res, err := http.Get(os.Getenv("INCIDENT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var s []*IncidentData
	if res.StatusCode == http.StatusOK {
		if err := json.Unmarshal(body, &s); err != nil {
			log.Fatal(err)
		}
	}
	return s
}

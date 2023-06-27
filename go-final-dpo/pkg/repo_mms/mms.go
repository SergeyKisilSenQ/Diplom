package repo_mms

import (
	"Diplom/go-final-dpo/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type StorageMMS [][]*MMSData

func NewStorageMMS() *StorageMMS {
	return &StorageMMS{}
}

func (*StorageMMS) GetMMSData() [][]*MMSData {
	sortedMMSData1 := make([]*MMSData, 0)
	sortedMMSData2 := make([]*MMSData, 0)
	res, err := http.Get(os.Getenv("MMS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var m []*MMSData
		if err := json.Unmarshal(body, &m); err != nil {
			log.Fatal(err)
		}
		for i, _ := range m {
			NewCountry, err := utils.CountryCheck(m[i].Country)
			if err != nil {
				log.Printf("%s: %s, in mms-data\n", err, m[i].Country)
				continue
			}
			NewProvider, err := utils.CheckProvider(m[i].Provider, strings.Split(os.Getenv("SMS_MMS_PROVIDER"), ", "))
			if err != nil {
				log.Printf("%s: %s, in mms-data\n", err, m[i].Provider)
				continue
			}
			NewMMSData1 := &MMSData{
				Country:      NewCountry,
				Provider:     NewProvider,
				Bandwidth:    m[i].Bandwidth,
				ResponseTime: m[i].Bandwidth,
			}
			NewMMSData2 := &MMSData{
				Country:      NewCountry,
				Provider:     NewProvider,
				Bandwidth:    m[i].Bandwidth,
				ResponseTime: m[i].Bandwidth,
			}
			sortedMMSData1 = append(sortedMMSData1, NewMMSData1)
			sortedMMSData2 = append(sortedMMSData2, NewMMSData2)
		}
	}
	sort.SliceStable(sortedMMSData1, func(i, j int) bool { return sortedMMSData1[i].Provider < sortedMMSData1[j].Provider })
	sort.SliceStable(sortedMMSData2, func(i, j int) bool { return sortedMMSData2[i].Country < sortedMMSData2[j].Country })
	result := make([][]*MMSData, 0)
	result = append(result, sortedMMSData1, sortedMMSData2)
	return result
}

package repo_mms

import (
	"Diplom/go-final-dpo/pkg/repo_country"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type StorageMMS map[int]*MMSData

func NewStorageMMS() StorageMMS {
	return make(map[int]*MMSData)
}

func (MD StorageMMS) Put(Country *MMSData) {
	MD[len(MD)] = Country
}

func (MD StorageMMS) GetMMS() {
	res, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	if res.StatusCode == http.StatusOK {
		var m []*MMSData
		if err := json.Unmarshal(body, &m); err != nil {
			log.Fatal(err)
		}
		for i, _ := range m {
			for j, _ := range NSC {
				if m[i].Country == NSC[j].NameCountry {
					if m[i].Provider == "Topolo" || m[i].Provider == "Rond" || m[i].Provider == "Kildy" {
						MD.Put(m[i])
					}
				}

			}
		}

		//for i, _ := range MD {
		//	fmt.Printf("%s, %s, %s, %s, \n", MD[i].Country, MD[i].Provider, MD[i].ResponseTime, MD[i].Bandwidth)
		//}
	} else {
		log.Fatal(err)
	}
}

func SortedMMSData(MD StorageMMS) [][]*MMSData {
	sortedMMSData1 := make([]*MMSData, len(MD))
	sortedMMSData2 := make([]*MMSData, len(MD))
	result := make([][]*MMSData, 0)
	for i, _ := range MD {
		if MD[i].Country == "RU" {
			MD[i].Country = "Russian Federation"
		}
		if MD[i].Country == "US" {
			MD[i].Country = "United States"
		}
		if MD[i].Country == "GB" {
			MD[i].Country = "Great Britain"
		}
		if MD[i].Country == "FR" {
			MD[i].Country = "France"
		}
		if MD[i].Country == "BL" {
			MD[i].Country = "Saint Barthelemy"
		}
		if MD[i].Country == "AT" {
			MD[i].Country = "Austria"
		}
		if MD[i].Country == "BG" {
			MD[i].Country = "Bulgaria"
		}
		if MD[i].Country == "DK" {
			MD[i].Country = "Denmark"
		}
		if MD[i].Country == "CA" {
			MD[i].Country = "Canada"
		}
		if MD[i].Country == "ES" {
			MD[i].Country = "Spain"
		}
		if MD[i].Country == "CH" {
			MD[i].Country = "Switzerland"
		}
		if MD[i].Country == "TR" {
			MD[i].Country = "Turkey"
		}
		if MD[i].Country == "PE" {
			MD[i].Country = "Peru"
		}
		if MD[i].Country == "NZ" {
			MD[i].Country = "New Zealand"
		}
		if MD[i].Country == "MC" {
			MD[i].Country = "Monaco"
		}
	}

	for i := 1; i < len(MD); i++ {
		j := i
		for j > 0 {
			if MD[j-1].Provider > MD[j].Provider {
				MD[j-1], MD[j] = MD[j], MD[j-1]
			}
			j = j - 1
		}
		for k, _ := range MD {
			sortedMMSData1[k] = MD[k]
		}
	}
	for i := 1; i < len(MD); i++ {
		j := i
		for j > 0 {
			if MD[j-1].Country > MD[j].Country {
				MD[j-1], MD[j] = MD[j], MD[j-1]
			}
			j = j - 1
		}
		for k, _ := range MD {
			sortedMMSData2[k] = MD[k]
		}
	}

	result = append(result, sortedMMSData1, sortedMMSData2)

	//fmt.Println(result)
	//fmt.Println(sortedMMSData1[0].Provider, sortedMMSData1[1].Provider, sortedMMSData1[4].Provider)
	//fmt.Println(sortedMMSData2[0].Country, sortedMMSData2[3].Country, sortedMMSData2[5].Country)
	return result
}

package repo_mms

import (
	"Diplom/go-final-dpo/pkg/repo_country"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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
	res, err := http.Get(os.Getenv("MMS_URL"))
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
				if m[i].Country == NSC[j].CodeCountry {
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
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	for i, _ := range MD {
		for j, _ := range NSC {
			if MD[i].Country == NSC[j].CodeCountry {
				MD[i].Country = NSC[j].NameCountry
			}
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

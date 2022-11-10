package repo_sms

import (
	"Diplom/go-final-dpo/pkg/repo_country"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type SMSData struct {
	Сountry      string `json:"country"`
	Bandwidth    int    `json:"bandwidth"`
	ResponseTime int    `json:"response_time"`
	Provider     string `json:"provider"`
}

type StorageSMS map[int]*SMSData

func NewStorageSMS() StorageSMS {
	return make(map[int]*SMSData)
}

func (SD StorageSMS) Put(Country *SMSData) {
	SD[len(SD)] = Country
}

func (SD StorageSMS) ReadFileSMS() {
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	r, err := ioutil.ReadFile(os.Getenv("SMS_FILE"))
	if err != nil {
		panic(err)
	}
	s := strings.Fields(string(r))
	for i := 0; i < len(s); i++ {
		var result string
		var e []string
		result = strings.ReplaceAll(s[i], ";", " ")
		e = strings.Fields(result)

		if len(e) == 4 {
			for k, _ := range NSC {
				if e[0] == NSC[k].CodeCountry {
					if e[3] == "Topolo" || e[3] == "Rond" || e[3] == "Kildy" {
						NewСountry := e[0]
						NewBandwidth, err := strconv.Atoi(e[1])
						if err != nil {
							panic(err)
						}
						NewResponseTime, err := strconv.Atoi(e[2])
						if err != nil {
							panic(err)
						}
						NewProvider := e[3]
						NewSMSData := SMSData{
							Сountry:      NewСountry,
							Bandwidth:    NewBandwidth,
							ResponseTime: NewResponseTime,
							Provider:     NewProvider,
						}
						SD.Put(&NewSMSData)
					}
				}
			}
		}
	}
}

func SortedSMSData(SD StorageSMS) [][]*SMSData {
	sortedSMSData1 := make([]*SMSData, len(SD))
	sortedSMSData2 := make([]*SMSData, len(SD))
	result := make([][]*SMSData, 0)
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	for i, _ := range SD {
		for j, _ := range NSC {
			if SD[i].Сountry == NSC[j].CodeCountry {
				SD[i].Сountry = NSC[j].NameCountry
			}
		}
	}
	for i := 1; i < len(SD); i++ {
		j := i
		for j > 0 {
			if SD[j-1].Provider > SD[j].Provider {
				SD[j-1], SD[j] = SD[j], SD[j-1]
			}
			j = j - 1
		}
		for k, _ := range SD {
			sortedSMSData1[k] = SD[k]
		}
	}
	for i := 1; i < len(SD); i++ {
		j := i
		for j > 0 {
			if SD[j-1].Сountry > SD[j].Сountry {
				SD[j-1], SD[j] = SD[j], SD[j-1]
			}
			j = j - 1
		}
		for k, _ := range SD {
			sortedSMSData2[k] = SD[k]
		}
	}

	result = append(result, sortedSMSData1, sortedSMSData2)

	//fmt.Println(result)
	//fmt.Println(sortedSMSData1)
	//fmt.Println(sortedSMSData2)
	return result
}

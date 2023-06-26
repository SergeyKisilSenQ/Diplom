package repo_email

import (
	"Diplom/go-final-dpo/iternal/repo_country"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type StorageEmail map[int]*EmailData

func NewStorageEmail() StorageEmail {
	return make(map[int]*EmailData)
}

func (ED StorageEmail) Put(Country *EmailData) {
	ED[len(ED)] = Country
}

func (ED StorageEmail) ReadFileEmail() {
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	r, err := ioutil.ReadFile(os.Getenv("EMAIL_FILE"))
	if err != nil {
		panic(err)
	}
	s := strings.Fields(string(r))
	for i := 0; i < len(s); i++ {
		var result string
		var e []string
		result = strings.ReplaceAll(s[i], ";", " ")
		e = strings.Fields(result)

		if len(e) == 3 {
			for k, _ := range NSC {
				if e[0] == NSC[k].CodeCountry {
					if e[1] == "Gmail" || e[1] == "Yahoo" || e[1] == "Hotmail" || e[1] == "MSN" || e[1] == "Orange" || e[1] == "Comcast" || e[1] == "AOL" || e[1] == "Live" || e[1] == "RediffMail" || e[1] == "GMX" || e[1] == "ProtonMail" || e[1] == "Yandex" || e[1] == "Mail.ru" {
						NewСountry := e[0]
						NewProvider := e[1]
						NewDeliveryTime, err := strconv.Atoi(e[2])
						if err != nil {
							panic(err)
						}
						NewEmailData := EmailData{
							Country:      NewСountry,
							Provider:     NewProvider,
							DeliveryTime: NewDeliveryTime,
						}
						ED.Put(&NewEmailData)
					}
				}
			}
		}
	}
	//for i, _ := range ED {
	//	fmt.Println(i, ED[i])
	//
	//}
}

func SortedEmailData(ED StorageEmail) map[string][][]*EmailData {
	result := make(map[string][][]*EmailData)
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)

	for i, _ := range NSC {
		tempData := make([]*EmailData, 0)
		sortedEmailDataF := make([]*EmailData, 3)
		sortedEmailDataS := make([]*EmailData, 3)
		slice := make([][]*EmailData, 0)
		for j, _ := range ED {
			if ED[j].Country == NSC[i].CodeCountry {
				ED[j].Country = NSC[i].NameCountry
				tempData = append(tempData, ED[j])
			}
		}
		for l := 1; l < len(tempData); l++ {
			j := l
			for j > 0 {
				if tempData[j-1].DeliveryTime > tempData[j].DeliveryTime {
					tempData[j-1], tempData[j] = tempData[j], tempData[j-1]
				}
				j = j - 1
			}
			for n := 0; n < 3; n++ {
				sortedEmailDataF[n] = tempData[n]
			}
		}
		for l := 1; l < len(tempData); l++ {
			j := l
			for j > 0 {
				if tempData[j-1].DeliveryTime < tempData[j].DeliveryTime {
					tempData[j-1], tempData[j] = tempData[j], tempData[j-1]
				}
				j = j - 1
			}
			for n := 0; n < 3; n++ {
				sortedEmailDataS[n] = tempData[n]
			}
		}
		slice = append(slice, sortedEmailDataF, sortedEmailDataS)
		if sortedEmailDataF[0] != nil && sortedEmailDataS[0] != nil {
			result[NSC[i].NameCountry] = slice
		}
	}

	//fmt.Println(result)
	return result
}

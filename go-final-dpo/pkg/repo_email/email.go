package repo_email

import (
	"Diplom/go-final-dpo/utils"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type StorageEmail map[string][][]*EmailData

func NewStorageEmail() *StorageEmail {
	return &StorageEmail{}
}

func (*StorageEmail) GetEmailData() map[string][][]*EmailData {
	result := make(map[string][][]*EmailData)
	readData, err := os.ReadFile(os.Getenv("EMAIL_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Fields(string(readData))
	tempEmailData := make(map[string][]*EmailData)
	for _, val := range str {
		strArray := strings.Split(val, ";")
		if len(strArray) != 3 {
			log.Println("Incorrect lenght data string")
			continue
		}
		NewСountry, err := utils.CountryCheck(strArray[0])
		if err != nil {
			log.Printf("%s: %s, in email-data\n", err, strArray[0])
			continue
		}
		NewProvider, err := utils.CheckProvider(strArray[1], strings.Split(os.Getenv("EMAIL_PROVIDER"), ", "))
		if err != nil {
			log.Printf("%s: %s, in email-data\n", err, strArray[1])
			continue
		}
		NewDeliveryTime, err := strconv.Atoi(strArray[2])
		if err != nil {
			log.Println(err)
			continue
		}
		NewEmailData := &EmailData{
			Country:      NewСountry,
			Provider:     NewProvider,
			DeliveryTime: NewDeliveryTime,
		}
		tempEmailData[NewСountry] = append(tempEmailData[NewСountry], NewEmailData)
	}

	for s, data := range tempEmailData {
		sortedEmailDataF := make([]*EmailData, 3)
		sortedEmailDataS := make([]*EmailData, 3)
		sort.SliceStable(data, func(i, j int) bool { return data[i].DeliveryTime < data[j].DeliveryTime })
		for n := 0; n < 3; n++ {
			sortedEmailDataF[n] = data[n]
		}
		sort.SliceStable(data, func(i, j int) bool { return data[i].DeliveryTime > data[j].DeliveryTime })
		for n := 0; n < 3; n++ {
			sortedEmailDataS[n] = data[n]
		}
		result[s] = append(result[s], sortedEmailDataF, sortedEmailDataS)
	}
	return result
}

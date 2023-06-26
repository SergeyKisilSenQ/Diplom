package repo_sms

import (
	"Diplom/go-final-dpo/utils"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SMSData struct {
	Сountry      string `json:"country"`
	Bandwidth    int    `json:"bandwidth"`
	ResponseTime int    `json:"response_time"`
	Provider     string `json:"provider"`
}

type StorageSMS [][]*SMSData

func NewStorageSMS() *StorageSMS {
	return &StorageSMS{}
}

func (SD *StorageSMS) GetSmsData() [][]*SMSData {
	sortedSMSData1 := make([]*SMSData, 0)
	sortedSMSData2 := make([]*SMSData, 0)
	readData, err := os.ReadFile(os.Getenv("SMS_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Fields(string(readData))
	for _, val := range str {
		strArray := strings.Split(val, ";")
		if len(strArray) != 4 {
			log.Println("Incorrect lenght data string")
			continue
		}
		NewCountry, err := utils.CountryCheck(strArray[0])
		if err != nil {
			log.Printf("%s: %s, in sms-data\n", err, strArray[0])
			continue
		}
		NewBandwidth, err := strconv.Atoi(strArray[1])
		if err != nil {
			log.Println(err)
			continue
		}
		NewResponseTime, err := strconv.Atoi(strArray[2])
		if err != nil {
			log.Println(err)
			continue
		}
		NewProvider, err := utils.CheckProvider(strArray[3], strings.Split(os.Getenv("SMS_MMS_PROVIDER"), ", "))
		if err != nil {
			log.Println(err)
			continue
		}
		NewSMSData1 := &SMSData{
			Сountry:      NewCountry,
			Bandwidth:    NewBandwidth,
			ResponseTime: NewResponseTime,
			Provider:     NewProvider,
		}
		NewSMSData2 := &SMSData{
			Сountry:      NewCountry,
			Bandwidth:    NewBandwidth,
			ResponseTime: NewResponseTime,
			Provider:     NewProvider,
		}
		sortedSMSData1 = append(sortedSMSData1, NewSMSData1)
		sortedSMSData2 = append(sortedSMSData2, NewSMSData2)
	}
	sort.SliceStable(sortedSMSData1, func(i, j int) bool { return sortedSMSData1[i].Provider < sortedSMSData1[j].Provider })
	sort.SliceStable(sortedSMSData2, func(i, j int) bool { return sortedSMSData2[i].Сountry < sortedSMSData2[j].Сountry })
	result := make([][]*SMSData, 0)
	result = append(result, sortedSMSData1, sortedSMSData2)
	return result
}

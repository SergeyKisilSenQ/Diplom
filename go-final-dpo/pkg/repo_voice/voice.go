package repo_voice

import (
	"Diplom/go-final-dpo/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

type VoiceData struct {
	Country             string  `json:"country"`
	Bandwidth           int     `json:"bandwidth"`
	ResponseTime        int     `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianCallDuration  int     `json:"median_of_calls_time"`
}

type StorageVoice []*VoiceData

func NewStorageVoice() *StorageVoice {
	return &StorageVoice{}
}

func (*StorageVoice) GetVoiceData() []*VoiceData {
	voiceData := make([]*VoiceData, 0)
	readData, err := os.ReadFile(os.Getenv("VOICE_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Fields(string(readData))
	for _, val := range str {
		strArray := strings.Split(val, ";")
		if len(strArray) != 8 {
			log.Println("Incorrect lenght data string")
			continue
		}
		NewCountry, err := utils.CountryCheck(strArray[0])
		if err != nil {
			log.Printf("%s: %s, in voice-data\n", err, strArray[0])
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
		NewProvider, err := utils.CheckProvider(strArray[3], strings.Split(os.Getenv("VOICE_PROVIDER"), ", "))
		if err != nil {
			log.Printf("%s: %s, in voice-data\n", err, strArray[3])
			continue
		}
		NewConnectionStability, err := strconv.ParseFloat(strArray[4], 32)
		if err != nil {
			log.Println(err)
			continue
		}
		NewTTFB, err := strconv.Atoi(strArray[5])
		if err != nil {
			log.Println(err)
			continue
		}
		NewVoicePurity, err := strconv.Atoi(strArray[6])
		if err != nil {
			log.Println(err)
			continue
		}
		NewMedianCallDuration, err := strconv.Atoi(strArray[7])
		if err != nil {
			log.Println(err)
			continue
		}
		NewVoiceData := &VoiceData{
			Country:             NewCountry,
			Bandwidth:           NewBandwidth,
			ResponseTime:        NewResponseTime,
			Provider:            NewProvider,
			ConnectionStability: float32(NewConnectionStability),
			TTFB:                NewTTFB,
			VoicePurity:         NewVoicePurity,
			MedianCallDuration:  NewMedianCallDuration,
		}
		voiceData = append(voiceData, NewVoiceData)
	}
	return voiceData
}

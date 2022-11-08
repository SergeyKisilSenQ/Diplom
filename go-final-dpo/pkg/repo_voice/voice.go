package repo_voice

import (
	"Diplom/go-final-dpo/pkg/repo_country"
	"io/ioutil"
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

type StorageVoice map[int]*VoiceData

func NewStorageVoice() StorageVoice {
	return make(map[int]*VoiceData)
}

func (VD StorageVoice) Put(Country *VoiceData) {
	VD[len(VD)] = Country
}

func (VD StorageVoice) ReadFileVoice() {
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	r, err := ioutil.ReadFile("simulator/voice.data")
	if err != nil {
		panic(err)
	}
	s := strings.Fields(string(r))
	for i := 0; i < len(s); i++ {
		var result string
		var e []string
		result = strings.ReplaceAll(s[i], ";", " ")
		e = strings.Fields(result)

		if len(e) == 8 {
			for k, _ := range NSC {
				if e[0] == NSC[k].NameCountry {
					if e[3] == "TransparentCalls" || e[3] == "E-Voice" || e[3] == "JustPhone" {
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
						NewConnectionStability, err := strconv.ParseFloat(e[4], 32)
						if err != nil {
							panic(err)
						}
						NewTTFB, err := strconv.Atoi(e[5])
						if err != nil {
							panic(err)
						}
						NewVoicePurity, err := strconv.Atoi(e[6])
						if err != nil {
							panic(err)
						}
						NewMedianCallDuration, err := strconv.Atoi(e[7])
						if err != nil {
							panic(err)
						}
						NewVoiceData := VoiceData{
							Country:             NewСountry,
							Bandwidth:           NewBandwidth,
							ResponseTime:        NewResponseTime,
							Provider:            NewProvider,
							ConnectionStability: float32(NewConnectionStability),
							TTFB:                NewTTFB,
							VoicePurity:         NewVoicePurity,
							MedianCallDuration:  NewMedianCallDuration,
						}
						VD.Put(&NewVoiceData)
					}
				}
			}
		}
	}
	//for i, _ := range VD {
	//	fmt.Printf("%s, %v, %v, %s, %v, %v, %v,%v  \n", VD[i].Country, VD[i].Bandwidth, VD[i].ResponseTime, VD[i].Provider, VD[i].ConnectionStability, VD[i].TTFB, VD[i].VoicePurity, VD[i].MedianCallDuration)
	//}
}

func AddVoiceData(VD StorageVoice) []*VoiceData {
	addVoiceData := make([]*VoiceData, len(VD))
	for i, _ := range VD {
		addVoiceData[i] = VD[i]
	}
	return addVoiceData
}

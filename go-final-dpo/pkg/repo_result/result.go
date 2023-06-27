package repo_result

import (
	"Diplom/go-final-dpo/pkg/repo_billing"
	"Diplom/go-final-dpo/pkg/repo_email"
	"Diplom/go-final-dpo/pkg/repo_incident"
	"Diplom/go-final-dpo/pkg/repo_mms"
	"Diplom/go-final-dpo/pkg/repo_sms"
	"Diplom/go-final-dpo/pkg/repo_support"
	"Diplom/go-final-dpo/pkg/repo_voice"
	"sync"
)

type ResultSetT struct {
	SMS       [][]*repo_sms.SMSData                `json:"sms"`
	MMS       [][]*repo_mms.MMSData                `json:"mms"`
	VoiceCall []*repo_voice.VoiceData              `json:"voice_call"`
	Email     map[string][][]*repo_email.EmailData `json:"email"`
	Billing   *repo_billing.BillingData            `json:"billing"`
	Support   []int                                `json:"support"`
	Incidents []*repo_incident.IncidentData        `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

func getResultData() *ResultSetT {
	var wg sync.WaitGroup
	var result *ResultSetT
	var smsData [][]*repo_sms.SMSData
	var mmsData [][]*repo_mms.MMSData
	var voiceCallData []*repo_voice.VoiceData
	var emailData map[string][][]*repo_email.EmailData
	var billingData *repo_billing.BillingData
	var supportData []int
	var incidentsData []*repo_incident.IncidentData
	wg.Add(7)

	go func() {
		defer wg.Done()
		smsData = repo_sms.NewStorageSMS().GetSmsData()
	}()
	go func() {
		defer wg.Done()
		MD := repo_mms.NewStorageMMS()
		MD.GetMMS()
		mmsData = repo_mms.SortedMMSData(MD)
	}()

	go func() {
		defer wg.Done()
		voiceCallData = repo_voice.NewStorageVoice().GetVoiceData()
	}()

	go func() {
		defer wg.Done()
		emailData = repo_email.NewStorageEmail().GetEmailData()
	}()

	go func() {
		defer wg.Done()
		billingData = repo_billing.NewStorageBilling().ReadFileBilling()
	}()

	go func() {
		defer wg.Done()
		SSD := repo_support.NewStorageSupport()
		SSD.GetSupport()
		supportData = repo_support.SupportDataStatus(SSD)
	}()

	go func() {
		defer wg.Done()
		ID := repo_incident.NewStorageIncident()
		ID.GetIncident()
		incidentsData = repo_incident.AddIncidentData(ID)
	}()

	wg.Wait()

	result = &ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voiceCallData,
		Email:     emailData,
		Billing:   billingData,
		Support:   supportData,
		Incidents: incidentsData,
	}

	return result
}

func GetResult() *ResultT {
	var result *ResultT
	r := getResultData()
	if len(r.SMS[0]) == 0 || len(r.MMS[0]) == 0 || len(r.VoiceCall) == 0 || len(r.Email) == 0 || r.Support[0] == 0 || len(r.Incidents) == 0 {
		result = &ResultT{
			Status: false,
			Error:  "error data",
		}
	} else {
		result = &ResultT{
			Status: true,
			Data:   *r,
			Error:  "",
		}
	}
	return result
}

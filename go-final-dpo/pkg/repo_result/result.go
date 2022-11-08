package repo_result

import (
	"Diplom/go-final-dpo/pkg/repo_billing"
	"Diplom/go-final-dpo/pkg/repo_email"
	"Diplom/go-final-dpo/pkg/repo_incident"
	"Diplom/go-final-dpo/pkg/repo_mms"
	"Diplom/go-final-dpo/pkg/repo_sms"
	"Diplom/go-final-dpo/pkg/repo_support"
	"Diplom/go-final-dpo/pkg/repo_voice"
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

	var result *ResultSetT

	SD := repo_sms.NewStorageSMS()
	SD.ReadFileSMS()
	MD := repo_mms.NewStorageMMS()
	MD.GetMMS()
	VD := repo_voice.NewStorageVoice()
	VD.ReadFileVoice()
	ED := repo_email.NewStorageEmail()
	ED.ReadFileEmail()
	BD := repo_billing.NewStorageBilling()
	BD.ReadFileBilling()
	SSD := repo_support.NewStorageSupport()
	SSD.GetSupport()
	ID := repo_incident.NewStorageIncident()
	ID.GetIncident()

	result = &ResultSetT{
		SMS:       repo_sms.SortedSMSData(SD),
		MMS:       repo_mms.SortedMMSData(MD),
		VoiceCall: repo_voice.AddVoiceData(VD),
		Email:     repo_email.SortedEmailData(ED),
		Billing:   BD[0],
		Support:   repo_support.SupportDataStatus(SSD),
		Incidents: repo_incident.AddIncidentData(ID),
	}

	return result
}

func GetResult() ResultT {
	return ResultT{
		Status: true,
		Data:   *getResultData(),
		Error:  "",
	}
}

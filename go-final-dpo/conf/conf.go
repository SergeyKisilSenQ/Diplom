package conf

import "os"

func GetConf() {
	os.Setenv("SMS_FILE", "simulator/sms.data")
	os.Setenv("SMS_MMS_PROVIDER", "Topolo, Rond, Kildy")
	os.Setenv("VOICE_PROVIDER", "TransparentCalls, E-Voice, JustPhone")
	os.Setenv("EMAIL_PROVIDER", "Gmail, Yahoo, Hotmail, MSN, Orange, Comcast, AOL, Live, RediffMail, GMX, ProtonMail, Yandex, Mail.ru")
	os.Setenv("MMS_URL", "http://127.0.0.1:8383/mms")
	os.Setenv("VOICE_FILE", "simulator/voice.data")
	os.Setenv("EMAIL_FILE", "simulator/email.data")
	os.Setenv("BILLING_FILE", "simulator/billing.data")
	os.Setenv("SUPPORT_URL", "http://127.0.0.1:8383/support")
	os.Setenv("INCIDENT_URL", "http://127.0.0.1:8383/accendent")
	os.Setenv("COUNTRY_FILE", "go-final-dpo/iternal/data_file/countries_codes.csv")
}

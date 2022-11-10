package conf

import "os"

func GetConf() {
	os.Setenv("SMS_FILE", "simulator/sms.data")
	os.Setenv("MMS_URL", "http://127.0.0.1:8383/mms")
	os.Setenv("VOICE_FILE", "simulator/voice.data")
	os.Setenv("EMAIL_FILE", "simulator/email.data")
	os.Setenv("BILLING_FILE", "simulator/billing.data")
	os.Setenv("SUPPORT_URL", "http://127.0.0.1:8383/support")
	os.Setenv("INCIDENT_URL", "http://127.0.0.1:8383/accendent")
	os.Setenv("COUNTRY_FILE", "go-final-dpo/countries_codes.csv")
}

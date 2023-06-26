package repo_country

import (
	"encoding/csv"
	"os"
	"strings"
)

type CountryStorage struct {
	CodeCountry string
	NameCountry string
}

type NewCountryStorage map[string]*CountryStorage

func CreateNewCountryStorage() NewCountryStorage {
	return make(map[string]*CountryStorage)
}
func (NCS NewCountryStorage) Put(Country *CountryStorage) {
	NCS[Country.CodeCountry] = Country
}

func ReadFile(nsc NewCountryStorage) {
	file, err := os.Open(os.Getenv("COUNTRY_FILE"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	for {
		record, e := reader.Read()
		if e != nil {
			break
		}
		NewNameCountry := strings.Trim(record[0], " ")
		NewCodeCountry := strings.Trim(record[1], " ")
		NewCountry := CountryStorage{
			CodeCountry: NewCodeCountry,
			NameCountry: NewNameCountry,
		}
		nsc.Put(&NewCountry)
	}
}

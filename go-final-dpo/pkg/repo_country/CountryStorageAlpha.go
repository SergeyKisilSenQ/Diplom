package repo_country

import (
	"encoding/csv"
	"os"
)

type CoutryStorage struct {
	NameCountry string
}
type NewCountryStorage map[string]*CoutryStorage

func CreateNewCountryStorage() NewCountryStorage {
	return make(map[string]*CoutryStorage)
}
func (NCS NewCountryStorage) Put(NameCountry *CoutryStorage) {
	NCS[NameCountry.NameCountry] = NameCountry
}

func ReadFile(nsc NewCountryStorage) (NCS NewCountryStorage) {
	file, err := os.Open("go-final-dpo/country.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 1
	for {
		record, e := reader.Read()
		if e != nil {
			break
		}
		NewNameCountry := record[0]
		NewCountry := CoutryStorage{
			NameCountry: NewNameCountry,
		}
		nsc.Put(&NewCountry)
	}
	return NCS
}

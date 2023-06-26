package utils

import (
	"Diplom/go-final-dpo/iternal/repo_country"
	"errors"
)

func CountryCheck(str string) (string, error) {
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	for i, _ := range NSC {
		if str == NSC[i].CodeCountry {
			return NSC[i].NameCountry, nil
		}
	}
	return "", errors.New("not found code country")
}

func CheckProvider(str string, checkArray []string) (string, error) {
	for _, s := range checkArray {
		if str == s {
			return s, nil
		}
	}
	return "", errors.New("incorrect provider")
}

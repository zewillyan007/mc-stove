package util

import "time"

func CalculateAge(birthDateString string) (float64, error) {
	now := time.Now()

	birthDate, err := time.Parse("2006-01-02", birthDateString)

	if err != nil {
		return 0, err
	}

	diff := now.Sub(birthDate)
	age := diff.Hours() / 24 / 365.25

	return age, nil
}

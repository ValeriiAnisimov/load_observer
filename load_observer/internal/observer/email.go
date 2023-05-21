package observer

import (
	"log"
	"sort"
	"strconv"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func GetherEmailData() ([]EmailData, error) {
	a, err := ParseCSVtoSlice("../simulator/email.data")
	if err != nil {
		return nil, err
	}
	b := FilterEmailData(&a)
	c := EmailToStruct(b)

	return c, nil
}

func FilterEmailData(data *[][]string) *[][]string {
	result := make([][]string, 0)

	for _, v := range *data {
		if !EmailDataValidation(v) {
			continue
		}
		result = append(result, v)
	}
	return &result
}

func EmailDataValidation(data []string) bool {
	validEmailProvidres := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Proton Mail", "Yandex", "Mail.ru"}
	validFieldsCount := 3

	if !IsValidCountryCode(data[0], countryCodes) {
		return false
	}
	if !IsValidFieldsCount(validFieldsCount, &data) {
		return false
	}
	if !IsValidProvider(data[1], &validEmailProvidres) {
		return false
	}
	return true
}

func EmailToStruct(data *[][]string) []EmailData {
	result := make([]EmailData, 0, len(*data))

	for _, v := range *data {
		EmailItem := EmailData{}

		EmailItem.Country = v[0]
		EmailItem.Provider = v[1]

		a, err := strconv.ParseInt(v[2], 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		EmailItem.DeliveryTime = int(a)

		result = append(result, EmailItem)
	}

	return result
}

func getResultEmail() (map[string][][]EmailData, error) {
	result := make(map[string][][]EmailData, 0)

	data, err := GetherEmailData()
	if err != nil {
		return nil, err
	}

	for countryCode := range countryCodes {
		slowestProviders, fastestProviders := groupedByCountryAndSpeedProviders(data, countryCode)
		if len(slowestProviders) != 0 && len(fastestProviders) != 0 {
			result[countryCodes[countryCode]] = [][]EmailData{slowestProviders, fastestProviders}
		}
	}

	return result, nil
}

func groupedByCountryAndSpeedProviders(data []EmailData, countryCode string) ([]EmailData, []EmailData) {

	groupedByCountry := make([]EmailData, 0)
	for _, item := range data {
		if item.Country == countryCode {
			groupedByCountry = append(groupedByCountry, item)
		}
	}

	sort.SliceStable(groupedByCountry, func(i, j int) bool {
		return groupedByCountry[i].DeliveryTime < groupedByCountry[j].DeliveryTime
	})

	if len(groupedByCountry) < 3 {
		return groupedByCountry, groupedByCountry
	}

	slowest := groupedByCountry[len(groupedByCountry)-3:]
	fastest := groupedByCountry[:3]

	return slowest, fastest
}

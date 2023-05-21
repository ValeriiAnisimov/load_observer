package observer

import "sort"

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func GetherSMSData() ([]SMSData, error) {
	a, err := ParseCSVtoSlice("../simulator/sms.data")
	if err != nil {
		return nil, err
	}
	b := FilterData(&a)
	c := SMSStringToStruct(b)

	return c, nil
}

func FilterData(data *[][]string) *[][]string {
	result := make([][]string, 0)

	for _, v := range *data {
		if !SMSDataValidation(v) {
			continue
		}
		result = append(result, v)
	}
	return &result
}

func SMSDataValidation(data []string) bool {
	validSMSProvidres := []string{"Topolo", "Rond", "Kildy"}
	validFieldsCount := 4

	if !IsValidCountryCode(data[0], countryCodes) {
		return false
	}
	if !IsValidFieldsCount(validFieldsCount, &data) {
		return false
	}
	if !IsValidProvider(data[3], &validSMSProvidres) {
		return false
	}
	return true
}

func SMSStringToStruct(data *[][]string) []SMSData {
	result := make([]SMSData, 0, len(*data))

	for _, v := range *data {
		smsItem := SMSData{}

		smsItem.Country = v[0]
		smsItem.Bandwidth = v[1]
		smsItem.ResponseTime = v[2]
		smsItem.Provider = v[3]

		result = append(result, smsItem)
	}

	return result
}

func getResultSMS() ([][]SMSData, error) {
	data, err := GetherSMSData()
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		v.Country = GetFullContryName(v.Country, countryCodes)
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Provider < data[j].Provider
	})
	sortedByProvider := data

	sortedByCountry := make([]SMSData, len(data))
	copy(sortedByCountry, data)

	sort.SliceStable(sortedByCountry, func(i, j int) bool {
		return sortedByCountry[i].Country < sortedByCountry[j].Country
	})

	result := [][]SMSData{sortedByProvider, sortedByCountry}

	return result, nil
}

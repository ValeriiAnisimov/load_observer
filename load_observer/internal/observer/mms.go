package observer

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func GetherMMSData() ([]MMSData, error) {
	a, err := GetMMSData("http://localhost:8383/mms")
	if err != nil {
		return nil, err
	}

	b := FilterMMSData(&a)

	return b, nil
}

func GetMMSData(addr string) ([]MMSData, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return []MMSData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []MMSData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []MMSData{}, err
	}

	data := make([]MMSData, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []MMSData{}, err
	}
	return data, nil
}

func FilterMMSData(data *[]MMSData) []MMSData {
	result := make([]MMSData, 0)

	for _, v := range *data {
		if !MMSDataValidation(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func MMSDataValidation(data MMSData) bool {
	validMMSProvidres := []string{"Topolo", "Rond", "Kildy"}

	if !IsValidCountryCode(data.Country, countryCodes) {
		return false
	}
	if !IsValidProvider(data.Provider, &validMMSProvidres) {
		return false
	}
	return true
}

func getResultMMS() ([][]MMSData, error) {
	data, err := GetherMMSData()
	if err != nil {
		return [][]MMSData{}, err
	}

	for _, v := range data {
		v.Country = GetFullContryName(v.Country, countryCodes)
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Provider < data[j].Provider
	})
	sortedByProvider := data

	sortedByCountry := make([]MMSData, len(data))
	copy(sortedByCountry, data)

	sort.SliceStable(sortedByCountry, func(i, j int) bool {
		return sortedByCountry[i].Country < sortedByCountry[j].Country
	})

	result := [][]MMSData{sortedByProvider, sortedByCountry}

	return result, nil
}

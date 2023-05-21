package observer

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func GetherIncidentData() ([]IncidentData, error) {
	a, err := GetIncidentData("http://localhost:8383/accendent")
	if err != nil {
		return nil, err
	}
	b := FilterIncidentData(&a)

	return b, nil
}

func GetIncidentData(addr string) ([]IncidentData, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return []IncidentData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []IncidentData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []IncidentData{}, err
	}

	data := make([]IncidentData, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []IncidentData{}, err
	}
	return data, nil
}

func FilterIncidentData(data *[]IncidentData) []IncidentData {
	result := make([]IncidentData, 0)

	for _, v := range *data {
		if !IncidentDataValidation(v) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func IncidentDataValidation(data IncidentData) bool {
	validStatusCodes := []string{"active", "closed"}
	if !IsValidIncindentStatusCode(data, validStatusCodes) {
		return false
	}
	return true
}

func getResultIncidents() ([]IncidentData, error) {
	data, err := GetherIncidentData()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(data, func(i, j int) bool {
		if data[i].Status == "active" {
			return true
		}
		return false
	})

	return data, nil
}

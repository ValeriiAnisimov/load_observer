package observer

import (
	"encoding/json"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func GetherSupportData() ([]SupportData, error) {
	a, err := GetSupportData("http://localhost:8383/support")
	if err != nil {
		return nil, err
	}

	return a, nil
}

func GetSupportData(addr string) ([]SupportData, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return []SupportData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []SupportData{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []SupportData{}, err
	}

	data := make([]SupportData, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []SupportData{}, err
	}
	return data, nil
}

func getResultSupport() ([]int, error) {
	data, err := GetherSupportData()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []int{}, err
	}

	activeTickets := 0
	for _, support := range data {
		activeTickets += support.ActiveTickets
	}

	avgTime := activeTickets * 60 / 18
	load := 1
	switch {
	case activeTickets >= 9 && activeTickets <= 16:
		load = 2
	case activeTickets > 16:
		load = 3
	}

	return []int{load, avgTime}, nil
}

package observer

import (
	"log"
	"strconv"
)

type VoiceCallData struct {
	Country        string  `json:"country"`
	Load           string  `json:"bandwidth"`
	ResponseTime   string  `json:"response_time"`
	Provider       string  `json:"provider"`
	Stability      float32 `json:"connection_stability"`
	TTFB           int     `json:"ttfb"`
	Purity         int     `json:"voice_purity"`
	MedianDuration int     `json:"median_of_calls_time"`
}

func GetherVoiceCallData() ([]VoiceCallData, error) {
	a, err := ParseCSVtoSlice("../simulator/voice.data")
	if err != nil {
		return nil, err
	}
	b := FilterVoiceCallData(&a)
	c := VoiceCVSToStruct(b)

	return c, nil
}

func FilterVoiceCallData(data *[][]string) *[][]string {
	result := make([][]string, 0)

	for _, v := range *data {
		if !VoiceCallDataValidation(v) {
			continue
		}
		result = append(result, v)
	}
	return &result
}

func VoiceCallDataValidation(data []string) bool {
	validVoiceCallProvidres := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	validFieldsCount := 8

	if !IsValidCountryCode(data[0], countryCodes) {
		return false
	}
	if !IsValidFieldsCount(validFieldsCount, &data) {
		return false
	}
	if !IsValidProvider(data[3], &validVoiceCallProvidres) {
		return false
	}
	return true
}

func VoiceCVSToStruct(data *[][]string) []VoiceCallData {
	result := make([]VoiceCallData, 0, len(*data))

	for _, v := range *data {
		VoiceCallItem := VoiceCallData{}

		VoiceCallItem.Country = v[0]
		VoiceCallItem.Load = v[1]
		VoiceCallItem.ResponseTime = v[2]
		VoiceCallItem.Provider = v[3]

		a, err := strconv.ParseFloat(v[4], 32)
		VoiceCallItem.Stability = float32(a)
		if err != nil {
			log.Println(err)
			continue
		}

		b, err := strconv.ParseInt(v[5], 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		VoiceCallItem.TTFB = int(b)

		c, err := strconv.ParseInt(v[6], 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		VoiceCallItem.Purity = int(c)

		d, err := strconv.ParseInt(v[7], 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		VoiceCallItem.MedianDuration = int(d)

		result = append(result, VoiceCallItem)
	}

	return result
}

func getResultVoice(voicePath string) ([]VoiceCallData, error) {
	voiceCallData, err := GetherVoiceCallData()
	if err != nil {
		return nil, err
	}

	return voiceCallData, nil
}

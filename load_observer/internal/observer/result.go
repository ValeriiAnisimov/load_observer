package observer

import (
	"log"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

func GetResultT() ResultT {
	resultSetT, err := GetResultSetT()
	resultT := ResultT{}

	resultT.Status = true
	resultT.Data = resultSetT
	resultT.Error = ""

	if err != nil {
		log.Printf("FAILED:%+v\n", err)
		resultT.Status = false
		resultT.Error = err.Error()
	}
	return resultT
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

func GetResultSetT() (ResultSetT, error) {
	var result ResultSetT

	var err error
	result.SMS, err = getResultSMS()
	if err != nil {
		return result, err
	}

	result.MMS, err = getResultMMS()
	if err != nil {
		return result, err
	}

	result.VoiceCall, err = GetherVoiceCallData()
	if err != nil {
		return ResultSetT{}, err
	}

	result.Email, err = getResultEmail()
	if err != nil {
		return ResultSetT{}, err
	}

	result.Billing, err = getResultBilling()
	if err != nil {
		return ResultSetT{}, err
	}

	result.Support, err = getResultSupport()
	if err != nil {
		return ResultSetT{}, err
	}

	result.Incidents, err = getResultIncidents()
	if err != nil {
		return ResultSetT{}, err
	}

	return result, nil
}

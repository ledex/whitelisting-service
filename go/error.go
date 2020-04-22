package swagger

import (
	"encoding/json"
	"net/http"
)

type LogType int

//go:generate easytags $GOFILE

const (
	Info    LogType = iota
	Warning LogType = iota
	Error   LogType = iota
)

type LogMessage struct {
	Type       LogType `json:"type"`
	Identifier string  `json:"identifier"`
	Details    string  `json:"details"`
}

func writeLogToWriter(w http.ResponseWriter, t LogType, i string, d string) {
	log := LogMessage{
		Type:       t,
		Identifier: i,
		Details:    d,
	}

	result, _ := json.MarshalIndent(log, "", " ")

	switch t {
	case Info:
	case Warning:
		w.WriteHeader(http.StatusOK)
		break
	case Error:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(result)
}

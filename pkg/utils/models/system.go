package models

import "encoding/json"

// swagger:model System
type System struct {
	// the status of the system
	// readOnly: true
	Status string `json:"status"`
	// a message describing the system health
	// readOnly: true
	Message string `json:"message"`
	// the current version the server is runnnig
	// readOnly: true
	Version string `json:"version"`
}

func (s *System) ToJSON() string {
	b, _ := json.Marshal(s)
	return string(b)
}

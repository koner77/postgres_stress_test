package domain

import (
	"encoding/json"
)

type IncomingLog struct {
	TestID           string          `json:"test_id"`
	RowNumber        int             `json:"row_number"`
	Request          json.RawMessage `json:"request"`
	Response         json.RawMessage `json:"response"`
	IP               string          `json:"ip"`
	FallbackProvider string          `json:"provider"`
	GeneralProvider  string          `json:"general_provider"`
	Backfill         uint8           `json:"backfill"`
	RequestAt        string          `json:"request_at"`
	ResponseAt       string          `json:"response_at"`
	ResponseTime     float32         `json:"response_time"`
	AddedAt          string          `json:"added_at"`
	FallbackReason   string          `json:"fallback_reason"`
	Events           json.RawMessage `json:"events"`
}

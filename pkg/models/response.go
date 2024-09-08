package models

type HostListResponse struct {
	Hosts      []VirtualMachine `json:"entities"`
	ApiVersion string           `json:"api_version"`
	Metadata   struct {
		TotalMatches int    `json:"total_matches"`
		Kind         string `json:"kind"`
		Length       int    `json:"length"`
		Offset       int    `json:"offset"`
	} `json:"metadata"`
}

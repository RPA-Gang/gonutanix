package gonut

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
)

const (
	hostUrl   string = "rsssydvpc001.apac.linkgroup.corp:9440"
	clusterId string = "0005f778-5751-e25f-65a1-3a68dd8c938f"
)

var (
	user   = "schneet@linkgroup.com"
	pass   = "Shitpass69@@"
	target = &url.URL{
		Scheme: "https",
		Host:   hostUrl,
		Path:   "/api/nutanix/v3/vms/list",
	}
)

func main() {
	response := getHosts()
	for _, host := range response.Hosts {
		log.Print(host.Status.Resources)
	}
}

func getData() (HostListResponse, []http.Cookie) {
	return HostListResponse{}, []http.Cookie{}
}

func getHosts() HostListResponse {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		target.String(),
		bytes.NewBufferString(`{"kind": "vm"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(user, pass)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseStruct := HostListResponse{}
	if err = json.Unmarshal(body, &responseStruct); err != nil {
		log.Fatal(err)
	}
	return responseStruct
}

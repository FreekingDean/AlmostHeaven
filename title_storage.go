package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

type TitleStorageResponse struct {
	Checksum    string `json:"checksum"`
	DownloadURL string `json:"download_url"`
	LastUpdated string `json:"last_updated"`
	Size        int    `json:"size"`
}

func TitleStorage(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	TitleStorageGateways(rec, nil)
	data, _ := ioutil.ReadAll(rec.Result().Body)
	resp := buildPlatformSuccess(TitleStorageResponse{
		Checksum:    fmt.Sprintf("%x", sha256.Sum256(data)),
		DownloadURL: "https://api.bethesda.net/titlestorage/actualstorage/gateways",
		LastUpdated: time.Now().String(),
		Size:        100,
	})
	DefaultJSONEncoder(w, resp)
}

type GatewayResp struct {
	Global  Global   `json:"global"`
	Regions []Region `json:"regions"`
}

type Global struct {
	Services []Service `json:"services"`
}

type Region struct {
	ID       int       `json:"id"`
	PingURL  string    `json:"ping_url"`
	Services []Service `json:"services"`
}

type Service struct {
	Name      string `json:"name"`
	PublicKey string `json:"pubkey"`
	URL       string `json:"url"`
}

func TitleStorageGateways(w http.ResponseWriter, _ *http.Request) {
	regionIDs := []int{11, 6, 7, 8, 2, 3, 5}
	regions := []Region{}
	for _, regionID := range regionIDs {
		region := Region{
			ID:      regionID,
			PingURL: fmt.Sprintf("https://api.bethesda.net/regions/%d/ping", regionID),
			Services: []Service{
				Service{
					Name:      "bps-gatewayreg",
					PublicKey: "nokey",
					URL:       fmt.Sprintf("https://api.behtesda.net/region/%d/gatewayreg", regionID),
				},
			},
		}
		regions = append(regions, region)
	}
	resp := GatewayResp{
		Global: Global{
			Services: []Service{
				Service{
					Name:      "bps-loginqueue",
					PublicKey: "loginqueue_key",
					URL:       "https://api.bethesda.net/loginqueue",
				},
				Service{
					Name:      "bps-gateway",
					PublicKey: "loginqueue_key",
					URL:       "https://api.bethesda.net/gateway",
				},
				Service{
					Name:      "bps-bigateway",
					PublicKey: "loginqueue_key",
					URL:       "https://api.bethesda.net/bigateway",
				},
				Service{
					Name:      "bps-pushy",
					PublicKey: "loginqueue_key",
					URL:       "https://api.bethesda.net/pushy",
				},
			},
		},
		Regions: regions,
	}
	DefaultJSONEncoder(w, resp)
}

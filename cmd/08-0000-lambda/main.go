package helloworld

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

//nolint:tagliatelle // The provided json structure from tdr's official site
type StandbyTimes []struct {
	FacilityID             string           `json:"FacilityID,omitempty"`
	FacilityName           string           `json:"FacilityName,omitempty"`
	FacilityKanaName       string           `json:"FacilityKanaName,omitempty"`
	NewFlg                 bool             `json:"NewFlg,omitempty"`
	FacilityURLSP          any              `json:"FacilityURLSP,omitempty"`
	FacilityStatusCD       any              `json:"FacilityStatusCD,omitempty"`
	FacilityStatus         any              `json:"FacilityStatus,omitempty"`
	StandbyTime            any              `json:"StandbyTime,omitempty"`
	OperatingHoursFromDate string           `json:"OperatingHoursFromDate,omitempty"`
	OperatingHoursFrom     string           `json:"OperatingHoursFrom,omitempty"`
	OperatingHoursToDate   string           `json:"OperatingHoursToDate,omitempty"`
	OperatingHoursTo       string           `json:"OperatingHoursTo,omitempty"`
	OperatingStatusCD      string           `json:"OperatingStatusCD,omitempty"`
	OperatingStatus        string           `json:"OperatingStatus,omitempty"`
	SunsetFlg              bool             `json:"SunsetFlg,omitempty"`
	DPAStatusCD            any              `json:"DPAStatusCD,omitempty"`
	DPAStatus              any              `json:"DPAStatus,omitempty"`
	PPStatusCD             any              `json:"PPStatusCD,omitempty"`
	PPStatus               any              `json:"PPStatus,omitempty"`
	Fsflg                  bool             `json:"Fsflg,omitempty"`
	FsStatusflg            any              `json:"FsStatusflg,omitempty"`
	FsStatus               any              `json:"FsStatus,omitempty"`
	FsStatusCD             any              `json:"FsStatusCD,omitempty"`
	FsStatusStartDate      any              `json:"FsStatusStartDate,omitempty"`
	FsStatusStartTime      any              `json:"FsStatusStartTime,omitempty"`
	FsStatusEndDate        any              `json:"FsStatusEndDate,omitempty"`
	FsStatusEndTime        any              `json:"FsStatusEndTime,omitempty"`
	UseLimitFlg            bool             `json:"UseLimitFlg,omitempty"`
	UseStandbyTimeStyle    bool             `json:"UseStandbyTimeStyle,omitempty"`
	OperatingChgFlg        bool             `json:"OperatingChgFlg,omitempty"`
	UpdateTime             string           `json:"UpdateTime,omitempty"`
	OperatingHours         []OperatingHours `json:"operatingHours,omitempty"`
}

//nolint:tagliatelle // The provided json structure from tdr's official site
type OperatingHours struct {
	OperatingHoursFromDate string `json:"OperatingHoursFromDate,omitempty"`
	OperatingHoursFrom     string `json:"OperatingHoursFrom,omitempty"`
	OperatingHoursToDate   string `json:"OperatingHoursToDate,omitempty"`
	OperatingHoursTo       string `json:"OperatingHoursTo,omitempty"`
	SunsetFlg              bool   `json:"SunsetFlg,omitempty"`
	OperatingStatusCD      string `json:"OperatingStatusCD,omitempty"`
	OperatingStatus        string `json:"OperatingStatus,omitempty"`
	OperatingChgFlg        bool   `json:"OperatingChgFlg,omitempty"`
}

type FacilityType int

const (
	Attraction FacilityType = iota
)

type ParkType int

const (
	TokyoDisneyLand ParkType = iota
	TokyoDisneySea
)

const (
	BaseURL       = "https://www.tokyodisneyresort.jp/_/realtime"
	TDLAttraction = "tdl_attraction.json"
	TDSAttraction = "tds_attraction.json"
)

var jsonMap = map[FacilityType]map[ParkType]string{
	Attraction: {
		TokyoDisneyLand: TDLAttraction,
		TokyoDisneySea:  TDSAttraction,
	},
}

func New(facility FacilityType, park ParkType) *url.URL {
	// https://www.tokyodisneyresort.jp/_/realtime/tdl_attraction.json?1726296151344
	base, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}

	base = base.JoinPath(jsonMap[facility][park])
	base.RawQuery = fmt.Sprint(time.Now().UnixNano())

	return base
}

func Req(ctx context.Context, park ParkType) ([]byte, error) {
	reader := bytes.NewReader(make([]byte, 0))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, New(Attraction, park).String(), reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "LearningGolang/0.1.0")
	req.Header.Add("Host", "shusann01116.dev")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode not 200, actual: %v", resp.StatusCode)
	}

	rawBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return rawBytes, nil
}

func HandleRequest(ctx context.Context) error {
	parks := []ParkType{TokyoDisneyLand, TokyoDisneySea}
	for _, p := range parks {
		slog.Info(fmt.Sprintf("%v", p))
		ctx, f := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
		defer f()

		respBody, err := Req(ctx, p)
		if err != nil {
			return err
		}

		var resp StandbyTimes
		if err := json.Unmarshal(respBody, &resp); err != nil {
			return err
		}

		slog.Info(fmt.Sprintf("len(resp) = %v", len(resp)))
		slog.Info(fmt.Sprintf("resp[0] = %+v", resp[0]))
	}
	return nil
}

func init() {
	functions.HTTP("HelloHTTP", helloHTTP)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func helloHTTP(w http.ResponseWriter, r *http.Request) {
	HandleRequest(r.Context())
}

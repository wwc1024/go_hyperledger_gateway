package dto

import (
	"go_gateway/public"

	"github.com/gin-gonic/gin"
)

type PanelGroupDataOutput struct {
	ChannelNum    int64 `json:"channelNum"`
	JoinedChannel int64 `json:"joinedChannel"`
	JoinService   int64 `json:"joinService"`
	PeerNum       int64 `json:"peerNum"`
}

type DashServiceStatItemOutput struct {
	Name        string `json:"name"`
	ServiceType string `json:"service_type"`
	Value       int64  `json:"value"`
}

type DashServiceStatOutput struct {
	Legend []string                    `json:"legend"`
	Data   []DashServiceStatItemOutput `json:"data"`
}

/*------dhtdashboard-------*/
type PanelGroupDataOutput2 struct {
	DhtNum    int64 `json:"dhtNum"`
	NowdhtID  int64 `json:"nowdhtId"`
	NowSample int64 `json:"nowSample"`
}

type DhtSetting struct {
	DhtID     int64 `json:"dhtId"`
	NowSample int64 `json:"nowSample"`
}

func (params *DhtSetting) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type DhtMap struct {
	DhtID   string `json:"dhtId"`
	Address string `json:"address"`
}

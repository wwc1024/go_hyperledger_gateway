package dto

import (
	_ "time"

	"go_gateway/public"

	"github.com/gin-gonic/gin"
)

func (param *ChannelAddInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ChannelAddInput struct {
	ChannelName string `json:"channel_name" form:"channel_name" comment:"通道名" example:"" validate:"required"` //通道名
	File        string `json:"file" form:"file" comment:"文件名" example:"" validate:"required"`
}

type ChannelListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                      //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` //每页条数
}

func (param *ChannelListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ChannelListItemInDb struct {
	ID        string `json:"id" form:"id"` //id
	Peer      string `json:"peer" form:"peer"`
	ChannelID string `json:"channel_id" form:"channel_id"` //通道ID
	Chaincode string `json:"chaincodelist" form:"chaincodelist"`
}
type ChannelListItemOutput struct {
	ID        string `json:"id" form:"id"` //id
	Peer      string `json:"peer" form:"peer"`
	ChannelID string `json:"channel_id" form:"channel_id"` //通道ID
	Chaincode string `json:"chaincode" form:"chaincode"`
	Detail    string `json:"detail" form:"detail"`
	PeerNum   string `json:"peerNum" form:"peerNum"`
}
type ChannelListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" validate:""` //总数
	List  []ChannelListItemOutput `json:"list" form:"list" comment:"列表" validate:""`   //列表
}

//合约list
func (param *ContentAddInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ContentAddInput struct {
	ID          int64  `json:"id" form:"id" comment:"id" validate:""`                        //id
	ContentName string `json:"content_name" form:"content_name" comment:"name" validate:""`  //合约名称
	Tap         string `json:"tap" form:"tap" comment:"tap" validate:""`                     //tap
	ServiceName string `json:"service_name" form:"service_name" comment:"sname" validate:""` //ServiceName
	ServiceType string `json:"service_type" form:"service_type" comment:"type" validate:""`  //ServiceType
	Detail      string `json:"detail" form:"detail" comment:"detail" validate:""`            //detail
}

type ContentListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                      //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` //每页条数
}

func (param *ContentListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ContentListItemOutput struct {
	ID          int64  `json:"id" form:"id"`                     //id
	ContentName string `json:"content_name" form:"content_name"` //合约名称
	Tap         string `json:"tap" form:"tap"`                   //tap
	ServiceName string `json:"service_name" form:"service_name"` //ServiceName
	ServiceType string `json:"service_type" form:"service_type"` //ServiceType
	Detail      string `json:"detail" form:"detail"`             //detail
}

type ContentListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" validate:""` //总数
	List  []ContentListItemOutput `json:"list" form:"list" comment:"列表" validate:""`   //列表
}

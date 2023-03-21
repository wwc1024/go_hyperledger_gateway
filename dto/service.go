package dto

import (
	"go_gateway/public"

	"github.com/gin-gonic/gin"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                      //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` //每页条数
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"通道ID" example:"1" validate:"required"` //服务ID
}

func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceListItemOutput struct {
	ID            int64  `json:"id" form:"id"`                         //id
	ContentName   string `json:"content_name" form:"content_name"`     //合约名称
	Tag           string `json:"tag" form:"tag"`                       //版本号
	ServiceType   string `json:"service_type" form:"service_type"`     //服务类型
	ServiceName   string `json:"service_name" form:"service_name"`     //服务名称
	ServiceDetail string `json:"service_detail" form:"service_detail"` //服务简介
	Port          int64  `json:"port" form:"port"`                     //端口
}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" validate:""` //总数
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表" validate:""`   //列表
}

type ServiceStatOutput struct {
	AllService  []int64 `json:"allService" form:"allService" comment:"总服务数"  validate:""`    //列表
	JoinService []int64 `json:"joinService" form:"joinService" comment:"已参与服务"  validate:""` //列表
}

type TemperatureLineData struct {
	DhtId       string   `json:"dhtId" form:"dhtId"`
	Now         []string `json:"time" form:"time"`
	NowTempList []string `json:"nowTempList" form:"nowTempList" comment:"当前温度"  validate:""` //列表
}

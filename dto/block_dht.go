package dto

import (
	"go_gateway/public"

	"github.com/gin-gonic/gin"
)

//近15条数据
type DhtInput struct {
	Info     string `json:"info" form:"info" comment:"查找信息" validate:""`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` //每页条数
}

type DhtInput2 struct {
	DhtId string `json:"dhtId" form:"dhtId" comment:"查找信息" validate:""`
}

//解析后结构体
type DhtOutput struct {
	ID          string `json:"id" form:"id"`
	DhtId       string `json:"dhtId" form:"dhtId"`
	Time        string `json:"time" form:"time"`
	Temperature string `json:"temperature" form:"temperature"`
	Humitidy    string `json:"humitidy" form:"humitidy"`
}

func (param *DhtOutput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

func (param *DhtInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

func (param *DhtInput2) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type DhtListOutput struct {
	DhtId string      `json:"dhtId" form:"dhtId" comment:"dhtId" validate:""`
	Total int64       `json:"total" form:"total" comment:"总数" validate:""` //总数
	List  []DhtOutput `json:"list" form:"list" comment:"列表" validate:""`   //列表
}

func PageList(c *gin.Context, param *DhtInput, outList []DhtOutput) (outList1 []DhtOutput) {
	outList1 = outList
	if param.Info == "dht1" {
		outList1 = []DhtOutput{}
		for _, listItem := range outList {
			if listItem.DhtId == "dht1" {
				outList1 = append(outList1, listItem)
			}
		}
	}
	if param.Info == "dht2" {
		outList1 = []DhtOutput{}
		for _, listItem := range outList {
			if listItem.DhtId == "dht2" {
				outList1 = append(outList1, listItem)
			}
		}
	}
	return outList1
}

package dto

import (
	"time"

	"go_gateway/public"

	"github.com/gin-gonic/gin"
)

type PeerListInput struct {
	Info     string `json:"info" form:"info" comment:"查找信息" validate:""`
	PageSize int    `json:"page_size" form:"page_size" comment:"页数" validate:"required,min=1,max=999"`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页码" validate:"required,min=1,max=999"`
}

type PeerListOutput struct {
	List  []PeerListItemOutput `json:"list" form:"list" comment:"节点列表"`
	Total int64                `json:"total" form:"total" comment:"节点总数"`
}

type PeerListItemOutput struct {
	Index      int       `json:"index"`
	ID         string    `json:"id"`
	Names      string    `json:"names"`
	Image      string    `json:"image"`
	ImageID    string    `json:"imageid"`
	Command    string    `json:"command"`
	Created    time.Time `json:"created"`
	Ports      string    `json:"ports"`
	State      string    `json:"state"`
	Status     string    `json:"status"`
	HostConfig string    `json:"hostConfig"`
}

func (params *PeerListInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type PeerDetailInput struct {
	ID int64 `json:"id" form:"id" comment:"节点ID" validate:"required"`
}

func (params *PeerDetailInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

// type StatisticsOutput struct {
// 	Today     []int64 `json:"today" form:"today" comment:"今日统计" validate:"required"`
// 	Yesterday []int64 `json:"yesterday" form:"yesterday" comment:"昨日统计" validate:"required"`
// }

type PeerAdd struct {
	AppID    string `json:"app_id" form:"app_id" comment:"租户id" validate:"required"`
	Name     string `json:"name" form:"name" comment:"租户名称" validate:"required"`
	Secret   string `json:"secret" form:"secret" comment:"密钥" validate:""`
	WhiteIPS string `json:"white_ips" form:"white_ips" comment:"ip白名单，支持前缀匹配"`
	Qpd      int64  `json:"qpd" form:"qpd" comment:"日请求量限制" validate:""`
	Qps      int64  `json:"qps" form:"qps" comment:"每秒请求量限制" validate:""`
}

func (params *PeerAdd) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type PeerUpdate struct {
	ID       int64  `json:"id" form:"id" gorm:"column:id" comment:"主键ID" validate:"required"`
	AppID    string `json:"app_id" form:"app_id" gorm:"column:app_id" comment:"租户id" validate:""`
	Name     string `json:"name" form:"name" gorm:"column:name" comment:"租户名称" validate:"required"`
	Secret   string `json:"secret" form:"secret" gorm:"column:secret" comment:"密钥" validate:"required"`
	WhiteIPS string `json:"white_ips" form:"white_ips" gorm:"column:white_ips" comment:"ip白名单，支持前缀匹配		"`
	Qpd      int64  `json:"qpd" form:"qpd" gorm:"column:qpd" comment:"日请求量限制"`
	Qps      int64  `json:"qps" form:"qps" gorm:"column:qps" comment:"每秒请求量限制"`
}

func (params *PeerUpdate) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

// type PeerListItemOutput struct {
// 	ID        int64     `json:"id" gorm:"primary_key"`
// 	Name      string    `json:"name" gorm:"column:name" description:"节点名称	"`
// 	CraeteAt  time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间	"`
// 	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
// 	Org       string    `json:"org" gorm:"column:org" description:"所属组织"`
// 	IP        string    `json:"ip" gorm:"column:ip" description:"ip"`
// 	State     string    `json:"state" gorm:"column:state" description:"状态"`
// 	Port      int64     `json:"port" gorm:"column:port" description:"端口"`
// 	IsDelete  int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
// }

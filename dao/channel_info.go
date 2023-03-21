package dao

import (
	"fmt"

	"go_gateway/dto"
	"go_gateway/public"

	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ChannelInfo struct {
	ID        string `gorm:"primary_key" `
	Peer      string `gorm:"column:peer" form:"peer"`
	ChannelID string `gorm:"column:channel_id" form:"channel_id"` //通道ID
	Chaincode string `gorm:"column:chaincode" form:"chaincode"`
	Detail    string `gorm:"column:detail" form:"detail"`   //通道描述
	PeerNum   string `gorm:"column:peernum" form:"peernum"` //节点数
}

func (t *ChannelInfo) TableName() string {
	return "gateway_channellist"
}

func (t *ChannelInfo) PageList(c *gin.Context, tx *gorm.DB, param *dto.ChannelListInput) ([]ChannelInfo, int64, error) {
	total := int64(0)
	list := []ChannelInfo{}
	offset := (param.PageNo - 1) * param.PageSize
	//模糊查询
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName())
	if param.Info != "" {
		query = query.Where("channel_id like ? or detail like ?", "%"+param.Info+"%", "%"+param.Info+"%")
	}
	if err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}

func (t *ChannelInfo) ChannelCount(c *gin.Context, tx *gorm.DB) (string, error) {
	total := 0
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName())
	query.Count(&total)
	return fmt.Sprint(total), nil
}

func (t *ChannelInfo) Find(c *gin.Context, tx *gorm.DB, search *ChannelInfo) (*ChannelInfo, error) {
	out := &ChannelInfo{}
	//直接查询结构体
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *ChannelInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}

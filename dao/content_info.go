package dao

import (
	"fmt"

	"go_gateway/dto"
	"go_gateway/public"

	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ContentInfo struct {
	ID          int64  `json:"id" gorm:"primary_key" description:"自增主键"`
	ContentName string `json:"content_name" gorm:"column:content_name" form:"content_name"` //合约名
	Tap         string `json:"tap" gorm:"column:tap" form:"tap"`                            //tap
	ServiceType string `json:"service_type" gorm:"column:service_type" form:"service_type"` //service_type
	ServiceName string `json:"service_name" gorm:"column:service_name" form:"service_name"` //service_name
	Detail      string `json:"detail" gorm:"column:detail" form:"detail"`                   //detail
}

func (t *ContentInfo) TableName() string {
	return "gateway_contentlist"
}

func (t *ContentInfo) PageList(c *gin.Context, tx *gorm.DB, param *dto.ContentListInput) ([]ContentInfo, int64, error) {
	total := int64(0)
	list := []ContentInfo{}
	offset := (param.PageNo - 1) * param.PageSize
	//模糊查询
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName())
	if param.Info != "" {
		query = query.Where("content_name like ? or detail like ?", "%"+param.Info+"%", "%"+param.Info+"%")
	}
	if err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}

func (t *ContentInfo) ContentCount(c *gin.Context, tx *gorm.DB) (string, error) {
	total := 0
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName())
	query.Count(&total)
	return fmt.Sprint(total), nil
}

func (t *ContentInfo) Find(c *gin.Context, tx *gorm.DB, search *ContentInfo) (*ContentInfo, error) {
	out := &ContentInfo{}
	//直接查询结构体
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *ContentInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}

func (t *ContentInfo) GroupByLoadType(c *gin.Context, tx *gorm.DB) ([]dto.DashServiceStatItemOutput, error) {
	list := []dto.DashServiceStatItemOutput{}
	query := tx.SetCtx(public.GetGinTraceContext(c))
	if err := query.Table(t.TableName()).Select("service_type, count(*) as value").Group("service_type").Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

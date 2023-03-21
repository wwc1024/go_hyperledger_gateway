package dao

import (
	"fmt"

	"go_gateway/dto"
	"go_gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceInfo struct {
	ID          int64  `json:"id" gorm:"primary_key" description:"自增主键"`
	ContentName string `json:"content_name" form:"content_name"  gorm:"column:content_name"` //合约名称
	Tag         string `json:"tag" form:"tag" gorm:"column:tag"`                             //版本号
	ServiceType string `json:"service_type" gorm:"column:service_type"`                      //服务类型
	ServiceName string `json:"service_name" gorm:"column:service_name"`                      //服务名称
	Detail      string `json:"detail" form:"detail"`                                         //服务简介
	Port        int64  `json:"port" form:"port"`                                             //端口
	IsDelete    int    `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_servicelist"
}

func (t *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := int64(0)
	list := []ServiceInfo{}
	offset := (param.PageNo - 1) * param.PageSize
	//模糊查询
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName()).Where("is_delete=0")
	if param.Info != "" {
		query = query.Where("content_name like ? or detail like ?", "%"+param.Info+"%", "%"+param.Info+"%")
	}
	if err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Limit(param.PageSize).Offset(offset).Count(&total)
	return list, total, nil
}

func (t *ServiceInfo) ServiceCount(c *gin.Context, tx *gorm.DB) (string, error) {
	total := 0
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName()).Where("is_delete=0")
	query.Count(&total)
	return fmt.Sprint(total), nil
}

func (t *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	//直接查询结构体
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}

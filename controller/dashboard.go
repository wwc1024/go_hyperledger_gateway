package controller

import (
	"go_gateway/dao"
	"go_gateway/dto"
	"go_gateway/golang_common/lib"
	"go_gateway/middleware"
	"go_gateway/public"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	service := &DashboardController{}
	group.GET("/panel_group_data", service.PanelGroupData)
	group.GET("/serviceNum_stat", service.ServiceNumline)
	group.GET("/serviceline_stat", service.ServiceStat)
}

// PanelGroupData godoc
// @Summary 指标统计
// @Description 指标统计
// @Tags 首页大盘
// @ID /dashboard/panel_group_data
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.PanelGroupDataOutput} "success"
// @Router /dashboard/panel_group_data [get]
func (service *DashboardController) PanelGroupData(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	channelInfo := &dao.ChannelInfo{}
	_, channelNum, err := channelInfo.PageList(c, tx, &dto.ChannelListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	peer := &dao.PeerList{}
	_, peernum, err := peer.PeerList(c, tx, &dto.PeerListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	out := &dto.PanelGroupDataOutput{
		ChannelNum:    channelNum,
		JoinedChannel: 1,
		JoinService:   2,
		PeerNum:       peernum,
	}
	middleware.ResponseSuccess(c, out)
}

// ServiceStat godoc
// @Summary 扇形图
// @Description 扇形图
// @Tags 首页大盘
// @ID /dashboard/service_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.DashServiceStatOutput} "success"
// @Router /dashboard/service_stat [get]
func (service *DashboardController) ServiceStat(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	serviceInfo := &dao.ContentInfo{}
	list, err := serviceInfo.GroupByLoadType(c, tx)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	legend := []string{}
	for index, item := range list {
		name, ok := public.LoadTypeMap[item.ServiceType]
		if !ok {
			middleware.ResponseError(c, 2003, errors.New("load_type not found"))
			return
		}
		list[index].Name = name
		legend = append(legend, name)
	}

	out := &dto.DashServiceStatOutput{
		Legend: legend,
		Data:   list,
	}
	middleware.ResponseSuccess(c, out)
}

// FlowStat godoc
// @Summary 服务统计
// @Description 服务统计
// @Tags 首页大盘
// @ID /dashboard/serviceNum_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.ServiceStatOutput} "success"
// @Router /dashboard/serviceNum_stat [get]
func (service *DashboardController) ServiceNumline(c *gin.Context) {
	allServiceList := []int64{220, 182, 191, 134, 150, 120, 110}
	joinServiceList := []int64{200, 180, 190, 130, 150, 100, 100}
	middleware.ResponseSuccess(c, &dto.ServiceStatOutput{
		AllService:  allServiceList,
		JoinService: joinServiceList,
	})
	// counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	// if err != nil {
	// 	middleware.ResponseError(c, 2001, err)
	// 	return
	// }
	// todayList := []int64{}
	// currentTime := time.Now()
	// for i := 0; i <= currentTime.Hour(); i++ {
	// 	dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, lib.TimeLocation)
	// 	hourData, _ := counter.GetHourData(dateTime)
	// 	todayList = append(todayList, hourData)
	// }

	// yesterdayList := []int64{}
	// yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
	// for i := 0; i <= 23; i++ {
	// 	dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0, 0, lib.TimeLocation)
	// 	hourData, _ := counter.GetHourData(dateTime)
	// 	yesterdayList = append(yesterdayList, hourData)
	// }
	// middleware.ResponseSuccess(c, &dto.ServiceStatOutput{
	// 	Today:     todayList,
	// 	Yesterday: yesterdayList,
	// })
}

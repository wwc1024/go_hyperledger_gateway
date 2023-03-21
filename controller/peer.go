package controller

import (
	"context"
	"fmt"
	_ "go_gateway/dao"
	"go_gateway/dto"
	_ "go_gateway/golang_common/lib"
	"go_gateway/iot_device"
	"go_gateway/middleware"
	"time"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//PeerControllerRegister admin路由注册
func PeerRegister(router *gin.RouterGroup) {
	admin := PeerController{}
	router.GET("/app_list", admin.PeerList)
	// router.GET("/peer_detail", admin.PeerDetail)
	// router.GET("/peer_stat", admin.PeerStatistics)
	// router.GET("/peer_delete", admin.PeerDelete)
	router.POST("/app_start", admin.PeerStart)
	router.POST("/app_updateproto", admin.PeerUpdateproto)
}

type PeerController struct {
}

// PeerList godoc
// @Summary 节点列表
// @Description 节点列表
// @Tags 节点管理
// @ID /app/app_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query string true "每页多少条"
// @Param page_no query string true "页码"
// @Success 200 {object} middleware.Response{data=dto.PeerListOutput} "success"
// @Router /app/app_list [get]
func (admin *PeerController) PeerList(c *gin.Context) {
	params := &dto.PeerListInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	// info := &dao.PeerList{}
	// list, total, err := info.PeerList(c, lib.GORMDefaultPool, params)
	// if err != nil {
	// 	middleware.ResponseError(c, 2002, err)
	// 	return
	// }
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("docker函数client.NewClientWithOpts failed")
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		fmt.Println("docker函数cli.ContainerListfailed")
	}

	//格式化输出信息
	outList := []dto.PeerListItemOutput{}
	for i, listItem := range containers {
		var portsIP string
		var portsPrivatePort string
		var portsPublicPort string
		var portsType string
		var Ports []string
		for _, item := range listItem.Ports {
			portsIP = item.IP
			portsPrivatePort = strconv.Itoa(int(item.PrivatePort))
			portsPublicPort = strconv.Itoa(int(item.PublicPort))
			portsType = item.Type
			Ports = append(Ports, fmt.Sprintf("%v %v %v %v", portsIP, portsPrivatePort, portsPublicPort, portsType))
		}
		PortsJson, err := json.Marshal(Ports)
		if err != nil {
			fmt.Printf("json.Marshal ports failed")
		}
		PortStr := string(PortsJson)
		PortStr = strings.Replace(PortStr, "[", "", 1)
		PortStr = strings.Replace(PortStr, "]", "", 1)
		outItem := dto.PeerListItemOutput{
			Index:      i + 1,
			ID:         listItem.ID,
			Command:    listItem.Command,
			Names:      listItem.Names[0],
			Ports:      PortStr,
			State:      listItem.State,
			Status:     listItem.Status,
			Created:    time.Unix(listItem.Created, 0),
			HostConfig: listItem.HostConfig.NetworkMode,
		}
		outList = append(outList, outItem)
	}
	fmt.Println(outList)
	out := &dto.PeerListOutput{
		Total: int64(len(containers)),
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}

// // PeerDetail godoc
// // @Summary 节点详情
// // @Description 节点详情
// // @Tags 节点管理
// // @ID /peer/peer_detail
// // @Accept  json
// // @Produce  json
// // @Param id query string true "节点ID"
// // @Success 200 {object} middleware.Response{data=dao.Peer} "success"
// // @Router /peer/peer_detail [get]
// func (admin *PeerController) PeerDetail(c *gin.Context) {
// 	params := &dto.PeerDetailInput{}
// 	if err := params.GetValidParams(c); err != nil {
// 		middleware.ResponseError(c, 2001, err)
// 		return
// 	}
// 	search := &dao.PeerList{
// 		ID: params.ID,
// 	}
// 	detail, err := search.Find(c, lib.GORMDefaultPool, search)
// 	if err != nil {
// 		middleware.ResponseError(c, 2002, err)
// 		return
// 	}
// 	middleware.ResponseSuccess(c, detail)
// 	return
// }

// // peerDelete godoc
// // @Summary 节点删除
// // @Description 节点删除
// // @Tags 节点管理
// // @ID /peer/peer_delete
// // @Accept  json
// // @Produce  json
// // @Param id query string true "节点ID"
// // @Success 200 {object} middleware.Response{data=string} "success"
// // @Router /peer/peer_delete [get]
// func (admin *PeerController) PeerDelete(c *gin.Context) {
// 	params := &dto.PeerDetailInput{}
// 	if err := params.GetValidParams(c); err != nil {
// 		middleware.ResponseError(c, 2001, err)
// 		return
// 	}
// 	search := &dao.PeerList{
// 		ID: params.ID,
// 	}
// 	info, err := search.Find(c, lib.GORMDefaultPool, search)
// 	if err != nil {
// 		middleware.ResponseError(c, 2002, err)
// 		return
// 	}
// 	info.IsDelete = 1
// 	if err := info.Save(c, lib.GORMDefaultPool); err != nil {
// 		middleware.ResponseError(c, 2003, err)
// 		return
// 	}
// 	middleware.ResponseSuccess(c, "")
// 	return
// }

// peerAdd godoc
// @Summary 节点添加
// @Description 节点添加
// @Tags 节点管理
// @ID /app/app_start
// @Accept  json
// @Produce  json
// @Param body body dto.PeerStart true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /app/app_startt [post]
func (admin *PeerController) PeerStart(c *gin.Context) {
	iot_device.AppStart()
	middleware.ResponseSuccess(c, "peerstart")
	return
}

// peerAdd godoc
// @Summary 节点updateproto
// @Description 节点updateproto
// @Tags 节点管理
// @ID /app/app_updateproto
// @Accept  json
// @Produce  json
// @Param body body dto.PeerUpdate true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /app/app_updateproto [post]
func (admin *PeerController) PeerUpdateproto(c *gin.Context) {
	iot_device.AppUpdateproto()
	middleware.ResponseSuccess(c, "")
	return
}

// // peerUpdate godoc
// // @Summary 节点更新
// // @Description 节点更新
// // @Tags 节点管理
// // @ID /peer/peer_update
// // @Accept  json
// // @Produce  json
// // @Param body body dto.peerUpdateHttpInput true "body"
// // @Success 200 {object} middleware.Response{data=string} "success"
// // @Router /peer/peer_update [post]
// func (admin *PeerController) PeerUpdate(c *gin.Context) {
// 	params := &dto.peerUpdateHttpInput{}
// 	if err := params.GetValidParams(c); err != nil {
// 		middleware.ResponseError(c, 2001, err)
// 		return
// 	}
// 	search := &dao.peer{
// 		ID: params.ID,
// 	}
// 	info, err := search.Find(c, lib.GORMDefaultPool, search)
// 	if err != nil {
// 		middleware.ResponseError(c, 2002, err)
// 		return
// 	}
// 	if params.Secret == "" {
// 		params.Secret = public.MD5(params.peerID)
// 	}
// 	info.Name = params.Name
// 	info.Secret = params.Secret
// 	info.WhiteIPS = params.WhiteIPS
// 	info.Qps = params.Qps
// 	info.Qpd = params.Qpd
// 	if err := info.Save(c, lib.GORMDefaultPool); err != nil {
// 		middleware.ResponseError(c, 2003, err)
// 		return
// 	}
// 	middleware.ResponseSuccess(c, "")
// 	return
// }

// // peerStatistics godoc
// // @Summary 节点统计
// // @Description 节点统计
// // @Tags 节点管理
// // @ID /peer/peer_stat
// // @Accept  json
// // @Produce  json
// // @Param id query string true "节点ID"
// // @Success 200 {object} middleware.Response{data=dto.StatisticsOutput} "success"
// // @Router /peer/peer_stat [get]
// func (admin *PeerController) PeerStatistics(c *gin.Context) {
// 	params := &dto.peerDetailInput{}
// 	if err := params.GetValidParams(c); err != nil {
// 		middleware.ResponseError(c, 2001, err)
// 		return
// 	}

// 	search := &dao.peer{
// 		ID: params.ID,
// 	}
// 	detail, err := search.Find(c, lib.GORMDefaultPool, search)
// 	if err != nil {
// 		middleware.ResponseError(c, 2002, err)
// 		return
// 	}

// 	//今日流量全天小时级访问统计
// 	todayStat := []int64{}
// 	counter, err := public.FlowCounterHandler.GetCounter(public.FlowpeerPrefix + detail.peerID)
// 	if err != nil {
// 		middleware.ResponseError(c, 2002, err)
// 		c.Abort()
// 		return
// 	}
// 	currentTime := time.Now()
// 	for i := 0; i <= time.Now().In(lib.TimeLocation).Hour(); i++ {
// 		dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, lib.TimeLocation)
// 		hourData, _ := counter.GetHourData(dateTime)
// 		todayStat = append(todayStat, hourData)
// 	}

// 	//昨日流量全天小时级访问统计
// 	yesterdayStat := []int64{}
// 	yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
// 	for i := 0; i <= 23; i++ {
// 		dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0, 0, lib.TimeLocation)
// 		hourData, _ := counter.GetHourData(dateTime)
// 		yesterdayStat = append(yesterdayStat, hourData)
// 	}
// 	stat := dto.StatisticsOutput{
// 		Today:     todayStat,
// 		Yesterday: yesterdayStat,
// 	}
// 	middleware.ResponseSuccess(c, stat)
// 	return
// }

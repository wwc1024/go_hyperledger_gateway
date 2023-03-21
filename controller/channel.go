package controller

import (
	"encoding/json"
	"fmt"
	"go_gateway/dao"
	"go_gateway/dto"
	"go_gateway/golang_common/lib"
	"go_gateway/iot_device"
	"go_gateway/middleware"
	"go_gateway/public"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	baidumap "github.com/menduo/gobaidumap"
)

type ChannelController struct{}

func ChannelRegister(group *gin.RouterGroup) {
	channel := &ChannelController{}
	group.GET("/channel_list", channel.ChannelList)
	group.POST("/channel_create", channel.ChannelCreate)
	group.POST("/channel_join", channel.ChannelJoin)
	group.POST("/channel_installCC", channel.ChannelInstallCC)
	group.POST("/channel_initCC", channel.ChannelInitCC)
	//通道数
	group.GET("/channel_count", channel.ChannelCount)
	//合约list
	group.GET("/content_list", channel.ContentList)

	//合约数
	group.GET("/content/dhtdashboard/content_num", channel.ContentCount)
	group.GET("/content/dhtdashboard/service_num", channel.ServiceCount)
	group.GET("/content/dhtdashboard/baidu_map", channel.BaiduMap)
	group.GET("/content/dhtdashboard/panel_group_data2", channel.PanelGroupData2)
	group.GET("/content/dhtdashboard/dht_get", channel.DhtGetAll)
	group.GET("/content/dhtdashboard/dhtdata_get", channel.DhtGetAllData)
	group.GET("/content/cardata_get", channel.CarGetAllData)
	group.GET("/content/dhtdashboard/mapdata1_get", channel.MapDate1Get)
	group.GET("/content/dhtdashboard/mapdata2_get", channel.MapDate2Get)
	group.POST("/content/dhtdashboard/dhtsetting", channel.DhtSetting)

	// group.GET("/channel_stat", channel.ChannelStat)

}

/*---------------通道部分--------------*/

// ChannelList godoc
// @Summary 通道列表
// @Description 通道列表
// @Tags 通道管理
// @ID /channel/channel_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ChannelListOutput} "success"
// @Router /channel/channel_list [get]
func (channel *ChannelController) ChannelList(c *gin.Context) {
	params := &dto.ChannelListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//0先写后读
	outdbList := []dto.ChannelListItemInDb{}
	outdbList = append(outdbList, dto.ChannelListItemInDb{
		ID:        fmt.Sprint(0),
		Peer:      "-",
		ChannelID: "-",
		Chaincode: "-",
	})
	//1要写的数据
	channellist, err := iot_device.QueryChannel()
	if err != nil {
		fmt.Printf("controler/channel failed to iot_device.QueryChannel(): %s\n", err)
		middleware.ResponseError(c, 2000, err)
		return
	}
	for i, channeldata := range channellist {
		outdbList = append(outdbList, dto.ChannelListItemInDb{
			ID:        fmt.Sprint(i + 1),
			Peer:      channeldata.Target,
			ChannelID: channeldata.Channel_id,
			Chaincode: channeldata.Package_id,
		})
	}
	//遍历数据传到db
	for _, outdb := range outdbList {
		channelInfoWrite := &dao.ChannelInfo{
			ID:        outdb.ID,
			Peer:      outdb.Peer,
			ChannelID: outdb.ChannelID,
			Chaincode: outdb.Chaincode,
			Detail:    "通道细节",
			PeerNum:   "通道节点数",
		}
		if err := channelInfoWrite.Save(c, tx); err != nil {
			middleware.ResponseError(c, 2003, err)
			return
		}
	}

	//从db中分页读取基本信息
	channelInfoRead := &dao.ChannelInfo{}
	list, total, err := channelInfoRead.PageList(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	//格式化输出信息
	outList := []dto.ChannelListItemOutput{}
	for _, listItem := range list {
		outItem := dto.ChannelListItemOutput{
			ID:        listItem.ID,
			Peer:      listItem.Peer,
			ChannelID: listItem.ChannelID,
			Chaincode: listItem.Chaincode,
			Detail:    listItem.Detail,
			PeerNum:   listItem.PeerNum,
		}
		outList = append(outList, outItem)
	}
	out := &dto.ChannelListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}

// channelCount godoc
// @Summary 通道数
// @Description 通道数
// @Tags 通道管理
// @ID /channel/channel_count
// @Accept  json
// @Produce  json
// @Param id query string true "通道ID"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /channel/channel_count [get]
func (channel *ChannelController) ChannelCount(c *gin.Context) {
	channellist, err := iot_device.QueryChannel()
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	middleware.ResponseSuccess(c, len(channellist))
}

// ChannelCreate godoc
// @Summary 通道创建
// @Description 通道创建
// @Tags 通道管理
// @ID /channel/channel_create
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{string} "success"
// @Router /channel/channel_create [post]
func (channel *ChannelController) ChannelCreate(c *gin.Context) {
	//当前仅仅创建mychanneltwo
	var msg string
	if iot_device.Flag == 0 {
		// Peer save channel 'mychanneltwo'
		iot_device.CreateChannel()
		//加入数据库
		tx, err := lib.GetGormPool("default")
		if err != nil {
			middleware.ResponseError(c, 2001, err)
			return
		}
		//虚假的info！----------
		info := &dao.ContentInfo{
			ID:          2,
			ContentName: "应安装fabdht2",
			Tap:         "1",
			ServiceType: "交易类",
			ServiceName: "车辆交易",
			Detail:      "fabdht2实现车辆交易",
		}
		if err := info.Save(c, tx); err != nil {
			middleware.ResponseError(c, 2003, err)
			return
		}
		msg = "创建通道succ + 写入数据ku"
	} else {
		msg = "通道已存在"
	}
	middleware.ResponseSuccess(c, msg)
}

// ChannelJoin godoc
// @Summary 通道加入
// @Description 通道加入
// @Tags 通道管理
// @ID /channel/channel_join
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{string} "success"
// @Router /channel/channel_join [post]
func (channel *ChannelController) ChannelJoin(c *gin.Context) {
	//加入mychanneltwo
	rc := iot_device.FabsdkConfFileRC()
	iot_device.JoinChannel(rc)
	msg := "这个节点已加入通道mychanneltwo"
	middleware.ResponseSuccess(c, msg)
}

// ChannelInstallCC godoc
// @Summary 通道安装连码
// @Router /channel/channel_installCC [post]
func (channel *ChannelController) ChannelInstallCC(c *gin.Context) {
	err := iot_device.InstallCC()
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	middleware.ResponseSuccess(c, "installCC in mychanneltwo succ")
}

// ChannelInstallCC godoc
// @Summary 通道安装连码
// @Router /channel/channel_initCC [post]
func (channel *ChannelController) ChannelInitCC(c *gin.Context) {
	iot_device.ChannelInitCC()
	middleware.ResponseSuccess(c, "installCC in mychanneltwo succ")
}

/*---------------合约部分--------------*/

// ContentList godoc
// @Summary 合约列表
// @Description 合约列表
// @Tags 通道管理
// @ID /channel/content_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ContentListOutput} "success"
// @Router /channel/content_list [get]
func (channel *ChannelController) ContentList(c *gin.Context) {
	params := &dto.ContentListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//从db中分页读取基本信息
	ContentInfo := &dao.ContentInfo{}
	list, total, err := ContentInfo.PageList(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//格式化输出信息
	outList := []dto.ContentListItemOutput{}
	for _, listItem := range list {
		outItem := dto.ContentListItemOutput{
			ID:          listItem.ID,
			ContentName: listItem.ContentName,
			Tap:         listItem.Tap,
			ServiceName: listItem.ServiceName,
			ServiceType: listItem.ServiceType,
			Detail:      listItem.Detail,
		}
		outList = append(outList, outItem)
	}
	out := &dto.ContentListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}

/*---------------图表部分--------------*/

// ContentCount godoc
// @Summary 合约数
// @Description 合约数
// @Tags 图表管理
// @ID /content/dhtdashboard/content_num
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{string} "success"
// @Router /content/dhtdashboard/content_num [get]
func (channel *ChannelController) ContentCount(c *gin.Context) {
	middleware.ResponseSuccess(c, "1")
}

// ServiceCount godoc
// @Summary 服务数
// @Description 服务数
// @Tags 图表管理
// @ID /content/dhtdashboard/service_num
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{string} "success"
// @Router /content/dhtdashboard/service_num [get]
func (channel *ChannelController) ServiceCount(c *gin.Context) {
	middleware.ResponseSuccess(c, "2")
}

// BaiduMap godoc
// @Summary 地图
// @Description 地图
// @Tags 图表管理
// @ID /content/dhtdashboard/baidu_map
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{string} "success"
// @Router /content/dhtdashboard/baidu_map [get]
func (channel *ChannelController) BaiduMap(c *gin.Context) {
	middleware.ResponseSuccess(c, "192.168.1.103")
}

/*---------------区块链部分--------------*/

// var channelclient = iot_device.FabsdkConfFile()

// DhtGetALL godoc
// @Summary 取得dhtid输出温度图片形式
// @Description 取得dhtid输出温度图片形式
// @Tags dht接口
// @ID /channel/content/dhtdashboard/dht_get
// @Accept  json
// @Produce  json
// @Param info query string false "关键词dht1"
// @Success 200 {object} middleware.Response{data=dto.TemperatureLineData} "success"
// @Router /channel/content/dhtdashboard/dht_get [get]
func (ser *ChannelController) DhtGetAll(c *gin.Context) {
	params := &dto.DhtInput2{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	list := new([]dto.DhtOutput)
	respayload, err := iot_device.QueryAllDhtsContent()
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	_ = json.Unmarshal(respayload, list)

	// nowid = []string{}
	nowlist = []string{}
	nowTempList = []string{}
	// // 	// 格式化输出信息
	outList2 := dto.TemperatureLineData{}
	for _, listItem := range *list {
		// if listItem.DhtId == params.DhtId {
		if listItem.DhtId == "dht2" {
			nowlist = append(nowlist, listItem.Time)
			nowTempList = append(nowTempList, listItem.Temperature)
		}
	}
	fmt.Printf("------------%v/n", nowTempList)
	// //例子
	// if public.DhtId == "dht2" {
	// 	for i := 0; i < 10; i++ {
	// 		nowlist = append(nowlist, fmt.Sprint(time.Now().Format("2006-01-02 15:04:05")+fmt.Sprint(i)))
	// 		nowTempList = append(nowTempList, fmt.Sprint(22+int64(i)))
	// 	}
	// }
	outList2 = dto.TemperatureLineData{
		DhtId:       fmt.Sprint(public.DhtId),
		Now:         nowlist,
		NowTempList: nowTempList,
	}

	middleware.ResponseSuccess(c, outList2)
}

// DhtGetALLdata godoc
// @Summary 取得所有dhtid输出温度数据形式
// @Description 取得所有dhtid输出温度数据形式
// @Tags dht接口
// @ID /channel/content/dhtdashboard/dhtdata_get
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int false "每页个数"
// @Param page_no query int falses "当前页数"
// @Success 200 {object} middleware.Response{data=dto.DhtOutput} "success"
// @Router /channel/content/dhtdashboard/dhtdata_get [get]
func (ser *ChannelController) DhtGetAllData(c *gin.Context) {
	params := &dto.DhtInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	list := new([]dto.DhtOutput)
	respayload, err := iot_device.QueryAllDhtsContent()
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	_ = json.Unmarshal(respayload, list)
	// fmt.Printf("0-------------------1 %v\n", list)

	//搜索
	outList1 := dto.PageList(c, params, *list)
	out := &dto.DhtListOutput{
		Total: int64(len(*list)),
		List:  outList1,
	}
	middleware.ResponseSuccess(c, out)
}

// PanelGroupData godoc
// @Summary dht传感器统计
// @Description 指标统计
// @Tags 图表管理
// @ID /channel/content/dhtdashboard/panel_group_data2
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.PanelGroupDataOutput2} "success"
// @Router /channel/content/dhtdashboard/panel_group_data2 [get]
func (service *ChannelController) PanelGroupData2(c *gin.Context) {
	out := &dto.PanelGroupDataOutput2{
		DhtNum:    2,
		NowdhtID:  public.DhtIdInt,
		NowSample: public.NowSample,
	}
	middleware.ResponseSuccess(c, out)
}

// Dhtsetting godoc
// @Summary Dhtsetting
// @Description Dhtsetting
// @Tags dht接口
// @ID /channel/content/dhtdashboard/dhtsetting
// @Accept  json
// @Produce  json
// @Param body body dto.DhtSetting true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /channel/content/dhtdashboard/dhtsetting [post]
func (channel *ChannelController) DhtSetting(c *gin.Context) {
	params := &dto.DhtSetting{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	public.DhtId = params.DhtID
	public.NowSample = params.NowSample
	middleware.ResponseSuccess(c, "")
}

func (channel *ChannelController) MapDate1Get(c *gin.Context) {
	out := &dto.DhtMap{
		DhtID:   "dht1",
		Address: "testAddress",
	}
	fmt.Println(out)
	middleware.ResponseSuccess(c, out)
}

func (channel *ChannelController) MapDate2Get(c *gin.Context) {
	//获取ip
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
	}
	defer resp.Body.Close()
	ip, _ := ioutil.ReadAll(resp.Body)
	ipAddress := string(ip)
	fmt.Println(ipAddress)
	//创建Map客户端
	bc := baidumap.NewBaiduMapClient("bs6e6vp4hKyL6mONieVt4hjjVsP7vPrP")
	IPToAddress, err := bc.GetAddressViaIP(ipAddress)

	out := &dto.DhtMap{
		DhtID:   fmt.Sprint(public.DhtId),
		Address: IPToAddress.Address,
	}
	fmt.Println(out)

	middleware.ResponseSuccess(c, out)
}

// GetALLCardata godoc
// @Summary 取得所有carid输出温度数据形式
// @Description 取得所有carid输出温度数据形式
// @Tags car接口
// @ID /channel/content/cardata_get
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int false "每页个数"
// @Param page_no query int falses "当前页数"
// @Success 200 {object} middleware.Response{data=dto.CarOutput} "success"
// @Router /channel/content/cardata_get [get]
func (ser *ChannelController) CarGetAllData(c *gin.Context) {
	list := make([]dto.CarOutput, 0)
	respayload, err := iot_device.QueryAllCarsContent()
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	// fmt.Printf("0-------------------1 %v\n", list)

	//搜索
	list = append(list, dto.CarOutput{
		ID:  "1",
		Msg: string(respayload),
	})
	out := &dto.CarListOutput{
		Total: int64(len(list)),
		List:  list,
	}
	middleware.ResponseSuccess(c, out)
}

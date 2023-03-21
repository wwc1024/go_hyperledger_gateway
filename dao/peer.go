package dao

import (
	"go_gateway/dto"
	"go_gateway/golang_common/lib"
	"go_gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"sync"
	"time"
)

type PeerList struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"column:name" description:"\b节点名称	"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间	"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	Org       string    `json:"org" gorm:"column:org" description:"所属组织"`
	Ip        string    `json:"ip" gorm:"column:ip" description:"所属ip"`
	State     string    `json:"state" gorm:"column:state" description:"状态"`
	Port      int64     `json:"port" gorm:"column:port" description:"端口"`
	IsDelete  int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

func (t *PeerList) TableName() string {
	return "gateway_peerlist"
}

func (t *PeerList) Find(c *gin.Context, tx *gorm.DB, search *PeerList) (*PeerList, error) {
	model := &PeerList{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(model).Error
	return model, err
}

func (t *PeerList) Save(c *gin.Context, tx *gorm.DB) error {
	if err := tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *PeerList) PeerList(c *gin.Context, tx *gorm.DB, params *dto.PeerListInput) ([]PeerList, int64, error) {
	var list []PeerList
	var count int64
	pageNo := params.PageNo
	pageSize := params.PageSize

	//limit offset,pagesize
	offset := (pageNo - 1) * pageSize
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName()).Select("*")
	query = query.Where("is_delete=?", 0)
	if params.Info != "" {
		query = query.Where(" (name like ? id like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, 0, err
	}
	return list, count, nil
}

var PeerListManagerHandler *PeerListManager

func init() {
	PeerListManagerHandler = NewPeerListManager()
}

type PeerListManager struct {
	PeerListMap   map[int64]*PeerList
	PeerListSlice []*PeerList
	Locker        sync.RWMutex
	init          sync.Once
	err           error
}

func NewPeerListManager() *PeerListManager {
	return &PeerListManager{
		PeerListMap:   map[int64]*PeerList{},
		PeerListSlice: []*PeerList{},
		Locker:        sync.RWMutex{},
		init:          sync.Once{},
	}
}

func (s *PeerListManager) GetPeerList() []*PeerList {
	return s.PeerListSlice
}

func (s *PeerListManager) LoadOnce() error {
	s.init.Do(func() {
		PeerListInfo := &PeerList{}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		tx, err := lib.GetGormPool("default")
		if err != nil {
			s.err = err
			return
		}
		params := &dto.PeerListInput{PageNo: 1, PageSize: 99999}
		list, _, err := PeerListInfo.PeerList(c, tx, params)
		if err != nil {
			s.err = err
			return
		}
		s.Locker.Lock()
		defer s.Locker.Unlock()
		for _, listItem := range list {
			tmpItem := listItem
			s.PeerListMap[listItem.ID] = &tmpItem
			s.PeerListSlice = append(s.PeerListSlice, &tmpItem)
		}
	})
	return s.err
}

package dto

type CarListOutput struct {
	CarId string      `json:"carId" form:"carId" comment:"carId" validate:""`
	Total int64       `json:"total" form:"total" comment:"总数" validate:""` //总数
	List  []CarOutput `json:"list" form:"list" comment:"列表" validate:""`   //列表
}

type CarOutput struct {
	ID  string `json:"id" form:"id"`
	Msg string `json:"msg" form:"msg"`
}

package models

import "time"

type Users struct {
	Id         int64     `gorm:"AUTO_INCREMENT" json:"id"`
	Username   string    `gorm:"type:varchar(50);index" json:"username"`
	Password   string    `gorm:"type:varchar(50)" json:"password"`
	Nickname   string    `gorm:"type:varchar(50)" json:"nickname"`     //昵称
	Agentid    int64     `gorm:"index"  json:"agent"`                  //上级代理的ID，0为普通代理
	Role       int       `json:"role"`                                 //用户角色系统管理3、总代理2、代理1、会员0
	Currency   string    `gorm:"type:varchar(50)" json:"currency"`     //货币类型 “RMB"
	Balance    float64   `gorm:"type:numeric(18,2)" json:"balance"`    //财务额度
	Credit     float64   `gorm:"type:numeric(18,2)" json:"credit"`     //信用额度
	Nowbalance float64   `gorm:"type:numeric(18,2)" json:"nowbalance"` //当前信用额度
	Btccoin    float64   `gorm:"type:numeric(18,8)" json:"btccoin"`    //btc
	Ethcoin    float64   `gorm:"type:numeric(18,8)" json:"btccoin"`    //eth
	Ltbcoin    float64   `gorm:"type:numeric(18,8)" json:"btccoin"`    //ltb
	Online     bool      `gorm:"default(true)" json:"online"`          //是否在线
	Login      bool      `gorm:"default(true)" json:"Login"`           //账号状态
	Odds       string    `gorm:"type:varchar(50)" json:"odds"`         //赔率盘口
	Status     bool      `gorm:"default(true)" json:"status"`          //是否收单
	Created    time.Time `json:"created"`                              //添加时间
	Updated    time.Time `json:"updated"`                              //更新时间
}

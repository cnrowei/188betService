package models

import "time"

type Users struct {
	Id         int64     `gorm:"AUTO_INCREMENT" json:"id"`
	Username   string    `gorm:"size:50" json:"username"`
	Password   string    `gorm:"size:50" json:"password"`
	Nickname   string    `gorm:"size:50" json:"nickname"` //昵称
	Agentid    int64     `gorm:"index"  json:"agent"`     //上级代理的ID，0为普通代理
	Role       int       `json:"role"`                    //用户角色系统管理3、总代理2、代理1、会员0
	Currency   string    `json:"currency"`                //货币类型 “RMB"
	Balance    float64   `json:"balance"`                 //财务额度
	Credit     float64   `json:"credit"`                  //信用额度
	Nowbalance float64   `json:"nowbalance"`              //当前信用额度
	Btccoin    float64   `json:"btccoin"`                 //btc
	Ethcoin    float64   `json:"btccoin"`                 //eth
	Ltbcoin    float64   `json:"btccoin"`                 //ltb
	Online     bool      `json:"online"`                  //是否在线
	Login      bool      `json:"Login"`                   //账号状态
	Odds       string    `json:"odds"`                    //赔率盘口
	Status     bool      `json:"status"`                  //是否收单
	Created    time.Time `json:"created"`                 //添加时间
	Updated    time.Time `json:"updated"`                 //更新时间
}

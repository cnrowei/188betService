package models

//体育运动
type Sprots struct {
	Sid  int    `json:"sid"`
	Sn   string `json:"sn"`
	Sen  string `json:"sen"`
	Mt   int    `json:"mt"`
	Son  int    `json:"son"`
	Tc   int    `json:"tc"`
	Tmrc int    `json:"tmrc"`
	Ipc  int    `json:"ipc"`
	Ec   int    `json:"ec"`
	Psc  int    `json:"psc"`
	Orc  int    `json:"orc"`
	Eec  int    `json:"eec"`
	Puc  []struct {
	} `json:"puc"`
	Pc int `json:"pc"`
}

//赛事
type Games struct {
	Ces []interface{} `json:"ces"`
	Cid int           `json:"cid"`
	Cn  string        `json:"cn"`
	Cen string        `json:"cen"`
}

/*
测试
*/
type Ces struct {
	Eid  int      `json:"eid"`
	En   string   `json:"en"`
	Gt   string   `json:"gt"`
	Tid  string   `json:"tid"`
	Ior  bool     `json:"ior"`
	Ht   string   `json:"ht"`
	At   string   `json:"at"`
	Esd  string   `json:"esd"`
	Est  string   `json:"est"`
	Hs   struct{} `json:"hs"`
	As   struct{} `json:"as"`
	Il   bool     `json:"il"`
	Mts  int      `json:"mts"`
	Isbg bool     `json:"isbg"`
	Pt   int      `json:"pt"`
	Isb  bool     `json:"isb"`
}

//球队
type Teams struct {
}

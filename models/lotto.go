package models

import (
	"fmt"
	"time"
)

//彩票列表
type Counters struct {
	Id                          int64  `gorm:"AUTO_INCREMENT" json:"id"`
	Name                        string `gorm:"type:varchar(50)" json:"name"`      //名称
	Official                    string `gorm:"type:varchar(250)" json:"official"` //盘口浏览地址
	Status                      int    `json:"status"`                            //状态
	Seq                         int    `json:"seq"`
	Ballcount                   int    `json:"ballCount"`
	Resultopenintervalsecond    int    `json:"resultOpenIntervalSecond"`
	Resultwaitingintervalsecond int    `json:"resultWaitingIntervalSecond"`
}

//期数
type Draws struct {
	Id              int64     `gorm:"AUTO_INCREMENT" json:"id"`
	Counterid       int64     `gorm:"index" json:"counterId"`
	Drawno          int64     `gorm:"index" json:"drawNo"` //期数
	Drawtime        time.Time `json:"drawTime"`            //"0001-01-01T00:00:00+08:00"
	Drawstatus      int       `json:"drawStatus"`          //开奖的状态
	Starttime       time.Time `json:"startTime"`           //开奖的时间
	Endtime         time.Time `json:"endTime"`             //结束的时间
	Betclosedmmss   string    `gorm:"type:varchar(50)" json:"betClosedMMSS"`
	Isclosemanually bool      `json:"isCloseManually"`
	Voidreason      int       `json:"voidReason"`
	Resultballs     string    `gorm:"type:varchar(50)" json:"resultBalls"` //开奖的数字
}

//投注的列表
type Selections struct {
	Id        int64   `gorm:"AUTO_INCREMENT" json:"id"`
	Counterid int64   `gorm:"index" json:"counterId"`
	Name      string  `gorm:"type:varchar(50)" json:"name"`     //标识词
	Odds      float64 `gorm:"type:numeric(10,2)" json:"odds"`   //赔率
	Minbet    float64 `gorm:"type:numeric(10,2)" json:"minBet"` //最小押注
	Maxbet    float64 `gorm:"type:numeric(10,2)" json:"maxBet"` //最大押注
}

//下注
type Wagers struct {
	Id           int64     `gorm:"index" json:"wagerNo"`                   //编号
	Uid          int64     `gorm:"index" json:"Uid"`                       //类型
	Counterid    int64     `gorm:"index" json:"counterId"`                 //类型
	Drawno       int64     `gorm:"index" json:"drawNo"`                    //期数
	Selections   string    `gorm:"type:json" json:"selections"`            //单注的信息
	Stake        float64   `gorm:"type:numeric(18,4)" json:"stake"`        //投注的金额
	Estwinning   float64   `gorm:"type:numeric(18,1)" json:"estWinning"`   //可盈利
	Issystempick bool      `json:"isSystemPick"`                           //系统选取
	Bettype      string    `gorm:"type:varchar(50)" json:"betType"`        //投注类型 num
	Bettext      string    `gorm:"type:varchar(50)" json:"Bettext"`        //投注中文
	Selection    string    `gorm:"type:varchar(50)" json:"selection"`      //投注类型 num
	Returnamount float64   `gorm:"type:numeric(10,1)" json:"returnAmount"` //返还金额
	Status       int       `json:"status"`                                 //-1输,0未结算,1为赢 输赢状态
	Createtime   time.Time `json:"createTime"`                             //下单时间
}

//类型
type Bets struct {
	Id        int64   `gorm:"AUTO_INCREMENT" json:"Id"`
	Wagerno   int64   `json:"Wagerno"`
	Counterid int64   `gorm:"index" json:"counterId"`
	Drawno    string  `gorm:"varchar(50) json:"drawNo"`
	Bettype   string  `gorm:"type:varchar(50)" json:"betType"`
	Selection string  `gorm:"type:varchar(50)" json:"selection"`
	Odds      float64 `gorm:"type:numeric(18,4)" json:"odds"`
}

//下注列表
func GetWagers(counterid int64, drawno int64) ([]*Wagers, error) {
	//list := &[]Counters{}
	var lists []*Wagers
	err := db.Where("counterid=? and drawno=? and status=?", counterid, drawno, 0).Find(&lists).Error
	return lists, err
}

//下注列表
func GetWagersStatus() ([]*Wagers, error) {
	//list := &[]Counters{}
	var lists []*Wagers
	err := db.Where("status=?", -1).Find(&lists).Error
	return lists, err
}

//更新输赢
func EditWagers(wagerno int64, status int) (int64, error) {
	info := &Wagers{}
	fmt.Println("edit status:", status)
	edit := &Wagers{Status: status}

	if err := db.Model(&info).Where("wagerno = ?", wagerno).UpdateColumns(edit).Error; err == nil {
		return edit.Id, nil
	} else {
		return -1, err
	}
}

//更新输赢
func UpWagers(wager *Wagers, status int) (int64, error) {
	if err := db.Model(&wager).Update("status", status).Error; err == nil {
		return wager.Id, nil
	} else {
		return 0, err
	}

	// info := &Wagers{}
	// // fmt.Println("edit status:", status)
	// // edit := &Wagers{Status: status}

	// if err := db.Model(&info).Where("wagerno = ?", wagerno).UpdateColumns(&wager).Error; err == nil {
	// 	return info.Wagerno, nil
	// } else {
	// 	return -1, err
	// }
}

//彩票列表
func GetCounters(desc string) ([]*Counters, error) {
	//list := &[]Counters{}
	var lists []*Counters
	err := db.Order(desc).Find(&lists).Error
	return lists, err
}

//投注列表
func GetSelections(Counterid int64) ([]*Selections, error) {
	var lists []*Selections
	err := db.Where("Counterid=?", Counterid).Find(&lists).Error
	return lists, err
}

//添加单注赔率
func NewSelection(sel *Selections) (int64, error) {
	if db.NewRecord(sel) {
		err := db.Create(sel).Error
		return sel.Id, err
	} else {
		return -1, nil
	}
}

//查找单注是否存在
func FindSelection(name string, counterID int64) (int64, error) {
	sel := &Selections{}
	err := db.Where("name=? and counterid=?", name, counterID).Find(sel).Error
	if err == nil {
		return sel.Id, nil
	} else {
		return 0, err
	}
}

//编辑
func EditSelection(id int64, sel *Selections) (int64, error) {

	//db.Model(&user).Update("name", "hello")
	//add := &Selections{Code: code, Descr: descr, Lang: lang}
	//has, err := DB.ID(id).Cols("odds", "mixBet", "maxBet").Update(&sel)
	return 1, nil
}

//获取期数
func GetDraw(drawNo, counterID int64) (*Draws, error) {
	info := &Draws{}
	err := db.Where("drawNo=? and counterid=?", drawNo, counterID).Find(info).Error
	if err != nil {
		return nil, err
	} else {
		return info, nil
	}
}

func CheckDraw(drawNo, counterID int64) bool {
	info := &Draws{}
	return db.Where("drawNo=? and counterid=?", drawNo, counterID).Find(info).RecordNotFound()
}

//添加单注赔率
func NewDraw(draw *Draws) (int64, error) {
	if db.NewRecord(draw) {
		err := db.Create(draw).Error
		return draw.Id, err
	} else {
		return -1, nil
	}
}

//更新时间
func EditDraw(drawNo int64, counterId int64, draw *Draws) {
	info := &Draws{}
	if err := db.Where("drawNo=? and counterId=?", drawNo, counterId).Find(info).Error; err == nil {
		if info.Resultballs == "" {
			db.Model(&info).Where("counterId = ?", counterId).UpdateColumns(draw)
		}
	}
}

//获取期数
func GetDrawno(nowTime time.Time, counterId int64) (*Draws, error) {
	info := &Draws{}
	err := db.Where("starttime<=? and endtime>=? and counterid=?", nowTime, nowTime, counterId).Find(info).Error
	if err == nil {

		fmt.Println(info.Id)
		info.Id = info.Id - 1
		err = db.Find(info).Error
	}
	return info, err
}

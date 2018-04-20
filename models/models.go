package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@localhost/%v?sslmode=disable", "root", "ishgishg", "bet18888"))
	if err != nil {
		panic(err)
	} else {

		// 创建数据库
		db.AutoMigrate(&Users{})
		db.Debug()
		db.LogMode(true)
	}

	//initData()
}

func initData() {

	user := Users{Id: 1, Username: "admin", Password: "498619c053f00fd6", Agentid: 0, Role: 3, Currency: "RMB", Odds: "A", Status: true, Login: true, Created: time.Now(), Updated: time.Now()}
	if has := db.NewRecord(&user); has {
		db.Create(&user)
	}

	counters := make([]Counters, 6)
	counters[0] = Counters{Id: 310, Name: "ShangHai", Official: "http://fucai.eastday.com/LotteryNew/app_SSSL.aspx", Status: 4, Seq: 0, Ballcount: 3, Resultopenintervalsecond: 0, Resultwaitingintervalsecond: 0}
	counters[1] = Counters{Id: 320, Name: "Chongqing", Official: "http://www.cqcp.net/game/ssc/", Status: 4, Seq: 2, Ballcount: 5, Resultopenintervalsecond: 0, Resultwaitingintervalsecond: 0}
	counters[2] = Counters{Id: 330, Name: "JiangXi", Official: "http://data.shishicai.cn/jxssc/haoma/", Status: 4, Seq: 3, Ballcount: 5, Resultopenintervalsecond: 0, Resultwaitingintervalsecond: 0}
	counters[3] = Counters{Id: 340, Name: "TianJing", Official: "http://www.tjflcpw.com/report/SSC_WinMessage.aspx", Status: 4, Seq: 4, Ballcount: 5, Resultopenintervalsecond: 0, Resultwaitingintervalsecond: 0}
	counters[4] = Counters{Id: 350, Name: "ChinaSwl3D", Official: "http://www.zhcw.com/3d/kaijiangshuju/index.shtml?type=0", Status: 4, Seq: 1, Ballcount: 3, Resultopenintervalsecond: 0, Resultwaitingintervalsecond: 0}
	counters[5] = Counters{Id: 360, Name: "XinJiang", Official: "http://www.xjflcp.com/game/sscIndex", Status: 2, Seq: 5, Ballcount: 5, Resultopenintervalsecond: 0, Resultwaitingintervalsecond: 0}

	for i := 0; i <= 5; i++ {
		if has := db.NewRecord(&counters[i]); has {
			db.Create(&counters[i])
		}
	}

	fmt.Println("##################  Service Start! #####################")
}

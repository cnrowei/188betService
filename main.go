package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"./query"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	id, ts := query.Get_Chongqing()
	query.WagersChongqing()

	fmt.Println(id, ts)
	timeCountDown(id, ts)

	StartTimer(func() {
		query.NewlottoChongqing()
		fmt.Println("添加第二天的赛事。", time.Now())
	})

	router.Run(":8888")

}

func startTimer1(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//启动的时候执行一次，以后每天晚上12点执行，怎么实现
func StartTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Minute * 24)

			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			next = next.AddDate(0, 0, 1)
			durmi, _ := time.ParseDuration("2m")
			next = next.Add(durmi)
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//倒计时
func timeCountDown(drawsno int64, drawstime time.Time) {
	var minute, second int

	now := time.Now()
	ws1 := time.Date(now.Year(), now.Month(), now.Day(), 22, 00, 0, 0, time.Local)
	ws2 := time.Date(now.Year(), now.Month(), now.Day(), 24, 00, 0, 0, time.Local)
	ls1 := time.Date(now.Year(), now.Month(), now.Day(), 0, 00, 0, 0, time.Local)
	ls2 := time.Date(now.Year(), now.Month(), now.Day(), 2, 00, 0, 0, time.Local)
	sw1 := time.Date(now.Year(), now.Month(), now.Day(), 10, 00, 0, 0, time.Local)
	sw2 := time.Date(now.Year(), now.Month(), now.Day(), 22, 00, 0, 0, time.Local)

	formate := "04:05"
	var strss string
	//晚上1
	if (now.Unix() >= ws1.Unix() && now.Unix() <= ws2.Unix()) || (now.Unix() >= ls1.Unix() && now.Unix() <= ls2.Unix()) {
		strss = "5m"
	}

	//10点到2点5分钟一期
	if now.Unix() >= sw1.Unix() && now.Unix() <= sw2.Unix() {
		strss = "10m"
	}

	if now.Unix() >= ls2.Unix() && now.Unix() <= sw1.Unix() {
		strss = "480m"
	}

	fmt.Println("strss", strss)
	fmt.Println("drawstime", drawstime)

	ss, _ := time.ParseDuration("-1s")
	sss, _ := time.ParseDuration(strss)

	dt := drawstime.Add(sss)
	subm := dt.Sub(now)

	fmt.Println("DT", dt)
	fmt.Println("SUBM", subm)

	minute = 0
	second = int(subm.Seconds())

	//fmt.Println("//////////", second)
	//ets := ts.Add(durmi)
	//var tss time.Time
	ts := time.Date(0, 0, 0, 0, minute, second, 0, time.Local)

	//fmt.Println("距离下一次开奖时间倒计时", second)
	var keys string = "D.Time"
	ticker := time.NewTicker(time.Millisecond * 1000)
	go func() {
		for _ = range ticker.C {

			ts = ts.Add(ss)

			fmt.Println("["+keys+"]", ts.Format(formate))

			if ts.Minute() == 0 && ts.Second() == 0 {

				id, _ := query.Get_Chongqing()

				if id == drawsno {
					minute = 0
					second = 10
					keys = "10 Second Get Data"
				} else {

					nw := time.Now()
					//晚上
					if (nw.Unix() >= ws1.Unix() && nw.Unix() <= ws2.Unix()) || (nw.Unix() >= ls1.Unix() && nw.Unix() <= ls2.Unix()) {
						minute = 5
						second = 0
					} else if now.Unix() >= sw1.Unix() && now.Unix() <= sw2.Unix() { //10点到2点5分钟一期
						minute = 10
						second = 0
					} else {

						//if now.Unix() >= ls2.Unix() && now.Unix() <= sw1.Unix() {
						minute = 480
						second = 0
					}
					keys = "D.Time"
					query.WagersChongqing()
				}

				ts = time.Date(0, 0, 0, 0, minute, second, 0, time.Local)
				fmt.Println("重新开始，距离下一次开奖时间倒计时:", ts)
			}
			//query.Get_Chongqing()
		}
	}()
}

// //检测系统时间，判断时间差
// func Heyispig() {
// 	ticker := time.NewTicker(time.Millisecond * 1000)
// 	fmt.Println("--------", time.Millisecond*1000)
// 	go func() {
// 		for _ = range ticker.C {

// 			fmt.Println("系统整在检测时间,每10分钟执行一次....", time.Now())

// 			now := time.Now()
// 			s := now.Minute()
// 			ss := now.Second()

// 			str := strconv.Itoa(s)

// 			if strings.Contains(str, "0") || strings.Contains(str, "5") && ss == 00 {
// 				ticker.Stop()
// 				Ok10timer()
// 				fmt.Println("执行。。。", time.Now())
// 			}

// 			// if s == 00 || s == 10 || s == 20 || s == 30 || s == 40 || s == 50 && ss == 33 {
// 			// 	ticker.Stop()
// 			// 	Ok10timer()
// 			// 	fmt.Println("执行。。。", time.Now())
// 			// 	//ticker = time.NewTicker(time.Minute * 1)
// 			// }

// 		}
// 	}()
// }

// //重庆时时彩 程序对整时间后，每10分钟读取一次数据
// func Ok10timer() {
// 	ticker := time.NewTicker(time.Minute * 1)
// 	go func() {
// 		for _ = range ticker.C {

// 			query.Get_Chongqing()
// 			query.WagersChongqing()

// 			fmt.Println("重庆时时彩////读取数据...", time.Now())
// 		}
// 	}()
// }

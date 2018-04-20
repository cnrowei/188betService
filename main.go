package main

import (
	"188betService/query"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//query.GetCQCP()
	query.Get_Chongqing()
	//query.Get_Xinjiang()
	//query.NewlottoXinjiang()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	Heyispig()

	startTimer(func() {
		query.NewlottoChongqing()
		fmt.Println("计算下一个零点.自动添加第二天赛事", time.Now())
	})

	router.Run(":8886")
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
func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Minute * 24)
			//beego.Error("next11", next)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			next = next.AddDate(0, 0, 1)
			durmi, _ := time.ParseDuration("2m")
			next = next.Add(durmi)

			//beego.Error("next22", next.Sub(now))
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//检测系统时间，判断时间差
func Heyispig() {
	ticker := time.NewTicker(time.Millisecond * 1000)
	fmt.Println("--------", time.Millisecond*1000)
	go func() {
		for _ = range ticker.C {

			fmt.Println("系统整在检测时间,每10分钟执行一次....", time.Now())

			now := time.Now()
			s := now.Minute()
			ss := now.Second()

			str := strconv.Itoa(s)

			if strings.Contains(str, "0") || strings.Contains(str, "5") && ss == 00 {
				ticker.Stop()
				Ok10timer()
				fmt.Println("执行。。。", time.Now())
			}

			// if s == 00 || s == 10 || s == 20 || s == 30 || s == 40 || s == 50 && ss == 33 {
			// 	ticker.Stop()
			// 	Ok10timer()
			// 	fmt.Println("执行。。。", time.Now())
			// 	//ticker = time.NewTicker(time.Minute * 1)
			// }

		}
	}()
}

//重庆时时彩 程序对整时间后，每10分钟读取一次数据
func Ok10timer() {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {

			query.Get_Chongqing()
		}
	}()
}

package query

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"../models"
)

//重庆判断输赢
func WagersChongqing() {
	fmt.Println("重庆结算输赢!")
	//未结算的列表
	var lists []*models.Wagers
	var win bool = false

	lists, err := models.GetWagersStatus()
	if err != nil {
		log.Panicln("Get Wagers Status listd data error!")
	}
	//循环
	for _, v := range lists {

		var balls string

		if dinfo, err := models.GetDraw(v.Drawno, v.Counterid); err == nil {
			balls = dinfo.Resultballs
		} else {
			fmt.Println("Get Draws Data [Resultballs] error!")
			balls = ""
		}

		if balls == "" {
			fmt.Println("Get Draws Data [Resultballs] is null!!!!!")
			continue
		}

		//比赛结果
		sballs := Ball(balls)

		bettypes := strings.Split(v.Bettype, "_")
		blen := len(bettypes)
		//判断是否是定位
		if strings.Contains(v.Bettype, "Fixed") {

			if strings.Contains(v.Bettype, "BS") { //判断是否是大小
				//获取结果
				win = BS(bettypes[2], v.Selection, sballs)
				fmt.Println("BS", win)
			} else if strings.Contains(v.Bettype, "OE") { //判断是否是单双
				//获取结果
				win = OE(bettypes[2], v.Selection, sballs)
				fmt.Println("OE", win)
			} else {

				if blen == 2 {

					uballs := SelectionBall(v.Selection)
					//定位输赢结果
					win = Fixed(bettypes[1], sballs, uballs)
				}
			}

		} else if strings.Contains(v.Bettype, "Digit") { //三字属性
			//组三结果
			win = Digit(bettypes[1], v.Selection, sballs)
		} else if strings.Contains(v.Bettype, "sum") {

			//百十个 | 和值
			selint, err := strconv.Atoi(v.Selection)
			if err != nil {
				selint = -1
			}

			win = SUM(selint, sballs)

		} else {

			//百十个 | 跨度
			selint, err := strconv.Atoi(v.Selection)
			if err != nil {
				selint = -1
			}
			win = SPAN(selint, sballs)
		}

		var winstatus int
		if win {
			winstatus = 1
		} else {
			winstatus = 0
		}

		//更新输赢状态
		id, err := models.UpWagers(v, winstatus)
		if err != nil {
			fmt.Println("Edit Wagers [status] Error!")
		}

		fmt.Println("ID %s UPDATE WIN %s", win, id)
	}

}

//重庆判断输赢
func WagerChongqing(balls string, counterid int64, drawno int64) {

	if balls == "[]" {
		log.Println("balls is null")
		return
	}

	//比赛结果
	sballs := Ball(balls)

	//未结算的列表
	var lists []*models.Wagers
	var win bool = false
	var err error
	lists, err = models.GetWagers(counterid, drawno)

	if lists != nil && err == nil {

		for _, v := range lists {
			bettypes := strings.Split(v.Bettype, "_")
			blen := len(bettypes)
			//判断是否是定位
			if strings.Contains(v.Bettype, "Fixed") {

				if strings.Contains(v.Bettype, "BS") { //判断是否是大小
					//获取结果
					win = BS(bettypes[2], v.Selection, sballs)

				} else if strings.Contains(v.Bettype, "OE") { //判断是否是单双
					//获取结果
					win = OE(bettypes[2], v.Selection, sballs)

				} else {

					if blen == 2 {
						uballs := Ball(v.Selection)
						//定位输赢结果
						win = Fixed(bettypes[1], sballs, uballs)
					}
				}

			} else if strings.Contains(v.Bettype, "Digit") { //三字属性
				//组三结果
				win = Digit(bettypes[1], v.Selection, sballs)
			} else if strings.Contains(v.Bettype, "sum") {

				//百十个 | 和值
				selint, err := strconv.Atoi(v.Selection)
				if err != nil {
					selint = -1
				}

				win = SUM(selint, sballs)

			} else {

				//百十个 | 跨度
				selint, err := strconv.Atoi(v.Selection)
				if err != nil {
					selint = -1
				}
				win = SPAN(selint, sballs)
			}

			var winstatus int
			if win {
				winstatus = 1
			} else {
				winstatus = 0
			}

			//更新输赢状态
			id, err := models.EditWagers(v.Id, winstatus)
			if err != nil {
				log.Println("Edit Wagers [status] Error!")
			}

			fmt.Println("ID %d UPDATE WIN %v", id, win)
		}
	}
}

func ReplaceString(ostr string, rstr string, nstr string) string {
	reg := regexp.MustCompile(rstr)
	return reg.ReplaceAllString(ostr, nstr)
}

//比赛结果[]int化
func Ball(balls string) [6]int {

	arrs := strings.Split("-1,"+balls, ",")
	var sarr [6]int

	for i, v := range arrs {
		if x, err := strconv.Atoi(v); err == nil {
			sarr[i] = x
		} else {
			sarr[i] = -1
		}
	}

	fmt.Println("ARR BALLS IS ", arrs)
	return sarr
}

//88XXX转换-1,-1,-1,-1,8,8
func SelectionBall(str string) [6]int {
	var arr string = ""

	for i, v := range str {
		s := SprintfV(v)
		if s == "X" {
			if i >= len(str)-1 {
				arr = arr + "-1"
			} else {
				arr = arr + "-1,"
			}
		} else {

			if i >= len(str)-1 {
				arr = arr + s
			} else {
				arr = arr + s + ","
			}
		}
	}

	arrs := strings.Split(arr+",-1", ",")
	var sarr [6]int

	fmt.Println("SelectionBall is ", arrs)

	//彩票结果转换为反向数组，按个、十、百、千、万排序
	for i := 0; i < 6; i++ {
		k := 5 - i
		if x, err := strconv.Atoi(arrs[k]); err == nil {
			sarr[i] = x
		} else {
			sarr[i] = 0
		}
	}

	return sarr
}

func SprintfV(v interface{}) string {
	return fmt.Sprintf("%c", v)
}

//单双
func OE(fixed, selection string, balls [6]int) bool {
	var arr []int
	var str string

	for _, v := range fixed {
		s := SprintfV(v)

		k, err := strconv.Atoi(s)

		if err != nil {
			k = -1
		}
		arr = append(arr, k)
	}

	sum := 0

	for _, v := range arr {
		fmt.Println("balls[v]", balls[v])
		sum = sum + balls[v]
	}

	//获取个位数字
	gwnum := sum % 10
	if gwnum%2 != 0 { //单
		str = "Odd"
	} else if gwnum%2 == 0 { //双
		str = "Even"
	} else {
		str = "Err"
	}

	if str == selection {
		return true
	} else {
		return false
	}
}

//大小
func BS(fixed, selection string, balls [6]int) bool {

	fmt.Println("balls:", balls)
	var arr []int
	var str string

	for _, v := range fixed {
		s := SprintfV(v)

		k, err := strconv.Atoi(s)

		if err != nil {
			k = -1
		}
		arr = append(arr, k)
	}

	sum := 0
	for _, v := range arr {
		sum = sum + balls[v]
	}

	//获取个位数字
	gwnum := sum % 10
	if gwnum >= 5 { //大
		str = "Big"
	} else if gwnum <= 4 { //小
		str = "Small"
	} else {
		str = "Err"
	}

	if str == selection {
		return true
	} else {
		return false
	}
}

//判断定位
func Fixed(fixed string, balls [6]int, uballs [6]int) bool {
	fmt.Println("balls:", balls)
	fmt.Println("uballs:", uballs)

	var arr []int
	var win bool = false

	for _, v := range fixed {
		s := SprintfV(v)

		k, err := strconv.Atoi(s)

		if err != nil {
			k = -1
		}
		arr = append(arr, k)
	}

	fmt.Println("fixed:", arr)

	for _, v := range arr {
		fmt.Println("balls[v]:", balls[v])
		fmt.Println("uballs[v]:", uballs[v])
		if balls[v] == uballs[v] {
			win = true
		} else {
			win = false
		}
	}

	return win

}

//组几
func Digit(fixed, selection string, balls [6]int) bool {
	var arr []int
	var sel []int
	//var win []int

	var err error

	win := []int{0, 0, 0}

	for i, v := range fixed {
		s := SprintfV(v)

		arr[i], err = strconv.Atoi(s)
		if err != nil {
			arr[i] = -1
		}
	}

	for i, v := range selection {
		s := SprintfV(v)

		sel[i], err = strconv.Atoi(s)
		if err != nil {
			sel[i] = -1
		}
	}

	//查找下注中是否数组
	for i, v := range arr {
		for x := range sel {
			if balls[v] == x {
				win[i] = 1
			}
		}
	}

	wins := false

	if len(sel) == 1 {
		if win[0] == 1 {
			wins = true
		}
	}

	if len(sel) == 2 {
		if win[1] == 1 && win[0] == 1 {
			wins = true
		}

	}

	if len(sel) == 3 {
		if win[1] == 1 && win[0] == 1 && win[2] == 1 {
			wins = true
		}
	}

	return wins
}

//和值.个十百
func SUM(num int, balls [6]int) bool {
	sum := balls[1] + balls[2] + balls[3]
	win := false

	switch num {
	case 7:
		if sum >= 7 && sum <= 20 {
			win = true
		}
	case 8:
		if sum >= 8 && sum <= 19 {
			win = true
		}
	case 9:
		if sum >= 9 && sum <= 18 {
			win = true
		}
	case 10:
		if sum >= 10 && sum <= 17 {
			win = true
		}
	case 11:
		if sum >= 11 && sum <= 16 {
			win = true
		}
	case 12:
		if sum >= 12 && sum <= 15 {
			win = true
		}
	case 13:
		if sum >= 13 && sum <= 14 {
			win = true
		}
	default:
		if sum >= 0 && sum <= 6 {
			win = true
		} else if sum >= 21 && sum <= 27 {
			win = true
		} else {
			win = false
		}
	}
	return win
}

//跨度
func SPAN(num int, balls [6]int) bool {
	max := balls[1]
	min := balls[1]

	//计算个十百最大数
	for i := 1; i < 4; i++ {
		if balls[i] > max {
			max = balls[i]
		}
	}

	//计算个十百最小数
	for i := 1; i < 4; i++ {
		if balls[i] < min {
			min = balls[i]
		}
	}

	result := max - min
	if result == num {
		return true
	} else {
		return false
	}
}

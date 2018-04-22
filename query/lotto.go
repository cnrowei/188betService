package query

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"../models"
)

func Readfile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)

}

type Chongqing struct {
	Drawno      string //期号
	Resultballs string //中奖号码
}

func Get_Xinjiang() {
	strUrl := "http://www.xjflcp.com/game/sscIndex"
	request, err := http.NewRequest("GET", strUrl, nil)
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Set("Accept-Encoding", "")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,ja;q=0.88")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Length", "8756")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Host", "www.xjflcp.com")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")

	if err != nil {
		fmt.Println("Error response for strUrl=%s got error=%s\n", strUrl, err.Error())
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("get DefaultClient for strUrl=%s got error=%s\n", strUrl, err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get ReadAll for strUrl=%s got error=%s\n", strUrl, err.Error())
	}

	testStr := string(body)

	fmt.Println(testStr)

}

func NewlottoXinjiang() {
	fmt.Println("New Draws 120 is Xinjiang......")

	//时间格式
	//formate := "2006-01-02T15:04:05+08:00"
	//time.ParseInLocation("2006-01-02 15:04:05", "2017-07-07 09:00:00", time.Local)

	var starttime time.Time
	var endtime time.Time

	var t int64 //期数编号
	var strt string
	var edt time.Time

	now := time.Now()
	ts := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, time.Local)

	strt = ts.Format("20060102000")
	t, _ = strconv.ParseInt(strt, 10, 64)

	durmi, _ := time.ParseDuration("+10m")
	start_durmi, _ := time.ParseDuration("-12m")
	end_durmi, _ := time.ParseDuration("-2m")

	//day_durmi, _ := time.ParseDuration("+24h")

	for i := 1; i <= 96; i++ {

		edt = ts.Add(durmi)
		ts = edt

		starttime = ts.Add(start_durmi)
		endtime = ts.Add(end_durmi)
		t++

		// if i == 1 {
		// 	starttime = starttime.Add(day_durmi)
		// }

		//fmt.Println("新疆彩数据添加成功", t, edt, starttime, endtime)

		//添加到数据
		drawtime, _ := time.ParseInLocation("2006-01-02 15:04:05", "0001-01-01T00:00:00+08:00", time.Local)

		draws := &models.Draws{
			Counterid:       360,
			Drawno:          t,
			Drawtime:        drawtime,
			Drawstatus:      0,
			Starttime:       starttime,
			Endtime:         endtime,
			Betclosedmmss:   "00:00",
			Isclosemanually: false,
			Voidreason:      0,
			Resultballs:     "[]",
		}

		if has := models.CheckDraw(t, 360); has {
			falg, _ := models.NewDraw(draws)
			fmt.Println("---------------------------------------------------------------------------")
			fmt.Println("新疆时时彩数据添加成功", t, edt, starttime, endtime, falg)
		}

	}
}

//抓取数据 Chongqingwww.cqcp.net
func Get_Chongqing() {
LABEL1:
	x := 1
	strUrl := "http://www.cqcp.net/game/ssc/"

	post_data := `ScriptManager1=UpdatePanel1%7ClnkBtnFirst&__EVENTTARGET=lnkBtnFirst&__EVENTARGUMENT=&__VIEWSTATE=yk2t2rtNVT3gJkFtFrZHlvffgels6hiqkfahIJnv4CszKxNAKr%2FYwvxdIwcr4%2BpXN5PD0wq68wfksaty4bInckf4zEb%2F%2FoRkTXTMmIdT8quEOK%2BPgIWRY94JBr4yXbLVc5iDtxXLY1evb11avNE7MOBv4VkmzpLWktqCYyjaBuwi%2BSJitrmHnNAK420yM3EbZgg3%2F%2Bl9qxoLGAefT58HQ72tj7uSwQsvQ2uvASQKMpmpFuUfY6YEve1xtq32cX7zQllvdq2fw4WIB5JCnYb9QJFEITL60Dd3kXX4uR2Q9ekGoJ3wuKV5hmOShpvS1cGNg1ckYDz2mQvyBHt6PtW%2B4DlxRk0FGrIiEP8peke6im5V%2BFdgqrZFyFuilW2QGttXuYF0IiKpP8lQl%2BRnP819qH0ZSNSksQun9zjGGA46HaSagtb2UO62sJ7RO%2BC7GPzD4fIPJqcwm2ttSOXHBCxzZDWdWLjeRyhPEyA4S2EuQ3I%2FTLbok9z02bH3gaarNdIezTT3yhOG5QyMEm6oUIUaSU9eiB%2FO%2BysggsK5yFE%2FNylEWyf4FVOXkORI2qmIyBVQTRJSXeMCvf3vB1W%2BIqr1m5DRU%2FNBRuQl1f19yaDF74eOgXPK5Nya4gbt0uAKZAaH974djXia0FcrMm98iKgGlE4xolySSxLpPzDxW1NEi99SMxC%2BpzuDZI02Rft0Y9dnFGvZ0yNbUoh%2FwYS93GNceapEF3Viv7DZiM5JgBGoDEoJGAfUnka8WZMcQDXN23y8Db3afez7UV4jbf5mHzd1ycg%2FSu7pwqOXWaF%2FPSpWpHt8tBpZocHa4QmHO0QCVVAs7YBJYv443o2op2l2vSl0S9%2Ft6zLsFJhhfkBhO0XciLxI7ymU6u0xpHYV3TafDVBwOi3Y9iSOqYzbh0WII7uFN1Uou4UEPzkgOXylRR8WgPExjMsYl9s3Rx0X2hlTsiBSBbYrJ8PXV9DghXHRnkRJfHW7dRG0%2BBJ5ymq2GAQXqlWibDw0vpUmbgD%2B5PiXQPhZmsQRtVzQEPIzTrwkruHFCvCqblJdoKPA4zP1jqt67eutoM4k6DQB7ducPIsfzoaVhI79EmEFmylKeHotO0Ocai5buHQtxfR35OzRQ1w5t%2BHBJjiAYQAn8WG3sB4BpY5LwbOpk8JBm9LqXTeVKABROCkg%2B6ifvfxU6QuyPk4T3cuaJOPyUk36jwmLqyTYgY0ksW2I6IhElHmaUFL8lp7EXtiOMSuGRv2pdEUxPCvrxii25TWIp%2F0MiKPvwT5d78chB%2BrF%2BMQe11Sbpe1MDEt5VH3RrgraAfnsqU92lZOny6Vp%2FRwdoiK8M4xgsFP8G4NV3b67eC%2FnOTjACCaBTAIg8zdp%2B5hCNhdxtqdAJkkrhaNKpZPhYEZL5vADtVoCIARx0g7yBYpQpZ%2B3c7dbF6OlmWrEFTHZHWLN85uc2vWUaMqIoF0KMaOEiYzOjfRoBG%2FeKMMuZG3gsxbBO490tsYTgxr%2FVK%2BqtDmwpm1nw8HfGqO0yxwiOoOAXOxe2hEMojEN0WoDsG7IVmSqo%2FbAR98ChueC7%2FJuQmz5ZNfop4pBLoD%2BG%2BkEcd2aRJ%2F8J0tMXA8iwHMnP84CJ6cNaUBZV3fQJoZLyp1TtClnoWdS0P0a2lsyD9ls5UaRBO%2BIvIi9IoCE%2Fq9lH4zwqGmdC8mJzqLzJZzEja9YE%2FDDUQev8SE6vzr1NUO6srb576aJ483SR3HTgY71hSEllR1hGjTSeJkY2NL%2By3HAHAdtJ7k7Jm%2FPGJg2Co3ve8FW2XdxhzPwEQE6piBsrFC5MnaSPqiENtqH7yBHmbvZ0vJGXq%2BD%2FfYZc7TjIoZexYkndv5lg%2BJ2I2gsM9NcXuVaKFhYpgait3qerLhUWTKSJPlVEnMK3kMA7IgxUNvvSTkGKLzjrlXGPl6VUdkdPbDHCq6B0AgpsHABv3T6t%2FPXrJXKi%2Bd%2BOv%2BSDq6CFMAzupRK7ggsMtGvyqwDqv4S7T47u7JCsLw7p067TpGSogcVO3iqjO8l9OkNuvB%2BapUdnTvl0aCMSwbrk9suaDkZvg0KJvVWOpHqXa3wiWa3h%2B0TWEg9sfNf9Yd%2B1ADEE1H6o87HtHIyGjBhyXqrkcp1W46TytAFoZQE9g6fdPitoS92JjJ1QHQEjmS6qaH%2FK4fqT9U9Q8roNFJJrJVCz9qglPaZzbYUo3Zvf%2FCIQTw%2BPW57LGj3hVr2f5r5BlCyVcQqlxHPxnmxEM%2BQom8lrq0%2F1Ezd9dONX4ylKVfhzQy8fCGdzLfxdZYQrKDHqzHLbrZuprO6BbXvof9zrXscMhgvfWouP3OrC%2FrbqO87h1FC1h0VpJYlWlqnMT2lOLTVvog6tPBnSWVKlw49ousomoCKgvVygAV1sK53FMJg2PmDhqCBiOfPALxCP4tmPIoSmAAA0akVJVk745GDaDVXqsH%2FW3SbefM%2FebXaf4pVRmSwWuXH4VmppHRFBZRa5eFAAulCWSgaZKD6aGnGyk7dLqCDqnqEO5NCsc87uocDPhVoqet6%2BY2XnW3EWei9rfK8y1RF283SQY6UWkwmua34Av0qBYwPyC2JeHdY6O9jSbCoMVQao%2Fa%2BiupBaA1WAAELq7qt4Y6PSem9EqrWR4zP8Q95U3DYdEqBP5evXi5d%2BneJoCcVw7Vwj%2FhVJjsPfSNTVsbpJCB210acr7bYmJaZT%2FslK5dHB2ZdDDEEw%2FKeJwl7X1xsMxW4%2BVRGe0mSSOdC%2FGRv%2FqZY5%2BcMvyaC9eDosJs93aRio%2F5c6EweJ4UemD7SJVoJ4RHX26ZdNs%2BnPUHn52%2FpKT53uRZJqazRLr77C6WHd%2FprMurEBfdgd6erDB97nGTmptrXxhu3gnnURAad9Bds5G57Cat8JenDzX5iEjPyLVqE%2BVi9512mAUJ9BSbvG4RorJiDQ361cPnw87NjC%2FxQHPD75S8Nxx%2BO2zU9MQGU39iWfNlvLdLASlG3zhxXHT%2Fk9voidDEt%2FDfglog5of6uzAV28H4vr%2F5WmRqgFf41qSlUN434%2FIF8dcSsTKe1gcUXub3Z3Vq0Q8Jjck39lD7fFd2dLzlDuGKHGzE8d%2FbtuAXnV5gCsnxortBY3jtAcVd3gT4huJ2%2BiXoMjkCKHn%2BuI7EA%2BPcgDkl9Jbvg8zdBiBm9VZM8yvCYU9Z%2B%2FpeYKr7Akf%2FGPHjIe8rONxrFPg2tMtE%2BsXPGkyCkToevll0mQVDzAnmbm30MYW6zM5%2FwWnxxcSyVzk%2F25XFVmHmNTaydTfjJ336p595vkK7%2F7HdOvTj9DHWH44rpa1lBOiFHjanxwBmDbYFbiD9wLhUVwJZL0ppwtFyNdpXD4Huo7yxBB5CMuGbD9gmX1TX4rHxYlGpNmn8yKjyuAeA7OBk4OU1pI7M%2F1Z5Lzz4F83VDb7MWbipfThFf0swPc95PjU8FvU0ou68KA4jCQoiAHEV2pMupLegpSXwP7WkZ7uTF6WN3aPSIV%2F4326E%2Fb8B%2FdjYI1q5uXKLkcOY5O4AYnltcDSsQhpKDvf2rr5nhIe1OkLCAqmkhYw0NpDbx7asilHFYap7GJkXNXBg5pXAjb4Zva7%2BhAP3hz88u02EXiO0nNrzqlFpFO%2FXg1pak%2BX2Ia3LSIpTENQnznOJ00oNbLXQ36eUkQhIv5BF%2FX3%2BbgQVepTOyjnAG6Ld3UN5ihJd%2FnQKEmHgcMBYgr35jVXijw%2Fhy9LO4azqUOCgPaY%2FEg93vwso0OcIfrwBeuM2YV90zgdctQpR5m%2FkpcDkZ1g8pX%2Bj22qrhiOdN67ZXqSOMdCDKE4qcm%2BPIrdhjsqSWDm1QygKd9ku8eKKBNHeapZty0o56lu03ZUVoh3M9TWtulD3kjqNokkS6yaTb9qWyr36zUp3Ix0i1wLqaqG1Yo97qAptkWCJjtc9YAD3ozN2772ubFzQ1zzx7MybwfdV%2Bg3XPlAkBiqJDprGNUx0wG5nR7HeOQHrDgpP4ErDexMct0QGL7m6w3lbKVHe3zt5zIqvYjLB3KZmrwJ3Ya5EaDb%2FXYOXHSLYo7srKQTIDGFrUU%2F388EXYWkSg8h0%2B0%2F0md6mscrXvDjz4NNNcbwKQ6PR7yT6%2BI62h%2B1otFYd9pdshmqOghOh4nPJnPr9QG4gGmtzKYA9mPRDbnhyA%2BVqSt7aEkT461znOIGtGGgVwU3%2FIpZx3ONUeJC9jnDUrsgsfehwHmHGNTekVEwdBqwSCNqXILs%2F8%2FSKJsswM%2BSMEAGChhA3UCo7IHindWgnAQUDzZPNk7wHLAlo%2FhmAeieZfWCMsAzej8fPC8hzKoEqK0s3ZOEiv79bWp%2BDgwn0VqvF54yOtp2PfvlVNDWK%2BnTJsLiFvnbBoJgp2JMrrQ9MG7j9FfKzJRUjVh97G2CgXhN1higcdtPx1I%2FnIg035UhN3HMflCys8hHldkCR2Ksh6WRI2VcCvsO5Ua%2Bfje%2Fc49Wg1CloP9PvS86KVtoD7tP7VXCdoBAuMIn4bqhGs%2F7S4I0uebDoXxjtQGDFDzgRGB%2FUfgfnMe1hWPtmolonwpiuRZLKQtePdBS0dBLkbviPsuc%2F3ZHBlmo%2B%2BJj1Jrxsv1sMD2z5d36gV1KfcXtNX6d57fuXXI8F%2F%2FdkUZl8q8vNqIpGfxhdx%2B7%2BClcjuXTfxxll8urcPwK7FpUQLDtFvAqXSda3dEef9%2F0wjJKfjk8OaxWvfLQj5T9c5Bwd9CtWXfrU4tWBEkpbN8k4noXIHlM3bKKPp8AGxlloCH3qW76OTtVdlvLk6cs%2Bl1e8lk2pdwwUh6y2sWEQ25cuYO3YnwOejiKBagj%2FL33z8tShKr2vf2UyA%2B45A%2FeAMLIPrBlfoKQEw9J5js90HNOEquQCTMSvGfnHA2AhmjdCp78gCDoZCEFQd8mipwvFPCNHgY%2Fd0PAAzJfI1hXPuyrzlZU1x8KOao3nAYN%2F6W%2B41XgJ7KLi2iuloC1qkZcwXX8zdomRbMck7GQJcw7twjc6Y0MSUHIcqpmD6G9vycKqNu95cZDKIIE%2Fb%2BY92%2ButFjGmyROm%2FLMt6g99jzf9ZRW8WtBQJzT6KCY4O7qYZ%2FzNa1VbCpx8yzGHEcA4U0tnTc5fLcyhVvlpe3gVe8OqehIEbpoDr9F4U2qeCbTuxQudxnySohQvui8xEHAxVUoz7JQFM0ukw75XDjqv1ypXioEwQsbYy%2Fd%2FmdQXzx0XAxocHtYsE7FKsEmlvarwVRRmQ8Co8AEz%2F7ZVbz6g%2B1lNrH1tn%2FOmY%2FmIG5hfaG8aCgBXJx%2BifJdmnD3Pwv80Z73e%2F5qyBq2TdLZ02tUurWVVOvHIs%2FBW42Ni2i7JykA7cNAhdtC67xtUxc9IiCj3Atr98nxZk4l0WZ%2BUuhQ0jrob67jLfZcSTmw3nzefF9u%2BydH1mNRGmQk0679LCR%2BHOevM%2FoQdK67JAMsZoMkOnAXLGcNAOdZkeMUl%2BgOUmbPgch%2B%2FSP1OpxzbX%2FPdRcyxdqe7clG9wIB%2Fjhoy6qXioTL15JVKLkW%2BBhJDsupCq4Kd74CmpRlKhNROZJStaK006q6sEN3gdFJ%2BMD5hcDXrOnEaAQe%2BbFIvmWGAnEEOrIh41KHfZ%2BJW0y72k7w8j3oEJCHuvlitPVSYjKxoIl94We4iWLDqh7uh%2BdcVWVujLHaGqosaX9ikZcr5icOEigKcHWBQn2NLG1xDTTDG%2BOdgVobDq6cc8BOBFhXhJwKJY9kDjFCYBKt3iJpOpqvea9KNnP%2BO1idk2qRn9nWkPQu1gXvTG49LIGb8hxIk0JkG%2FdOdp0a6sRnqT7%2Bko1KY5t4x5DOGm%2B9Ol18ujs0ofUa4wKCpZAOtYIJDRaEKulqmL7RlxgIz71IqpIqK5x8hIFGYpc%2BVmzB28kk3V1BjEjmQa2E077QEo6NID8tTDaXDaX1QYDr3DFwh%2BJM6oMqLCzuacu3VOYfMQU%2Fq7IzPHRQbNrLp899aD1oBglSY42f75o2rsAsz%2BZzP%2ByY0EgXbpckCmQ5%2BqTw5q24JJr5wVdjGkXi%2F9HA4JHnl3sT2mU9UaP4tfjzaF4UJAbw4bMzlx%2F4ujwF2E8mGaZ1m4sxXWvZDIBlkXiWWOJjVONJSIKoK3S0XmR1weNWSW%2FKzodbAtbZ055WrogwcsgLnGJY5gK3i2W9ROYKO24HeRrnpVVADKo7fBdI8gkiXARZymaWAqUBQIbR8xpK4leuQ8HeNpTPdXYEdeCnN3M8x4mCYLz9Raos97e70L2b2cjPVaTRDrIT5pTvbtEuTEbWf%2BRk35euTFTaAHYoP2ti4bebsqarEPW6Y0UdLWcCtjePloDeRNDMOGbltP2E6GnL2DxPeOnaepYo%2FlCyfwKRaCCx7gSNk8A0k5Ez2V4CDot35d2%2BBuJYiT2paE5tB7zKVcnM9ssvtiMnJRa12MQRxz0Riy5VZfGiHsHmauRLeJ25nSP3CvCUwNN2bNP35YYy%2BzJbCW%2FdvnT07V5wHfMD2TdfqlsqjITGxp1MzEGl7nTbNcx6z09w6mbygtE4oDW7G7NPZJ%2BYP0ZAZKfnUjR6yKypQR%2B9S0j39k4dbQGw2drzMfu9I63QHTSckvN3r7agheT641IBUzZ2NJWDoqoDLTMtuNAPUpA%2Ff%2FSiPPQQYVNnXiV1bAl1SR1oOZpEFZe%2FMYxeIkw1b8EfO4TF2yPBBbWZXEqhlklQhCZVEgLcXQ12ahn%2BI8Sx40SrrDYP4F4KqfG1wbooEN4dNcr0MpztbA3BKnQQcH42QyAcJVu1rPbERVk64zmqbZA3EK3nFebKJHLvm%2FuVRYNuRbpldGDl6Hvm3WqScEPJxdzipsrrrGBXgUXpjW9Pzx9aR%2FbWqZN2ieNVDiraY%2BwKqBNOoQS6ajqZw25IlEPM1hImpE8ZQmYBeWZFsy4rHdgCiJXX6AZ2pZWqFPIENQURR%2BpUgTtaxO3EkgnLXATsYMVl4eSPw05SYqv5o2QxMAfS1QmpC%2B51%2BfpinU%2BTgV%2FkrVZOX6iyOQmVK0VTQXjSuoD8m96H00AdXDZMdy1k1sBQfQac1XwTI%2FibVYiAF5AXFFHmiPgRWrTYsUZdkY%2BbsRnDHdMYOL8lJPyUbIp1sl0DIgT1A3GZdiR268CNpQqsupUdbCV1N2Diq%2FJEGLCJJQ8d4LL2zwBD6mT%2FHjsVMas3HRTapBH7wwN61mTyNd762B0GfZjpl1pPe%2BV2Xq3I7zBQCN6CAFETFJbOiOiEia2MGK%2FOl%2F9AmcLfzLu6x9aNN%2Bb3QBBwp0atjHpT6O1vSt%2Bv1rAB7UudfbBF%2Fm335T2%2FD5iRBf19Gb%2FiBI0CnVV%2Fe2I9nY%2BxEIheRzz2NCz0xTptAuhUvEUlil%2B3bvn0kNIQUJi9b9ga23BtxAbxfKm8uHlHTQMeMynaAXjvVkhrV10ME5c4RFLYOAv0SXQ09yg2nNu%2B4fZNDaIs2hZhjLSAkUflYa%2FTVsiiEwjPca3NS8QGYTYYhbz5sELZ29m5tmvdrvzXc8275noDOmkRqh94FmRvLn3UcpPHhuG0C6%2BfaqxwbVwtZdcit%2F5AKGhfVbhg3R9W3BbrlL289LTwys3gMMa3%2B%2BUxtLcU2p9pYt31AzXOdCGpe5MklGpEeEWFHdMfxZZ6JlkhEZVRpGDqz12sVRZSj%2FYvNPNY5luhpcJ%2Bwb%2B00Ci4Q0bhofzNKIxAl%2F1zcvgbmcZwLVtjLXpzCxXnL2v0nnhR75DxzXgUj7WZQet%2F8N0adY72CpY946LMOrqe3enbf7WFmN%2BAcCrQ1TpxavvkThUalUg9epbED0oFazp7CepA5PdhYRzsOqVNuypO%2FpAC3cB3CujSsna%2BMLuwHZ9MRNxpRmevE35r%2FEjw%2FCN4fHM%2BuKM%2BvLw%2B%2FeMWaFpZL9pQup4bxoIduim0GqkJORub09JZEJONG%2FQ80hpq%2FUD3FfQZ2fmGdrD%2BE2stYM93Mk5%2BqGAGMjZwd6N8jU1I2zAxVXhD1uOsEEU%2FKfdXecuakk&__VIEWSTATEGENERATOR=7B9E3A74&__EVENTVALIDATION=3DdaK0VavnudXI0nrp1Y%2FwRCiZim%2Bx%2F6RH0SgGRJXAhLKIyL6D%2FwGG%2Fj8klI18fBo26G4ml9B2p%2BBeF%2BiSN6NTdHVmymfTIqCZOl8RYhbyzyDmbRsBGgbZH0kCMpaKaltsfOacP1ZK%2FlaJnuiKsOFpbrAKU%3D&__ASYNCPOST=true&`

	request, err := http.NewRequest("POST", strUrl, strings.NewReader(post_data))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Accept-Encoding", "")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,ja;q=0.8")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Length", "8756")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Host", "www.cqcp.net")
	request.Header.Set("Origin", "http://www.cqcp.net")
	//request.Header.Set("Cookie:", "ASP.NET_SessionId=fgz03kmkbe03u1uapopbejyr; UM_distinctid=162c8d253d2e5f-04f55d7fbea179-336b7b05-4b9600-162c8d253d33b8; CNZZDATA725126=cnzz_eid%3D969395611-1523785873-null%26ntime%3D1524028890")
	request.Header.Set("Referer", "http://www.cqcp.net/game/ssc/")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	// 完成后断开连接
	request.Header.Set("Connection", "close")

	if err != nil {
		fmt.Println("get response for strUrl=%s got error=%s\n", strUrl, err.Error())
	}

	// 设置 TimeOut
	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(30 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*30)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	//var resp *http.Response
	resp, err := DefaultClient.Do(request)
	if err != nil {
		fmt.Println("get DefaultClient for strUrl=%s got error=%s\n", strUrl, err.Error())
	}
	// 保证I/O正常关闭
	defer resp.Body.Close()

	// 判断返回状态
	if resp.StatusCode == http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("get ReadAll for strUrl=%s got error=%s\n", strUrl, err.Error())
		}

		RegexChongqing(string(body))
	} else {
		x++
		fmt.Println(fmt.Sprintf("获取页面状态失败，重新开始获取尝试%d次........", x))
		goto LABEL1
	}

	//bodyStr := string(body)
	//fmt.Println("GBK to UTF-8: ", bodyStr)
	//var dec mahonia.Decoder
	//var enc mahonia.Encoder
	//dec = mahonia.NewDecoder("gb2312")
	//ret := dec.ConvertString(testStr)
	//fmt.Println("GBK to UTF-8: ", ret)
}

func RegexChongqing(bodyStr string) {

	//分析代码
	regex1, _ := regexp.Compile("<li style='width:65px;'>(?sU:.*)</li>")
	regex2 := regexp.MustCompile("<li style='width:80px;'>(?sU:.*)</li>")
	//rxp_a := regexp.MustCompile(`<li style='width:80px; border-right:0px;'>大小单双</li></ul><ul>(.*?)</li></ul></div>`)
	//rxp_b := regexp.MustCompile(`<li style='width:65px;'>(?sU:.*)</li>`)
	//rxp_c := regexp.MustCompile(`<li style='width:80px;'>4-1-0-7-9</li>`)
	des1 := regex1.FindAllStringSubmatch(bodyStr, -1)
	des2 := regex2.FindAllStringSubmatch(bodyStr, -1)
	//des := regex1.FindStringSubmatch(ret)
	//code_a := rxp_a.FindAllStringSubmatch(ret, -1)

	if len(des1) < 1 {
		fmt.Println(errors.New("desc error get"))
	}

	var resultballs string
	for i, v := range des1 {
		if i > 0 {

			li1 := strings.Replace(v[0], "<li style='width:65px;'>", "", -1)
			li1 = strings.Replace(li1, "</li>", "", -1)

			li2 := strings.Replace(des2[i][0], "<li style='width:80px;'>", "", -1)
			li2 = strings.Replace(li2, "</li>", "", -1)

			li1 = "20" + li1

			list2 := strings.Split(li2, "-")

			for i, v := range list2 {
				resultballs += v
				if i < len(list2)-1 {
					resultballs += ","
				}
			}

			//打印添加结果
			fmt.Println(li1, resultballs)

			draws := &models.Draws{Resultballs: "[" + resultballs + "]"}

			dno, err := strconv.ParseInt(li1, 10, 64)

			if err == nil {
				models.EditDraw(dno, 320, draws)
			} else {
				fmt.Println(err.Error())
			}

			resultballs = ""
		}
	}

}

//每天重庆棋盘
func NewlottoChongqing() {

	fmt.Print("New Draws 120 is Chongqing......")
	//时间格式
	//formate := "2006-01-02T15:04:05+08:00"
	//time.ParseInLocation("2006-01-02 15:04:05", "2017-07-07 09:00:00", time.Local)

	var starttime time.Time
	var endtime time.Time

	var t int64 //期数编号
	var strt string
	var edt time.Time

	now := time.Now()
	ts := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	ls := time.Date(now.Year(), now.Month(), now.Day(), 1, 54, 0, 0, time.Local)

	zs := time.Date(now.Year(), now.Month(), now.Day(), 9, 45, 0, 0, time.Local)

	bs := time.Date(now.Year(), now.Month(), now.Day(), 9, 50, 0, 0, time.Local)
	//晚上
	ws := time.Date(now.Year(), now.Month(), now.Day(), 21, 55, 0, 0, time.Local)

	strt = ts.Format("20060102000")
	t, _ = strconv.ParseInt(strt, 10, 64)

	durmi, _ := time.ParseDuration("+5m")
	start_durmi, _ := time.ParseDuration("-6m")
	end_durmi, _ := time.ParseDuration("-1m")
	day_durmi, _ := time.ParseDuration("+24h")

	for i := 1; i <= 120; i++ {
	LABEL:
		if ts.Unix() > ls.Unix() && ts.Unix() <= zs.Unix() {
			edt = ts.Add(durmi)
			ts = edt
			goto LABEL
		}

		if ts.Unix() <= ws.Unix() && ts.Unix() >= bs.Unix() {
			durmi, _ = time.ParseDuration("+10m")
			start_durmi, _ = time.ParseDuration("-11m")
			//end_durmi, _ = time.ParseDuration("-5m")
		} else {
			durmi, _ = time.ParseDuration("+5m")
			start_durmi, _ = time.ParseDuration("-6m")
		}

		edt = ts.Add(durmi)
		ts = edt

		starttime = ts.Add(start_durmi)
		endtime = ts.Add(end_durmi)
		t++

		if i == 1 {
			starttime = starttime.Add(day_durmi)
		}

		//fmt.Println("重庆时时彩数据添加成功", t, edt, ws, starttime, endtime)

		//添加到数据
		drawtime, _ := time.ParseInLocation("2006-01-02 15:04:05", "0001-01-01T00:00:00+08:00", time.Local)

		draws := &models.Draws{
			Counterid:       320,
			Drawno:          t,
			Drawtime:        drawtime,
			Drawstatus:      0,
			Starttime:       starttime,
			Endtime:         endtime,
			Betclosedmmss:   "00:00",
			Isclosemanually: false,
			Voidreason:      0,
			Resultballs:     "[]",
		}

		if has := models.CheckDraw(t, 320); has {
			falg, _ := models.NewDraw(draws)
			fmt.Println("---------------------------------------------------------------------------")
			fmt.Println("重庆时时彩数据添加成功", t, edt, ws, starttime, endtime, falg)
		}

	}

}

/*
ScriptManager1=UpdatePanel1%7ClnkBtnFirst
&__EVENTTARGET=lnkBtnFirst
&__EVENTARGUMENT=
&__VIEWSTATE=WEbm%2B543tuXtFfJ3UZY3j1t2t65KnxycuAVUi5VecmFytQjobRYazzx14iuc9QZqAMJmS2QW6QJQ2wKgnsQQ0ypPhSI2mgLPFBryrZwnBubfQ1qsVtSWf9JoGVUSyOO2vHlLwXzZY5ODZn0xBTU0aeAeKzNyWKUs%2FKwCqT%2Bj67si1rFy7v%2B5YsXPjGYCjVkitrU87kM3RyvSTae9tByWfMi0noKxgvzGiEPR8402CdgOBPTns%2BZgTlykTja6BF6pSInH156pwYfYreYbJvqTTcCucv0LhfESzZ%2BuRbRMToTzqJ%2Bka6az4pxkfpgnPl34NhKT3IzlgzHQ1Tu0558PkRTWiVLzvfmw5EtUvbuiw9tRq5ZtVx%2FvW%2FcCUiket3vEdUx79pF161AAGpyvcPAw2DSvBQiqlj4BkWInhsjn3KEQvk71%2BkaEn%2BRwg84xIlXo3UKxiYPhaGCudkuXj79nSc6k2u5zPDGo%2BAZlt72fFS2sIFqMR%2FwghKrWh67gmPf2fn4AJH2R9QTzXWVfNgMoF1b6lFyND7FQPs3I6PPwDEilqpanaY2X9jUTBDpNDuWThSgicbaNQfRmw%2BRtbqZBWEn3GEoYGF82FOtbT9ZC1%2Bo2HB8ek2%2F7WQoohnQaPtr%2Bx%2ByJeqdaGLX%2BegTwWOaoQPRWukCfgjx0kio%2BOqHQU4rmT0Sgoq4APYJRK4peM9cWwAl92vN4hBdbYphWpS1kIV%2FnxzxA3w13%2FPoBk6gKxQx0ApO3fzNJ%2BvOqoQ89hDy%2BE0LV147BzOHBwWI8Rd2Rd%2BCVtUduBD1s6L3s%2FFD7pTdWCCdGyqVZTKOzRiwKutHzkDKuZZxNGkUQyisV0Ma8Hlp3jI4zC8RO6OzDO9ZPYLf4Kg3xh%2BHo3jC2GrT1Kl2B0F%2BR2vWwShBVuGctoGEyumI1NEPxE4x0sRrYvgNo7ie%2B2MyrAeVL41yaLhAzl25O7J2aSpU6Qxfp9H2pVQQu3CE%2F73nka0DYIw%2F2N8WHH8ZqQFrY7SuVSC9NTE980QaHCxWTebcvJOMfXvVOIlgSVq%2B19LRX9ldCY1pMrtcT2OYwGgny0riJMhTO%2FSHsVwVtORl8uzrzSLAURe%2BWXrK8SmXuw150I1%2FOFqftZm49ZHon6po1UGNFYIv%2FntHxTdvifleKmtBYRGnbfhZiNHhN%2BveyDUihJVo2FbuCjuVRx2sjIjRrd3uj8yupaUEV0%2BzxGZ6r1pfPhmjmBSBQ1lPCnnpOh8pVWhK55YGGOTfUBeS442xBHr%2BPBTGzK2CJotH6YwZC6Wzhg9wcBQLmeX0YbH1b3XNRg%2FGrJRf3Ood%2BJNt0fVmUM18Dyo0o7jtLHZ82HJJaUSbIMDhSfkMHq9IOsrEYSIbeTKHDUNoZdsa69KNguoIkA7axz2ietDUXmKOdM7Z6V70444BNKpB1prcvD1kXLMt3WiMVzTnhC1cJsf6nu%2FClzGFmkG9uJnETCjanCMe30B1nGlKc%2FG8XBhIC4W%2FDB%2BhTMAFXebEdw2RMpszkR3fG2oNZhDC8C8GBB5YPdgwcuDxhNN69ZQZkbAAfC8SHSCY7ZqHzEl9FWNL1z7MeMHYPbcGyO8VJs%2FisB82Qoi%2BXiYkCSf42z8F0RtgXZVvZr6Ka%2B3XzwZNgSReVz2Jtc72Kxbf2vleKq%2Bjt8iksyg7klDnpD5dUWpWl6166s1%2BQHa5fsGOrd4%2FCnEv78DD8bEPLuaT0uYXT3a34olaggJiOQ3FMWj4NL1nXJBM4Wnf6KfASd%2BvcvkRaaY8hpJ%2FxUtXpeXnvtFpg%2BRLJlBa5r9CcIgdf3vflArDd%2B70UWwAi2R%2FqlpYbSdvGkpZF%2BWrS8jCwkGr1%2BpswfXHdlu5wBR0VSZ0AkgzXj2JEgMmzrWNJ%2Fp00IBw%2FRZdL0glhYRwNVFq8Mj9oVNHlfOT7FG3zIz1TzcuJxUpoQgfHJXgRRnSDkoX5fHTw%2FwG6D3PZUndoQ%2ButSxmPET0txd6fNGHJXo06moYUnvJ0pLCOgFU355dmKHkR9H0bP6G1%2BMkoyfXFP7fnqLAJFeAWRTWLsS31dw6%2FpqtRRtllt4ma3Qjv1KThWTY%2FYuOhez98vaciEajwnccxVZk2L3odR4yynjCabmoTDk1Ib%2FCDmzHqWNAmDfrILhtBIWhEWqYZaBw9tqnWGnhJZ2gXIa9qRyo52btzxlfYBjvmfe7IX21in2xocHlw3cYQyDjHrgEmwZCaKRMCRm4F9vkdP7byPX4rmqXQoEj57dK3ENOX0G7MhribUd5RDC4MeHTNdMHN3RlpJ%2FKz8ERaG4Z1oXGfeDAo2p8353oa0Hz0cFTWh7%2BC9gXeAHCUrR2709FlAfoOVHZ66YpzLdBRhxwrs%2BZw2HWVO8rQqnNOHIJwmfps6Egq6d%2BaUyyDvPGkHOPNWiTBgpjJqZOT5dk1HcqD%2BWpFpmQW27HWIDNbo%2BRoThd3aOwiqBOK0OJ6DZYXZvbe9HcmxhtaNtt8yCsLgv%2BnPAO1zCchwg3MndR%2BbIKfMT5Zxjlfjfw9ix7CtNjThqL%2FiTZ8%2FliyHYJxHZftXtRkTeL57X%2BY39FtIOAtSwO%2FwXIOWUjOA6tdxtbiCwuWRZNv%2BsUSvaPcg9k5TTYeK1H%2F7gUsDU37CcSoVCNykXeUx7t9cCdfVR1fHcs%2BXxKRkzWBbTbdMxVP689W5ACWXPOXRdPGy6rFjBFI8EPPf4SuaK6yIfOgYWaRbh4AShjTLcEqWCRFs0YvBj3UHz6KVMEb%2Brc%2FPg%2F8btQPa47wDc1Dfmmevmpjgf%2FAeWjSZTfYnODsuLufx8UcUsVnExyEJv8hHwkXcrABwZDr%2F3suohKwUMUsQ9pfyyrmcu6EPcLawizmMIQAAYe%2Bu9s3Kj9NV8tImOeppSu0g%2FCTB644%2BctSZld9K7a6oGzC%2B%2FpEiy43yhHOrjrlP2dP7o391l9nVRdXjQpkvE%2FKwdo9yDOQ%2Bensf8niQ2%2B3UhSbuc5c0Erlu48VenXG1KerMmkjls6wNz50owxrNRFiakCM4kV%2B0eRpxPC7KyK5H%2BanKxN5toQTSqSIh%2FYMDJJ8ffDgCldsv4lnLoA7DrkYfPlCM1uKWe5gKO406ds5kkoY906oJwCciUV9x55uREdv3%2B3v2sCn3f6eiuxzsK14CcDIi6qQ6%2BExPPJclTCMFgy1AyL8cjBhfK4Np5dhq7s0b0jM8CSbPAoj%2FeyhcvRYMOcLHaOKU7AWAStrJNfhyfSTjy7SljWXkBeMRzRut4EVz1vonAYdNOdXLuii%2BHH990TxaA2hMmbpD4PHdhZ5%2F3FbBDazwgiQY10gIT1ppyG8EOfTGYobSVV8BxjAgkinWG0RmY4sx%2BV5%2BQM25IVTx6tsxCZ7zpSvdTBCUhcnTI5Q9wLNuOzahLKtcCL0DGwByLdQZdDBa1U%2F49Ukg4VEh3Z68QFpTmR0%2B%2FhoU25JQtW03TloDihaEjr7zH%2BiQprrqW%2Bn%2BRD4G7MgNTEjNUoSFe1LS%2B%2FvyIg0Eh1jKXfUqEjv%2BkytSmxFIG8KRuaJiMjjTz11EjrOd0jvo1gu608BzbYXR3mCnGAta%2BC6CdQ3VuKdQdZ5iDHUxvAG2GFwvP6rT%2FW9lxF%2BzUsB2qfYcouglrbYszln%2FUlUJFUz3%2Bep0coKt6aX4EqWyxwtTk6%2BSYEPYIH7QMYlk3QOd%2FTHEPKYg4LMiZ4lPvt%2B2%2Bj7wO7Q7%2FG%2FRdsoDHU9LVFxbMRSU0f8Nl4bwWhx7cXAL8LWYo8CAWdEvSU3dzM%2FVZ3NRfo27qKdWQm%2By5EWD%2BKipTjXwSPh5cowaO4QUmitiBdgoI%2BloEk0xkpvk5Lhxn8f5kEhrGLxsVrJ9cJN5O9rUFEWQFhlHEVFBnPdxVGfIPXNqlyHa8EORulUA7gKgdkTfdKjerERUq44v%2B7ZNu2VVfFNGPcB7ObCPOnVl2RnCWUz5ryFiyIeU2Jguq2%2BXKdXqSlX6mig56NVLp6xANnFlbR6sjaaK%2Bt3alAtoM2FZblBHKnGNgOkhv%2B6aQuLhh6AiIcCGgBmXO6T9Y84RVHs1k%2ByK9wSb0TtLI71Ccb3ri0mecqC02SUKT8MlMY8Qd1%2BlCnXXL8n3eo5I4MdUIttW2z04Cv5SMRnxAXdOEvApaWRQCVvx4fy6GuXIsaEQ3fc%2BcJt1KrCUf%2B8%2Bn%2FT0LkWnEAwiAowaMjq5FNeiM8HUXUcy0R863jx5mv9Li53XnrQpjIlwqMU4kAVL6BFGrfhIg5GHviGy89SoZOSQZ%2Frgi%2FH%2FdmaLHh5Ti%2FXtAeLDPmGn0OffwIckd2SOYiYQF0foIZmzf0COw%2FWCca1NWlO9QUR5cdnZtt7EbKrKzi7EtIVxjcNsh2nQqcbCZi04tK8OGF89D%2FMBe4efOvMvNra7RuhSbQ%2FeJociRKwcsK%2FcqQL6Cx6Eqspp2Ui%2FXGyxnhDPLpdVJAP0HRkbux%2FS9igRk9PxZaxrnH2sukHbBEFW9ZsLoOXwKNeBM%2F2qMQvnKGX9VIUfHP4pYVSZN0Qz2gNqHPozVzuE5uSElM2TtS2UUOKEAca%2BayzD8syomki0K9C7%2BLYemabl3YSBMDTlE5NHT8FuxMNQQaZcRGorH0%2BwChNhlZkUizjNCV5ZgTCsmDuYPG1e8wU0KuYiKwB8%2F%2B2TbJzIeoEu1io19LMe%2FAu%2BMmJ5iIih2sypRGayT7EzlXYsiJAD%2BzHWCnxZpUpVqowFTyI0FYpA5BEYmPwPfBrk%2BT2uXThS2kGruWuMQo%2BCcRugUrAGQLXua1N9Xh3g0YOpXQAhU9Z2W%2FXV91uofeNoC0hTktCv%2Fw1GKHMKFEXewARHbqB1IM1dWL6qxl3XsMU%2F22FioZ698V0ILRoNmUuvpoz%2BLyDX7kUK6Pd%2BKMn14%2BAe%2FxdWB9RKwdIUXaiNoPHcu3auQ%2Bnrx6ZIrfNGUG9wYfYpHnbGEbclv86KNTzJ6tKX3d0sT9ggakoC4hVoC50ulYPM84fJiBq2CF2mN1G1YLkJ%2FaXrm4VT%2BcXU5LW%2F2JdocqH9GRExIg2Z7VNB0FoBpfl%2BiZQyhIO6j5ItwlcS4pV5KQq7aBA5gDlb%2Fmdf3UppzEYWSZLliegzAoX9m1pJTOjYZMEZalxBfSMclvotQ9pASXMdsMf6mi0aux4WVBNJgLRN8osVyXKev9SBV2dPXQjTH9S41ddmBkP8%2B1EtQvE5cNC9FTIBhrqkYmELAXjmPBBcGV6GsuVKtbNQraab%2FlBipoQeX5LTbcgPfKRKUcOIX%2FHoNs2tBaYyDldWM2jNL2QLszAyeqDuz7oKX%2B%2FF09PxcbxhAKaJ7RyRWhJhTRi4MN64VNRTjy5NesxvnlbRfpWgasHoIpX3%2FvAnvB%2BK3P6M9gS9CTYToSVFtrOPAtwXTRdC%2FP9Q5Z%2Fr7mzZ%2FQrktF8w8fzieRw1uxzOfd6Y0Fs3du2qKBaI9SkI92AtVacm5nP%2BAwiR%2FMhzJYDEaXOAp98SCOtiVn63ur2GFGPioqfJCjpXPIIDZnr7bY4vPP7rzPoiQ1%2Bn1A2W2%2FIcZgblazeS9IxaOGN1T4ei7ZdB6VN%2F%2BXmb9FcotWICkmafYE%2F%2Bv7nL2r9K0Mm4xY%2BVKLJX4XYA6iC2mT0m0%2Bszz1opx8sZGgmywj%2FA%2BBS05sGgNLtYDJl4VErjl13F22Y8YDPmd8vUfhsK8Qkl%2FZujW%2BrLptqYEYApAVvOpysyTDlraiN1R5nG29NKNGcxAH5ZuDrQU8nZMDSiPsKadyfhl9yH0i0Vrk%2FVMM%2Brxd%2BA7BOSDH57M%2B0PsyZt4%2F4orBKhRQHljAxOdSezIiAy8pdKGntPouQQCidGZmjOqWuTk9BLdPqEvfOG4T4h%2FSod6Y2M7wn9O0s2s3C8BeWQh9RODWCnT2lWTKSe1oN%2BG0FSsHG6S9GKfGRyH7V7wv81FB7IMZgWqLF%2BaBYAGgF%2B1OKlL3BfMuGhgWxPcc0xZA1w2Bqg%2FCUz8ikachRr4ZocDbeBusXxXmJZAWN3wAK3d58esUxkqtfq6m44FXs0kDZWNByAovjQPzxM7tDeBb8sOdq6NLURIG1m42jVNLRWxPVVGmsoXD1GIiwyU7%2FS%2Bau25jCY%2BFC7GWajkLLnxwYNBcQ61dNCmMx72v0vRN%2F%2Bj9cNxo6JWLr5K5Nz9klvBtfEt4pfq%2B4aK036dB6sRXusKUwzwejb7vxZ%2FywYuY4fOkkOsSmhyIulrO4f2cAg7VHbFasnKrpGbXz1X9kyq2InuT32pyzjlFLqMiOa7r0knzcz4OTwHwRXBvgBMpy9cp6UM80TqMROtxH74MZpZ82uBBfUed%2FpazPK8rC%2B0pDJTxaqDPraaroNrAdEpgZUIj%2BFaqJYBo1AZ6c6JWZvgYkSKUGZacfkXkwAhptPFuQilT%2B7up6HPsm3a6AYduGlzoekhbE%2FsxVb2jr7cuWE3T8K46Ubyq9ogLRnr3S3WWb2PtY703KHBvGzH%2BUG1955M0fFTziMzipvcbkjVJ%2F5z62BsqkwvP5KgR%2F4Frs%2BqSWmhU9qLPeIAzx9x0Zps%2FChpMjrq%2FLmoEwVVvHq11VGintQj8dolSi4Fis6PvwdDINwSC2%2FMm3kfSIEQC1179%2BzLngG7NX7U9iQFXZpL2mqIPUHBOTAqQcsnXuui2oQBAndkaXNPmTuL%2FRtondz989iZ3sePtWUtjK9IxgVRNGgh2olSzJ2iZZpuwXLbN50m5%2FZzfe%2ByQfAT0GiwhO%2BxL9HFgMn6eurVKFtQhdMgBYfEL7DDtyurWYwmnPut6s5fE5PDGVD%2FOeDuRXvIzp0fCI%2BrSRwFmWXraL4fExj%2FMUh3gCLL0bMuIwoSnLxKiFmM1u30Z8GxwIHDnNNdXFBkOkOLzo0jqqsndiczdo1lQqElkanCdCz0yzxQiomEFAN0MKmCRlCTvI5dAmBMdXgOPiQLxUf4BhUsHINv7ILa34U8vahNJ%2FcPE2xoXaOhU2pjm4TKWJ0x5tcc%2BSdcoEMBv2e5nYKnaOR3H7wrKn%2F9onRWwx4tB1m8KwNQfJsMKack5SQFW99XC6wt3sHmMUwcyRi47BTQEgATNJnOwmRPYimZWiGol7s3bIeAyJDpYoEwyHxpzeCiloMYdmgvk5As7QKG%2BvClT%2Bv8p%2FVf6Vm7rM9LooeARHULwvimC8gVajOaqSmysbBF5UsoYWW3zrFzqm1CdJckqUICUCIoXb2Mp94i5rjTgtBbjE3Y2ifyxNH7%2Fz%2FM6uMIsudSDfMW9tYUQWs%2B09DOoEqdLDPOCT9IKb2x3W1TVaudanIQFb6b596vomQ8hEJILQkrEMEY%2BSrc%2FU6uTEZ%2Bmo6faDUrbh5vny4E9RQ40FFQZA%2Fg6B%2FXmekHUUTDMScQzYxvefd5FT50Smm2HHiNaUhy%2B4VlEQ9jPu6BogVvqkRUXtp%2Bl0NqPS7eZlZBRZgUxdNu1hwC88x%2BysK1%2BDkgCgvCWM4vSWFe%2BHXFuoiuxeXy3BL8mSA2PHMXeMHQffHHeXWtDMOOpNnnBEaPSl%2Fdxn0HdnOeK1qp7bR3uHKwe1sIWWQayivKALwx5iVm3Lf27K29QwgC36yxmSBbwEH4K8TQ4EEUO%2FS8kUEWn8unu%2BItUALvITeTMFk7LktHYSFhGVKaISNO0%2BBDT0rhMjKt6yNXnmWbZ%2F6s3hO6aGtlcZR6JkjsyCOV3YNFOGgxROjlJPFnUj0y7%2F9Md0BIYhW0LWYvkgcT%2FFq%2FcMY7JQ%2Bq3LYwR%2BN1L2Ah24%2FltPoFMXQJFbrgBLFZpERmHI4%2FGfh2iFTSQyyn%2FIrLPXzoOauW1P61ku%2FEmqi2sH2hepiOHfZXjW3x9DO2mZqtmDymabJnN%2FjWLU7VCWK7qrJCdA%2FiAD1vrbkIORboZYaDVwPgdHOePxNAjq3JUFhyWFMYwGwE6rqMyw2Fzpuan65o3ai6qumauHRcCpfsufKHpv9mb0BcnwEIvxmdwaxBvusumnuapnFo60t01q0FCK3cRsr7hd
&__VIEWSTATEGENERATOR=7B9E3A74
&__EVENTVALIDATION=TYvrNB%2Bpy6ohPKd33%2BodwfH5YHjNA3%2BZsyMWgOAOLWTKOwscm5POGShbOd%2BKh8096HVlE4Ri0jKIoiXNo4QWEdYqa6n1a1zQbN4AK7Nt734vA1mKdtIXFhugFNWg7IAbL3OB0K6kaZPOV%2FqrPe4DfJUCTZM%3D
&__ASYNCPOST=true&

*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	//"fmt"
	//"encoding/json"
	//"io/ioutil"
	//"os"
	//"fmt"
	//"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	"net/http"
	//"188bet/models"
)

func main() {
	centraServicePost()
}

// func main(){
// 	/
// 	router := gin.Default()
//     router.POST("/form_post", func(c *gin.Context) {
//         message := c.PostForm("message")
//         nick := c.DefaultPostForm("nick", "anonymous")

//         c.JSON(http.StatusOK, gin.H{
//             "status":  gin.H{
//                 "status_code": http.StatusOK,
//                 "status":      "ok",
//             },
//             "message": message,
//             "nick":    nick,
//         })
//     })
// }

func centraServicePost() {

	c := &http.Client{}
	//req, _ := http.NewRequest("GET", login_url, nil)
	//res, _ := c.Do(req)
	//post数据
	postValues := url.Values{}
	postValues.Add("IsFirstLoad", "true")
	postValues.Add("VersionL", "-1")
	postValues.Add("VersionU", "0")
	postValues.Add("VersionS", "-1")
	postValues.Add("VersionF", "-1")
	postValues.Add("VersionH", "1:0,2:0,3:0,9:0,13:0,18:0,21:0,23:0")
	postValues.Add("VersionT", "-1")

	postValues.Add("IsEventMenu", "false")
	postValues.Add("SportID", "1")
	postValues.Add("CompetitionID", "-1")
	postValues.Add("reqUrl", "/zh-cn/sports/")
	postValues.Add("oIsInplayAll", "false")
	postValues.Add("oVersion", "3,181066|10,623")
	postValues.Add("oIsFirstLoad", "false")
	postValues.Add("oSortBy", "1")
	postValues.Add("oOddsType", "0")
	postValues.Add("oPageNo", "0")
	postValues.Add("LiveCenterEventId", "0")
	postValues.Add("LiveCenterSportId", "0")

	postURL := "https://sb.oneeightyeightbet.com/zh-cn/Service/CentralService?GetData&ts=1521010315794"

	res, _ := c.PostForm(postURL, postValues)
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(string(data))
}

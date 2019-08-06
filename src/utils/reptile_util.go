package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/redis.v4"
)

var loginUrl string = "https://manage.xxx.com.cn/home/login"
var dataUrl string = "https://manage.xxx.com.cn/?pageNo=1&pageSize=100"
var userName string = "xxx"
var passWord string = "xxx"
var redisServer string = "192.168.1.2"
var redisPort string = "6379"
var redisPass string = ""
var redisDB int = 0

func main() {
	productList := make(map[string]string)
	productList["椰油"] = "CL_Spot"
	productList["咖啡"] = "COFFEE"
	productList["工业铜"] = "COPPER"
	volumeList := make(map[string]int)
	volumeList["u_CL_Spot"] = 0
	volumeList["d_CL_Spot"] = 0
	volumeList["u_COFFEE"] = 0
	volumeList["d_COFFEE"] = 0
	volumeList["u_COPPER"] = 0
	volumeList["d_COPPER"] = 0
	jsessionid := getCookie()
	doLogin(jsessionid)

	request, err := http.NewRequest("GET", dataUrl, nil)
	request.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: jsessionid})
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	doc.Find("table").Eq(1).Find("tr").Each(func(i int, tr *goquery.Selection) {
		td := tr.Find("td")
		name := td.Eq(3).Text()
		dir := td.Eq(6).Text()
		if val, ok := productList[name]; ok {
			buyNum, _ := strconv.Atoi(td.Eq(7).Text())
			buyUnit, _ := strconv.Atoi(td.Eq(8).Text())
			num := buyNum * buyUnit
			cacheKey := ""
			if dir == "买" {
				cacheKey = fmt.Sprintf("u_%s", val)
			} else if dir == "卖" {
				cacheKey = fmt.Sprintf("d_%s", val)
			}
			volumeList[cacheKey] += num
		}
	})
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisServer, redisPort),
		Password: redisPass,
		DB:       redisDB,
	})
	for k, v := range volumeList {
		strKey := fmt.Sprintf("redis_order_count_%s", k)
		redisClient.Set(strKey, int(v), 0)
	}
	fmt.Println("puti volume get success")
}

func getCookie() string {
	jsessionid := ""
	response, err := http.Get(loginUrl)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer response.Body.Close()
	for _, val := range response.Cookies() {
		if val.Name == "JSESSIONID" {
			jsessionid = val.Value
		}
	}
	return jsessionid
}

func doLogin(jsessionid string) bool {
	data := url.Values{}
	data.Set("userName", userName)
	data.Add("password", passWord)
	request, _ := http.NewRequest("POST", loginUrl, strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	request.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: jsessionid})
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer response.Body.Close()
	return true
}

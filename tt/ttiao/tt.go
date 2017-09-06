package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	proxy "github.com/myself659/xiciproxy"
	"github.com/parnurzeal/gorequest"
	"github.com/myself659/csvdata"
)

//通道定义
var gchid = make(chan string, 32)     //分发ID
var gchdata = make(chan DataDesc, 32) // 写入数据库

type fileDesc struct {
	Login_status bool       `json:"login_status"` // 是否登陆
	Message      string     `json:"message"`      // 消息
	Data         []DataDesc `json:"data"`         // 具体数据，获取的重点
}

type DataDesc struct {
	Title           string `json:"title"`           // 标题
	Go_detail_count int    `json:"go_detail_count"` // 阅读数
	Comments_count  int    `json:"comments_count"`  // 评论数
	Display_url     string `json:"display_url"`     // url
	Behot_time      int64  `json:"behot_time"`      // 发布时间
	Chinese_tag     string `json:"chinese_tag"`     // 中文标签
	Source          string `json:"source"`          // 来源，作者
	Source_url      string `json:"source_url"`
	SrcID           string //`json:"-"`
}

// 将一个目录下ID文件进行分发
func deliverIdFromDir(name string) {

}

// 分发ID
func deliverIdFromfile(name string) {
	defer close(gchid)
	fr, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fr.Close()
	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		if s != "" {
			//fmt.Println(s)
			gchid <- s
		}
	}
	fmt.Println("deliverIdFromfile")
}

// 用户cookie库
func httpGet(url string) (*http.Response, []error) {

	cookies := `uuid="w:43e2e16a7fe14cd49bef554de069f635"; UM_distinctid=15c9a3c0060d6a-0e87f19271ca43-143d6d5d-384000-15c9a3c0061ca6; sso_login_status=1; cp=594BA66C49770E1; __utma=24953151.1210624272.1497236244.1498120797.1498130052.3; __utmc=24953151; __utmz=24953151.1498103544.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); login_flag=02d25c8fbb14f81de8b59d4c9ec36d5a; sessionid=c7c4e1d7be9494a636c75ed4ec5db4b0; uid_tt=e742c919b21dc0a292ce77f30c8cc4a8; sid_tt=c7c4e1d7be9494a636c75ed4ec5db4b0; sid_guard="c7c4e1d7be9494a636c75ed4ec5db4b0|1499666449|2591573|Wed\054 09-Aug-2017 05:53:42 GMT"; csrftoken=78d8471020e22cc2ec20791fa4164b15; CNZZDATA1259612802=490756214-1497230969-%7C1499685996; _ga=GA1.2.1210624272.1497236244; tt_webid=6430580691792971266`
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
	request := gorequest.New()
	resp, _, errs := request.Get(url).Set("User-Agent", ua).Set("Cookie", cookies).End()
	fmt.Println(errs)
	return resp, errs
}

// 处理一个url 返回个数与最小时间值
func idItemDo(id string, url string) (int, int) {
	// 需要支持代理
	//resp, errs := httpGet(url)
	resp, errs := gproxy.Get(url)
	if errs != nil {
		return 0, 0
	}
	defer resp.Body.Close()
	// 保存文件
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0
	}
	//fmt.Println("body:", string(b))
	fdesc := fileDesc{}
	err = json.Unmarshal(b, &fdesc)
	if err != nil {
		fmt.Println("json:", err)
		return 0, 0
	}
	num := len(fdesc.Data)
	hot := 0
	if num > 0 {
		hot = int(fdesc.Data[num-1].Behot_time)
	}
	fmt.Println(fdesc)
	for _, item := range fdesc.Data {
		item.SrcID = id
		gchdata <- item
	}

	return num, hot
}

// id 处理的协程
func idLoop() {
	for {
		id, ok := <-gchid
		if ok == false {
			<-time.After(500 * time.Second)
			return
		}
		idDo(id)
	}
}

func strDate() string {
	y, m, d := time.Now().Date()
	return string(y) + "-" + string(m) + "-" + string(d)
}
func fileName(id string, hot int) string {
	return id + "-" + strDate() + "-" + strconv.Itoa(hot) + ".txt"
}

// 完成一个id的处理
func idDo(id string) {
	after := 0
	count := 300

	for {
		url := idUrl(id, count, after)
		fmt.Println("url:", url)
		num, hot := idItemDo(id, url)
		if num < count {
			break
		}
		fmt.Println("num:", num)
		after = hot
	}
}

// 获取url

func idUrl(id string, count int, after int) string {
	sc := strconv.Itoa(count)
	hot := strconv.Itoa(after)
	url := "http://www.toutiao.com/c/user/article/?page_type=1&user_id=" + id + "&max_behot_time=" + hot + "&count=" + sc

	/*
	 http://www.toutiao.com/c/user/article/?page_type=1&user_id=51558403490&max_behot_time=1499664205&count=20
	*/

	return url
}

var ggromdb *gorm.DB
var gproxy *proxy.ProxyPool

func dbsave() {
	for {
		item := <-gchdata
		errs := ggromdb.Create(&item).GetErrors()
		if len(errs) > 0{
			ggromdb.Model(&item).Update(&item)
		}
	}
}
func dbgrominit() bool {
	user := "root"
	psword := "jbigc"
	ipport := "127.0.0.1:3306"
	para := user + ":" + psword + "@tcp" + "(" + ipport + ")" + "/" + "tt" + "?charset=utf8&parseTime=True"
	fmt.Println(para)
	var err error
	ggromdb, err = gorm.Open("mysql", para)
	if err != nil {
		fmt.Println(err)
		return false
	}
	ggromdb.AutoMigrate(&DataDesc{})
	ggromdb.LogMode(true)
	go dbsave()

	return true
}
func main() {
	fmt.Println("start")
	gproxy = proxy.NewProxyPool()

	ok := dbgrominit()
	if ok != true {
		return
	}
	//fashion5w()
	
	fname := "ids.txt"
	go deliverIdFromfile(fname)
	idLoop()
	

	/*
		id := "51558403490"
		idDo(id)
	*/
}




func unixtime(second int64) time.Time{
	return  time.Unix(second, 0)
}

func fashion5w() {
	db := ggromdb 
	datas := []DataDesc{}
	
	//db.Find(&datas, "source_url = ")
	db.Table("data_descs").Where("go_detail_count > ? and  chinese_tag like ?", 50000, "%时尚%").Find(&datas)
	name := "fashion5w.csv"
	fd, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer fd.Close()
	csvenc := csvdata.NewEncoder(fd)
	for _,temp := range  datas {
		//fmt.Println(temp)
		csvenc.Encode(temp)
	}
}
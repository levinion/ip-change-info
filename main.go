package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"time"
	"github.com/jordan-wright/email"
)

var (
	//发件人
	From = "username <xxx@126.com>"
	//收件人
	To = []string{"xxx@163.com"}
	//邮件标题
	Subject = "您的 ip 有了新的变化"
	//服务器地址
	addr = "smtp.126.com:25"
	//授权
	Auth = smtp.PlainAuth("", "xxx@126.com", "授权码", "smtp.126.com")
)

func NewCustomTick(interval int) *time.Ticker {
	return time.NewTicker(time.Duration(interval) * time.Minute)
}

func GetIP() string {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	ip, _ := io.ReadAll(resp.Body)
	return string(ip)
}

func SendEmail(oldIP string, newIP string) {
	content := fmt.Sprintf(`你的 ip 从 %s 改变为 %s ，请关注。`, oldIP, newIP)
	e := email.NewEmail()
	e.From = From
	e.To = To
	e.Subject = Subject
	e.Text = []byte(content)
	err := e.Send(addr, Auth)
	if err != nil {
		log.Fatal(err)
	}
}

func EmailTest() {
	ip := GetIP()
	content := fmt.Sprintf(`你收到这封邮件代表你的配置设置正确。你现有的 ip 是： %s`, ip)
	e := email.NewEmail()
	e.From = From
	e.To = To
	e.Subject = "这是一封测试邮件"
	e.Text = []byte(content)
	err := e.Send(addr, Auth)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckIPChange(oldIP string) string {
	newIP := GetIP()
	if newIP != oldIP {
		SendEmail(oldIP, newIP)
		oldIP = newIP
	}
	return oldIP
}

func main() {
	oldIP := GetIP()
	//检查时间间隔,默认每60分钟检查一次
	t := NewCustomTick(60)
	log.Println("Server starting...")
	EmailTest()
	for {
		<-t.C
		oldIP = CheckIPChange(oldIP)
	}
}

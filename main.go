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
		Subject = ""
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
	content := fmt.Sprintf(`Your IP has been changed from %s to %s`, oldIP, newIP)
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

func EmailTest(){
	content := fmt.Sprintln(`Email send successfully`)
	e := email.NewEmail()
	e.From = From
	e.To = To
	e.Subject = "Your server starting..."
	e.Text = []byte(content)
	err := e.Send(addr, Auth)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckIPChange(oldIP *string) {
	newIP := GetIP()
	if newIP != *oldIP {
		SendEmail(*oldIP, newIP)
		oldIP = &newIP
	}
}

func main() {
	oldIP := GetIP()
	//检查时间间隔,默认每15分钟检查一次
	t := NewCustomTick(15)
	log.Println("Server starting...")
	EmailTest()
	for {
		<-t.C
		CheckIPChange(&oldIP)
	}
}
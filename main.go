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

func NewCustomTick(interval int) *time.Ticker {
	return time.NewTicker(time.Duration(interval) * time.Second)
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
	e.From = "maruka <levinion@126.com>"
	e.To = []string{"levinion@163.com"}
	e.Subject = "Your IP has a change!"
	e.Text = []byte(content)
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "levinion@126.com", "XWJRKRFQVZHPCXHC", "smtp.126.com"))
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
	oldIP := ""
	t := NewCustomTick(1)
	for {
		<-t.C
		CheckIPChange(&oldIP)
	}
}
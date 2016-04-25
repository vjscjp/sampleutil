package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"
	f "github.com/MustWin/gomarathon"
	"github.com/gin-gonic/gin"
)

const (
	MARATHON_URL  = "MARATHON_URL"
	MARATHON_USER = "MARATHON_USER"
	MARATHON_PWD  = "MARATHON_PWD"

	m_url  = "https://shipped-tx3-control-01.tx3.shipped-cisco.com/marathon"
	m_user = "synthetic-mon"
	m_pwd  = "VpYdy5abudqkk3Ts"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.Run(":8888")
	
}

func InitMarathon() (client *f.Client, err error) {

	url := os.Getenv(MARATHON_URL)
	if len(url) < 1 {
		os.Setenv(MARATHON_URL, m_url)
		url = m_url
	}

	user := os.Getenv(MARATHON_USER)
	if len(url) < 1 {
		os.Setenv(MARATHON_USER, m_user)
		user = m_user
	}

	pwd := os.Getenv(MARATHON_PWD)
	if len(url) < 1 {
		os.Setenv(MARATHON_PWD, m_pwd)
		pwd = m_pwd
	}

	auth := f.HttpBasicAuth{user, pwd}
	client, err = f.NewClient(url, &auth, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("Invalid Details : " + err.Error())
	}
	return
}

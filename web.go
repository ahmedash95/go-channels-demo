package main

import (
	"github.com/ahmedash95/go-channels/emails"
	"github.com/ahmedash95/go-channels/queue"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var mailService *emails.EmailService

//InitWebServer ... intialize gin web server
func InitWebServer() *gin.Engine {
	mailService = emails.NewEmailService(queue.JobQueue)

	r := gin.Default()
	r.GET("/", homeHandler)
	r.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/email", sendEmailHandler)

	return r
}

func homeHandler(c *gin.Context) {
	c.String(200, "Simple Email Service")
}

func sendEmailHandler(c *gin.Context) {
	emailTo := c.Query("to")
	emailFrom := c.Query("from")
	emailSubject := c.Query("subject")
	emailContent := c.Query("content")

	email := emails.Email{
		To:      emailTo,
		From:    emailFrom,
		Subject: emailSubject,
		Content: emailContent,
	}

	mailService.Send(email)

	c.String(200, "Email will be sent soon :)")
}

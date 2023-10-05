package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())

}

type birthdayMessage struct {
	birthDayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthDayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your %s report is ready, You've sent %v times", sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("=============================")
}

func main() {
	test(birthdayMessage{
		birthDayTime:  time.Date(1998, 11, 2, 0, 0, 0, 0, time.Now().Local().Location()),
		recipientName: "Joyce",
	})

	test(sendingReport{
		reportName:    "Ola john",
		numberOfSends: 2,
	})
}

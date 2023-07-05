package main

import "fmt"

func main() {
	sendSoFar := 430
	const sendToAdd = 25

	sendSoFar = incrementSends(sendSoFar, sendToAdd)
	fmt.Println("you've sent", sendSoFar, "messages")
}

func incrementSends(sendSoFar int, sendToAdd int) int {
	sendSoFar = sendSoFar + sendToAdd
	return sendSoFar
}

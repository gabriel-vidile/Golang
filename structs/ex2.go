package main

import (
	"fmt"
)

//Declarando os structs

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func main() {
	test(
		messageToSend{
			message: "Olá adorei conhecer go",
			sender: user{
				name:   "Gabriel",
				number: 231312321312,
			},
			recipient: user{
				name:   "Beatriz",
				number: 313123123123,
			},
		},
	)
}

func test(mToSend messageToSend) {
	if !canSendMessage(mToSend) {
		fmt.Print("Não foi possivel enviar a mensagem")
		return
	}
	fmt.Printf("Mensagem enviada: '%s' de: %v - '%s' \n para  %v - '%s' ",
		mToSend.message,
		mToSend.sender.number,
		mToSend.sender.name,
		mToSend.recipient.number,
		mToSend.recipient.name,
	)

}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false

	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false

	}
	return true
}

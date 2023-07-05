package main

import "fmt"

//Criando um struct que representa um SMS
type messageToSend struct {
	Message     string
	PhoneNumber int
}

func main() {
	test(messageToSend{
		PhoneNumber: 12345565665,
		Message:     "Ola tudo bem com você?",
	})
}

//Criando um "Serviço" que envia a mensagem
func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.Message, m.PhoneNumber)
	fmt.Println("================================")
}

package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func main() {
	test(sender{
		rateLimit: 2,
		user: user{
			name:   "Olav",
			number: 2312312321312,
		},
	})
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("sender number", s.number)
	fmt.Println("sender rate limit", s.rateLimit)
}

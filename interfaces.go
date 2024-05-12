package main

import "fmt"

type Logger interface {
	Info() string
}

type User struct {
	Name string
}

func (t User) Info() string {
	return t.Name
}

func main() {
	user1 := User{"Jon"}

	fmt.Println(user1.Info())

}

package main

import "fmt"

// struct embedding
type Human struct {
	Name    string
	Surname string
}

func (h Human) Introduce() {
	fmt.Printf("Hello! My name is %s\n", h.Name)
}

type Action struct {
	Human
	Action string
}

func (a Action) Do() {
	fmt.Printf("I'm %s %s and I'm doing %s\n", a.Name, a.Surname, a.Action)
}

func main() {
	human := Human{Name: "John", Surname: "Doe"}
	action := Action{Human: human, Action: "running"}
	action.Introduce()
	action.Do()

}

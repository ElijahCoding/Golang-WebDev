package main

import "fmt"

type Cat struct {}

func (c Cat) Speak() {
	fmt.Println("meow")
}

type Dog struct {}

func (d Dog) Speak() {
	fmt.Println("woof")
}

type Husky struct {
	Speaker
}

type SpeakerPrefix struct {
	Speaker
}

func (sp SpeakerPrefix) Speak() {
	fmt.Print("Prefix: ")
	sp.Speaker.Speak()
}

type Speaker interface {
	Speak()
}

func main()  {
	h := Husky{SpeakerPrefix{Cat{}}}
	h.Speak()
}
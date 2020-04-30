package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// seperation of concerns
type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeperator)), 0644)
}

func (p *Persistence) Load(j *Journal, filename string) {
	// ...
}

func (p *Persistence) LoadFromWeb(j *Journal, filename string) {
	// ...
}

func main()  {
	j := Journal{}
	j.AddEntry("test one")
	j.AddEntry("test two")
	fmt.Println(j.String())

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
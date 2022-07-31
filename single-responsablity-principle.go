//Single Responsbility Principle
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Blogs struct {
	count int32
	pages []string
}

func (b *Blogs) String() string {
	return strings.Join(b.pages, "\n")
}

func (b *Blogs) AddPage(page string) {
	b.count += 1
	page = fmt.Sprintf("%d - %s", b.count, page)
	b.pages = append(b.pages, page)
}

type Persistent struct{}

func (p *Persistent) SavePage(b *Blogs, filename string) {
	data := []byte(strings.Join(b.pages, "\n"))
	ioutil.WriteFile(filename, data, 0644)
}

func (p *Persistent) LoadPage(b *Blogs, filename string) {
	data, _ := ioutil.ReadFile(filename)
	b.pages = strings.Split(string(data), "\n")
}

func main() {
	bg := Blogs{}
	bg.AddPage("page one")
	bg.AddPage("page two")

	ps := Persistent{}
	ps.SavePage(&bg, "pages")
	ps.LoadPage(&bg, "pages")
 
	fmt.Println(bg.String())

}

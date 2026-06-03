package author

import "fmt"

type Author struct {
	Name    string
	Contact string
}

func NewAuthor(name, contract string) *Author {
	return &Author{Name: name, Contact: contract}
}

func (a *Author) WriteChapter(chapterTitle, content string) {
	fmt.Printf("Author %s is writing a chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

func (a *Author) ReviewChapter(chapterTitle, content string) {
	fmt.Printf("Author %s is reviewing a chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

func (a *Author) FinalizeChapter(chapterTitle string) {
	fmt.Printf("Author %s is finalised the chapter titled '%s'\n", a.Name, chapterTitle)

}

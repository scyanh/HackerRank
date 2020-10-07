package interfaces

import (
	"fmt"
)

// Inteface to print Title and Author from two objects
type Article struct {
	Title  string
	Author string
}

type Book struct {
	Title  string
	Author string
}

type Stringly interface {
	ConvertToString() string
}

func (article Article) ConvertToString() string {
	return article.Title + " " + article.Author
}
func (book Book) ConvertToString() string {
	return book.Title + " " + book.Author
}


func ObjToString() {
	article := Article{
		Title:  "Hello",
		Author: "World",
	}
	fmt.Println(article.ConvertToString())
	Print(article)
	book := Book{
		Title:  "Book Title",
		Author: "Book Author",
	}
	Print(book)
}
func Print(s Stringly) {
	fmt.Println(s.ConvertToString())
}

// Interface to calculate area from circle and square
type Circle struct {
	Radious float64
}

type Square struct {
	Width float64
}

type DoArea interface  {
	Area() float64
}

func (square Square)Area() float64{
	return square.Width*square.Width
}
func (circle Circle) Area() float64{
	return circle.Radious*circle.Radious
}

func CircleAndSquare()  {
	circle:=Circle{
		Radious: 2,
	}
	fmt.Println(circle.Area())
	square:=Square{
		Width: 2,
	}
	fmt.Println(square.Area())

	Compare(1,circle)
}

func Compare(n float64, area DoArea){
	fmt.Println("totalArea=",n+area.Area())
}

func StartThis() {
	//ObjToString()
	CircleAndSquare()
}

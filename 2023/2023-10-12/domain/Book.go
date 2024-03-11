package domain

type Book struct {
	Title  string
	Author string
}

type BusinessBook struct {
	Book
	BusinessType string
}

func (b *BusinessBook) Author() string {
	return b.Book.Author
}

type ProgrammingBook struct {
	Book
	Language string
}

func (p *ProgrammingBook) Author() string {
	return p.Book.Author
}

//type BookInterface interface {
//	Author() string
//}

type BookInterface interface {
	BusinessBook | ProgrammingBook
}

// GetAuthor BookInterfaceを実装している型を引数に取る
func GetAuthor[T BookInterface](b T) T {
	return b
}

func Authors() {

	// businessBook, programmingBook はそれぞれBookInterfaceを実装している
	businessBook := BusinessBook{
		Book: Book{
			Title:  "Business Book",
			Author: "Business Author",
		},
		BusinessType: "Business",
	}
	programmingBook := ProgrammingBook{
		Book: Book{
			Title:  "Programming Book",
			Author: "Programming Author",
		},
		Language: "Go",
	}

	//businessBookAuthor := GetAuthor(&businessBook)
	//programmingBookAuthor := GetAuthor(&programmingBook)

	book := Book{Author: "Book Author", Title: "Book Title"}

	businessBookAuthor := GetAuthor(businessBook)
	programmingBookAuthor := GetAuthor(programmingBook)
	// Book 型は BookInterface を実装していないのでコンパイルエラー
	// bookAuthor := GetAuthor(book)

	println(businessBookAuthor.Author())
	println(programmingBookAuthor.Author())

}

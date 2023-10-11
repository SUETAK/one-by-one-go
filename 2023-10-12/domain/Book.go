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

type BookInterface interface {
	Author() string
}

func GetAuthor[T BookInterface](b T) string {
	return b.Author()
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

	businessBookAuthor := GetAuthor[BookInterface](&businessBook)
	programmingBookAuthor := GetAuthor[BookInterface](&programmingBook)

	println(businessBookAuthor)
	println(programmingBookAuthor)

}

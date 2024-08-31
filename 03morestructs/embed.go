package morestructs

import "fmt"

type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	Book
	ISBN13 string
}

func (o OreillyBook) GetOreillyURL() string {
	return "https://www.oreilly.co.jp/books/" + o.ISBN13 + "/"
}

func usageExample() {
	ob := OreillyBook{
		ISBN13: "9784873119038",
		Book: Book{
			Title: "Real World https",
		},
	}
	// 埋め込んだ構造体 (Book) のメソッドが呼べる
	fmt.Println(ob.GetAmazonURL())
	// OreillyBook のメソッドが呼べる
	fmt.Println(ob.GetOreillyURL())
	// 明示的な呼び出しも OK
	fmt.Println(ob.Book.GetAmazonURL())
}

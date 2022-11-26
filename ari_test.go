package ari

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAddToCustomStruct(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	if list.Size() != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}
}

func TestAddToCustomStructPointer(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[*Book]{}
	list.Add(&Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(&Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(&Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	if list.Size() != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}
}

func TestAddToCustomStructPointerWithNil(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[*Book]{}
	list.Add(&Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(&Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(&Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(nil)
	if list.Size() != 4 {
		t.Errorf("Expected list size to be 4, got %d", list.Size())
	}
}

func TestFindElement(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	bk := list.Find(func(element Book) bool {
		return element.Title == "The Lord of the Rings"
	})
	if bk == nil {
		t.Errorf("Expected to find a book")
		return
	}
	if bk.Title != "The Lord of the Rings" {
		t.Errorf("Expected to find 'The Lord of the Rings', got '%s'", bk.Title)
	}

}
func TestFindNonExitingElement(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	bk := list.Find(func(element Book) bool {
		return element.Title == "Lord of the flies"
	})
	if bk != nil {
		t.Errorf("Expected to not find a book")
		return
	}
}

func TestEveryElement(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	// Check if all books are by Tolkien
	allTolkien := list.Every(func(element Book) bool {
		return element.Author == "J.R.R. Tolkien"
	})
	if !allTolkien {
		t.Errorf("Expected all books to be by Tolkien")
		return
	}
}

func TestSomeElements(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 2,
	})
	// Check some of the books are a second edition
	someSecondEdition := list.Some(func(element Book) bool {
		return element.Edition == 2
	})
	if !someSecondEdition {
		t.Errorf("Expected some books to be second edition")
		return
	}
}
func TestFailedSomeElements(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 2,
	})
	// Check some of the books are a third edition
	someThirdEdition := list.Some(func(element Book) bool {
		return element.Edition == 3
	})
	if someThirdEdition {
		t.Errorf("Expected no books to be third edition")
		return
	}
}

func TestFilterElements(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 2,
	})
	// Filter out all books that are not second edition
	filtered := list.Filter(func(element Book) bool {
		return element.Edition == 2
	})
	if filtered.Size() != 1 {
		t.Errorf("Expected one book to be second edition")
		return
	}
}

func TestMapElements(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 2,
	})
	// Map all books to their ISBN
	isbns := list.Map(func(element Book) Book {
		return Book{
			ISBN: element.ISBN,
		}
	})
	if isbns.Size() != 3 {
		t.Errorf("Expected three ISBNs")
		return
	}
}

func TestForEach(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 2,
	})
	// print all books' titles
	list.ForEach(func(element Book) {
		fmt.Println(element.Title)
	})
}
func TestForEachAsync(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 2,
	})
	// print all books' titles
	list.ForEachAsync(func(element Book) {
		// genearate a random number between 1 and 5
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(5) + 1
		fmt.Println("sleeping for", randNum, "seconds")
		// sleep for a random amount of time
		time.Sleep(time.Duration(randNum) * time.Second)
		fmt.Println(element.Title)
	})
	t.Log("Finished")
}

func TestReduce(t *testing.T) {
	type Book struct {
		Title   string
		Author  string
		ISBN    string
		Edition int
	}
	list := ArrayList[Book]{}
	list.Add(Book{
		Title:   "The Hobbit",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Lord of the Rings",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 1,
	})
	list.Add(Book{
		Title:   "The Silmarillion",
		Author:  "J.R.R. Tolkien",
		ISBN:    "9780547928227",
		Edition: 2,
	})
	// reduce all books to a single string
	str := ""
	books := Reduce(list.ToArray(),
		func(accumulator string, element Book) string {
			return accumulator + element.Title + ", "
		}, str)
	if books != "The Hobbit, The Lord of the Rings, The Silmarillion, " {
		t.Errorf("Expected all books to be reduced")
		return
	}
}

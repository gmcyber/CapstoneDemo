package main
import (
  "database/sql"
  "log"
  _ "github.com/mattn/go-sqlite3"
)

type Book struct {
  id int
  name string
  author string
}

func main() {
  db, err := sql.Open("sqlite3", "./books.db")
  log.Println(db)
  if err != nil {
    log.Println(err)
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
  if err != nil {
    log.Println("Error creating books")
  } else {
    log.Println("Created Books")
  }
  statement.Exec()
  statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
  statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
  log.Println("Inserted a book")
  rows, _ := db.Query("SELECT id, name, author from books")
  var tempBook Book
  for rows.Next() {
    rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
    log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
  }

  statement, _ = db.Prepare("update books set name=? where id=?")
  statement.Exec("The Tale of Two Cities", 1)
  log.Println("Update")

  statement, _ = db.Prepare("delete from books where id = ?")
  statement.Exec(1)
  log.Println("Deleted")



}

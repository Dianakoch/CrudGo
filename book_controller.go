package main

func createBook(Book Book) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO book (name, genre, year) VALUES (?, ?, ?)", Book.Name, Book.Genre, Book.Year)
	return err
}

func deleteBook(id int64) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM book WHERE id = ?", id)
	return err
}

func updateBook(Book Book) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE book SET name = ?, genre = ?, year = ? WHERE id = ?", Book.Name, Book.Genre, Book.Year, Book.Id)
	return err
}
func getBooks() ([]Book, error) {
	Books := []Book{}
	bd, err := getDB()
	if err != nil {
		return Books, err
	}
	rows, err := bd.Query("SELECT id, name, genre, year FROM book")
	if err != nil {
		return Books, err
	}
	for rows.Next() {
		var Book Book
		err = rows.Scan(&Book.Id, &Book.Name, &Book.Genre, &Book.Year)
		if err != nil {
			return Books, err
		}
		Books = append(Books, Book)
	}
	return Books, nil
}

func getBookById(id int64) (Book, error) {
	var Book Book
	bd, err := getDB()
	if err != nil {
		return Book, err
	}
	row := bd.QueryRow("SELECT id, name, genre, year FROM book WHERE id = ?", id)
	err = row.Scan(&Book.Id, &Book.Name, &Book.Genre, &Book.Year)
	if err != nil {
		return Book, err
	}
	return Book, nil
}

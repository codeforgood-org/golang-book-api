package models

// Book represents a book in the library
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Validate checks if the book data is valid
func (b *Book) Validate() error {
	if b.Title == "" {
		return ErrInvalidTitle
	}
	if b.Author == "" {
		return ErrInvalidAuthor
	}
	return nil
}

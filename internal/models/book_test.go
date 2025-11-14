package models

import "testing"

func TestBook_Validate(t *testing.T) {
	tests := []struct {
		name    string
		book    Book
		wantErr error
	}{
		{
			name: "valid book",
			book: Book{
				Title:  "Test Book",
				Author: "Test Author",
			},
			wantErr: nil,
		},
		{
			name: "missing title",
			book: Book{
				Title:  "",
				Author: "Test Author",
			},
			wantErr: ErrInvalidTitle,
		},
		{
			name: "missing author",
			book: Book{
				Title:  "Test Book",
				Author: "",
			},
			wantErr: ErrInvalidAuthor,
		},
		{
			name: "missing both",
			book: Book{
				Title:  "",
				Author: "",
			},
			wantErr: ErrInvalidTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.book.Validate()
			if err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

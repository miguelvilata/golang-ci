package library

// Book type
type Book struct {
	Title       string `json:"title" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Year        string `json:"year" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

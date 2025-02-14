package main

type Status string //we can create our cuttom type
// in go like the enum
const (
	Read    Status = "read"
	Reading Status = "reading"
	ToRead  Status = "to_read"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"` //prevents marshalling to JSON
}

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status Status `json:"status"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	UserID int    `json:"userId"`
}

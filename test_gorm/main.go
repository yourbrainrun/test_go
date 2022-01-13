package main

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	Id     int          `gorm:"column:id"`
	Name   string       `gorm:"column:name"`
	Create sql.NullTime `gorm:"column:create"`
}

func main() {
	user := User{Id: 1, Name: "hua", Create: sql.NullTime{
		Time:  time.Now(),
		Valid: false,
	}}

	fmt.Println(user)
}

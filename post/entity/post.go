package post_entity

import (
	"fmt"
	"github.com/ujwaldhakal/go-blog-example/db"
	"time"
)

type Post struct {
	ID          uint
	Title        string
	Description  string
	UserId    	string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Post) Create() error {
	con := db.GetConnection()

	fmt.Println("upto here",p)
	result := con.Create(p)

	return result.Error
}
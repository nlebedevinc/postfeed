package models

import "github.com/google/uuid"

type Post struct {
	Id     string `json:"UID"`
	Author string `json:"author"`
	Post   string `json:"post"`
}

func (p Post) Key() string {
	return "post:" + p.Id
}

func (p *Post) GenerateId() string {
	p.Id = uuid.New().String()
	return p.Id
}

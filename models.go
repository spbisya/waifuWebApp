package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Waifu struct {
	Id        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
  Characters string `db:"characters" json:"characters"`
}

type Greeting struct{
  Id int64 `db:"id" json:"id"`
  Characters string `db:"characters" json:"characters"`
  Texts string `db:"texts" json:"texts"`
}

type Accost struct {
  Id int64 `db:"id" json:"id"`
  Characters string `db:"characters" json:"characters"`
  Texts string `db:"texts" json:"texts"`
}

type Question struct {
  Id int64 `db:"id" json:"id"`
  Characters string `db:"characters" json:"characters"`
  Texts string `db:"texts" json:"texts"`
}

type User struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Email string `json:"email"`
	Waifu string `json:"waifu"`
}

type Token struct {
	User string `json:"user"`
	Token string `json:"token"`
}

type Post struct {
  Id int64 `db:"id" json:"id"`
  Title string `db:"title" json:"title"`
  Text string `db:"text" json:"text"`
	CreatedAt string `db:"created" json:"created"`
	Truncated bool `db:"truncated" json:"truncated"`
	Tags string `db:"tags" json:"tags"`
	Summary string `db:"summary" json:"summary"`
}

type PostForSite struct{
	Post Post
	Tags []string
}

type PostForSingle struct{
	Post Post
	Tags []string
	Right bool
	Left bool
	RightId int64
	LeftId int64
}

type Page struct {
	Number int
	Current bool
}

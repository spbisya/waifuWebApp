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

package models

type Contact struct {
    Id    int    `json:"id" form:"id"`
    First string `json:"first" form:"first"`
    Last  string `json:"last" form:"last"`
    Phone string `json:"phone" form:"phone"`
    Email string `json:"email" form:"email"`
}

package dto

type UserDto struct{
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}
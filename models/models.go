package models

import "github.com/lib/pq"

type User struct {
	ID     uint           `json:"id"`
	Name   string         `json:"name"`
	Orders pq.StringArray `json:"orders"`
}

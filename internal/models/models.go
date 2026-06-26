package models

import "time"

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

type Ticket struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
    OwnerID     int       `json:"owner_id"`
    CreatedAt   time.Time `json:"created_at"`
}

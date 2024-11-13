package bookservices

import "time"

type Book struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type BookUpdateRequest struct {
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

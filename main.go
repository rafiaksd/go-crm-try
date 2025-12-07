package main

import (
	"fmt"
	"time"
)

type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
}

type CRM struct {
	customers map[int]Customer
	NextID    int
}

func NewCRM() *CRM {
	return &CRM{customers: make(map[int]Customer), NextID: 1}
}

func main() {
	fmt.Println("First push")
}

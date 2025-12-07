package main

import (
	"fmt"
	"strconv"
	"strings"
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

func (crm *CRM) addCustomer(firstname, lastname, email, phone string) {
	cust := Customer{
		ID:        crm.NextID,
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
	}

	crm.customers[crm.NextID] = cust
	crm.NextID++
	fmt.Printf("Customer {%s} - {%s} added successfully\n", strconv.Itoa(cust.ID), cust.FirstName)
}

func (crm CRM) ViewCustomer(id int) {
	customer, exists := crm.customers[id]

	if !exists {
		fmt.Println("Customer not found!")
		return
	}

	fmt.Println("Customer ID: ", customer.ID)
	fmt.Println("Name: ", customer.FirstName, customer.LastName)
	fmt.Println("Email:", customer.Email)
	fmt.Println("Created at:", customer.CreatedAt.Format("2006-01-02 15:04")) // 2 January 2006 3:04 pm
}

func (crm CRM) ViewAllCustomers() {
	if len(crm.customers) == 0 {
		fmt.Println("No customers foudn")
		return
	}

	for _, cust := range crm.customers {
		fmt.Printf("\nID: %d Name: %s %s Email: %s Phone: %s\n", cust.ID, cust.FirstName, cust.LastName, cust.Email, cust.Phone)
	}
}

func (crm *CRM) UpdateCustomer(id int, firstName, lastName, email, phone string) {
	customer, exists := crm.customers[id]

	if !exists {
		fmt.Printf("Customer with ID:%d, dont exist\n", id)
		return
	}

	customer.FirstName = firstName
	customer.LastName = lastName
	customer.Email = email
	customer.Phone = phone

	// save updated customer
	crm.customers[id] = customer

	fmt.Printf("Customer - ID: %d - updated successfully\n", id)
}

func (crm *CRM) DeleteCustomer(id int) {
	_, exists := crm.customers[id]
	if !exists {
		fmt.Println("Customer not found")
		return
	}

	delete(crm.customers, id)
	fmt.Println("Customer deleted successfully")
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	crm := NewCRM()

	// main loop for interacting with the CRM

	// make it run indefinitely
	for {
		fmt.Println("\nCRM System Menu:")
		fmt.Println("1. Add Customer")
		fmt.Println("2. View Customer")
		fmt.Println("3. View All Customers")
		fmt.Println("4. Update Customer")
		fmt.Println("5. Delete Customer")
		fmt.Println("6. Exit")

		option := getInput("Choose an option (1-6)")

		switch strings.TrimSpace(option) {
		case "1":
			// add customer
			firstName := getInput("Enter cust. first name:")
			last := getInput("Enter cust. last name:")
			e := getInput("Enter email:")
			num := getInput("Enter phone:")
			crm.addCustomer(firstName, last, e, num)
		case "2":
			// view customer of id
			id := getInput("Enter customer ID to view:")

			customerID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid customer id")
				break
			}

			crm.ViewCustomer(customerID)
		case "3":
			crm.ViewAllCustomers()
		case "4":
			//update customer
			id := getInput("Enter customer id")
			customerID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid customer id")
				break
			}

			first := getInput("Enter first name:")
			last := getInput("Enter last name:")
			email := getInput("Enter email:")
			num := getInput("Enter phone num:")

			crm.UpdateCustomer(customerID, first, last, email, num)
		case "5":
			// Delete customer
			id := "Enter Customer iD to DeLeTe:"
			customerID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Invalid customer id")
				break
			}

			crm.DeleteCustomer(customerID)
		case "6":
			fmt.Println("Exiting CRM System!")
			return
		default:
			fmt.Println("Invalid option! Please choose between 1-6")
		}
	}
}

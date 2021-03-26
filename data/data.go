package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"workspace/model"
)

func GetUsersFromFile() []*model.User {
	var users []*model.User
	userJsonFile, userErr := os.Open("data/users.json")

	if userErr != nil{
		fmt.Println("Lỗi khi mở file nè")
	}
	userByteVal, _ := ioutil.ReadAll(userJsonFile)
	json.Unmarshal([]byte(userByteVal), &users)
	defer userJsonFile.Close()
	return users
}

func GetTicketsFromFile() []*model.Ticket {
	var tickets []*model.Ticket
	ticketJsonFile, ticketErr := os.Open("data/tickets.json")

	if ticketErr != nil{
		fmt.Println("Lỗi khi mở file nè")
	}
	userByteVal, _ := ioutil.ReadAll(ticketJsonFile)
	json.Unmarshal([]byte(userByteVal), &tickets)
	defer ticketJsonFile.Close()
	return tickets
}

func GetOrganizationsFromFile() []*model.Organization {
	var organizations []*model.Organization
	organizationJsonFile, organizationErr := os.Open("data/organizations.json")

	if organizationErr != nil{
		fmt.Println("Lỗi khi mở file nè")
	}
	userByteVal, _ := ioutil.ReadAll(organizationJsonFile)
	json.Unmarshal([]byte(userByteVal), &organizations)
	defer organizationJsonFile.Close()
	return organizations
}
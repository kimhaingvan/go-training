package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"workspace/model"
)

var usersFilePath = "data/users.json"
var ticketsFilePath = "data/tickets.json"
var organizationsFilePath = "data/organizations.json"

func GetUsersFromFile() []*model.User {
	var users []*model.User
	userJsonFile, err := os.Open(usersFilePath)

	if err != nil {
		fmt.Println("Lỗi khi mở file nè")
	}

	userByteVal, _ := ioutil.ReadAll(userJsonFile)
	defer userJsonFile.Close()
	json.Unmarshal([]byte(userByteVal), users)
	return users
}

func GetTicketsFromFile() []*model.Ticket {
	var tickets []*model.Ticket
	ticketJsonFile, err := os.Open(ticketsFilePath)

	if err != nil {
		fmt.Println("Lỗi khi mở file nè")
	}

	ticketByteVal, _ := ioutil.ReadAll(ticketJsonFile)
	defer ticketJsonFile.Close()
	json.Unmarshal([]byte(ticketByteVal), tickets)
	return tickets
}

func GetOrganizationsFromFile() []*model.Organization {
	var organizations []*model.Organization
	organizationJsonFile, err := os.Open(organizationsFilePath)

	if err != nil {
		fmt.Println("Lỗi khi mở file nè")
	}

	organizationByteVal, _ := ioutil.ReadAll(organizationJsonFile)
	defer organizationJsonFile.Close()
	json.Unmarshal([]byte(organizationByteVal), organizations)
	return organizations
}
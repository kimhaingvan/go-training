package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"workspace/data"
	"workspace/model"
	"workspace/service"
)

func main() {
	var users, tickets, organizations = getAndMappingData()
	reader := bufio.NewReader(os.Stdin)
	for {
		service.ClearTerminal()
		optionStr := service.DisplayAndReadOption()
		switch optionStr {
		case "1":
			service.ClearTerminal()
			dataTypeStr := service.DisplayAndReadDataType()
			field, value := service.ReadSearchFieldAndVal()
			service.ClearTerminal()
			switch dataTypeStr {
			case "1":
				users := service.SearchUsers(value, field, users)
				for i, user := range users {
					fmt.Printf(`Result %v: 
%v
-----------------------------------------------------------`, i, user)
				}
			case "2":
				tickets := service.SearchTickets(value, field, tickets)
				for i, ticket := range tickets {
					fmt.Printf(`Result %v: 
%v
-----------------------------------------------------------`, i, ticket)
				}
			case "3":
				organizations := service.SearchOrganizations(value, field, organizations)
				for i, organization := range organizations {
					fmt.Printf(`Result %v: 
%v
-----------------------------------------------------------`, i, organization)
				}
			}
			enter, _ := reader.ReadString('\n')
			service.ClearTerminal()
			fmt.Print(enter)
		case "2":
			service.ClearTerminal()
			service.DisplaySearchableFields()
			enter, _ := reader.ReadString('\n')
			service.ClearTerminal()
			fmt.Print(enter)
		case "quit":
			break
		default:
			fmt.Println("default naf")
		}

		if strings.Compare(optionStr, "quit") == 0{
			break
		}
	}

}

func getAndMappingData() ([]*model.User, []*model.Ticket, []*model.Organization ){
	var users []*model.User = data.GetUsersFromFile()
	var tickets []*model.Ticket = data.GetTicketsFromFile()
	var organizations []*model.Organization = data.GetOrganizationsFromFile()

	service.MappingData(tickets, organizations, users)
	return users, tickets, organizations
}

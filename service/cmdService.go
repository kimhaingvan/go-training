package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"workspace/model"
)

func ReadSearchFieldAndVal() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(`Enter search field: `)
	searchField, _ := reader.ReadString('\n')
	searchField = strings.Replace(searchField, "\n", "", -1)

	fmt.Printf(`Enter search value: `)
	searchValue, _ := reader.ReadString('\n')
	searchValue = strings.Replace(searchValue, "\n", "", -1)
	return searchField, searchValue
}
func DisplayAndReadDataType() string {
	fmt.Printf(`Select: 1) Users  or  2) Tickets  or  3) Organizations
--------------------------------------------
Enter your option: `)
	reader := bufio.NewReader(os.Stdin)
	optionStr, _ := reader.ReadString('\n')
	optionStr = strings.Replace(optionStr, "\n", "", -1)
	return optionStr
}
func DisplayAndReadOption() string {
	DisplayOptions()
	reader := bufio.NewReader(os.Stdin)
	optionStr, _ := reader.ReadString('\n')
	optionStr = strings.Replace(optionStr, "\n", "", -1)
	return optionStr
}
func DisplayOptions() {
	fmt.Printf(`Select search options:
* Press 1 to search
* Press 2 to view a list of searchable fields
* Type 'quit' to quit
--------------------
Enter your option: `)
}
func DisplaySearchableFields() {
	fmt.Printf(`-----------------------------------------------------------------
Search Users with: `)
	var user model.User
	user.DisplayUserFields()
	fmt.Printf(`-----------------------------------------------------------------
Search Tickets with: `)
	var ticket model.Ticket
	ticket.DisplayTicketFields()

	fmt.Printf(`-----------------------------------------------------------------
Search Organizations with: `)
	var organization model.Organization
	organization.DisplayOrganizationFields()
}
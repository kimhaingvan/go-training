package service

import (
	"fmt"
	"strconv"
	"workspace/model"
)

func MappingData(tickets []*model.Ticket, organizations []*model.Organization, users []*model.User) {
	userMap := make(map[int]*model.User)
	for  _, user := range users {
		userMap[user.Id] = user
	}

	ticketMap := make(map[string]*model.Ticket)
	for _, ticket := range tickets {
		ticketMap[ticket.Id] = ticket
	}

	organizationMap := make(map[int]*model.Organization)
	for _, organization := range organizations {
		organizationMap[organization.Id] = organization
	}

	for _, ticket := range tickets {
		ticket.Submitter = userMap[ticket.SubmitterId]
		ticket.Assignee = userMap[ticket.AssigneeId]
		ticket.Organization = organizationMap[ticket.OrganizationId]
	}

	for _, user := range users {
		user.Organization = organizationMap[user.OrganizationId]
	}
}

func SearchTickets(textSearch, fieldSearch string, tickets []*model.Ticket) []*model.Ticket {
	var ticketsRes []*model.Ticket
	switch fieldSearch {
	case "_id":
		for _, user := range tickets {
			if user.Id == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "url":
		for _, user := range tickets {
			if user.Url == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "external_id":
		for _, user := range tickets {
			if user.ExternalId == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "created_at":
		for _, user := range tickets {
			if user.CreatedAt == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "type":
		for _, user := range tickets {
			if user.Type == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "subject":
		for _, user := range tickets {
			if user.Subject == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "description":
		for _, user := range tickets {
			if user.Description == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "priority":
		for _, user := range tickets {
			if user.Priority == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "status":
		for _, user := range tickets {
			if user.Status == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "submitterId":
		searchVal, _ := strconv.Atoi(textSearch)
		for _, user := range tickets {
			if user.SubmitterId == searchVal {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "assgigneeId":
		searchVal, _ := strconv.Atoi(textSearch)
		for _, user := range tickets {
			if user.AssigneeId == searchVal {
				ticketsRes = append(ticketsRes, user)
			}
		}

	case "organizationId":
		searchVal, _ := strconv.Atoi(textSearch)
		for _, user := range tickets {
			if user.OrganizationId == searchVal {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "DueAt":
		for _, user := range tickets {
			if user.DueAt == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	case "via":
		for _, user := range tickets {
			if user.Via == textSearch {
				ticketsRes = append(ticketsRes, user)
			}
		}
	default:
		fmt.Println("Error: Field text is not valid.")
	}
	return ticketsRes
}
func SearchOrganizations(textSearch, fieldSearch string, organizations []*model.Organization) []*model.Organization {
	var organizationRes []*model.Organization
	switch fieldSearch {
	case "_id":
		searchVal, _ := strconv.Atoi(textSearch)
		for _, user := range organizations {
			if user.Id == searchVal {
				organizationRes = append(organizationRes, user)
			}
		}
	case "url":
		for _, user := range organizations {
			if user.Url == textSearch {
				organizationRes = append(organizationRes, user)
			}
		}
	case "external_id":
		for _, user := range organizations {
			if user.ExternalId == textSearch {
				organizationRes = append(organizationRes, user)
			}
		}
	case "shared_tickets":
		searchVal, _ := strconv.ParseBool(textSearch)
		for _, user := range organizations {
			if user.SharedTickets == searchVal {
				organizationRes = append(organizationRes, user)
			}
		}
	case "name":
		for _, user := range organizations {
			if user.Name == textSearch {
				organizationRes = append(organizationRes, user)
			}
		}
	case "domain_names":
		for _, user := range organizations {
			if user.Name == textSearch {
				organizationRes = append(organizationRes, user)
			}
		}
	case "created_at":
		for _, user := range organizations {
			if user.CreatedAt == textSearch {
				organizationRes = append(organizationRes, user)
			}
		}
	case "details":
		for _, user := range organizations {
			if user.Details == textSearch {
				organizationRes = append(organizationRes, user)
			}
		}
	default:
		fmt.Println("Error: Field text is not valid.")
	}
	return organizationRes
}
func SearchUsers(textSearch, fieldSearch string, users []*model.User) []*model.User {
	var usersRes []*model.User
	switch fieldSearch {
	case "_id":
		searchVal, _ := strconv.Atoi(textSearch)
		for _, user := range users {
			if user.Id == searchVal {
				usersRes = append(usersRes, user)
			}
		}
	case "url":
		for _, user := range users {
			if user.Url == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "external_id":
		for _, user := range users {
			if user.ExternalId == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "name":
		for _, user := range users {
			if user.Name == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "alias":
		for _, user := range users {
			if user.Alias == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "created_at":
		for _, user := range users {
			if user.CreatedAt == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "active":
		searchVal, _ := strconv.ParseBool(textSearch)
		for _, user := range users {
			if user.Active == searchVal {
				usersRes = append(usersRes, user)
			}
		}
	case "shared":
		searchVal, _ := strconv.ParseBool(textSearch)
		for _, user := range users {
			if user.Shared == searchVal {
				usersRes = append(usersRes, user)
			}
		}
	case "verified":
		searchVal, _ := strconv.ParseBool(textSearch)
		for _, user := range users {
			if user.Verified == searchVal {
				usersRes = append(usersRes, user)
			}
		}
	case "locale":
		for _, user := range users {
			if user.Locale == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "timezone":
		for _, user := range users {
			if user.Timezone == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "last_login_at":
		for _, user := range users {
			if user.LastLoginAt == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "email":
		for _, user := range users {
			if user.Email == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "phone":
		for _, user := range users {
			if user.Phone == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "signature":
		for _, user := range users {
			if user.Signature == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	case "organization_id":
		searchVal, _ := strconv.Atoi(textSearch)
		for _, user := range users {
			if user.OrganizationId == searchVal {
				usersRes = append(usersRes, user)
			}
		}
	case "suspended":
		searchVal, _ := strconv.ParseBool(textSearch)
		for _, user := range users {
			if user.Suspended == searchVal {
				usersRes = append(usersRes, user)
			}
		}
	case "role":
		for _, user := range users {
			if user.Role == textSearch {
				usersRes = append(usersRes, user)
			}
		}
	default:
		fmt.Printf(`Error: Field text is not valid.
=> Enter to continue `)
	}
	return usersRes
}

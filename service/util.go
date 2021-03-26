package service

import "workspace/model"

func MappingData(tickets []*model.Ticket, organizations []*model.Organization, users []*model.User) {
	userMap := make(map[int]*model.User)
	for _, user := range users {
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
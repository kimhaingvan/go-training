package model

import (
	"fmt"
)

type User struct {
	Id             int      `json:"_id"`
	Url            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	Organization   *Organization
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`

}

func (u User) String() string {
	return fmt.Sprintf(`id: %v
name: %v
alias: %v
CreatedAt: %v
Active: %v
Verified: %v
Shared: %v
Locale: %v
Timezone: %v
LastLoginAt: %v
Email: %v
Phone: %v
Signature: %v
Organization_name: %v
Suspended: %v
Role: %v
`, u.Id, u.Name, u.Alias, u.CreatedAt, u.Active, u.Verified, u.Shared, u.Locale, u.Timezone, u.LastLoginAt, u.Email, u.Phone, u.Signature, u.Organization.Name, u.Suspended, u.Role)
}

func (u User) DisplayUserFields() {
	fmt.Printf(`_id
url
external_id
name
alias
created_at
active
verified
shared
locale
timezone
last_login_at
email
phone
signature
organization_id
tags
suspended
role
`)
}


type Ticket struct {
	Id             string   `json:"_id"`
	Url            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterId    int      `json:"submitter_id"`
	Submitter	   *User
	Assignee       *User
	AssigneeId     int      `json:"assignee_id"`
	Organization   *Organization
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}


func (t Ticket) String() string {
	return fmt.Sprintf(`id: %v
url: %v
Organization_name: %v
Assignee_name: %v
Submitter_name: %v
ExternalId: %v
CreatedAt: %v
Type: %v
Subject: %v
Description: %v
Priority: %v
Status: %v
HasIncidents: %v
DueAt: %v
Via: %v
`, t.Id, t.Organization.Name, t.Assignee.Name, t.Submitter.Name, t.ExternalId, t.CreatedAt, t.Type, t.Subject, t.Description, t.Priority, t.Status, t.HasIncidents, t.DueAt, t.Via)

}
func (t Ticket) DisplayTicketFields()  {
	fmt.Printf(`_id
url
external_id
type
created_at
subject
description
priority
status
submitter_id
assignee_id
organization_id
tags
has_incidents
due_at
via
`)
}

type Organization struct {
	Id            int      `json:"_id"`
	Url           string   `json:"url"`
	ExternalId    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	Tags          []string `json:"tags"`
	SharedTickets bool     `json:"shared_tickets"`
}


func (o Organization) String() string {
	return fmt.Sprintf(`id: %v
url: %v
ExternalId: %v
Name: %v
DomainNames: %v
CreatedAt: %v
Details: %v
SharedTickets: %v
Tags: %v
`, o.Id, o.Url, o.ExternalId, o.Name, o.DomainNames, o.CreatedAt, o.SharedTickets, o.Tags)
}


func (o Organization) DisplayOrganizationFields()  {
	fmt.Printf(`_id                                        
url                                        
external_id                                 
name                                       
domain_names                              
created_at                               
details                                   
shared_tickets                          
tags
`)
}

func MappingData(tickets []*Ticket , organizations []*Organization, users []*User) {
	userMap := make(map[int]*User)
	for  _, user := range users {
		userMap[user.Id] = user
	}

	ticketMap := make(map[string]*Ticket)
	for _, ticket := range tickets {
		ticketMap[ticket.Id] = ticket
	}

	organizationMap := make(map[int]*Organization)
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
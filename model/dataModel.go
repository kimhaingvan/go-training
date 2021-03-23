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
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

func (u User) String() string {
	return fmt.Sprintf(`id: %v
name: %v

`, u.Id, u.Name)
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
	AssigneeId     int      `json:"assignee_id"`
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

func (t Ticket)DisplayTicketFields()  {
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
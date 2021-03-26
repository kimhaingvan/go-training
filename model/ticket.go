package model

import "fmt"

type Ticket struct {
	Id             string `json:"_id"`
	Url            string `json:"url"`
	ExternalId     string `json:"external_id"`
	CreatedAt      string `json:"created_at"`
	Type           string `json:"type"`
	Subject        string `json:"subject"`
	Description    string `json:"description"`
	Priority       string `json:"priority"`
	Status         string `json:"status"`
	SubmitterId    int    `json:"submitter_id"`
	Submitter      *User
	Assignee       *User
	AssigneeId     int `json:"assignee_id"`
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
func (t Ticket) DisplayTicketFields() {
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


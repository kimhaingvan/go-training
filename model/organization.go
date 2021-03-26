package model

import "fmt"

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

func (o Organization) DisplayOrganizationFields() {
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


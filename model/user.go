package model

import "fmt"

type User struct {
	Id             int    `json:"_id"`
	Url            string `json:"url"`
	ExternalId     string `json:"external_id"`
	Name           string `json:"name"`
	Alias          string `json:"alias"`
	CreatedAt      string `json:"created_at"`
	Active         bool   `json:"active"`
	Verified       bool   `json:"verified"`
	Shared         bool   `json:"shared"`
	Locale         string `json:"locale"`
	Timezone       string `json:"timezone"`
	LastLoginAt    string `json:"last_login_at"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Signature      string `json:"signature"`
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



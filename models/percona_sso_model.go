// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"database/sql/driver"
	"time"

	"gopkg.in/reform.v1"
)

//go:generate reform

// PerconaSSODetails stores everything what we need to issue access_token from Percona SSO API.
// It is intended to have only one row in this table as PMM can be connected to Portal only once.
//reform:percona_sso_details
type PerconaSSODetails struct {
	ClientID       string                 `reform:"client_id,pk"`
	ClientSecret   string                 `reform:"client_secret"`
	IssuerURL      string                 `reform:"issuer_url"`
	Scope          string                 `reform:"scope"`
	AccessToken    *PerconaSSOAccessToken `reform:"access_token"`
	OrganizationID string                 `reform:"organization_id"`

	CreatedAt time.Time `reform:"created_at"`
}

// PerconaSSODetailsInsert stores everything what we can set. Other properties missing compared to
// PerconaSSODetails will be added automatically.
type PerconaSSODetailsInsert struct {
	ClientID       string
	ClientSecret   string
	IssuerURL      string
	Scope          string
	OrganizationID string
}

// PerconaSSOAccessToken represents structure for special access token options.
type PerconaSSOAccessToken struct {
	TokenType   string    `json:"token_type"`
	ExpiresIn   int32     `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
	AccessToken string    `json:"access_token"`
	Scope       string    `json:"scope"`
}

// Value implements database/sql/driver.Valuer interface. Should be defined on the value.
func (t PerconaSSOAccessToken) Value() (driver.Value, error) { return jsonValue(t) }

// Scan implements database/sql.Scanner interface. Should be defined on the pointer.
func (t *PerconaSSOAccessToken) Scan(src interface{}) error { return jsonScan(t, src) }

// BeforeInsert implements reform.BeforeInserter interface.
func (s *PerconaSSODetails) BeforeInsert() error {
	s.CreatedAt = Now().UTC()
	return nil
}

// BeforeUpdate implements reform.BeforeUpdater interface.
func (s *PerconaSSODetails) BeforeUpdate() error {
	s.CreatedAt = Now().UTC()
	return nil
}

// check interfaces.
var (
	_ reform.BeforeInserter = (*PerconaSSODetails)(nil)
	_ reform.BeforeUpdater  = (*PerconaSSODetails)(nil)
)

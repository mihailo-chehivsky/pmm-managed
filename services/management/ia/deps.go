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

package ia

import (
	"context"

	"github.com/percona/pmm/api/alertmanager/ammodels"
)

//go:generate mockery -name=alertManager -case=snake -inpkg -testonly
//go:generate mockery -name=vmAlert -case=snake -inpkg -testonly

// alertManager is is a subset of methods of alertmanager.Service used by this package.
// We use it instead of real type for testing and to avoid dependency cycle.
type alertManager interface {
	GetAlerts(ctx context.Context) ([]*ammodels.GettableAlert, error)
	Silence(ctx context.Context, ids []string) error
	SilenceAll(ctx context.Context) error
	Unsilence(ctx context.Context, ids []string) error
	UnsilenceAll(ctx context.Context) error
	RequestConfigurationUpdate()
}

// vmAlert is is a subset of methods of vmalert.Service used by this package.
// We use it instead of real type for testing and to avoid dependency cycle.
type vmAlert interface {
	RequestConfigurationUpdate()
}

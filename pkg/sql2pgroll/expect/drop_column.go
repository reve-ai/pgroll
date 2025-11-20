// SPDX-License-Identifier: Apache-2.0

package expect

import (
	"github.com/reve-ai/pgroll/pkg/migrations"
	"github.com/reve-ai/pgroll/pkg/sql2pgroll"
)

var DropColumnOp1 = &migrations.OpDropColumn{
	Table:  "foo",
	Column: "bar",
	Down:   sql2pgroll.PlaceHolderSQL,
}

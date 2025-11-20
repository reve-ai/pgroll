// SPDX-License-Identifier: Apache-2.0

package expect

import "github.com/reve-ai/pgroll/pkg/migrations"

var RenameConstraintOp1 = &migrations.OpRenameConstraint{
	Table: "foo",
	From:  "bar",
	To:    "baz",
}

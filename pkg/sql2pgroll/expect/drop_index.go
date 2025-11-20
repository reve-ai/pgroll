// SPDX-License-Identifier: Apache-2.0

package expect

import (
	"github.com/reve-ai/pgroll/pkg/migrations"
)

var DropIndexOp1 = &migrations.OpDropIndex{
	Name: "foo",
}

var DropIndexOp2 = &migrations.OpDropIndex{
	Name: "myschema.foo",
}

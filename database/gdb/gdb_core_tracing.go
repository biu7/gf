// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//

package gdb

import (
	"context"
)

const (
	tracingInstrumentName       = "github.com/gogf/gf/database/gdb"
	tracingAttrDbType           = "db.type"
	tracingAttrDbHost           = "db.host"
	tracingAttrDbPort           = "db.port"
	tracingAttrDbName           = "db.name"
	tracingAttrDbUser           = "db.user"
	tracingAttrDbLink           = "db.link"
	tracingAttrDbGroup          = "db.group"
	tracingEventDbExecution     = "db.execution"
	tracingEventDbExecutionSql  = "db.execution.sql"
	tracingEventDbExecutionCost = "db.execution.cost"
	tracingEventDbExecutionType = "db.execution.type"
)

// addSqlToTracing adds sql information to tracer if it's enabled.
func (c *Core) addSqlToTracing(ctx context.Context, sql *Sql) {}

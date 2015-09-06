// Copyright 2015 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tidb

import (
	"strings"
	"errors"

	"github.com/go-xorm/core"
)

var (
	_ core.Dialect = (*tidb)(nil)

	DBType core.DbType = "tidb"
)

func init() {
	core.RegisterDriver(string(DBType), &tidbDriver{})
	core.RegisterDialect(DBType, func() core.Dialect {
		return &tidb{}
	})
}

type tidbDriver struct {
}

func (p *tidbDriver) Parse(driverName, dataSourceName string) (*core.Uri, error) {
	params := strings.Split(dataSourceName, "://")
	if len(params) < 2 {
		return nil, errors.New("params error")
	}

	if params[0] != "goleveldb" {
		return nil, errors.New(params[0] + " is not supported yet.")
	}

	uri := &core.Uri{
		DbType: DBType,
		DbName: params[1],
	}

	return uri, nil
}

tidb driver and dialect for github.com/go-xorm/xorm
========

Currently, we can support tidb for allmost all the operations.

# How to use

Just like other supports of xorm, but you should import the three packages:

```Go
import (
    _ "github.com/pingcap/tidb"
    _ "github.com/go-xorm/tidb"
    "github.com/go-xorm/xorm"
)

// for goleveldb as store
xorm.NewEngine("tidb", "goleveldb://./tidb.db")
// for memory as store
xorm.NewEngine("tidb", "memory://tidb")
// for boltdb as store
xorm.NewEngine("tidb", "boltdb://./tidb.db")
```
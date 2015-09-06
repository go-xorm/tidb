tidb driver and dialect for github.com/go-xorm/xorm
========

** experiment support **

STATUS: build pass but some tests failed.

Currently, we can support tidb for some operations.

# How to use

Just like other supports of xorm, but you should import the three packages:

Since github.com/cznic/ql# has not been resolved, we just use github.com/lunny/ql instead.

```Go
import (
    _ "github.com/pingcap/tidb"
    _ "github.com/go-xorm/tidb"
    "github.com/go-xorm/xorm"
)

// for open a file
xorm.NewEngine("tidb", "goleveldb://./tidb.db")
```
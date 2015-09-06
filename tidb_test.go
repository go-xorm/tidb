package tidb

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/tests"
	"github.com/go-xorm/xorm"
	_ "github.com/pingcap/tidb"
)

var showTestSql = true

func newTidbEngine() (*xorm.Engine, error) {
	os.Remove("./tidb.db")
	return xorm.NewEngine("tidb", "goleveldb://./tidb.db")
}

func newTidbDriverDB() (*sql.DB, error) {
	os.Remove("./tidb.db")
	return sql.Open("tidb", "goleveldb://./tidb.db")
}

func newCache() core.Cacher {
	return xorm.NewLRUCacher2(xorm.NewMemoryStore(), time.Hour, 1000)
}

func setEngine(engine *xorm.Engine, useCache bool) {
	if useCache {
		engine.SetDefaultCacher(newCache())
	}
	engine.ShowSQL = showTestSql
	engine.ShowErr = showTestSql
	engine.ShowWarn = showTestSql
	engine.ShowDebug = showTestSql
}

func TestTidbNoCache(t *testing.T) {
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()
	setEngine(engine, false)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
	tests.BaseTestAll3(engine, t)
}

func TestTidbWithCache(t *testing.T) {
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	setEngine(engine, true)

	tests.BaseTestAll(engine, t)
	tests.BaseTestAll2(engine, t)
}

const (
	createTableQl = "CREATE TABLE IF NOT EXISTS `big_struct` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `name` TEXT NULL, `title` TEXT NULL, `age` TEXT NULL, `alias` TEXT NULL, `nick_name` TEXT NULL);"
	dropTableQl   = "DROP TABLE IF EXISTS `big_struct`;"
)

func BenchmarkTidbDriverInsert(t *testing.B) {
	tests.DoBenchDriver(newTidbDriverDB, createTableQl, dropTableQl,
		tests.DoBenchDriverInsert, t)
}

func BenchmarkTidbDriverFind(t *testing.B) {
	tests.DoBenchDriver(newTidbDriverDB, createTableQl, dropTableQl,
		tests.DoBenchDriverFind, t)
}

func BenchmarkTidbNoCacheInsert(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	tests.DoBenchInsert(engine, t)
}

func BenchmarkTidbNoCacheFind(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	//engine.ShowSQL = true
	tests.DoBenchFind(engine, t)
}

func BenchmarkTidbNoCacheFindPtr(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()
	//engine.ShowSQL = true
	tests.DoBenchFindPtr(engine, t)
}

func BenchmarkTidbCacheInsert(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	engine.SetDefaultCacher(newCache())
	tests.DoBenchInsert(engine, t)
}

func BenchmarkTidbCacheFind(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	engine.SetDefaultCacher(newCache())
	tests.DoBenchFind(engine, t)
}

func BenchmarkTidbCacheFindPtr(t *testing.B) {
	t.StopTimer()
	engine, err := newTidbEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	engine.SetDefaultCacher(newCache())
	tests.DoBenchFindPtr(engine, t)
}

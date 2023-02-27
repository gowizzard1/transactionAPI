package gorm

import (
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	db "transationAPI/database"
)

func New(dialector gorm.Dialector, c gorm.Config, maxOpenConns, maxIdleConns, connMaxLifeTime int) (db.Database, error) {
	connection, err := gorm.Open(dialector, &c)
	if err != nil {
		return nil, err
	}
	connection = connection.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	sqlDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	// Important.
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifeTime) * time.Second)

	return &gormDB2{db: connection}, err
}

// wrap2s gorm.DB in an interface
func wrap2(db *gorm.DB) db.Database {
	return &gormDB2{db}
}

type gormDB2 struct {
	db *gorm.DB
}

func (g *gormDB2) AddError(err error) error {
	return g.db.AddError(err)
}

func (g *gormDB2) Attrs(attrs ...interface{}) db.Database {
	return wrap2(g.db.Attrs(attrs...))
}

func (g *gormDB2) Assign(attrs ...interface{}) db.Database {
	return wrap2(g.db.Assign(attrs...))
}

func (g *gormDB2) Association(column string) *gorm.Association {
	return g.db.Association(column)
}
func (g *gormDB2) AutoMigrate(values ...interface{}) error {
	return g.db.AutoMigrate(values...)
}

func (g *gormDB2) Begin() db.Database {
	return wrap2(g.db.Begin())
}

func (g *gormDB2) Count(value *int64) db.Database {
	return wrap2(g.db.Count(value))
}

func (g *gormDB2) Commit() db.Database {
	return wrap2(g.db.Commit())
}

func (g *gormDB2) CreateInBatches(value interface{}, batchSize int) db.Database {
	return wrap2(g.db.CreateInBatches(value, batchSize))
}

func (g *gormDB2) Create(value interface{}) db.Database {
	return wrap2(g.db.Create(value))
}

func (g *gormDB2) DB() (*sql.DB, error) {
	return g.db.DB()
}

func (g *gormDB2) Debug() db.Database {
	return wrap2(g.db.Debug())
}

func (g *gormDB2) Delete(value interface{}, where ...interface{}) db.Database {
	return wrap2(g.db.Delete(value, where...))
}

func (g *gormDB2) Distinct(args ...interface{}) (tx db.Database) {
	return wrap2(g.db.Distinct(args...))
}

func (g *gormDB2) Error() error {
	return g.db.Error
}

func (g *gormDB2) Exec(sql string, values ...interface{}) db.Database {
	return wrap2(g.db.Exec(sql, values...))
}

func (g *gormDB2) Find(out interface{}, where ...interface{}) db.Database {
	return wrap2(g.db.Find(out, where...))
}

func (g *gormDB2) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) db.Database {
	return wrap2(g.db.FindInBatches(dest, batchSize, fc))
}

func (g *gormDB2) First(out interface{}, where ...interface{}) db.Database {
	return wrap2(g.db.First(out, where...))
}

func (g *gormDB2) FirstOrCreate(out interface{}, where ...interface{}) db.Database {
	return wrap2(g.db.FirstOrCreate(out, where...))
}

func (g *gormDB2) FirstOrInit(out interface{}, where ...interface{}) db.Database {
	return wrap2(g.db.FirstOrInit(out, where...))
}

func (g *gormDB2) Get(name string) (interface{}, bool) {
	return g.db.Get(name)
}

func (g *gormDB2) Group(query string) db.Database {
	return wrap2(g.db.Group(query))
}

func (g *gormDB2) Having(query interface{}, values ...interface{}) db.Database {
	return wrap2(g.db.Having(query, values...))
}

func (g *gormDB2) InstanceGet(key string) (interface{}, bool) {
	return g.db.InstanceGet(key)
}

func (g *gormDB2) InstanceSet(key string, value interface{}) db.Database {
	return wrap2(g.db.InstanceSet(key, value))
}

func (g *gormDB2) Joins(query string, args ...interface{}) db.Database {
	return wrap2(g.db.Joins(query, args...))
}

func (g *gormDB2) Last(out interface{}, where ...interface{}) db.Database {
	return wrap2(g.db.Last(out, where...))
}
func (g *gormDB2) Limit(limit int) db.Database {
	return wrap2(g.db.Limit(limit))
}

func (g *gormDB2) LogMode(level logger.LogLevel) db.Database {
	g.db.Logger = g.db.Logger.LogMode(level)
	return wrap2(g.db)
}

func (g *gormDB2) Model(value interface{}) db.Database {
	return wrap2(g.db.Model(value))
}

func (g *gormDB2) Not(query interface{}, args ...interface{}) db.Database {
	return wrap2(g.db.Not(query, args...))
}

func (g *gormDB2) Offset(offset int) db.Database {
	return wrap2(g.db.Offset(offset))
}

func (g *gormDB2) Omit(columns ...string) db.Database {
	return wrap2(g.db.Omit(columns...))
}

func (g *gormDB2) Or(query interface{}, args ...interface{}) db.Database {
	return wrap2(g.db.Or(query, args...))
}

func (g *gormDB2) Order(value interface{}) db.Database {
	return wrap2(g.db.Order(value))
}

func (g *gormDB2) Pluck(column string, value interface{}) db.Database {
	return wrap2(g.db.Pluck(column, value))
}

func (g *gormDB2) Preload(column string, conditions ...interface{}) db.Database {
	return wrap2(g.db.Preload(column, conditions...))
}

func (g *gormDB2) Raw(sql string, values ...interface{}) db.Database {
	return wrap2(g.db.Raw(sql, values...))
}

func (g *gormDB2) RecordNotFound() bool {
	return errors.Is(g.Error(), gorm.ErrRecordNotFound)
}

func (g *gormDB2) Rollback() db.Database {
	return wrap2(g.db.Rollback())
}

func (g *gormDB2) RollbackTo(name string) db.Database {
	return wrap2(g.db.RollbackTo(name))
}

func (g *gormDB2) Row() *sql.Row {
	return g.db.Row()
}

func (g *gormDB2) Rows() (*sql.Rows, error) {
	return g.db.Rows()
}

func (g *gormDB2) RowsAffected() int64 {
	return g.db.RowsAffected
}

func (g *gormDB2) Save(value interface{}) db.Database {
	return wrap2(g.db.Save(value))
}

func (g *gormDB2) SavePoint(name string) db.Database {
	return wrap2(g.db.SavePoint(name))
}

func (g *gormDB2) Scan(dest interface{}) db.Database {
	return wrap2(g.db.Scan(dest))
}

func (g *gormDB2) ScanRows(rows *sql.Rows, result interface{}) error {
	return g.db.ScanRows(rows, result)
}

func (g *gormDB2) Select(query interface{}, args ...interface{}) db.Database {
	return wrap2(g.db.Select(query, args...))
}

func (g *gormDB2) Session(config *gorm.Session) db.Database {
	return wrap2(g.db.Session(config))
}

func (g *gormDB2) Set(name string, value interface{}) db.Database {
	return wrap2(g.db.Set(name, value))
}

func (g *gormDB2) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return g.db.SetupJoinTable(model, field, joinTable)
}

func (g *gormDB2) Scopes(funcs ...func(*gorm.DB) *gorm.DB) db.Database {
	return wrap2(g.db.Scopes(funcs...))
}

func (g *gormDB2) Table(name string) db.Database {
	return wrap2(g.db.Table(name))
}

func (g *gormDB2) Take(out interface{}, where ...interface{}) db.Database {
	return wrap2(g.db.Take(out, where))
}

func (g *gormDB2) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return g.db.Transaction(fc, opts...)
}

func (g *gormDB2) Unscoped() db.Database {
	return wrap2(g.db.Unscoped())
}

func (g *gormDB2) Update(column string, value interface{}) db.Database {
	return wrap2(g.db.Update(column, value))
}

func (g *gormDB2) Updates(values interface{}) db.Database {
	return wrap2(g.db.Updates(values))
}

func (g *gormDB2) UpdateColumn(column string, value interface{}) db.Database {
	return wrap2(g.db.UpdateColumn(column, value))
}

func (g *gormDB2) UpdateColumns(values interface{}) db.Database {
	return wrap2(g.db.UpdateColumns(values))
}

func (g *gormDB2) Use(plugin gorm.Plugin) error {
	return g.db.Use(plugin)
}

func (g *gormDB2) Where(query interface{}, args ...interface{}) db.Database {
	return wrap2(g.db.Where(query, args...))
}

func (g *gormDB2) WithContext(ctx context.Context) db.Database {
	return wrap2(g.db.WithContext(ctx))
}

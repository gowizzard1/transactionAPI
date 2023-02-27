package db

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Database interface {
	AddError(err error) error
	AutoMigrate(values ...interface{}) error
	Association(column string) *gorm.Association
	Attrs(attrs ...interface{}) Database
	Assign(attrs ...interface{}) Database
	Begin() Database
	Commit() Database
	Count(value *int64) Database
	Create(value interface{}) Database
	CreateInBatches(value interface{}, batchSize int) (tx Database)
	DB() (*sql.DB, error)
	Debug() Database
	Distinct(args ...interface{}) Database
	Delete(value interface{}, where ...interface{}) Database
	Error() error
	Exec(sql string, values ...interface{}) Database
	Get(name string) (interface{}, bool)
	Group(query string) Database
	Find(out interface{}, where ...interface{}) Database
	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) Database
	First(out interface{}, where ...interface{}) Database
	FirstOrCreate(out interface{}, where ...interface{}) Database
	FirstOrInit(out interface{}, where ...interface{}) Database
	Having(query interface{}, values ...interface{}) Database
	InstanceSet(key string, value interface{}) Database
	InstanceGet(key string) (interface{}, bool)
	Joins(query string, args ...interface{}) Database
	Last(out interface{}, where ...interface{}) Database
	Limit(limit int) Database
	Model(value interface{}) Database
	Not(query interface{}, args ...interface{}) Database
	Offset(offset int) Database
	Omit(columns ...string) Database
	Or(query interface{}, args ...interface{}) Database
	Order(value interface{}) Database
	Pluck(column string, value interface{}) Database
	Preload(column string, conditions ...interface{}) Database
	Raw(sql string, values ...interface{}) Database
	RecordNotFound() bool
	Rollback() Database
	RollbackTo(name string) Database
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	RowsAffected() int64
	Save(value interface{}) Database
	SavePoint(name string) Database
	Scan(dest interface{}) Database
	ScanRows(rows *sql.Rows, result interface{}) error
	Select(query interface{}, args ...interface{}) Database
	Session(config *gorm.Session) Database
	Set(name string, value interface{}) Database
	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) Database
	Take(out interface{}, where ...interface{}) Database
	Table(name string) Database
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
	Unscoped() Database
	Update(column string, value interface{}) Database
	Updates(values interface{}) Database
	UpdateColumn(column string, value interface{}) Database
	UpdateColumns(values interface{}) Database
	Use(plugin gorm.Plugin) error
	WithContext(ctx context.Context) Database
	Where(query interface{}, args ...interface{}) Database
}

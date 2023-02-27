package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Deprecated DB contains information for current db connection
type DBInterface interface {
	New() DBInterface
	Close() error
	DB() *sql.DB
	CommonDB() gorm.SQLCommon
	Dialect() gorm.Dialect
	Callback() *gorm.Callback
	LogMode(enable bool) DBInterface
	BlockGlobalUpdate(enable bool) DBInterface
	HasBlockGlobalUpdate() bool
	SingularTable(enable bool)
	NewScope(value interface{}) *gorm.Scope
	Where(query interface{}, args ...interface{}) DBInterface
	Or(query interface{}, args ...interface{}) DBInterface
	Not(query interface{}, args ...interface{}) DBInterface
	Limit(limit interface{}) DBInterface
	Offset(offset interface{}) DBInterface
	Order(value interface{}, reorder ...bool) DBInterface
	Select(query interface{}, args ...interface{}) DBInterface
	Omit(columns ...string) DBInterface
	Group(query string) DBInterface
	Having(query interface{}, values ...interface{}) DBInterface
	Joins(query string, args ...interface{}) DBInterface
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) DBInterface
	Unscoped() DBInterface
	Attrs(attrs ...interface{}) DBInterface
	Assign(attrs ...interface{}) DBInterface
	First(out interface{}, where ...interface{}) DBInterface
	Take(out interface{}, where ...interface{}) DBInterface
	Last(out interface{}, where ...interface{}) DBInterface
	Find(out interface{}, where ...interface{}) DBInterface
	Preloads(out interface{}) DBInterface
	Scan(dest interface{}) DBInterface
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) DBInterface
	Count(value interface{}) DBInterface
	Related(value interface{}, foreignKeys ...string) DBInterface
	FirstOrInit(out interface{}, where ...interface{}) DBInterface
	FirstOrCreate(out interface{}, where ...interface{}) DBInterface
	Update(attrs ...interface{}) DBInterface
	Updates(values interface{}, ignoreProtectedAttrs ...bool) DBInterface
	UpdateColumn(attrs ...interface{}) DBInterface
	UpdateColumns(values interface{}) DBInterface
	Save(value interface{}) DBInterface
	Create(value interface{}) DBInterface
	Delete(value interface{}, where ...interface{}) DBInterface
	Raw(sql string, values ...interface{}) DBInterface
	Exec(sql string, values ...interface{}) DBInterface
	Model(value interface{}) DBInterface
	Table(name string) DBInterface
	Debug() DBInterface
	Begin() DBInterface
	Commit() DBInterface
	Rollback() DBInterface
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	CreateTable(models ...interface{}) DBInterface
	DropTable(values ...interface{}) DBInterface
	DropTableIfExists(values ...interface{}) DBInterface
	HasTable(value interface{}) bool
	AutoMigrate(values ...interface{}) DBInterface
	ModifyColumn(column string, typ string) DBInterface
	DropColumn(column string) DBInterface
	AddIndex(indexName string, columns ...string) DBInterface
	AddUniqueIndex(indexName string, columns ...string) DBInterface
	RemoveIndex(indexName string) DBInterface
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) DBInterface
	RemoveForeignKey(field string, dest string) DBInterface
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) DBInterface
	Set(name string, value interface{}) DBInterface
	InstantSet(name string, value interface{}) DBInterface
	Get(name string) (value interface{}, ok bool)
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	AddError(err error) error
	GetErrors() []error
	Error() error
	RowsAffected() int64
}

func NewDB(host string, port string, user string, password string, dbname string, debug bool, maxOpenConns int, maxIdleConns int, connMaxLifeTime int) (DBInterface, error) {
	connectionString := "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"
	connection := fmt.Sprintf(connectionString, host, port, user, dbname, password)

	db, err := Openw("postgres", connection)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")

	// Important.
	maxLifeTime := time.Duration(connMaxLifeTime) * time.Second
	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetConnMaxLifetime(maxLifeTime)
	db.LogMode(debug)

	return db, nil
}

type gormDB struct {
	db *gorm.DB
}

// Openw is a drop-in replacement for Open()
func Openw(dialect string, args ...interface{}) (db DBInterface, err error) {
	gormDB, err := gorm.Open(dialect, args...)
	return wrap(gormDB), err
}

func (gorm *gormDB) Close() error {
	return gorm.db.Close()
}

func (gorm *gormDB) DB() *sql.DB {
	return gorm.db.DB()
}

func (gorm *gormDB) New() DBInterface {
	return wrap(gorm.db.New())
}

func (gorm *gormDB) NewScope(value interface{}) *gorm.Scope {
	return gorm.db.NewScope(value)
}

func (gorm *gormDB) CommonDB() gorm.SQLCommon {
	return gorm.db.CommonDB()
}

func (gorm *gormDB) Callback() *gorm.Callback {
	return gorm.db.Callback()
}

func (gorm *gormDB) SetLogger(log gorm.Logger) {
	gorm.db.SetLogger(log)
}

func (gorm *gormDB) LogMode(enable bool) DBInterface {
	return wrap(gorm.db.LogMode(enable))
}

func (gorm *gormDB) BlockGlobalUpdate(enable bool) DBInterface {
	return wrap(gorm.db.BlockGlobalUpdate(enable))
}

func (gorm *gormDB) Take(out interface{}, where ...interface{}) DBInterface {
	return wrap(gorm.db.Take(out, where))
}

func (gorm *gormDB) RemoveForeignKey(field string, dest string) DBInterface {
	return wrap(gorm.db.RemoveForeignKey(field, dest))
}

func (gorm *gormDB) Preloads(out interface{}) DBInterface {
	return wrap(gorm.db.Preloads(out))
}

func (gorm *gormDB) HasBlockGlobalUpdate() bool {
	return gorm.db.HasBlockGlobalUpdate()
}

func (gorm *gormDB) Dialect() gorm.Dialect {
	return gorm.db.Dialect()
}

func (gorm *gormDB) SingularTable(enable bool) {
	gorm.db.SingularTable(enable)
}

func (gorm *gormDB) Where(query interface{}, args ...interface{}) DBInterface {
	return wrap(gorm.db.Where(query, args...))
}

func (gorm *gormDB) Or(query interface{}, args ...interface{}) DBInterface {
	return wrap(gorm.db.Or(query, args...))
}

func (gorm *gormDB) Not(query interface{}, args ...interface{}) DBInterface {
	return wrap(gorm.db.Not(query, args...))
}

func (gorm *gormDB) Limit(value interface{}) DBInterface {
	return wrap(gorm.db.Limit(value))
}

func (gorm *gormDB) Offset(value interface{}) DBInterface {
	return wrap(gorm.db.Offset(value))
}

func (gorm *gormDB) Order(value interface{}, reorder ...bool) DBInterface {
	return wrap(gorm.db.Order(value, reorder...))
}

func (gorm *gormDB) Select(query interface{}, args ...interface{}) DBInterface {
	return wrap(gorm.db.Select(query, args...))
}

func (gorm *gormDB) Omit(columns ...string) DBInterface {
	return wrap(gorm.db.Omit(columns...))
}

func (gorm *gormDB) Group(query string) DBInterface {
	return wrap(gorm.db.Group(query))
}

func (gorm *gormDB) Having(query interface{}, values ...interface{}) DBInterface {
	return wrap(gorm.db.Having(query, values...))
}

func (gorm *gormDB) Joins(query string, args ...interface{}) DBInterface {
	return wrap(gorm.db.Joins(query, args...))
}

func (gorm *gormDB) Scopes(funcs ...func(*gorm.DB) *gorm.DB) DBInterface {
	return wrap(gorm.db.Scopes(funcs...))
}

func (gorm *gormDB) Unscoped() DBInterface {
	return wrap(gorm.db.Unscoped())
}

func (gorm *gormDB) Attrs(attrs ...interface{}) DBInterface {
	return wrap(gorm.db.Attrs(attrs...))
}

func (gorm *gormDB) Assign(attrs ...interface{}) DBInterface {
	return wrap(gorm.db.Assign(attrs...))
}

func (gorm *gormDB) First(out interface{}, where ...interface{}) DBInterface {
	return wrap(gorm.db.First(out, where...))
}

func (gorm *gormDB) Last(out interface{}, where ...interface{}) DBInterface {
	return wrap(gorm.db.Last(out, where...))
}

func (gorm *gormDB) Find(out interface{}, where ...interface{}) DBInterface {
	return wrap(gorm.db.Find(out, where...))
}

func (gorm *gormDB) Scan(dest interface{}) DBInterface {
	return wrap(gorm.db.Scan(dest))
}

func (gorm *gormDB) Row() *sql.Row {
	return gorm.db.Row()
}

func (gorm *gormDB) Rows() (*sql.Rows, error) {
	return gorm.db.Rows()
}

func (gorm *gormDB) ScanRows(rows *sql.Rows, result interface{}) error {
	return gorm.db.ScanRows(rows, result)
}

func (gorm *gormDB) Pluck(column string, value interface{}) DBInterface {
	return wrap(gorm.db.Pluck(column, value))
}

func (gorm *gormDB) Count(value interface{}) DBInterface {
	return wrap(gorm.db.Count(value))
}

func (gorm *gormDB) Related(value interface{}, foreignKeys ...string) DBInterface {
	return wrap(gorm.db.Related(value, foreignKeys...))
}

func (gorm *gormDB) FirstOrInit(out interface{}, where ...interface{}) DBInterface {
	return wrap(gorm.db.FirstOrInit(out, where...))
}

func (gorm *gormDB) FirstOrCreate(out interface{}, where ...interface{}) DBInterface {
	return wrap(gorm.db.FirstOrCreate(out, where...))
}

func (gorm *gormDB) Update(attrs ...interface{}) DBInterface {
	return wrap(gorm.db.Update(attrs...))
}

func (gorm *gormDB) Updates(values interface{}, ignoreProtectedAttrs ...bool) DBInterface {
	return wrap(gorm.db.Updates(values, ignoreProtectedAttrs...))
}

func (gorm *gormDB) UpdateColumn(attrs ...interface{}) DBInterface {
	return wrap(gorm.db.UpdateColumn(attrs...))
}

func (gorm *gormDB) UpdateColumns(values interface{}) DBInterface {
	return wrap(gorm.db.UpdateColumns(values))
}

func (gorm *gormDB) Save(value interface{}) DBInterface {
	return wrap(gorm.db.Save(value))
}

func (gorm *gormDB) Create(value interface{}) DBInterface {
	return wrap(gorm.db.Create(value))
}

func (gorm *gormDB) Delete(value interface{}, where ...interface{}) DBInterface {
	return wrap(gorm.db.Delete(value, where...))
}

func (gorm *gormDB) Raw(sql string, values ...interface{}) DBInterface {
	return wrap(gorm.db.Raw(sql, values...))
}

func (gorm *gormDB) Exec(sql string, values ...interface{}) DBInterface {
	return wrap(gorm.db.Exec(sql, values...))
}

func (gorm *gormDB) Model(value interface{}) DBInterface {
	return wrap(gorm.db.Model(value))
}

func (gorm *gormDB) Table(name string) DBInterface {
	return wrap(gorm.db.Table(name))
}

func (gorm *gormDB) Debug() DBInterface {
	return wrap(gorm.db.Debug())
}

func (gorm *gormDB) Begin() DBInterface {
	return wrap(gorm.db.Begin())
}

func (gorm *gormDB) Commit() DBInterface {
	return wrap(gorm.db.Commit())
}

func (gorm *gormDB) Rollback() DBInterface {
	return wrap(gorm.db.Rollback())
}

func (gorm *gormDB) NewRecord(value interface{}) bool {
	return gorm.db.NewRecord(value)
}

func (gorm *gormDB) RecordNotFound() bool {
	return gorm.db.RecordNotFound()
}

func (gorm *gormDB) CreateTable(values ...interface{}) DBInterface {
	return wrap(gorm.db.CreateTable(values...))
}

func (gorm *gormDB) DropTable(values ...interface{}) DBInterface {
	return wrap(gorm.db.DropTable(values...))
}

func (gorm *gormDB) DropTableIfExists(values ...interface{}) DBInterface {
	return wrap(gorm.db.DropTableIfExists(values...))
}

func (gorm *gormDB) HasTable(value interface{}) bool {
	return gorm.db.HasTable(value)
}

func (gorm *gormDB) AutoMigrate(values ...interface{}) DBInterface {
	return wrap(gorm.db.AutoMigrate(values...))
}

func (gorm *gormDB) ModifyColumn(column string, typ string) DBInterface {
	return wrap(gorm.db.ModifyColumn(column, typ))
}

func (gorm *gormDB) DropColumn(column string) DBInterface {
	return wrap(gorm.db.DropColumn(column))
}

func (gorm *gormDB) AddIndex(indexName string, columns ...string) DBInterface {
	return wrap(gorm.db.AddIndex(indexName, columns...))
}

func (gorm *gormDB) AddUniqueIndex(indexName string, columns ...string) DBInterface {
	return wrap(gorm.db.AddUniqueIndex(indexName, columns...))
}

func (gorm *gormDB) RemoveIndex(indexName string) DBInterface {
	return wrap(gorm.db.RemoveIndex(indexName))
}

func (gorm *gormDB) Association(column string) *gorm.Association {
	return gorm.db.Association(column)
}

func (gorm *gormDB) Preload(column string, conditions ...interface{}) DBInterface {
	return wrap(gorm.db.Preload(column, conditions...))
}

func (gorm *gormDB) Set(name string, value interface{}) DBInterface {
	return wrap(gorm.db.Set(name, value))
}

func (gorm *gormDB) InstantSet(name string, value interface{}) DBInterface {
	return wrap(gorm.db.InstantSet(name, value))
}

func (gorm *gormDB) Get(name string) (interface{}, bool) {
	return gorm.db.Get(name)
}

func (gorm *gormDB) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	gorm.db.SetJoinTableHandler(source, column, handler)
}

func (gorm *gormDB) AddForeignKey(field string, dest string, onDelete string, onUpdate string) DBInterface {
	return wrap(gorm.db.AddForeignKey(field, dest, onDelete, onUpdate))
}

func (gorm *gormDB) AddError(err error) error {
	return gorm.db.AddError(err)
}

func (gorm *gormDB) GetErrors() (errors []error) {
	return gorm.db.GetErrors()
}

func (gorm *gormDB) RowsAffected() int64 {
	return gorm.db.RowsAffected
}

func (gorm *gormDB) Error() error {
	return gorm.db.Error
}

// Wraps gorm.DB in an interface
func wrap(db *gorm.DB) DBInterface {
	return &gormDB{db}
}

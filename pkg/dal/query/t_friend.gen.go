// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"mini-min-tiktok/pkg/dal/model"
)

func newTFriend(db *gorm.DB, opts ...gen.DOOption) tFriend {
	_tFriend := tFriend{}

	_tFriend.tFriendDo.UseDB(db, opts...)
	_tFriend.tFriendDo.UseModel(&model.TFriend{})

	tableName := _tFriend.tFriendDo.TableName()
	_tFriend.ALL = field.NewAsterisk(tableName)
	_tFriend.ID = field.NewInt64(tableName, "id")
	_tFriend.UserID = field.NewInt64(tableName, "user_id")
	_tFriend.FriendID = field.NewInt64(tableName, "friend_id")

	_tFriend.fillFieldMap()

	return _tFriend
}

type tFriend struct {
	tFriendDo tFriendDo

	ALL      field.Asterisk
	ID       field.Int64 // 主键id
	UserID   field.Int64 // 用户id
	FriendID field.Int64 // 好友id

	fieldMap map[string]field.Expr
}

func (t tFriend) Table(newTableName string) *tFriend {
	t.tFriendDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tFriend) As(alias string) *tFriend {
	t.tFriendDo.DO = *(t.tFriendDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tFriend) updateTableName(table string) *tFriend {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.UserID = field.NewInt64(table, "user_id")
	t.FriendID = field.NewInt64(table, "friend_id")

	t.fillFieldMap()

	return t
}

func (t *tFriend) WithContext(ctx context.Context) ITFriendDo { return t.tFriendDo.WithContext(ctx) }

func (t tFriend) TableName() string { return t.tFriendDo.TableName() }

func (t tFriend) Alias() string { return t.tFriendDo.Alias() }

func (t *tFriend) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tFriend) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["id"] = t.ID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["friend_id"] = t.FriendID
}

func (t tFriend) clone(db *gorm.DB) tFriend {
	t.tFriendDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tFriend) replaceDB(db *gorm.DB) tFriend {
	t.tFriendDo.ReplaceDB(db)
	return t
}

type tFriendDo struct{ gen.DO }

type ITFriendDo interface {
	gen.SubQuery
	Debug() ITFriendDo
	WithContext(ctx context.Context) ITFriendDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITFriendDo
	WriteDB() ITFriendDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITFriendDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITFriendDo
	Not(conds ...gen.Condition) ITFriendDo
	Or(conds ...gen.Condition) ITFriendDo
	Select(conds ...field.Expr) ITFriendDo
	Where(conds ...gen.Condition) ITFriendDo
	Order(conds ...field.Expr) ITFriendDo
	Distinct(cols ...field.Expr) ITFriendDo
	Omit(cols ...field.Expr) ITFriendDo
	Join(table schema.Tabler, on ...field.Expr) ITFriendDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITFriendDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITFriendDo
	Group(cols ...field.Expr) ITFriendDo
	Having(conds ...gen.Condition) ITFriendDo
	Limit(limit int) ITFriendDo
	Offset(offset int) ITFriendDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITFriendDo
	Unscoped() ITFriendDo
	Create(values ...*model.TFriend) error
	CreateInBatches(values []*model.TFriend, batchSize int) error
	Save(values ...*model.TFriend) error
	First() (*model.TFriend, error)
	Take() (*model.TFriend, error)
	Last() (*model.TFriend, error)
	Find() ([]*model.TFriend, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TFriend, err error)
	FindInBatches(result *[]*model.TFriend, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TFriend) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITFriendDo
	Assign(attrs ...field.AssignExpr) ITFriendDo
	Joins(fields ...field.RelationField) ITFriendDo
	Preload(fields ...field.RelationField) ITFriendDo
	FirstOrInit() (*model.TFriend, error)
	FirstOrCreate() (*model.TFriend, error)
	FindByPage(offset int, limit int) (result []*model.TFriend, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITFriendDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tFriendDo) Debug() ITFriendDo {
	return t.withDO(t.DO.Debug())
}

func (t tFriendDo) WithContext(ctx context.Context) ITFriendDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tFriendDo) ReadDB() ITFriendDo {
	return t.Clauses(dbresolver.Read)
}

func (t tFriendDo) WriteDB() ITFriendDo {
	return t.Clauses(dbresolver.Write)
}

func (t tFriendDo) Session(config *gorm.Session) ITFriendDo {
	return t.withDO(t.DO.Session(config))
}

func (t tFriendDo) Clauses(conds ...clause.Expression) ITFriendDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tFriendDo) Returning(value interface{}, columns ...string) ITFriendDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tFriendDo) Not(conds ...gen.Condition) ITFriendDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tFriendDo) Or(conds ...gen.Condition) ITFriendDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tFriendDo) Select(conds ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tFriendDo) Where(conds ...gen.Condition) ITFriendDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tFriendDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITFriendDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tFriendDo) Order(conds ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tFriendDo) Distinct(cols ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tFriendDo) Omit(cols ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tFriendDo) Join(table schema.Tabler, on ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tFriendDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tFriendDo) RightJoin(table schema.Tabler, on ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tFriendDo) Group(cols ...field.Expr) ITFriendDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tFriendDo) Having(conds ...gen.Condition) ITFriendDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tFriendDo) Limit(limit int) ITFriendDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tFriendDo) Offset(offset int) ITFriendDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tFriendDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITFriendDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tFriendDo) Unscoped() ITFriendDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tFriendDo) Create(values ...*model.TFriend) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tFriendDo) CreateInBatches(values []*model.TFriend, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tFriendDo) Save(values ...*model.TFriend) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tFriendDo) First() (*model.TFriend, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFriend), nil
	}
}

func (t tFriendDo) Take() (*model.TFriend, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFriend), nil
	}
}

func (t tFriendDo) Last() (*model.TFriend, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFriend), nil
	}
}

func (t tFriendDo) Find() ([]*model.TFriend, error) {
	result, err := t.DO.Find()
	return result.([]*model.TFriend), err
}

func (t tFriendDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TFriend, err error) {
	buf := make([]*model.TFriend, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tFriendDo) FindInBatches(result *[]*model.TFriend, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tFriendDo) Attrs(attrs ...field.AssignExpr) ITFriendDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tFriendDo) Assign(attrs ...field.AssignExpr) ITFriendDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tFriendDo) Joins(fields ...field.RelationField) ITFriendDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tFriendDo) Preload(fields ...field.RelationField) ITFriendDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tFriendDo) FirstOrInit() (*model.TFriend, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFriend), nil
	}
}

func (t tFriendDo) FirstOrCreate() (*model.TFriend, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFriend), nil
	}
}

func (t tFriendDo) FindByPage(offset int, limit int) (result []*model.TFriend, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tFriendDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tFriendDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tFriendDo) Delete(models ...*model.TFriend) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tFriendDo) withDO(do gen.Dao) *tFriendDo {
	t.DO = *do.(*gen.DO)
	return t
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                  = new(Query)
	Article            *article
	User               *user
	UserIntegralRecord *userIntegralRecord
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Article = &Q.Article
	User = &Q.User
	UserIntegralRecord = &Q.UserIntegralRecord
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		Article:            newArticle(db, opts...),
		User:               newUser(db, opts...),
		UserIntegralRecord: newUserIntegralRecord(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Article            article
	User               user
	UserIntegralRecord userIntegralRecord
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Article:            q.Article.clone(db),
		User:               q.User.clone(db),
		UserIntegralRecord: q.UserIntegralRecord.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Article:            q.Article.replaceDB(db),
		User:               q.User.replaceDB(db),
		UserIntegralRecord: q.UserIntegralRecord.replaceDB(db),
	}
}

type queryCtx struct {
	Article            IArticleDo
	User               IUserDo
	UserIntegralRecord IUserIntegralRecordDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Article:            q.Article.WithContext(ctx),
		User:               q.User.WithContext(ctx),
		UserIntegralRecord: q.UserIntegralRecord.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}

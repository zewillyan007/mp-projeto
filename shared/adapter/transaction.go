package adapter

import "gorm.io/gorm"

type Transaction struct {
	Db *gorm.DB
}

func BeginTransaction(db *gorm.DB) *Transaction {
	return &Transaction{Db: db.Begin()}
}

func (o *Transaction) Commit() *gorm.DB {
	return o.Db.Commit()
}

func (o *Transaction) Rollback() *gorm.DB {
	return o.Db.Rollback()
}

func (o *Transaction) GetTransaction() *gorm.DB {
	return o.Db
}

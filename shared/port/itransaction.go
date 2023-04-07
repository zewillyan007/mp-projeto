package port

import "gorm.io/gorm"

type ITransaction interface {
	GetTransaction() *gorm.DB
}

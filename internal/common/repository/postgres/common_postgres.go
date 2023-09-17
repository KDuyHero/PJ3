package common_postgres

import (
	"context"
	"mobile-ecommerce/internal/core/domain"

	"gorm.io/gorm"
)

type commonRepository struct {
	gormDB *gorm.DB
}

func NewCommonRepository(gormDB *gorm.DB) domain.CommonRepository {
	return &commonRepository{
		gormDB: gormDB,
	}
}

func (c *commonRepository) Transaction(ctx context.Context, txFunc func(saf context.Context) error) (err error) {
	// tx := c.gormDB.Begin()
	// if !errors.Is(tx.Error, nil) {
	// 	return tx.Error
	// }

	// defer func() {
	// 	if p := recover(); p != nil {
	// 		log.Print("recover")
	// 		tx.Rollback()
	// 	} else if !errors.Is(err, nil) {
	// 		log.Print("rollback")
	// 		tx.Rollback()
	// 	} else {
	// 		err = tx.Commit().Error
	// 	}
	// }()

	// ctx.
	// err = txFunc(tx)
	// return err

	return nil
}

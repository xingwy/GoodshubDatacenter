package middleware

import (
	"context"
	"fmt"
	"goodshub-datacenter/constants"
	"goodshub-datacenter/utils"

	"gorm.io/gorm"
)

// 使用事务
func WithTransaction(ctx context.Context, db *gorm.DB, fn func(ctx context.Context) (any, error)) (data any, err error) {
	tx := db.Begin()
	// create transaction & save events in ctx
	ctx = context.WithValue(ctx, constants.TRANSACTION_KEY, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("panic occurred: %v", r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
			if err != nil {
				utils.LogError(constants.LOG_ID__TRANSACTION, "COMMIT", err)
			}
		}
	}()
	data, err = fn(ctx)
	if err != nil {
		utils.LogError(constants.LOG_ID__COMMON, err)
	}
	return

}

// 不使用事务
func WithoutTransaction(ctx context.Context, db *gorm.DB, fn func(ctx context.Context) (any, error)) (data any, err error) {
	// create transaction & save events in ctx
	ctx = context.WithValue(ctx, constants.TRANSACTION_KEY, db)
	data, err = fn(ctx)
	return
}

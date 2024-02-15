package repository

import (
	"context"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type BaseRepository struct {
	DB *gorm.DB
}

var txKey = struct{}{}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{DB: db}
}

func (r *BaseRepository) ReadWriteTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := r.DB.WithContext(ctx).Begin()

	ctx = context.WithValue(ctx, &txKey, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := fn(ctx); err != nil {
		if rErr := tx.Rollback().Error; rErr != nil {
			return xerrors.Errorf("rollback failed: %w", rErr)
		}
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return xerrors.Errorf("commit failed: %w", err)
	}

	return nil
}

func getTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}

// If WithTx Transaction is in progress, extract tx from context. Otherwise, returns *gorm.DB with normal context.
func (r *BaseRepository) WithTx(ctx context.Context) *gorm.DB {
	tx, ok := getTx(ctx)
	if ok {
		return tx
	}
	return r.DB.WithContext(ctx)
}

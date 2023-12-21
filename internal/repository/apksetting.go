package repository

import (
	"context"
	"simpel-gateway/internal/model"
	"simpel-gateway/pkg/util"

	"gorm.io/gorm"
)

type ApkSetting interface {
	Store(ctx context.Context, data model.ApkSetting) error
	FindOne(ctx context.Context, selectedFields string) (model.ApkSetting, error)
}

type apksetting struct {
	Db *gorm.DB
}

func NewApkSettingRepository(db *gorm.DB) ApkSetting {
	return &apksetting{
		Db: db,
	}
}

func (r *apksetting) Store(ctx context.Context, data model.ApkSetting) error {
	tx := r.Db.WithContext(ctx)

	// Mulai transaksi
	err := tx.Transaction(func(tx *gorm.DB) error {
		// Hapus semua entri yang ada
		if err := tx.Where("1 = 1").Delete(&model.ApkSetting{}).Error; err != nil {
			return err
		}

		// Simpan data baru
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *apksetting) FindOne(ctx context.Context, selectedFields string) (model.ApkSetting, error) {
	var res model.ApkSetting

	db := r.Db.WithContext(ctx).Model(&model.ApkSetting{})

	db = util.SetSelectFields(db, selectedFields)

	// Gunakan First() tanpa kondisi WHERE untuk mengambil data pertama
	if err := db.First(&res).Error; err != nil {
		return model.ApkSetting{}, err
	}

	return res, nil
}

package model

import "time"

type ApkSetting struct {
	ApkMinVersion   string    `gorm:"varchar"`
	ApkDownloadLink string    `gorm:"varchar"`
	UserAgreement   string    `gorm:"varchar"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
}

func (ApkSetting) TableName() string {
	return "apk_setting"
}

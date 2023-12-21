package apksetting

import (
	"context"
	"simpel-gateway/internal/dto"
	"simpel-gateway/internal/factory"
	"simpel-gateway/internal/model"
	"simpel-gateway/internal/repository"
	"simpel-gateway/pkg/constants"
)

type service struct {
	apkSettingRepository repository.ApkSetting
}

type Service interface {
	Store(ctx context.Context, payload dto.PayloadApkSetting) error
	FindOne(ctx context.Context) (dto.FindOneApkSetting, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		apkSettingRepository: f.ApkSettingRepository,
	}
}

func (s *service) Store(ctx context.Context, payload dto.PayloadApkSetting) error {
	dataStore := model.ApkSetting{
		ApkMinVersion:   payload.ApkMinVersion,
		ApkDownloadLink: payload.ApkDownloadLink,
		UserAgreement:   payload.UserAgreement,
	}
	s.apkSettingRepository.Store(ctx, dataStore)
	return nil
}

func (s *service) FindOne(ctx context.Context) (dto.FindOneApkSetting, error) {
	data, err := s.apkSettingRepository.FindOne(ctx, "apk_min_version, apk_download_link, user_agreement")
	if err != nil {
		return dto.FindOneApkSetting{}, constants.ErrorNoFoundDataApkSetting
	}

	res := dto.FindOneApkSetting{
		ApkMinVersion:   data.ApkMinVersion,
		ApkDownloadLink: data.ApkDownloadLink,
		UserAgreement:   data.UserAgreement,
	}

	return res, nil
}

package dto

type (
	PayloadApkSetting struct {
		ApkMinVersion   string `json:"apk_min_version"`
		ApkDownloadLink string `json:"apk_download_link"`
		UserAgreement   string `json:"user_agreement"`
	}

	FindOneApkSetting struct {
		ApkMinVersion   string `json:"apk_min_version"`
		ApkDownloadLink string `json:"apk_download_link"`
		UserAgreement   string `json:"user_agreement"`
	}
)

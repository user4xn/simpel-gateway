package dto

type (
	PayloadServer struct {
		ServerIp   string `json:"server_ip" binding:"required"`
		ServerName string `json:"server_name" binding:"required"`
	}

	ListServer struct {
		ServerIp   string `json:"server_ip"`
		ServerName string `json:"server_name"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}

	CheckServerIp struct {
		IsRegistered bool `json:"is_registered"`
	}

	PayloadCheckServerIp struct {
		ServerIp string `json:"server_ip" binding:"required"`
	}
)

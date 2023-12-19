package model

import "time"

type Server struct {
	ServerIp   string    `gorm:"varchar"`
	ServerName string    `gorm:"varchar"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Server) TableName() string {
	return "servers"
}

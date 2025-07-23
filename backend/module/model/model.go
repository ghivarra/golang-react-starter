package model

import (
	"time"

	"gorm.io/gorm"
)

type Module struct {
	Name      string    `gorm:"unique;size:200;not null"`
	Alias     string    `gorm:"size:200;not null"`
	CreatedAt time.Time `gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type Role struct {
	ID           uint64    `gorm:"primaryKey;->"`
	Name         string    `gorm:"unique;size:60;not null"`
	IsSuperadmin int       `gorm:"default:0;not null"`
	CreatedAt    time.Time `gorm:"<-:create;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type RoleModuleList struct {
	ID         uint64 `gorm:"primaryKey;->"`
	RoleID     uint64 `gorm:"column:role_id;index"`
	Role       Role   `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ModuleName uint64 `gorm:"column:module_name;index"`
	Module     Module `gorm:"foreignKey:ModuleName;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	ID        uint64         `gorm:"primaryKey;->"`
	Name      string         `gorm:"size:100;not null"`
	Username  string         `gorm:"unique;size:100;not null"`
	Email     string         `gorm:"unique;size:100;not null"`
	Password  string         `gorm:"size:200;not null"`
	RoleID    uint64         `gorm:"column:role_id;index"`
	Role      Role           `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"default:null;"`
}

type TokenRefresh struct {
	ID        string    `gorm:"primaryKey"`
	UserID    uint64    `gorm:"column:user_id;index"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExpiredAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"<-:create;autoCreateTime"`
}

type TokenRevoked struct {
	ID        string    `gorm:"primaryKey"`
	ExpiredAt time.Time `gorm:"not null"`
	RevokedAt time.Time `gorm:"<-:create;autoCreateTime"`
}

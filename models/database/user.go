package database

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID 				uuid.UUID 	`gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name 			string 		`gorm:"type:varchar(255);not null;unique"`
	Email          	string    	`gorm:"type:varchar(50);not null;unique"`
	ProfilePicture 	string    	`gorm:"type:varchar(255)"`  
	IsActive       	bool      	`gorm:"type:boolean;not null"`                      
	CreatedBy      	string    	`gorm:"type:varchar(50);not null;default:'system'"`            
	UpdatedBy      	string    	`gorm:"type:varchar(50);not null;default:'system'"`              
	CreatedAt      	time.Time 	`gorm:"autoCreateTime;not null;default:now()"`                   
	UpdatedAt      	time.Time 	`gorm:"autoUpdateTime;not null;default:now()"` 
	
	Otps 			[]Otps 		`gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
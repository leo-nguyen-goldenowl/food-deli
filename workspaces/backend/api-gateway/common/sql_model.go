package common

import (
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid();"`
	Status    int        `json:"status" gorm:"column:status;default:1;not null;"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;"`
}

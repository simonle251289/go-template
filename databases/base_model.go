package databases

import "time"

type BaseModel struct {
	ID          string    `json:"id" gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();"`
	IsEnable    bool      `json:"isEnable" gorm:"column:is_enable;default:true;not null"`
	IsDeleted   bool      `json:"isDeleted" gorm:"column:is_deleted;default:false;not null"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:created_date;type:timestamptz;default:NOW();not null;"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"column:updated_date;type:timestamptz;default:NOW();not null;"`
}

package databases

import (
	"fmt"
	"strings"
	"time"
)

type AppTime time.Time

type AppMarshaller interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

func (appT *AppTime) MarshalJSON() ([]byte, error) {
	timeString := fmt.Sprintf("\"%s\"", time.Time(*appT).UTC().Format(time.RFC3339))
	return []byte(timeString), nil
}

func (appT *AppTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse(time.RFC3339, value) //parse time
	if err != nil {
		return err
	}
	*appT = AppTime(t) //set result using the pointer
	return nil
}

type BaseModel struct {
	ID          string  `json:"id" gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();"`
	IsEnable    bool    `json:"isEnable" gorm:"column:is_enable;default:true;not null"`
	IsDeleted   bool    `json:"isDeleted" gorm:"column:is_deleted;default:false;not null"`
	CreatedDate AppTime `json:"createdDate" gorm:"column:created_date;type:timestamptz;default:NOW();not null;"`
	UpdatedDate AppTime `json:"updatedDate" gorm:"column:updated_date;type:timestamptz;default:NOW();not null;"`
}

package databases

type TemplateModel struct {
	BaseModel `json:",inline"` //Set inline, if you want to display result on a struct
	Name      string           `json:"name" gorm:"not null;varchar(100)"`
}

func (TemplateModel) TableName() string {
	return "template"
}

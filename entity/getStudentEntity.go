package entity

type StudentEntity struct {
	RollNo int    `gorm:"primary_key; column:roll_no; autoIncreament"`
	Name   string `gorm:"column:name;type:varchar(50)"`
}

func (StudentEntity) TableName() string {
	return "student"
}

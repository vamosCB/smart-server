package test

import (
	"context"
	"smart-server/model/mysql"

	"gorm.io/gorm"
)

// Student 学生信息表
type Student struct {
	mysql.BaseColumn
	Name string `json:"name" comment:"学生名称"`
	Sex  string `json:"sex" comment:"学生性别"`
	Ext  string `json:"ext" comment:"额外信息"`
}

// GetStudentInfoByID 通过ID获取学生信息
func GetStudentInfoByID(ctx context.Context, ID int) (*Student, error) {
	var result Student
	err := mysql.DB.First(&result, ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}

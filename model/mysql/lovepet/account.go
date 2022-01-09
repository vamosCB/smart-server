package lovepet

import (
	"context"
	"smart-server/model/mysql"

	"gorm.io/gorm"
)

// Account 账户信息表
type Account struct {
	mysql.BaseColumn
	AccountNumber string `json:"account_number" comment:"账户ID"`
	NickName      string `json:"nick_name" comment:"账户昵称"`
	PhoneEncrypt  string `json:"phone_encrypt" comment:"加密手机号"`
	PwdEncrypt    string `json:"pwd_encrypt" comment:"账户加密密码"`
	Sex           string `json:"sex" comment:"性别"`
	Age           int    `json:"age" comment:"年龄"`
	Desc          string `json:"desc" comment:"账户描述"`
	ProvinceID    int    `json:"province_id" comment:"省份ID"`
	CityID        int    `json:"city_id" comment:"城市ID"`
}

// GetAccountInfoByID 通过ID获取账户信息
func GetAccountInfoByID(ctx context.Context, ID int) (*Account, error) {
	var result Account
	err := mysql.DB.First(&result, ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}

// CreateAccount 创建账户
func CreateAccount(ctx context.Context, account Account) error {
	return mysql.DB.Create(&account).Error
}

// GetAccountInfoByAccountNumber 通过账户ID获取账户信息
func GetAccountInfoByAccountNumber(ctx context.Context, aNumber string) (*Account, error) {
	var result Account
	where := map[string]string{
		"account_number": aNumber,
	}
	err := mysql.DB.Model(&Account{}).Where(where).Find(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}

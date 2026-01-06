package kurohelperdb

import "gorm.io/gorm"

func EnsureUser(userID, userName string) (*User, error) {
	var user User
	if err := Dbs.Where("id = ?", userID).FirstOrCreate(&user, User{ID: userID, Name: userName}).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func EnsureUserTx(tx *gorm.DB, userID, userName string) (*User, error) {
	var user User
	if err := tx.Where("id = ?", userID).FirstOrCreate(&user, User{ID: userID, Name: userName}).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 取得指定使用者資料
func GetUser(userID string) (User, error) {
	var user User

	err := Dbs.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// 取得所有使用者資料
func GetAllUser() ([]User, error) {
	var user []User

	err := Dbs.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

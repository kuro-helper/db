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

// 依據 userID 取得單一使用者資料
func GetUser(userID string) (User, error) {
	var user User

	err := Dbs.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// 與 GetUser 相同，但透過呼叫端提供的交易 tx 進行查詢
func GetUserTx(tx *gorm.DB, userID string) (User, error) {
	var user User

	err := tx.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// 取得所有使用者資料
func GetUsers() ([]User, error) {
	var user []User

	err := Dbs.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// 與 GetUsers 相同，但透過呼叫端提供的交易 tx 進行查詢
func GetUsersTx(tx *gorm.DB) ([]User, error) {
	var user []User

	err := tx.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

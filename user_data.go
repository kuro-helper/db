package kurohelperdb

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

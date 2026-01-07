package kurohelperdb

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func SelectUserInWish(userID string) ([]UserInWish, error) {
	var inWish []UserInWish

	err := Dbs.
		// Preload("User").
		Preload("GameErogs").
		Preload("GameErogs.BrandErogs").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&inWish).Error

	if err != nil {
		return nil, err
	}
	return inWish, nil
}

func GetUserInWish(userID string, gameErogsID int) (UserInWish, error) {
	var userInWish UserInWish

	err := Dbs.First(&userInWish, "user_id = ? AND game_erogs_id = ?", userID, gameErogsID).Error
	if err != nil {
		return userInWish, err
	}

	return userInWish, nil
}

func CreateUserInWish(userID string, gameErogsID int) error {
	userInWish := UserInWish{
		UserID:      userID,
		GameErogsID: gameErogsID,
	}

	if err := Dbs.Create(&userInWish).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUniqueViolation
		}
		return err
	}

	return nil
}

func DeleteUserInWish(userID string, gameErogsID int) error {
	err := Dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserInWish{}).Error

	return err
}

func CreateUserInWishTx(tx *gorm.DB, userID string, gameErogsID int) error {
	userInWish := UserInWish{
		UserID:      userID,
		GameErogsID: gameErogsID,
	}

	if err := tx.Create(&userInWish).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUniqueViolation
		}
		return err
	}

	return nil
}

func DeleteUserInWishTx(tx *gorm.DB, userID string, gameErogsID int) error {
	err := tx.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserInWish{}).Error

	return err
}

func FindUserInWishByUserAndGameNameLike(userID string, gameName string) (UserInWish, error) {
	var result UserInWish

	err := Dbs.
		Model(&UserInWish{}).
		Joins("JOIN game_erogs ON game_erogs.id = user_in_wishes.game_erogs_id").
		Where("user_in_wishes.user_id = ?", userID).
		Where("game_erogs.name ILIKE ?", "%"+gameName+"%").
		Preload("GameErogs").
		First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

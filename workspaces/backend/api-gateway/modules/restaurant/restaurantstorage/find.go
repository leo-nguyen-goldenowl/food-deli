package restaurantstorage

import (
	"context"
	"errors"

	"api-gateway/common"
	"api-gateway/modules/restaurant/restaurantmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	
	var result *restaurantmodel.Restaurant

	if err := db.Where(conditions).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return result, nil
}

package restaurantlikebiz

import (
	"context"

	"api-gateway/common"
	"api-gateway/modules/restaurantlike/restaurantlikemodel"
)

type ListUsersLikeRestaurantStore interface {
	GetUsersLikeRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
	) ([]common.SimpleUser, error)
}

type listUsersLikeRestaurantBiz struct {
	store ListUsersLikeRestaurantStore
}

func NewListUsersLikeRestaurantBiz(store ListUsersLikeRestaurantStore) *listUsersLikeRestaurantBiz {
	return &listUsersLikeRestaurantBiz{store: store}
}

func (biz *listUsersLikeRestaurantBiz) ListUsers(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}

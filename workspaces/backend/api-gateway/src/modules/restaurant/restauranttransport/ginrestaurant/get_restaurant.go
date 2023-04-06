package ginrestaurant

import (
	"net/http"

	"api-gateway/src/common"
	"api-gateway/src/component"
	"api-gateway/src/modules/restaurant/restaurantbiz"
	"api-gateway/src/modules/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurant(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
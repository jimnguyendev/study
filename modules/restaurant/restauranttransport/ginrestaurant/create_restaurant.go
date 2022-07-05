package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/study/common"
	"github.com/study/component"
	"github.com/study/modules/restaurant/restaurantbiz"
	"github.com/study/modules/restaurant/restaurantmodel"
	"github.com/study/modules/restaurant/restaurantstorage"
	"net/http"
)

// appCtx => không tăng tham số
func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/study/common"
	"github.com/study/component"
	"github.com/study/modules/restaurant/restaurantbiz"
	"github.com/study/modules/restaurant/restaurantmodel"
	"github.com/study/modules/restaurant/restaurantstorage"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

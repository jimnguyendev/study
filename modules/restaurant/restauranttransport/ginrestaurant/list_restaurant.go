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

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

# Note

Storage => Biz => transport

handle panic for transport (chỉ panic ở tầng ngoài cùng - transport)

### Panic

A panic is caused either by a runtime error, or an explicit call to the built-in panic function.

#### Stack traces
A stack trace – a report of all active stack frames – is typically printed to the console when a panic occurs. Stack traces can be very useful for debugging

=> Hoạt động dựa trên current stack trace của nó

Server sẽ chết luôn
```go
func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			panic("a")
		}()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
```

### Tại sao chúng ta phải có một cái upload image nằm riêng

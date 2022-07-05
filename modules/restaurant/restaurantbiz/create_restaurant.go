package restaurantbiz

import (
	"context"
	"github.com/study/modules/restaurant/restaurantmodel"
)

// interface mà khai báo ở đâu thì sử dụng ở đó
type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

// Thay vì chúng ta gọi chay chúng ta nên biến nó thành 1 object để chúng ta có thể control được ở bên trong
// Khi đấy hàm của mình không nằm trơ trọi ở bên ngoài mà nó phải là một method của một object đó
// Trong object đó nó có những thứ thuộc về nội tại của chính nó và nó sẽ thay đổi nội tại đó ở chính nó mà thôi
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

// Tại sao createRestaurantBiz phải là con trỏ
// => tối ưu
func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)

	return err
}

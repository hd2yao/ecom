package cart

import (
    "fmt"

    "github.com/hd2yao/ecom/types"
)

func getCartItemsIDs(items []types.CartCheckoutItem) ([]int, error) {
    productIDs := make([]int, len(items))
    for i, item := range items {
        if item.Quantity <= 0 {
            return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
        }
        productIDs[i] = item.ProductID
    }
    return productIDs, nil
}

func checkIfCartIsInStock(cartItems []types.CartCheckoutItem, products map[int]types.Product) error {
    if len(cartItems) == 0 {
        return fmt.Errorf("cart is empty")
    }

    for _, item := range cartItems {
        // sql 中是使用 IN 来根据 id 查询产品信息的，因此有可能会出现数据库中不存在的 id
        product, ok := products[item.ProductID]
        if !ok {
            return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
        }

        if product.Quantity < item.Quantity {
            return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
        }
    }

    return nil
}

func calculateTotalPrice(cartItems []types.CartCheckoutItem, products map[int]types.Product) float64 {
    var total float64

    for _, item := range cartItems {
        product := products[item.ProductID]
        total += product.Price * float64(item.Quantity)
    }

    return total
}

func (h *Handler) createOrder(products []types.Product, items []types.CartCheckoutItem, userID int) (int, float64, error) {
    // create a map of products for easier access
    productsMap := make(map[int]types.Product)
    for _, product := range products {
        productsMap[product.ID] = product
    }

    // check if all products are actually in stock
    if err := checkIfCartIsInStock(items, productsMap); err != nil {
        return 0, 0, err
    }

    // calculate total price
    totalPrice := calculateTotalPrice(items, productsMap)

    // reduce quantity of products in our db
    for _, item := range items {
        product := productsMap[item.ProductID]
        // 如果有多个请求进来，那么可能会出现问题
        // 更好的解决办法，是拆分到多表
        product.Quantity -= item.Quantity
        h.productStore.UpdateProduct(product)
    }

    // create order record
    orderID, err := h.store.CreateOrder(types.Order{
        UserID:  userID,
        Total:   totalPrice,
        Status:  "pending",
        Address: "some address", // 可以从用户地址表中获取地址
    })
    if err != nil {
        return 0, 0, err
    }

    // create order items records
    for _, item := range items {
        h.store.CreateOrderItem(types.OrderItem{
            OrderID:   orderID,
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
            Price:     productsMap[item.ProductID].Price,
        })
    }

    return orderID, totalPrice, nil
}

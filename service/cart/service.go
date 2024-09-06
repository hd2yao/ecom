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

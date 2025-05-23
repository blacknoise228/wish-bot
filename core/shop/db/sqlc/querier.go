// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateShop(ctx context.Context, arg CreateShopParams) (Shop, error)
	CreateShopAdmin(ctx context.Context, arg CreateShopAdminParams) error
	DeleteProduct(ctx context.Context, arg DeleteProductParams) error
	DeleteShop(ctx context.Context, arg DeleteShopParams) error
	DeleteShopAdmin(ctx context.Context, adminID int64) error
	GetDimOrderStatusByID(ctx context.Context, id int32) (DimOrderStatus, error)
	GetOrder(ctx context.Context, id uuid.UUID) (Order, error)
	GetOrdersByAdmin(ctx context.Context, adminID int64) ([]GetOrdersByAdminRow, error)
	GetOrdersByShop(ctx context.Context, shopID uuid.UUID) ([]GetOrdersByShopRow, error)
	GetProductByID(ctx context.Context, id uuid.UUID) (GetProductByIDRow, error)
	GetProductCategoryByName(ctx context.Context, categoryName string) (DimProductCategory, error)
	GetProductStatusByName(ctx context.Context, statusName string) (DimProductStatus, error)
	GetProducts(ctx context.Context, shopID uuid.UUID) ([]Product, error)
	GetRandomAdminByShopID(ctx context.Context, shopID uuid.UUID) (ShopAdmin, error)
	GetShopAdminsByAdminID(ctx context.Context, adminID int64) (ShopAdmin, error)
	GetShopAdminsByShopID(ctx context.Context, shopID uuid.UUID) ([]ShopAdmin, error)
	GetShopByID(ctx context.Context, id uuid.UUID) (Shop, error)
	GetShopByToken(ctx context.Context, token string) (Shop, error)
	GetShops(ctx context.Context) ([]Shop, error)
	GetUserInfo(ctx context.Context, chatID int64) (UserInfo, error)
	UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateProductStatus(ctx context.Context, arg UpdateProductStatusParams) error
	UpdateShop(ctx context.Context, arg UpdateShopParams) (Shop, error)
}

var _ Querier = (*Queries)(nil)

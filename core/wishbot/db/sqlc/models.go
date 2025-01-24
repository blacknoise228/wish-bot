// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type DimFriendStatus struct {
	ID         int32  `json:"id"`
	StatusName string `json:"status_name"`
}

type DimOrderStatus struct {
	ID         int32  `json:"id"`
	StatusCode string `json:"status_code"`
	StatusName string `json:"status_name"`
}

type DimProductCategory struct {
	ID           int32  `json:"id"`
	CategoryCode string `json:"category_code"`
	CategoryName string `json:"category_name"`
}

type DimProductStatus struct {
	ID         int32  `json:"id"`
	StatusCode string `json:"status_code"`
	StatusName string `json:"status_name"`
}

type DimWishStatus struct {
	ID         int32  `json:"id"`
	StatusName string `json:"status_name"`
}

type Friend struct {
	ChatID    int64     `json:"chat_id"`
	FriendID  int64     `json:"friend_id"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID            uuid.UUID        `json:"id"`
	Price         float64          `json:"price"`
	Status        int32            `json:"status"`
	CustomerID    int64            `json:"customer_id"`
	CustomerLogin string           `json:"customer_login"`
	ConsigneeID   int64            `json:"consignee_id"`
	ProductID     uuid.UUID        `json:"product_id"`
	AdminID       int64            `json:"admin_id"`
	ShopID        uuid.UUID        `json:"shop_id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}

type Product struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Price       float64          `json:"price"`
	Description string           `json:"description"`
	Image       string           `json:"image"`
	CategoryID  int32            `json:"category_id"`
	Status      int32            `json:"status"`
	ShopID      uuid.UUID        `json:"shop_id"`
	AdminID     int64            `json:"admin_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type Shop struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
	Image       string           `json:"image"`
	Token       string           `json:"token"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type ShopAdmin struct {
	AdminID int64     `json:"admin_id"`
	ShopID  uuid.UUID `json:"shop_id"`
}

type User struct {
	Username  string    `json:"username"`
	ChatID    int64     `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserInfo struct {
	ChatID      int64            `json:"chat_id"`
	Address     string           `json:"address"`
	Phone       string           `json:"phone"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type Wish struct {
	ID        int32     `json:"id"`
	ChatID    int64     `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
	ProductID uuid.UUID `json:"product_id"`
	Status    int32     `json:"status"`
}

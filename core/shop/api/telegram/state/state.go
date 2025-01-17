package state

const (
	AddAdmin              = "add_admin"
	CreateProduct         = "create_product"
	AddProductName        = "add_product_name"
	AddProductDesc        = "add_product_desc"
	AddProductPrice       = "add_product_price"
	AddProductCategory    = "add_product_category"
	AddProductImage       = "add_product_image"
	UpdateProduct         = "update_product"
	UpdateProductName     = "update_product_name"
	UpdateProductDesc     = "update_product_desc"
	UpdateProductPrice    = "update_product_price"
	UpdateProductCategory = "update_product_category"
	UpdateProductImage    = "update_product_image"
	UpdateProductStatus   = "update_product_status"
	SendPaymentLink       = "send_payment_link"
)

var UserStates = make(map[int64]string)

func SetUserState(chatID int64, usrstate string) {
	UserStates[chatID] = usrstate
}

func GetUserState(chatID int64) string {
	return UserStates[chatID]
}

func ClearUserState(chatID int64) {
	delete(UserStates, chatID)
}

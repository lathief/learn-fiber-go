package constant

import (
	"fmt"
)

var (
	ErrProductNotFound     = fmt.Errorf("Product not found")
	ErrOrderNotFound       = fmt.Errorf("Order not found")
	ErrCategoryNotFound    = fmt.Errorf("Category not found")
	ErrCartNotFound        = fmt.Errorf("Cart not found")
	ErrUserNotFound        = fmt.Errorf("User not found")
	ErrRoleNotFound        = fmt.Errorf("Role not found")
	ErrUsernameExists      = fmt.Errorf("Username has been used")
	ErrEmailExists         = fmt.Errorf("Email has been used")
	ErrPasswordNotMatch    = fmt.Errorf("Password does not match")
	ErrTokenInvalid        = fmt.Errorf("Token invalid")
	ErrUserNeedLogin       = fmt.Errorf("Login to proceed")
	ErrNotAllowedAccess    = fmt.Errorf("You are not allowed to access this data")
	ErrInternalServerError = fmt.Errorf("Internal server error")
)

func JoinMsgError(err error, msg string) string {
	return fmt.Sprintf("%s: %s", err.Error(), msg)
}

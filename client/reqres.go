// Package client provides the client-side functionality for interacting with the server.
package client

// Importing the decimal package for handling decimal numbers.
import "github.com/shopspring/decimal"

// Response is a generic struct that represents the server response for any type of request.
// It contains fields for success status, error message, and result data.
type Response[T any] struct {
	Success bool   `json:"success"` // Indicates if the request was successful.
	Error   string `json:"error"`   // Contains the error message, if any.
	Result  []T    `json:"result"`  // Contains the result data, if any.
}

// OnboardingRequest represents the data required for a client onboarding request.
// It includes fields for client status, wallet address, and signature.
type OnboardingRequest struct {
	IsClient  bool   `json:"is_client,omitempty"`                            // Indicates if the requester is a client.
	Wallet    string `json:"wallet,omitempty" binding:"len=42,required"`     // The wallet address of the client.
	Signature string `json:"signature,omitempty" binding:"len=132,required"` // The signature of the client.
}

// SecretRefreshRequest represents the data required to refresh the client's secret.
// It includes a field for the refresh token.
type SecretRefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // The refresh token of the client.
}

// OrderListRequest represents the data required to list the client's orders.
// It includes fields for market ID, timestamps, status, order ID, client order ID, and order type.
type OrderListRequest struct {
	MarketId      string   `json:"market_id" binding:"omitempty"`                                                                                                                      // The market ID of the order.
	TimeStamp     uint64   `json:"start_time,default=0" binding:"omitempty,min=0"`                                                                                                     // The start time of the order.
	EndTime       uint64   `json:"end_time,default=0" binding:"omitempty,min=0"`                                                                                                       // The end time of the order.
	Status        []string `json:"status" binding:"omitempty,dive,oneof=processing open closed rejected canceled canceling amending cancelingall placed"`                              // The status of the order.
	OrderId       string   `json:"order_id" binding:"omitempty"`                                                                                                                       // The order ID.
	ClientOrderId string   `json:"client_order_id" binding:"omitempty"`                                                                                                                // The client order ID.
	OrderType     []string `json:"order_type" binding:"omitempty,dive,oneof=limit market stop_loss take_profit stop_loss_limit take_profit_limit stop_market stop_limit cancel amend"` // The type of the order.
}

// OrderCreateRequest represents the data required to create a new order.
// It includes fields for market ID, type, side, price, size, client order ID, trigger price, size percent, and time in force.
type OrderCreateRequest struct {
	MarketId      string   `json:"market_id" binding:"required"`                                                                                                                               // The market ID of the order.
	Type          string   `json:"type" binding:"oneof=limit market stop_loss take_profit stop_loss_limit take_profit_limit stop_market stop_limit cancel amend,required"`                     // The type of the order.
	Side          string   `json:"side" binding:"required_unless=Type stop_loss Type take_profit Type stop_loss_limit Type take_profit_limit,omitempty,oneof=short long"`                      // The side of the order.
	Price         *float64 `json:"price" binding:"required_if=Type limit Type stop_limit Type stop_loss_limit Type take_profit_limit,omitempty"`                                               // The price of the order.
	Size          *float64 `json:"size" binding:"required_unless=Type stop_loss Type take_profit Type stop_loss_limit Type take_profit_limit,omitempty"`                                       // The size of the order.
	ClientOrderId *string  `json:"client_order_id" binding:"omitempty"`                                                                                                                        // The client order ID.
	TriggerPrice  *float64 `json:"trigger_price" binding:"required_if=Type stop_loss Type take_profit Type stop_loss_limit Type take_profit_limit Type stop_market Type stop_limit,omitempty"` // The trigger price of the order.
	SizePercent   *float64 `json:"size_percent" binding:"required_if=Type stop_loss Type take_profit Type stop_loss_limit Type take_profit_limit,omitempty"`                                   // The size percent of the order.
	TimeInForce   *string  `json:"time_in_force" binding:"omitempty,oneof=good_till_cancel immediate_or_cancel fill_or_kill post_only"`                                                        // The time in force of the order.
}

// OrderAmendRequest represents the data required to amend an existing order.
// It includes fields for order ID, market ID, price, size, trigger price, and size percent.
type OrderAmendRequest struct {
	OrderId      string   `json:"order_id" binding:"required"`       // The order ID.
	MarketId     string   `json:"market_id" binding:"required"`      // The market ID of the order.
	Price        *float64 `json:"price" binding:"omitempty"`         // The new price of the order.
	Size         *float64 `json:"size" binding:"omitempty"`          // The new size of the order.
	TriggerPrice *float64 `json:"trigger_price" binding:"omitempty"` // The new trigger price of the order.
	SizePercent  *float64 `json:"size_percent" binding:"omitempty"`  // The new size percent of the order.
}

// OrderCancelRequest represents the data required to cancel an existing order.
// It includes fields for order ID, market ID, and client order ID.
type OrderCancelRequest struct {
	OrderId       string `json:"order_id" binding:"omitempty"`        // The order ID.
	MarketId      string `json:"market_id" binding:"required"`        // The market ID of the order.
	ClientOrderId string `json:"client_order_id" binding:"omitempty"` // The client order ID.
}

// OrderCreateResponse represents the server response for a create order request.
// It includes fields for order ID, market ID, profile ID, status, size, price, side, type, liquidation status, client order ID, trigger price, size percent, and time in force.
type OrderCreateResponse struct {
	OrderId       string           `json:"id"`              // The order ID.
	MarketId      string           `json:"market_id"`       // The market ID of the order.
	ProfileId     uint             `json:"profile_id"`      // The profile ID of the client.
	Status        string           `json:"status"`          // The status of the order.
	Size          *decimal.Decimal `json:"size"`            // The size of the order.
	Price         *decimal.Decimal `json:"price"`           // The price of the order.
	Side          string           `json:"side"`            // The side of the order.
	Type          string           `json:"type"`            // The type of the order.
	IsLiquidation bool             `json:"is_liquidation"`  // Indicates if the order is a liquidation order.
	ClientOrderId *string          `json:"client_order_id"` // The client order ID.
	TriggerPrice  *decimal.Decimal `json:"trigger_price"`   // The trigger price of the order.
	SizePercent   *decimal.Decimal `json:"size_percent"`    // The size percent of the order.
	TimeInForce   *string          `json:"time_in_force"`   // The time in force of the order.
}

// OrderCancelResponse represents the server response for a cancel order request.
// It includes fields for order ID, market ID, profile ID, status, and client order ID.
type OrderCancelResponse struct {
	OrderId       string `json:"id"`              // The order ID.
	MarketId      string `json:"market_id"`       // The market ID of the order.
	ProfileId     uint   `json:"profile_id"`      // The profile ID of the client.
	Status        string `json:"status"`          // The status of the order.
	ClientOrderId string `json:"client_order_id"` // The client order ID.
}

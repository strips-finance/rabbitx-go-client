// Package model contains constants used in the model package.
package model

// Constants for profile statuses, types, and errors.
const (
	// PROFILE_STATUS_ACTIVE represents an active profile status.
	PROFILE_STATUS_ACTIVE = "active"

	// PROFILE_STATUS_LIQUIDATING represents a liquidating profile status.
	PROFILE_STATUS_LIQUIDATING = "liquidating"

	// PROFILE_TYPE_TRADER represents a trader profile type.
	PROFILE_TYPE_TRADER = "trader"

	// PROFILE_TYPE_INSURANCE represents an insurance profile type.
	PROFILE_TYPE_INSURANCE = "insurance"

	// PROFILE_NOT_FOUND_ERROR represents a profile not found error.
	PROFILE_NOT_FOUND_ERROR = "profile_not_found"
)

// Constants for order types.
const (
	// LONG represents a long order type.
	LONG = "long"

	// SHORT represents a short order type.
	SHORT = "short"

	// LIMIT represents a limit order type.
	LIMIT = "limit"

	// MARKET represents a market order type.
	MARKET = "market"

	// STOP_LOSS represents a stop loss order type.
	STOP_LOSS = "stop_loss"

	// TAKE_PROFIT represents a take profit order type.
	TAKE_PROFIT = "take_profit"
)

// Constants for stop order types.
const (
	// STOP_LIMIT represents a stop limit order type.
	STOP_LIMIT = "stop_limit"

	// STOP_MARKET represents a stop market order type.
	STOP_MARKET = "stop_market"
)

// Constants for account prefix and balance operations statuses.
const (
	// ACCOUNT_PREFIX represents the prefix for an account.
	ACCOUNT_PREFIX = "account@"

	// BALANCE_OPS_STATUS_PENDING represents a pending balance operation status.
	BALANCE_OPS_STATUS_PENDING = "pending"

	// BALANCE_OPS_STATUS_SUCCESS represents a successful balance operation status.
	BALANCE_OPS_STATUS_SUCCESS = "success"

	// BALANCE_OPS_STATUS_FAILED represents a failed balance operation status.
	BALANCE_OPS_STATUS_FAILED = "failed"

	// BALANCE_OPS_STATUS_TRANSFERING represents a transferring balance operation status.
	BALANCE_OPS_STATUS_TRANSFERING = "transferring"

	// BALANCE_OPS_STATUS_UNKNOWN represents an unknown balance operation status.
	BALANCE_OPS_STATUS_UNKNOWN = "unknown"
)

// Constants for order statuses.
const (
	// UNKNOWN represents an unknown order status.
	UNKNOWN = "unknown"

	// PROCESSING represents a processing order status.
	PROCESSING = "processing"

	// PLACED represents a placed order status.
	PLACED = "placed"

	// OPEN represents an open order status.
	OPEN = "open"

	// CLOSED represents a closed order status.
	CLOSED = "closed"

	// REJECTED represents a rejected order status.
	REJECTED = "rejected"

	// CANCELED represents a canceled order status.
	CANCELED = "canceled"

	// CANCELING represents a canceling order status.
	CANCELING = "canceling"

	// AMENDING represents an amending order status.
	AMENDING = "amending"

	// CANCELINGALL represents a canceling all order status.
	CANCELINGALL = "cancelingall"
)

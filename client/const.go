// Package client contains constants used for API communication.
package client

// Constants for API headers and paths.
const (
	// API_SECRET_SIGNATURE_HEADER is the header for the API secret signature.
	API_SECRET_SIGNATURE_HEADER = "RBT-SIGNATURE"

	// API_SECRET_TIMESTAMP_HEADER is the header for the API secret timestamp.
	API_SECRET_TIMESTAMP_HEADER = "RBT-TS"

	// PK_SIGNATURE_HEADER is the header for the PK signature.
	PK_SIGNATURE_HEADER = "RBT-PK-SIGNATURE"

	// PK_TIMESTAMP_HEADER is the header for the PK timestamp.
	PK_TIMESTAMP_HEADER = "RBT-PK-TS"

	// API_KEY_HEADER is the header for the API key.
	API_KEY_HEADER = "RBT-API-KEY"

	// SIGNATURE_LIFETIME is the lifetime of the signature in seconds.
	SIGNATURE_LIFETIME = 300

	// PATH_ONBOARDING is the API path for onboarding.
	PATH_ONBOARDING = "/onboarding"

	// PATH_MARKETS is the API path for markets.
	PATH_MARKETS = "/markets"

	// PATH_ORDERS is the API path for orders.
	PATH_ORDERS = "/orders"

	// PATH_ORDERS_CANCEL_ALL is the API path to cancel all orders.
	PATH_ORDERS_CANCEL_ALL = "/orders/cancel_all"

	// PATH_ORDERS_LIST is the API path to list orders.
	PATH_ORDERS_LIST = "/orders/list"

	// PATH_JWT is the API path for JWT.
	PATH_JWT = "/jwt"

	// PATH_ACCOUNT is the API path for account.
	PATH_ACCOUNT = "/account"

	// PATH_POSITIONS is the API path to list positions.
	PATH_POSITIONS = "/positions/list"

	// PATH_DEPOSIT is the API path for deposit.
	PATH_DEPOSIT = "/balanceops/deposit"

	// PATH_WITHDRAW is the API path for withdrawal.
	PATH_WITHDRAW = "/balanceops/withdraw"

	// PATH_CANCEL_WITHDRAWAL is the API path to cancel withdrawal.
	PATH_CANCEL_WITHDRAWAL = "/balanceops/cancel"

	// PATH_CLAIM_WITHDRAWAL is the API path to claim withdrawal.
	PATH_CLAIM_WITHDRAWAL = "/balanceops/claim"

	// PATH_SECRETS is the API path for secrets.
	PATH_SECRETS = "/secrets"

	// PATH_SECRETS_REFRESH is the API path to refresh secrets.
	PATH_SECRETS_REFRESH = "/secrets/refresh"
)

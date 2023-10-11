// Package model contains all data types returned by rabbitx and error codes.
package model

import "github.com/shopspring/decimal"

// ProfileNotification represents a notification related to a profile.
type ProfileNotification struct {
	NotifType   string  `json:"type"`                  // The type of the notification.
	Title       *string `json:"title,omitempty"`       // The title of the notification.
	Description *string `json:"description,omitempty"` // The description of the notification.
}

// ProfileCache represents a cached profile.
type ProfileCache struct {
	ProfileID           uint                       `json:"id"`                              // The ID of the profile.
	ProfileType         *string                    `json:"profile_type,omitempty"`          // The type of the profile.
	Status              *string                    `json:"status,omitempty"`                // The status of the profile.
	Wallet              *string                    `json:"wallet,omitempty"`                // The wallet of the profile.
	LastUpdate          *int64                     `json:"last_update,omitempty"`           // The last update time of the profile.
	Balance             *decimal.Decimal           `json:"balance,omitempty"`               // The balance of the profile.
	AccountEquity       *decimal.Decimal           `json:"account_equity,omitempty"`        // The account equity of the profile.
	TotalPositionMargin *decimal.Decimal           `json:"total_position_margin,omitempty"` // The total position margin of the profile.
	TotalOrderMargin    *decimal.Decimal           `json:"total_order_margin,omitempty"`    // The total order margin of the profile.
	TotalNotional       *decimal.Decimal           `json:"total_notional,omitempty"`        // The total notional of the profile.
	AccountMargin       *decimal.Decimal           `json:"account_margin,omitempty"`        // The account margin of the profile.
	WithdrawbleBalance  *decimal.Decimal           `json:"withdrawable_balance,omitempty"`  // The withdrawable balance of the profile.
	CumUnrealizedPnl    *decimal.Decimal           `json:"cum_unrealized_pnl,omitempty"`    // The cumulative unrealized profit and loss of the profile.
	Health              *decimal.Decimal           `json:"health,omitempty"`                // The health of the profile.
	AccountLeverage     *decimal.Decimal           `json:"account_leverage,omitempty"`      // The account leverage of the profile.
	CumTradingVolume    *decimal.Decimal           `json:"cum_trading_volume,omitempty"`    // The cumulative trading volume of the profile.
	Leverage            map[string]decimal.Decimal `json:"leverage,omitempty"`              // The leverage of the profile.
	LastLiqCheck        *int64                     `json:"last_liq_check,omitempty"`        // The last liquidation check of the profile.
	ShardId             string                     `json:"-"`                               // The shard ID of the profile.
	ArchiveId           int                        `json:"-"`                               // The archive ID of the profile.
}

// ProfileData represents a profile with its positions, orders, and notifications.
type ProfileData struct {
	ProfileCache                         // The cached profile.
	Positions     []*PositionData        `json:"positions,omitempty"`             // The positions of the profile.
	Orders        []*OrderData           `json:"orders,omitempty"`                // The orders of the profile.
	Notifications []*ProfileNotification `json:"profile_notifications,omitempty"` // The notifications of the profile.
}

// ExtendedProfileData represents a profile with its extended positions, orders, and notifications.
type ExtendedProfileData struct {
	ProfileCache                          // The cached profile.
	Positions     []*ExtendedPositionData `json:"positions,omitempty"`             // The extended positions of the profile.
	Orders        []*OrderData            `json:"orders,omitempty"`                // The orders of the profile.
	Notifications []*ProfileNotification  `json:"profile_notifications,omitempty"` // The notifications of the profile.
}

// OrderData represents an order.
type OrderData struct {
	OrderId         string           `json:"id"`                          // The ID of the order.
	ProfileID       uint             `json:"profile_id"`                  // The profile ID of the order.
	MarketID        string           `json:"market_id"`                   // The market ID of the order.
	OrderType       string           `json:"order_type"`                  // The type of the order.
	Status          string           `json:"status"`                      // The status of the order.
	Price           *decimal.Decimal `json:"price,omitempty"`             // The price of the order.
	Size            *decimal.Decimal `json:"size,omitempty"`              // The size of the order.
	InitialSize     *decimal.Decimal `json:"initial_size,omitempty"`      // The initial size of the order.
	TotalFilledSize *decimal.Decimal `json:"total_filled_size,omitempty"` // The total filled size of the order.
	Side            string           `json:"side"`                        // The side of the order.
	Timestamp       int64            `json:"timestamp"`                   // The timestamp of the order.
	Reason          string           `json:"reason"`                      // The reason of the order.
	ClientOrderId   *string          `json:"client_order_id,omitempty"`   // The client order ID of the order.
	TriggerPrice    *decimal.Decimal `json:"trigger_price,omitempty"`     // The trigger price of the order.
	SizePercent     *decimal.Decimal `json:"size_percent,omitempty"`      // The size percent of the order.
	TimeInForce     string           `json:"time_in_force"`               // The time in force of the order.
	CreatedAt       int64            `json:"created_at"`                  // The creation time of the order.
	UpdatedAt       int64            `json:"updated_at"`                  // The update time of the order.
	ShardId         string           `json:"-"`                           // The shard ID of the order.
	ArchiveId       int              `json:"-"`                           // The archive ID of the order.
}

// PositionData represents a position.
type PositionData struct {
	PositionID        string           `json:"id"`                          // The ID of the position.
	MarketID          string           `json:"market_id"`                   // The market ID of the position.
	ProfileID         uint             `json:"profile_id"`                  // The profile ID of the position.
	Size              decimal.Decimal  `json:"size"`                        // The size of the position.
	Side              string           `json:"side"`                        // The side of the position.
	EntryPrice        decimal.Decimal  `json:"entry_price"`                 // The entry price of the position.
	UnrealizedPnlFair *decimal.Decimal `json:"unrealized_pnl,omitempty"`    // The unrealized profit and loss of the position.
	NotionalFair      *decimal.Decimal `json:"notional,omitempty"`          // The notional of the position.
	Margin            *decimal.Decimal `json:"margin,omitempty"`            // The margin of the position.
	LiquidationPrice  *decimal.Decimal `json:"liquidation_price,omitempty"` // The liquidation price of the position.
	FairPrice         *decimal.Decimal `json:"fair_price,omitempty"`        // The fair price of the position.
	ShardId           string           `json:"-"`                           // The shard ID of the position.
	ArchiveId         int              `json:"-"`                           // The archive ID of the position.
}

// ExtendedPositionData represents an extended position.
type ExtendedPositionData struct {
	PositionData            // The position.
	StopLoss     *OrderData `json:"stop_loss"`   // The stop loss order of the position.
	TakeProfit   *OrderData `json:"take_profit"` // The take profit order of the position.
}

// MarketData represents market data.
type MarketData struct {
	MarketID          string           `json:"id"`                           // The ID of the market.
	Status            *string          `json:"status,omitempty"`             // The status of the market.
	MinInitialMargin  *decimal.Decimal `json:"min_initial_margin,omitempty"` // The minimum initial margin of the market.
	ForcedMargin      *decimal.Decimal `json:"forced_margin,omitempty"`      // The forced margin of the market.
	LiquidationMargin *decimal.Decimal `json:"liquidation_margin,omitempty"` // The liquidation margin of the market.
	MinTick           *decimal.Decimal `json:"min_tick,omitempty"`           // The minimum tick of the market.
	MinOrder          *decimal.Decimal `json:"min_order,omitempty"`          // The minimum order of the market.
	BestBid           *decimal.Decimal `json:"best_bid,omitempty"`           // The best bid of the market.
	BestAsk           *decimal.Decimal `json:"best_ask,omitempty"`           // The best ask of the market.
	MarketPrice       *decimal.Decimal `json:"market_price,omitempty"`       // The market price of the market.
	IndexPrice        *decimal.Decimal `json:"index_price,omitempty"`        // The index price of the market.
	LastTradePrice    *decimal.Decimal `json:"last_trade_price,omitempty"`   // The last trade price of the market.
	FairPrice         *decimal.Decimal `json:"fair_price,omitempty"`         // The fair price of the market.

	InstantFundingRate *decimal.Decimal `json:"instant_funding_rate,omitempty"`    // The instant funding rate of the market.
	LastFundingRate    *decimal.Decimal `json:"last_funding_rate_basis,omitempty"` // The last funding rate of the market.

	LastUpdateTime        int64            `json:"last_update_time,omitempty"`         // The last update time of the market.
	LastUpdateSequence    int64            `json:"last_update_sequence,omitempty"`     // The last update sequence of the market.
	AverageDailyVolumeQ   *decimal.Decimal `json:"average_daily_volume_q,omitempty"`   // The average daily volume of the market.
	LastFundingUpdateTime int64            `json:"last_funding_update_time,omitempty"` // The last funding update time of the market.
	IconUrl               string           `json:"icon_url"`                           // The icon URL of the market.
	MarketTitle           string           `json:"market_title"`                       // The title of the market.

	ShardId   string `json:"-"` // The shard ID of the market.
	ArchiveId int    `json:"-"` // The archive ID of the market.
}

// TradeData represents a trade.
type TradeData struct {
	TradeId     string          `json:"id"`          // The ID of the trade.
	MarketId    string          `json:"market_id"`   // The market ID of the trade.
	Timestamp   uint64          `json:"timestamp"`   // The timestamp of the trade.
	Price       decimal.Decimal `json:"price"`       // The price of the trade.
	Size        decimal.Decimal `json:"size"`        // The size of the trade.
	Liquidation bool            `json:"liquidation"` // Whether the trade is a liquidation.
	TakerSide   string          `json:"taker_side"`  // The taker side of the trade.
	ShardId     string          `json:"-"`           // The shard ID of the trade.
	ArchiveId   uint64          `json:"-"`           // The archive ID of the trade.
}

// FillData represents a fill.
type FillData struct {
	Id            string          `json:"id"`                        // The ID of the fill.
	ProfileId     uint            `json:"profile_id"`                // The profile ID of the fill.
	MarketId      string          `json:"market_id"`                 // The market ID of the fill.
	OrderId       string          `json:"order_id"`                  // The order ID of the fill.
	Timestamp     int64           `json:"timestamp"`                 // The timestamp of the fill.
	TradeId       string          `json:"trade_id"`                  // The trade ID of the fill.
	Price         decimal.Decimal `json:"price"`                     // The price of the fill.
	Size          decimal.Decimal `json:"size"`                      // The size of the fill.
	Side          string          `json:"side"`                      // The side of the fill.
	IsMaker       bool            `json:"is_maker"`                  // Whether the fill is a maker.
	Fee           decimal.Decimal `json:"fee"`                       // The fee of the fill.
	Liquidation   bool            `json:"liquidation"`               // Whether the fill is a liquidation.
	ClientOrderId *string         `json:"client_order_id,omitempty"` // The client order ID of the fill.

	ShardId   string `json:"-"` // The shard ID of the fill.
	ArchiveId int    `json:"-"` // The archive ID of the fill.
}

// OrderbookData represents an order book.
type OrderbookData struct {
	MarketID  string              `json:"market_id"`      // The market ID of the order book.
	Bids      [][]decimal.Decimal `json:"bids,omitempty"` // The bids of the order book.
	Asks      [][]decimal.Decimal `json:"asks,omitempty"` // The asks of the order book.
	Sequence  uint                `json:"sequence"`       // The sequence of the order book.
	Timestamp int64               `json:"timestamp"`      // The timestamp of the order book.
}

// UntypedOrderbookData represents an untyped order book.
type UntypedOrderbookData struct {
	MarketID  string      `json:"market_id"`      // The market ID of the order book.
	Bids      interface{} `json:"bids,omitempty"` // The bids of the order book.
	Asks      interface{} `json:"asks,omitempty"` // The asks of the order book.
	Sequence  uint        `json:"sequence"`       // The sequence of the order book.
	Timestamp int64       `json:"timestamp"`      // The timestamp of the order book.
}

// Profile represents a profile.
type Profile struct {
	ProfileId uint   `json:"id"`           // The ID of the profile.
	Type      string `json:"profile_type"` // The type of the profile.
	Status    string `json:"status"`       // The status of the profile.
	Wallet    string `json:"wallet"`       // The wallet of the profile.
	CreatedAt int64  `json:"created_at"`   // The creation time of the profile.
}

// BalanceOps represents a balance operation.
type BalanceOps struct {
	OpsId     string          `json:"id"`         // The ID of the balance operation.
	Status    string          `json:"status"`     // The status of the balance operation.
	Reason    string          `json:"reason"`     // The reason of the balance operation.
	Txhash    string          `json:"txhash"`     // The transaction hash of the balance operation.
	ProfileId uint            `json:"profile_id"` // The profile ID of the balance operation.
	Wallet    string          `json:"wallet"`     // The wallet of the balance operation.
	Type      string          `json:"ops_type"`   // The type of the balance operation.
	Id2       string          `json:"ops_id2"`    // The secondary ID of the balance operation.
	Amount    decimal.Decimal `json:"amount"`     // The amount of the balance operation.
	Timestamp int64           `json:"timestamp"`  // The timestamp of the balance operation.
	DueBlock  uint            `json:"due_block"`  // The due block of the balance operation.
	ShardId   string          `json:"shard_id"`   // The shard ID of the balance operation.
	ArchiveId int             `json:"-"`          // The archive ID of the balance operation.
}

// CandleData represents a candle.
type CandleData struct {
	Time   int64           `json:"time"`   // The time of the candle.
	Low    decimal.Decimal `json:"low"`    // The low of the candle.
	High   decimal.Decimal `json:"high"`   // The high of the candle.
	Open   decimal.Decimal `json:"open"`   // The open of the candle.
	Close  decimal.Decimal `json:"close"`  // The close of the candle.
	Volume decimal.Decimal `json:"volume"` // The volume of the candle.
}

// ExchangeData represents exchange data.
type ExchangeData struct {
	Id           int64           `json:"id"`            // The ID of the exchange.
	TradingFee   decimal.Decimal `json:"trading_fee"`   // The trading fee of the exchange.
	TotalBalance decimal.Decimal `json:"total_balance"` // The total balance of the exchange.
}

// OnboardMarketMakerResult represents the result of onboarding a market maker.
type OnboardMarketMakerResult struct {
	Profile   *ProfileData `json:"profile"`   // The profile of the market maker.
	APISecret *APISecret   `json:"apiSecret"` // The API secret of the market maker.
	Jwt       string       `json:"jwt"`       // The JWT of the market maker.
}

// APISecret represents an API secret.
type APISecret struct {
	Key        string `json:"Key"`        // The key of the API secret.
	ProfileID  uint   `json:"ProfileID"`  // The profile ID of the API secret.
	Secret     string `json:"Secret"`     // The secret of the API secret.
	Tag        string `json:"Tag"`        // The tag of the API secret.
	Expiration uint   `json:"Expiration"` // The expiration of the API secret.
	Status     string `json:"Status"`     // The status of the API secret.
}

// Secret represents a secret.
type Secret struct {
	APISecret     *APISecret `json:"api_secret"`      // The API secret of the secret.
	JwtPrivate    string     `json:"jwt_private"`     // The private JWT of the secret.
	JwtPublic     string     `json:"jwt_public"`      // The public JWT of the secret.
	RefreshToken  string     `json:"refresh_token"`   // The refresh token of the secret.
	AllowedIpList []string   `json:"allowed_ip_list"` // The allowed IP list of the secret.
	CreatedAt     int64      `json:"created_at"`      // The creation time of the secret.
}

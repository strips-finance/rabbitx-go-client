// Package bot contains the main logic for the trading bot.
package bot

// Importing necessary libraries.
import (
	"rabbitx-client/client"
	"rabbitx-client/model"
	"strings"
	"sync"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

// Constants for default market ID, price tick and size tick.
const (
	DEFAULT_MARKET_ID  = "ETH-USD"
	DEFAULT_PRICE_TICK = 0.1
	DEFAULT_SIZE_TICK  = 0.001
)

// EventData struct holds the websocket channel and data.
type EventData struct {
	WsChannel string
	Data      interface{}
}

// WatchDog struct holds the market ID, client, orders, data channel, done channel, best bid and best ask.
type WatchDog struct {
	marketId string
	muOrder  sync.RWMutex
	client   *client.RbClient
	orders   map[string]string
	dataCh   chan EventData
	done     chan struct{}
	muMarket sync.RWMutex
	bestBid  decimal.Decimal
	bestAsk  decimal.Decimal
}

// NewWatchDog function initializes a new WatchDog.
func NewWatchDog(client *client.RbClient, marketId string, dataCh chan EventData, done chan struct{}) *WatchDog {
	return &WatchDog{
		marketId: marketId,
		client:   client,
		dataCh:   dataCh,
		done:     done,
		orders:   make(map[string]string),
	}
}

// Run function starts the WatchDog.
func (wd *WatchDog) Run() error {
	err := wd.loadOrders()
	if err != nil {
		return err
	}

	go wd.listener([]string{"account"})

	cancelCh := make(chan string, 10000)
	go wd.strategy(cancelCh)

	go wd.monitorOrders(cancelCh)

	return nil
}

// loadOrders function loads the orders from the client.
func (wd *WatchDog) loadOrders() error {
	orders, err := wd.client.ListOrders(&client.OrderListRequest{
		MarketId: wd.marketId,
	})
	if err != nil {
		return err
	}

	wd.muOrder.Lock()
	defer wd.muOrder.Unlock()
	for _, order := range orders {
		if order.Status == model.OPEN {
			wd.orders[order.OrderId] = order.Status
		}
	}

	return nil
}

// strategy function runs the trading strategy.
func (wd *WatchDog) strategy(cancelCh chan string) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	go wd.cancelOrders(cancelCh)

	for {
		select {
		case <-ticker.C:
			wd.placeOrder()
		case <-wd.done:
			return
		}
	}
}

// placeOrder function places an order.
func (wd *WatchDog) placeOrder() {
	wd.muMarket.RLock()
	price := wd.bestBid.Mul(decimal.NewFromFloat(0.94))
	wd.muMarket.RUnlock()

	if price.LessThanOrEqual(decimal.Zero) {
		return
	}

	order, err := wd.client.CreateOrder(&client.OrderCreateRequest{
		MarketId: wd.marketId,
		Type:     model.LIMIT,
		Side:     model.LONG,
		Price:    roundToNearestTick(price.InexactFloat64(), DEFAULT_PRICE_TICK),
		Size:     roundToNearestTick(DEFAULT_SIZE_TICK, DEFAULT_SIZE_TICK),
	})
	if err != nil {
		logrus.Error("Failed to create order: ", err)
		return
	}

	wd.muOrder.Lock()
	wd.orders[order.OrderId] = order.Status
	wd.muOrder.Unlock()

	logrus.Infof("Order created id : %s", order.OrderId)
}

// cancelOrders function cancels the orders.
func (wd *WatchDog) cancelOrders(cancelCh chan string) {
	for {
		select {
		case oid := <-cancelCh:
			wd.cancelOrder(oid)
		case <-wd.done:
			return
		}
	}
}

// cancelOrder function cancels a specific order.
func (wd *WatchDog) cancelOrder(oid string) {
	order, err := wd.client.CancelOrder(&client.OrderCancelRequest{
		OrderId:  oid,
		MarketId: wd.marketId,
	})
	if err != nil {
		logrus.Error("Failed to cancel order: ", err)
		return
	}

	wd.muOrder.Lock()
	wd.orders[order.OrderId] = order.Status
	wd.muOrder.Unlock()

	logrus.Infof("Order created id : %s", order.OrderId)
}

// monitorOrders function monitors the orders.
func (wd *WatchDog) monitorOrders(cancelCh chan string) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	cleanTicker := time.NewTicker(300 * time.Second)
	defer cleanTicker.Stop()

	for {
		select {
		case <-ticker.C:
			wd.checkOrders(cancelCh)
		case <-cleanTicker.C:
			wd.cleanOrders()
		case <-wd.done:
			return
		}
	}
}

// checkOrders function checks the orders.
func (wd *WatchDog) checkOrders(cancelCh chan string) {
	wd.muOrder.RLock()
	defer wd.muOrder.RUnlock()
	for id, status := range wd.orders {
		if status == model.OPEN || status == model.CANCELING {
			cancelCh <- id
		}
	}
}

// cleanOrders function cleans the orders.
func (wd *WatchDog) cleanOrders() {
	wd.muOrder.Lock()
	defer wd.muOrder.Unlock()
	for id, status := range wd.orders {
		if status == model.CANCELED {
			delete(wd.orders, id)
		}
	}
}

// listener function starts the listener.
func (wd *WatchDog) listener(display []string) {
	logrus.Info("listener started")

	for {
		select {
		case data := <-wd.dataCh:
			wd.handleData(data, display)
		case <-wd.done:
			logrus.Info("listener stopped")
			return
		}
	}
}

// handleData function handles the data from the listener.
func (wd *WatchDog) handleData(data EventData, display []string) {
	var channel string
	if strings.Contains(data.WsChannel, "@") {
		channel = strings.Split(data.WsChannel, "@")[0]
	} else if strings.Contains(data.WsChannel, ":") {
		channel = strings.Split(data.WsChannel, ":")[0]
	} else {
		logrus.Warnf("Unknown channel format = %s", data.WsChannel)
		return
	}

	show := false
	if slices.Contains(display, channel) {
		show = true
	}

	switch channel {
	case "market":
		wd.handleMarketData(data, show)
	case "account":
		wd.handleAccountData(data, show)
	case "orderbook":
		decodeAndPrintData[model.OrderbookData](channel, data.Data.([]byte), show)
	case "trade":
		decodeAndPrintData[model.TradeData](channel, data.Data.([]byte), show)
	default:
		logrus.
			WithField("channel", channel).
			Error("Unknown channel")
	}
}

// handleMarketData function handles the market data.
func (wd *WatchDog) handleMarketData(data EventData, show bool) {
	res := decodeAndPrintData[model.MarketData](data.WsChannel, data.Data.([]byte), show)
	if res == nil {
		return
	}

	wd.muMarket.Lock()
	defer wd.muMarket.Unlock()

	if res.BestAsk != nil && res.BestAsk.Abs().GreaterThan(decimal.Zero) {
		wd.bestAsk = *res.BestAsk
	}

	if res.BestBid != nil && res.BestBid.Abs().GreaterThan(decimal.Zero) {
		wd.bestBid = *res.BestBid
	}
}

// handleAccountData function handles the account data.
func (wd *WatchDog) handleAccountData(data EventData, show bool) {
	res := decodeAndPrintData[model.ProfileData](data.WsChannel, data.Data.([]byte), show)
	if res == nil {
		return
	}

	wd.muOrder.Lock()
	defer wd.muOrder.Unlock()
	for _, order := range res.Orders {
		wd.orders[order.OrderId] = order.Status
	}
}

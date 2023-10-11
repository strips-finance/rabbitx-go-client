// Package bot - dummy bot that executes trading on rabbitx
package bot

import (
	"rabbitx-client/client"
	"strconv"
	"strings"
	"time"

	"github.com/centrifugal/centrifuge-go"
	"github.com/sirupsen/logrus"
)

var channelPrefix = []string{"market:", "orderbook:", "trade:", "account@"}

// DummyBot is a struct that represents a dummy bot for executing trades on RabbitX.
// It contains the marketId, client, wsUrl, jwtPrivate, profileID, wsClient and done.
type DummyBot struct {
	marketId   string
	client     *client.RbClient
	wsUrl      string
	jwtPrivate string
	profileID  uint
	wsClient   *centrifuge.Client
	done       chan struct{}
}

// NewBot is a function that creates a new DummyBot.
// It takes a client, wsUrl and jwtPrivate as parameters and returns a pointer to a DummyBot.
func NewBot(client *client.RbClient, wsUrl, jwtPrivate string) *DummyBot {
	return &DummyBot{
		client:     client,
		wsUrl:      wsUrl,
		jwtPrivate: jwtPrivate,
		done:       make(chan struct{}),
	}
}

// Run is a method of DummyBot that runs the bot for a specific marketID.
// It sets the profile ID, connects the client to the websocket, subscribes to open channels and runs the watchdog.
// It returns an error if any.
func (b *DummyBot) Run(marketID string) error {
	// Set profile Id
	profileData, err := b.client.GetProfile()
	if err != nil {
		return err
	}

	b.marketId = marketID
	b.profileID = profileData.ProfileID
	logrus.Infof("ProfileId = %d detected", b.profileID)

	// Connect client to websocket
	b.wsClient = centrifuge.NewJsonClient(
		b.wsUrl,
		centrifuge.Config{
			Token:            b.jwtPrivate,
			ReadTimeout:      10 * time.Second,
			WriteTimeout:     10 * time.Second,
			HandshakeTimeout: 10 * time.Second,
		},
	)

	b.wsClient.OnConnecting(func(e centrifuge.ConnectingEvent) {
		logrus.Infof("Connecting - %d (%s) url: %s", e.Code, e.Reason, b.wsUrl)
	})

	b.wsClient.OnConnected(func(e centrifuge.ConnectedEvent) {
		logrus.Infof("Connected with ID %s", e.ClientID)
	})

	b.wsClient.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		logrus.Infof("Disconnected: %d (%s)", e.Code, e.Reason)
	})

	b.wsClient.OnError(func(e centrifuge.ErrorEvent) {
		logrus.Errorf("Websocket server connection Error: %s", e.Error.Error())
	})

	err = b.wsClient.Connect()
	if err != nil {
		return err
	}

	logrus.Info("Subscribing...")

	// Create a channel for event data
	dataCh := make(chan EventData)

	// Subscribe to open channels
	for _, prefix := range channelPrefix {
		var channel string
		if strings.HasPrefix(prefix, "account@") {
			channel = prefix + strconv.Itoa(int(b.profileID))
		} else {
			channel = prefix + b.marketId
		}

		logrus.Infof("Subscribing to channel: %s", channel)
		sub, err := b.wsClient.NewSubscription(channel, centrifuge.SubscriptionConfig{
			Recoverable: true,
		})
		if err != nil {
			return err
		}

		sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
			logrus.Info("Subscribed successfully to channel: ", channel)
		})

		sub.OnPublication(func(e centrifuge.PublicationEvent) {
			//Potential block here -that's why we need goroutine
			//OnPublication handler can't be block, or it will block the whole event loop
			go func() {
				dataCh <- EventData{
					WsChannel: channel,
					Data:      e.Data,
				}
			}()
		})

		err = sub.Subscribe()
		if err != nil {
			return err
		}
	}

	wd := NewWatchDog(b.client, b.marketId, dataCh, b.done)

	return wd.Run()
}

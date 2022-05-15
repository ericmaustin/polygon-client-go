# Polygon Go Client

[![Build][build-img]][build]

Originally a fork of the official Go client library for the [Polygon](https://polygon.io/) REST and WebSocket API.
This fork **may** be divergent enough from the original client library to be considered incompatible so use at your own
risk.

**This fork contains additional models and methods for OPTIONS api.**

`go get github.com/ericmaustin/polygon-client-go`

This client is still in pre-release. The public interface is relatively stable at this point but is still liable to
change slightly until we release v1. It also makes use of Go generics and thus requires Go 1.18.

See the [docs](https://polygon.io/docs/stocks/getting-started) for more details on our API.

## REST API Client

To get started, you'll need to import two main packages.

```golang
import (
polygon "github.com/polygon-io/client-go/rest"
"github.com/polygon-io/client-go/rest/models"
)
```

Next, create a new client with your [API key](https://polygon.io/dashboard/signup).

```golang
c := polygon.New("YOUR_API_KEY")
```

### Using the client

After creating the client, making calls to the Polygon API is simple.

```golang
params := models.GetAllTickersSnapshotParams{
Locale:     models.US,
MarketType: models.Stocks,
}.WithTickers("AAPL,MSFT")

res, err := c.GetAllTickersSnapshot(context.Background(), params)
if err != nil {
log.Fatal(err)
}
log.Print(res) // do something with the result
```

### Pagination

Our list methods return iterators that handle pagination for you.

```golang
// create a new iterator
params := models.ListTradesParams{Ticker: "AAPL"}.
WithTimestamp(models.GTE, models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).
WithOrder(models.Asc)
iter := c.ListTrades(context.Background(), params)

// iter.Next() advances the iterator to the next value in the list
for iter.Next() {
log.Print(iter.Item()) // do something with the current value
}

// if the loop breaks, it has either reached the end of the list or an error has occurred
// you can check if something went wrong with iter.Err()
if iter.Err() != nil {
log.Fatal(iter.Err())
}
```

### Request options

Advanced users may want to add additional headers or query params to a given request.

```golang
params := &models.GetGroupedDailyAggsParams{
Locale:     models.US,
MarketType: models.Stocks,
Date:       models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)),
}

res, err := c.GetGroupedDailyAggs(context.Background(), params,
models.APIKey("YOUR_OTHER_API_KEY"),
models.Header("X-CUSTOM-HEADER", "VALUE"),
models.QueryParam("adjusted", strconv.FormatBool(true)))
if err != nil {
log.Fatal(err)
}
log.Print(res) // do something with the result
```

## WebSocket Client

Import the WebSocket client and models packages to get started.

```golang
import (
polygonws "github.com/polygon-io/client-go/websocket"
"github.com/polygon-io/client-go/websocket/models"
)
```

Next, create a new client with your API key and a couple other config options.

```golang
// create a new client
c, err := polygonws.New(polygonws.Config{
APIKey:    "YOUR_API_KEY",
Feed:      polygonws.RealTime,
Market:    polygonws.Stocks,
})
if err != nil {
log.Fatal(err)
}
defer c.Close() // the user of this client must close it

// connect to the server
if err := c.Connect(); err != nil {
log.Fatal(err)
}
```

The client automatically reconnects to the server when the connection is dropped. By default, it will attempt to
reconnect indefinitely but the number of retries is configurable. When the client successfully reconnects, it
automatically resubscribes to any topics that were set before the disconnect.

### Using the client

After creating a client, subscribe to one or more topics and start accessing data. Currently, all of the data is pushed
to a single output channel.

```golang
// passing a topic by itself will subscribe to all tickers
if err := c.Subscribe(polygonws.StocksSecAggs); err != nil {
log.Fatal(err)
}
if err := c.Subscribe(polygonws.StocksTrades, "TSLA", "GME"); err != nil {
log.Fatal(err)
}

for {
select {
case err := <-c.Error(): // check for any fatal errors (e.g. auth failed)
log.Fatal(err)
case out, more := <-c.Output(): // read the next data message
if !more {
return
}

switch out.(type) {
case models.EquityAgg:
log.Print(out) // do something with the agg
case models.EquityTrade:
log.Print(out) // do something with the trade
}
}
}
```

See the [full example](./websocket/example/main.go) for more details on how to use this client effectively.

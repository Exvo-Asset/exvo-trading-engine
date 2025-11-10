package book

import (
	"sort"
	"sync"
)

type Order struct {
	ID     string
	Side   string // "BUY" or "SELL"
	Price  float64
	Qty    int64
	Symbol string
}

type OrderBook struct {
	sync.Mutex
	bids []Order // BUY orders (sorted desc price)
	asks []Order // SELL orders (sorted asc price)
}

func NewBook() *OrderBook {
	return &OrderBook{
		bids: make([]Order, 0),
		asks: make([]Order, 0),
	}
}

// Submit adds order to book and attempts naive matching.
// Returns ack string (accepted/partial/filled)
func (b *OrderBook) Submit(o Order) string {
	b.Lock()
	defer b.Unlock()

	if o.Side == "BUY" {
		// attempt to match against asks
		for i := 0; i < len(b.asks) && o.Qty > 0; {
			if o.Price >= b.asks[i].Price {
				// match
				if o.Qty >= b.asks[i].Qty {
					// consume ask
					o.Qty -= b.asks[i].Qty
					// remove ask
					b.asks = append(b.asks[:i], b.asks[i+1:]...)
					continue
				} else {
					// partial fill
					b.asks[i].Qty -= o.Qty
					o.Qty = 0
					break
				}
			}
			i++
		}
		if o.Qty > 0 {
			b.bids = append(b.bids, o)
			sort.Slice(b.bids, func(i, j int) bool {
				return b.bids[i].Price > b.bids[j].Price
			})
			if o.Qty == 0 {
				return "filled"
			}
			return "accepted"
		}
		return "filled"
	} else {
		// SELL side: match against bids
		for i := 0; i < len(b.bids) && o.Qty > 0; {
			if o.Price <= b.bids[i].Price {
				if o.Qty >= b.bids[i].Qty {
					o.Qty -= b.bids[i].Qty
					b.bids = append(b.bids[:i], b.bids[i+1:]...)
					continue
				} else {
					b.bids[i].Qty -= o.Qty
					o.Qty = 0
					break
				}
			}
			i++
		}
		if o.Qty > 0 {
			b.asks = append(b.asks, o)
			sort.Slice(b.asks, func(i, j int) bool {
				return b.asks[i].Price < b.asks[j].Price
			})
			if o.Qty == 0 {
				return "filled"
			}
			return "accepted"
		}
		return "filled"
	}
}

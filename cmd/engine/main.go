package main

import (
	"fmt"
	"github.com/yourorg/exvo-trading-engine/internal/book"
)

func main() {
	ob := book.NewBook()

	fmt.Println("Starting exvo-trading-engine (local simulation)\nSubmitting sample orders...")

	ack := ob.Submit(book.Order{ID: "o1", Side: "SELL", Symbol: "AAPL", Price: 150.0, Qty: 100})
	fmt.Println("Submit o1:", ack)

	ack = ob.Submit(book.Order{ID: "o2", Side: "BUY", Symbol: "AAPL", Price: 151.0, Qty: 50})
	fmt.Println("Submit o2:", ack)

	ack = ob.Submit(book.Order{ID: "o3", Side: "BUY", Symbol: "AAPL", Price: 150.0, Qty: 60})
	fmt.Println("Submit o3:", ack)

	fmt.Println("Done.")
}

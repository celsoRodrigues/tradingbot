package main

import (
	"context"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2"
)

var (
	binanceKey    string
	binanceSecret string
)

func init() {
	//binance.UseTestnet = true
	binanceKey = os.Getenv("BINANCE_KEY")
	binanceSecret = os.Getenv("BINANCE_SECRET")
}

func main() {
	fmt.Println("boom")

	// Generate default client
	client := binance.NewClient(binanceKey, binanceSecret)

	// order, err := client.NewCreateOrderService().Symbol("ADAGBP").
	// 	Side(binance.SideTypeBuy).Type(binance.OrderTypeMarket).QuoteOrderQty("25").Do(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(order)

	orders, err := client.NewListOrdersService().Symbol("ADAGBP").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}

	infoService := client.NewExchangeInfoService()

	res, err := infoService.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for _, x := range res.Symbols {
		if x.Symbol == "ADAGBP" {
			fmt.Println("MAX PRICE", x.Filters[0]["maxPrice"], "MIN PRICE", x.Filters[0]["minPrice"])
		}

	}

}

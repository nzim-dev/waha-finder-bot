package main

import (
	"context"
	"net/url"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	item "github.com/nzim-dev/waha-frinder/bot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func buildAddress(requst string) func() string {
	address := viper.GetString("searchRoute")

	return func() string {
		u, _ := url.Parse(address)
		q := u.Query()
		q.Set("query", requst)
		u.RawQuery = q.Encode()

		address = u.String()
		return address
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	//init config
	if err := initConfig(); err != nil {
		logrus.Fatalf("err init config: %s", err.Error())
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	//SEARCH COMES FROM BOT!!!
	search := "astartes"

	address := buildAddress(search)
	//navigate
	if err := chromedp.Run(ctx, chromedp.Navigate(address())); err != nil {
		logrus.Fatalln(err)
	}

	//get titles and links
	var header []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Nodes(`.unified-search__result__header a`, &header, chromedp.ByQueryAll)); err != nil {
		logrus.Fatalln(err)
	}

	res := make([]item.Item, 0)
	for _, head := range header {
		res = append(res, item.Item{
			Title: head.AttributeValue("data-title"),
			Link:  head.AttributeValue("href"),
		})
	}

	logrus.Println(res)
}

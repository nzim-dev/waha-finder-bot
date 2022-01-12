package scrapper

import (
	"context"
	"net/url"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	info "github.com/nzim-dev/waha-finder-bot"
	item "github.com/nzim-dev/waha-finder-bot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

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

func processOutput(infoNodes []*cdp.Node) []item.Item {
	res := make([]item.Item, 0)

	for _, head := range infoNodes {
		res = append(res, item.Item{
			Title: head.AttributeValue("data-title"),
			Link:  head.AttributeValue("href"),
		})
	}

	return res
}

func ScrapData(credentials *info.Request) []item.Item {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	address := buildAddress(credentials.SearchRequest)

	//navigate
	if err := chromedp.Run(ctx, chromedp.Navigate(address())); err != nil {
		logrus.Fatalln(err)
	}

	//get titles and links
	var infoNodes []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Nodes(`.unified-search__result__header a`, &infoNodes, chromedp.ByQueryAll)); err != nil {
		logrus.Fatalln(err)
	}

	output := processOutput(infoNodes)
	return output[:credentials.LengthOfOutput]
}

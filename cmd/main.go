package main

import (
	scrapper "github.com/nzim-dev/waha-frinder/bot/pkg/scrapper"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := scrapper.InitConfig(); err != nil {
		logrus.Fatalf("err init config: %s", err.Error())
	}
}

func main() {
	res := scrapper.ScrapData()
	logrus.Println(res)
}

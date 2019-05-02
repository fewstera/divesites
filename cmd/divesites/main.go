package main

import (
	"fmt"

	"github.com/fewstera/divesites/pkg/commandhandler"
	"github.com/fewstera/divesites/pkg/delivery/web"
	"github.com/fewstera/divesites/pkg/memoryeventstore"
	"github.com/fewstera/divesites/pkg/site"
)

func main() {
	es := memoryeventstore.New()
	siteRepo := site.NewRepository(es)
	ch := commandhandler.New(es, siteRepo)

	err := web.StartServer(":8080", ch, siteRepo)
	if err != nil {
		panic(fmt.Errorf("starting web server: %s", err))
	}

}

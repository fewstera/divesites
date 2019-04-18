package main

import (
	"fmt"

	"github.com/fewstera/divesites/pkg/commandhandler"
	"github.com/fewstera/divesites/pkg/memoryeventstore"
	"github.com/fewstera/divesites/pkg/site"
	"github.com/google/uuid"
)

func main() {
	es := memoryeventstore.New()
	ch := commandhandler.New(es)
	siterepo := site.NewRepository(es)

	siteId := uuid.New().String()
	err := ch.Handle(&site.CreateCommand{siteId, "Baygitano", "West Bay", 20})
	if err != nil {
		panic(fmt.Errorf("handling create command: %s", err))
	}

	s, err := siterepo.Get(siteId)
	if err != nil {
		panic(fmt.Errorf("fetching site: %s", err))
	}

	fmt.Printf("%#v\n", s)
}

package main

import (
	"github.com/fewstera/divesites/pkg/commandhandler"
	"github.com/fewstera/divesites/pkg/site"
	"github.com/google/uuid"
)

func main() {
	ch := &commandhandler.CommandHandler{}

	siteId := uuid.New().String()
	ch.Handle(&site.CreateCommand{siteId, "Baygitano", "West Bay", 20})
	//ch.Handle(&site.AddReportCommand{"Baygitano", "West Bay", 20})
	//ch.Handle(&site.CreateCommand{"Baygitano", "West Bay", 20})
}

package commandhandler

import (
	"github.com/fewstera/divesites"
	"github.com/fewstera/divesites/pkg/eventstore"
	"github.com/fewstera/divesites/pkg/site"
)

type CommandHandler struct {
	es             eventstore.EventStore
	siteRepository *site.Repository
}

func New(es eventstore.EventStore, siteRepository *site.Repository) *CommandHandler {
	return &CommandHandler{
		es:             es,
		siteRepository: siteRepository,
	}
}

func (ch *CommandHandler) Handle(command divesites.Command) error {
	switch c := command.(type) {
	case *site.CreateCommand:
		ch.handleSiteCreate(c)
	case *site.AddReportCommand:
		ch.handleSiteAddReport(c)
	}

	return nil
}

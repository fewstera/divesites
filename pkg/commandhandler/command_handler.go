package commandhandler

import (
	"fmt"

	"github.com/fewstera/divesites"
	"github.com/fewstera/divesites/pkg/eventstore"
	"github.com/fewstera/divesites/pkg/site"
)

type CommandHandler struct {
	es eventstore.EventStore
}

func New(es eventstore.EventStore) *CommandHandler {
	return &CommandHandler{
		es: es,
	}
}

func (ch *CommandHandler) Handle(command divesites.Command) error {
	switch c := command.(type) {
	case *site.CreateCommand:
		s := site.New(c.ID, c.Name, c.Location, c.Depth)
		if err := ch.es.Store(s.UncommitedChanges()); err != nil {
			return fmt.Errorf("commiting changes to event store: %s", err)
		}
		s.ClearChanges()
	}

	return nil
}

package commandhandler

import (
	"fmt"

	"github.com/fewstera/divesites"
	"github.com/fewstera/divesites/pkg/site"
	"github.com/mergermarket/intel-cms-service/pkg/eventstore"
)

type CommandHandler struct {
	es eventstore.EventStore
}

func New(es eventstore.EventStore) {
	return &CommandHandler{
		eventStore: es,
	}
}

func (*CommandHandler) Handle(command divesites.Command) {
	switch c := command.(type) {
	case *site.CreateCommand:
		s := site.New(c.ID, c.Name, c.Location, c.Depth)
		if err := es.Store(s.UncommitedChanges()); err != nil {
			return fmt.Errorf("commiting changes to event store: %s", err)
		}

		return nil
	}
}

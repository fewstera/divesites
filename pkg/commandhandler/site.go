package commandhandler

import (
	"fmt"

	"github.com/fewstera/divesites/pkg/site"
)

func (ch *CommandHandler) handleSiteCreate(c *site.CreateCommand) error {
	s := site.New(c.ID, c.Name, c.Location, c.Depth)
	if err := ch.es.Store(s.UncommitedChanges()); err != nil {
		return fmt.Errorf("commiting changes to event store: %s", err)
	}
	s.ClearChanges()

	return nil
}

func (ch *CommandHandler) handleSiteAddReport(c *site.AddReportCommand) error {
	s, err := ch.siteRepository.Get(c.SiteID)
	if err != nil {
		return fmt.Errorf("retrieving site: %s", err)
	}

	report := site.Report{
		Reporter:   c.Reporter,
		Visibility: c.Visibility,
		Rating:     c.Rating,
		Notes:      c.Notes,
	}

	s.AddReport(report)
	if err := ch.es.Store(s.UncommitedChanges()); err != nil {
		return fmt.Errorf("commiting changes to event store: %s", err)
	}
	s.ClearChanges()

	return nil
}

package site

type CreateCommand struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Depth    float32 `json:"depth"`
}

func (*CreateCommand) CommandType() string { return "SITE_CREATE" }

type AddReportCommand struct {
	SiteID     string
	Reporter   string
	Visibility *float32
	Rating     int
	Notes      *string
}

func (*AddReportCommand) CommandType() string { return "SITE_ADD_REPORT" }

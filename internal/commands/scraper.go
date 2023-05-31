package commands

type Scraper interface {
	Scrape() error
}

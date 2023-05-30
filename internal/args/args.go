package args

type Args struct {
	Help   bool
	File   string
	Scrape []Scrape
}

type Scrape struct {
	URL       string `json:"url"`
	Extension string `json:"extension"`
	Folder    string `json:"folder"`
}

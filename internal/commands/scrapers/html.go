package scrapers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/Pedrommb91/go-img-scrapper/internal/args"
	"github.com/Pedrommb91/go-img-scrapper/internal/commands"
	"github.com/Pedrommb91/go-img-scrapper/pkg/downloader"
	"github.com/gocolly/colly/v2"
)

type htmlScraper struct {
	scrape args.Scrape
}

func NewHTMLScraper(scrape args.Scrape) commands.Scraper {
	return htmlScraper{scrape: scrape}
}

func (s htmlScraper) Scrape() error {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)
	c.SetRequestTimeout(120 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		fmt.Println("Visiting:", r.URL)
	})

	c.OnHTML("img", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		fmt.Println(link)

		err := os.MkdirAll(s.scrape.Folder, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = downloader.DownloadFile(link, s.scrape.Folder+"/"+filepath.Base(link))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Dowload succeed")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
		fmt.Println(r.StatusCode)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit(s.scrape.URL)
	c.Wait()
	return nil
}

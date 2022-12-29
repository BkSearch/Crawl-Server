package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/BkSearch/Crawl-Server/common"
	"github.com/gocolly/colly"
)

var (
	port                                                           int
	host, userRead, userWrite, passwordRead, passwordWrite, dbName string
)

func loadConfig() {
	//load host
	host = os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	// load port
	port, _ = strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 5432
	}
	// read user
	userRead = os.Getenv("UserRead")
	if userRead == "" {
		log.Fatal("Invalid read user")
	}

	passwordRead = os.Getenv("PasswordRead")

	// write user
	userWrite = os.Getenv("UserWrite")
	if userWrite == "" {
		log.Fatal("Invalid write user")
	}
	passwordWrite = os.Getenv("PasswordWrite")
	// load db
	dbName = os.Getenv("DBName")
}

type Topic struct {
	questions common.Question
	answer    []common.Answer
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("stackoverflow.com"),
	)

	recursive_c := colly.NewCollector(
		colly.AllowedDomains("stackoverflow.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// var topic []Topic

	c.OnHTML("#questions", func(e *colly.HTMLElement) {
		e.ForEach(".s-post-summary.js-post-summary", func(i int, el *colly.HTMLElement) {
			fmt.Println(i)
			fmt.Println(el.Attr("data-post-id"))
			fmt.Println(el.ChildText(".s-post-summary--content-title"))
			link := el.ChildAttr("a[href]", "href")
			fmt.Println(link)
			fmt.Println(el.ChildText("div.s-post-summary--stats-item.s-post-summary--stats-item__emphasized > span.s-post-summary--stats-item-number"))
			recursive_c.Visit("https://stackoverflow.com" + link)
		})
		// var questions common.Question
		// questions.ID =
		// fmt.Println(e.ChildText(".s-post-summary--content-title"))
		//   fmt.Println(e.ChildAttr("a[href]", "href"))
		//   fmt.Println(e.Attr(""))
	})


	recursive_c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	recursive_c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	recursive_c.OnHTML("#answers", func(e *colly.HTMLElement) {
    e.ForEach("div.answer.js-answer", func (_ int, el *colly.HTMLElement) {
      fmt.Printf("Answer Id %v: \n",el.Attr("data-answerid"))
      fmt.Printf("Vote: %v\n",el.ChildText("div.votecell.post-layout--left > div > div.js-vote-count.flex--item.d-flex.fd-column.ai-center.fc-black-500.fs-title"))
      fmt.Println(el.ChildText("div > div.answercell.post-layout--right > div.s-prose.js-post-body"))
    })
	})

	recursive_c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

  for i := 1; i <= 10; i++ {
    var link ="https://stackoverflow.com/questions?tab=votes&page=" + strconv.Itoa(i) 
	  c.Visit(link)
  }
}

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type playerInfo struct {
	Name string `json:"name"`
}

func writeFile(name string, data string) {
	/*
		Writes data onto file
	*/
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func getPlayerLinks() []string {

	c := colly.NewCollector()
	playerLinks := []string{}

	c.OnHTML(".row playerList", func(e *colly.HTMLElement) {
		fmt.Println("found section")
		e.ForEach(".row playerList", func(_ int, el *colly.HTMLElement) {
			fmt.Println("found a a tag")
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.nba.com/players")

	return playerLinks
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	// c := colly.NewCollector()

	// // On every a element which has href attribute call callback
	// c.OnHTML("#block-league-content > section.row.nba-player-index__row > section:nth-child(1) > a:nth-child(2) > p", func(e *colly.HTMLElement) {
	// 	testPlayer := playerInfo{Name: e.Text}
	// 	testJSON, _ := json.Marshal(testPlayer)

	// 	// Print player name
	// 	// fmt.Printf("Player name: %q\n", e.Text)
	// 	// fmt.Printf("%s \n", testPlayer.Name)
	// 	// fmt.Printf("%s \n", string(testJSON))
	// 	writeFile("output.json", string(testJSON))
	// })

	// // Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// c.OnError(func(_ *colly.Response, err error) {
	// 	log.Println("Something went wrong:", err)
	// })

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited", r.Request.URL)
	// })

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("Finished", r.Request.URL)
	// })

	// // Start scraping on https://hackerspaces.org
	// c.Visit("https://www.nba.com/players")
	getPlayerLinks()
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type playerInfo struct {
	Name   string `json:"name"`
	League string `json:"league"`
	Played string `json:"played"`
	Starts string `json:"starts"`
	Min    string `json:"min"`
	FGM    string `json:"fgm"`
	FGA    string `json:"fga"`
	FGP    string `json:"fgp"`
	THPM   string `json:"thpm"`
	THPA   string `json:"thpa"`
	THPP   string `json:"thpp"`
	TWPM   string `json:"twpm"`
	TWPA   string `json:"twpa"`
	TWPP   string `json:"twpp"`
	EFGP   string `json:"efgp"`
	FTM    string `json:"ftm"`
	FTA    string `json:"fta"`
	FTP    string `json:"ftp"`
	ORB    string `json:"orb"`
	DRB    string `json:"drb"`
	AST    string `json:"ast"`
	STL    string `json:"stl"`
	BLK    string `json:"blk"`
	TOV    string `json:"tov"`
	Fouls  string `json:"fouls"`
	PTS    string `json:"pts"`
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
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	links := []string{}

	for _, letter := range alphabet {
		// Instantiate default collector
		c := colly.NewCollector()

		// On every a element which has href attribute call callback
		c.OnHTML("#players", func(e *colly.HTMLElement) {
			e.ForEach("#players tr a", func(_ int, el *colly.HTMLElement) {
				if strings.Contains(el.Attr("href"), "players") {
					links = append(links, el.Attr("href"))
				}
			})
		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())

		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
		})

		c.OnScraped(func(r *colly.Response) {
			fmt.Println("Finished", r.Request.URL)
		})

		c.Visit("https://www.basketball-reference.com/players/" + letter)

	}

	return links
}

func getPlayerInfo() []playerInfo {
	playerLinks := getPlayerLinks()
	players := []playerInfo{}

	for _, link := range playerLinks {

		player := playerInfo{}

		// Instantiate default collector
		c := colly.NewCollector()

		// On every a element which has href attribute call callback
		c.OnHTML("h1 > span", func(e *colly.HTMLElement) {
			player.Name = e.Text
		})

		c.OnHTML("#per_game tfoot", func(e *colly.HTMLElement) {
			e.ForEach("#per_game tfoot tr td", func(_ int, el *colly.HTMLElement) {

				switch el.Attr("data-stat") {
				case "lg_id":
					if player.League == "" {
						player.League = el.Text
					}
				case "g":
					if player.Played == "" {
						player.Played = el.Text
					}
				case "gs":
					if player.Starts == "" {
						player.Starts = el.Text
					}
				case "mp_per_g":
					if player.Min == "" {
						player.Min = el.Text
					}
				case "fg_per_g":
					if player.FGM == "" {
						player.FGM = el.Text
					}
				case "fga_per_g":
					if player.FGA == "" {
						player.FGA = el.Text
					}
				case "fg_pct":
					if player.FGP == "" {
						player.FGP = el.Text
					}
				case "fg3_per_g":
					if player.THPM == "" {
						player.THPM = el.Text
					}
				case "fg3a_per_g":
					if player.THPA == "" {
						player.THPA = el.Text
					}
				case "fg3_pct":
					if player.THPP == "" {
						player.THPP = el.Text
					}
				case "fg2_per_g":
					if player.TWPM == "" {
						player.TWPM = el.Text
					}
				case "fg2a_per_g":
					if player.TWPA == "" {
						player.TWPA = el.Text
					}
				case "fg2_pct":
					if player.TWPP == "" {
						player.TWPP = el.Text
					}
				case "efg_pct":
					if player.EFGP == "" {
						player.EFGP = el.Text
					}
				case "ft_per_g":
					if player.FTM == "" {
						player.FTM = el.Text
					}
				case "fta_per_g":
					if player.FTA == "" {
						player.FTA = el.Text
					}
				case "ft_pct":
					if player.FTP == "" {
						player.FTP = el.Text
					}
				case "orb_per_g":
					if player.ORB == "" {
						player.ORB = el.Text
					}
				case "drb_per_g":
					if player.DRB == "" {
						player.DRB = el.Text
					}
				case "ast_per_g":
					if player.AST == "" {
						player.AST = el.Text
					}
				case "stl_per_g":
					if player.STL == "" {
						player.STL = el.Text
					}
				case "blk_per_g":
					if player.BLK == "" {
						player.BLK = el.Text
					}
				case "tov_per_g":
					if player.TOV == "" {
						player.TOV = el.Text
					}
				case "pf_per_g":
					if player.Fouls == "" {
						player.Fouls = el.Text
					}
				case "pts_per_g":
					if player.PTS == "" {
						player.PTS = el.Text
					}
				}
			})
		})

		// Before making a request print "Visiting ..."
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())

		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
		})

		c.OnScraped(func(r *colly.Response) {
			players = append(players, player)
			fmt.Println("Finished", r.Request.URL)
		})

		c.Visit("https://www.basketball-reference.com" + link)

	}
	return players

}

func main() {
	players := getPlayerInfo()
	JSON, _ := json.Marshal(players)
	writeFile("output.json", string(JSON))
}

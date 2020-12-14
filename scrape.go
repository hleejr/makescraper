package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	"github.com/gocolly/colly"
)

type playerInfo struct {
	Name     string `json:"name"`
	Position string `json:"pos"`
	Image    string `json:"img"`
	League   string `json:"league"`
	Played   string `json:"played"`
	Starts   string `json:"starts"`
	Min      string `json:"min"`
	FGM      string `json:"fgm"`
	FGA      string `json:"fga"`
	FGP      string `json:"fgp"`
	THPM     string `json:"thpm"`
	THPA     string `json:"thpa"`
	THPP     string `json:"thpp"`
	TWPM     string `json:"twpm"`
	TWPA     string `json:"twpa"`
	TWPP     string `json:"twpp"`
	EFGP     string `json:"efgp"`
	FTM      string `json:"ftm"`
	FTA      string `json:"fta"`
	FTP      string `json:"ftp"`
	ORB      string `json:"orb"`
	DRB      string `json:"drb"`
	TRB      string `json:"trb"`
	AST      string `json:"ast"`
	STL      string `json:"stl"`
	BLK      string `json:"blk"`
	TOV      string `json:"tov"`
	Fouls    string `json:"fouls"`
	PTS      string `json:"pts"`
	Awards   string `json:"awards"`
}

type nflPlayerInfo struct {
	Name     string `json:"name"`
	Position string `json:"pos"`
	Image    string `json:"img"`
	Comp     string `json:"completions"`
	Attempts string `json:"attempts"`
	PYRD     string `json:"passing-yards"`
	PTD      string `json:"passing-tds"`
	INTT     string `json:"int-thrown"`
	Rate     string `json:"passer-rating"`
	Sacked   string `json:"sacked"`
	LYRD     string `json:"lost-yards"`
	Targets  string `json:"targets"`
	REC      string `json:"receptions"`
	Yards    string `json:"yards"`
	TD       string `json:"touchdowns"`
	Rush     string `json:"rushes"`
	RYRD     string `json:"rush-yards"`
	RTD      string `json:"rush-tds"`
	Returns  string `json:"returns"`
	KTD      string `json:"kick-tds"`
	Tackles  string `json:"tackles"`
	Ast      string `json:"ast-tackles"`
	Sacks    string `json:"sacks"`
	QBH      string `json:"qb-hits"`
	INT      string `json:"interceptions"`
	IYRD     string `json:"int-yards"`
	ITD      string `json:"int-tds"`
	PD       string `json:"pass-deflections"`
	Fumbles  string `json:"fumbles"`
	FL       string `json:"fumbles-lost"`
	FF       string `json:"fumbles-forced"`
	FR       string `json:"fumbles-recovered"`
	FYRD     string `json:"fumble-yards"`
	FTD      string `json:"fumble-tds"`
}

func testrun() {
	player := nflPlayerInfo{}

	for i := 0; i < len(structs.Map(nflPlayerInfo{})); i++ {

		val := reflect.Indirect(reflect.ValueOf(player))
		fmt.Println(val.Type().Field(i).Name)

	}
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

func nflPlayerLinks() []string {
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	links := []string{}

	for _, letter := range alphabet {
		// Instantiate default collector
		c := colly.NewCollector()

		// On every a element which has href attribute call callback
		c.OnHTML("#div_players", func(e *colly.HTMLElement) {
			e.ForEach("p a", func(_ int, el *colly.HTMLElement) {
				if strings.Contains(el.Attr("href"), "players") {
					link := el.Attr("href")

					if strings.HasSuffix(link, ".htm") {
						link = link[:len(link)-len(".htm")]
					}

					links = append(links, link)
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

		c.Visit("https://www.pro-football-reference.com/players/" + letter)

	}
	// fmt.Println("Finished", links)
	return links

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
	fmt.Println("Finished", links)
	return links
}

func nflPlayers() []nflPlayerInfo {
	playerLinks := nflPlayerLinks()
	players := []nflPlayerInfo{}

	for _, link := range playerLinks {

		player := nflPlayerInfo{}

		// Instantiate default collector
		c := colly.NewCollector()

		c.OnHTML("h1 > span", func(e *colly.HTMLElement) {
			player.Name = e.Text
		})
		c.OnHTML("#meta > div.media-item > img", func(e *colly.HTMLElement) {
			player.Image = e.Attr("src")
		})
		c.OnHTML("#meta > div:nth-child(2) > p:nth-child(3)", func(e *colly.HTMLElement) {
			player.Position = e.Text
		})
		c.OnHTML("#stats tfoot tr", func(e *colly.HTMLElement) {
			e.ForEach("#stats tfoot tr td", func(_ int, el *colly.HTMLElement) {

				switch el.Attr("data-stat") {
				case "pass_cmp":
					if player.Comp == "" {
						player.Comp = el.Text
					}
				case "pass_att":
					if player.Attempts == "" {
						player.Attempts = el.Text
					}
				case "pass_yds":
					if player.PYRD == "" {
						player.PYRD = el.Text
					}
				case "pass_td":
					if player.PTD == "" {
						player.PTD = el.Text
					}
				case "pass_int":
					if player.INTT == "" {
						player.INTT = el.Text
					}
				case "pass_rating":
					if player.Rate == "" {
						player.Rate = el.Text
					}
				case "pass_sacked":
					if player.Sacked == "" {
						player.Sacked = el.Text
					}
				case "pass_sacked_yds":
					if player.LYRD == "" {
						player.LYRD = el.Text
					}
				case "targets":
					if player.Targets == "" {
						player.Targets = el.Text
					}
				case "rec":
					if player.REC == "" {
						player.REC = el.Text
					}
				case "rec_yds":
					if player.Yards == "" {
						player.Yards = el.Text
					}
				case "rec_td":
					if player.TD == "" {
						player.TD = el.Text
					}
				case "rush_att":
					if player.Rush == "" {
						player.Rush = el.Text
					}
				case "rush_yds":
					if player.RYRD == "" {
						player.RYRD = el.Text
					}
				case "rush_td":
					if player.RTD == "" {
						player.RTD = el.Text
					}
				case "kick_ret":
					if player.Returns == "" {
						player.Returns = el.Text
					}
				case "kick_ret_td":
					if player.KTD == "" {
						player.Returns = el.Text
					}
				case "tackles_solo":
					if player.Tackles == "" {
						player.Tackles = el.Text
					}
				case "tackles_assist":
					if player.Ast == "" {
						player.Ast = el.Text
					}
				case "sacks":
					if player.Sacks == "" {
						player.Sacks = el.Text
					}
				case "qb_hits":
					if player.QBH == "" {
						player.QBH = el.Text
					}
				case "def_int":
					if player.INT == "" {
						player.INT = el.Text
					}
				case "def_int_yds":
					if player.IYRD == "" {
						player.IYRD = el.Text
					}
				case "def_int_td":
					if player.ITD == "" {
						player.ITD = el.Text
					}
				case "pass_defended":
					if player.PD == "" {
						player.PD = el.Text
					}
				case "fumbles":
					if player.Fumbles == "" {
						player.Fumbles = el.Text
					}
				case "fumbles_lost":
					if player.FL == "" {
						player.FL = el.Text
					}
				case "fumbles_forced":
					if player.FF == "" {
						player.FF = el.Text
					}
				case "fumbles_rec":
					if player.FR == "" {
						player.FR = el.Text
					}
				case "fumbles_rec_yds":
					if player.FYRD == "" {
						player.FYRD = el.Text
					}
				case "fumbles_rec_td":
					if player.FTD == "" {
						player.FTD = el.Text
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

		c.Visit("https://www.pro-football-reference.com/" + link + "/gamelog")

	}

	fmt.Println("Finished", players)
	return players
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

		c.OnHTML("#meta > div.media-item > img", func(e *colly.HTMLElement) {
			player.Image = e.Attr("src")
		})

		c.OnHTML("#per_game tbody", func(e *colly.HTMLElement) {
			e.ForEach("#per_game tbody tr td", func(_ int, el *colly.HTMLElement) {

				switch el.Attr("data-stat") {
				case "pos":
					if player.Position == "" {
						player.Position = el.Text
					}
				}
			})
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
				case "trb_per_g":
					if player.TRB == "" {
						player.TRB = el.Text
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

		c.OnHTML("#bling li", func(e *colly.HTMLElement) {
			if player.Awards == "" {
				player.Awards = e.Text
			} else {
				player.Awards = player.Awards + "\n" + e.Text
			}
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

func getCurrentPlayers() []string {
	actives := []string{}

	c := colly.NewCollector()

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("table tbody tr td a", func(_ int, el *colly.HTMLElement) {
			if strings.Contains(el.Attr("href"), "player") {
				actives = append(actives, el.Text)
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

	c.Visit("https://basketball.realgm.com/nba/players")

	return actives
}

func main() {
	// players := getPlayerInfo()
	players := nflPlayers()
	// actives := getCurrentPlayers()
	JSON, _ := json.Marshal(players)
	// writeFile("output.json", string(JSON))
	// writeFile("output2.json", string(JSON))
	writeFile("output3.json", string(JSON))
	// testrun()
	
}

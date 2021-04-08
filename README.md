# 🕷 makescraper

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started
3. [Resources](#resources)

## Project Structure

```bash
📂 makescraper
├── README.md
└── scrape.go
```

## Getting Started

Run each command line-by-line in your terminal to set up the utility:
```
$ git clone https://github.com/hleejr/makescraper.git
$ cd makesite
$ git remote rm origin
```
Then add the link to your repository as the new origin

#### Usage

- Golang's `struct` is used to store the scraped data
- On certain Colly functions like `c.OnHTML` you must refactor them to use the CSS selector(s) from the website you wish to scrape.
- Verify the data you want by printing it to `stdout`.
- Serialize the `struct` you created to JSON. Print the JSON to `stdout` to validate it.
- Write scraped data to a file named `output.json`.

## Resources

#### Scraping

- [**Colly** - Docs](http://go-colly.org/docs/): Check out the sidebar for 20+ examples!

#### Serializing & Saving

- [JSON to Struct](https://mholt.github.io/json-to-go/): Paste any JSON data and convert it into a Go structure that will support storing that data.
- [GoByExample - JSON](https://gobyexample.com/json): Covers Go's built-in support for JSON encoding and decoding to and from built-in and custom data types (structs).
- [GoByExample - Writing Files](https://gobyexample.com/writing-files): Covers creating new files and writing to them.

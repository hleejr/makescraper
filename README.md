[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/hleejr/makescraper">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Makescraper</h3>

  <p align="center">
    Golang webscraping tool using CSS identifiers to pull information, store it as structs, and convert in to json
    <br />
    <a href="https://github.com/hleejr/makescraper"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/hleejr/makescraper">View Demo</a>
    ·
    <a href="https://github.com/hleejr/makescraper/issues">Report Bug</a>
    ·
    <a href="https://github.com/hleejr/makescraper/issues">Request Feature</a>
  </p>
</p>

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Resources](#resources)

## Project Structure

```bash
📂 makescraper
├── 📂 images
├── README.md
└── scrape.go
```

## Getting Started

Run each command line-by-line in your terminal to set up the utility:
```
$ git clone https://github.com/hleejr/makescraper.git
$ cd makescraper
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

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/hleejr/makescraper.svg?style=for-the-badge
[contributors-url]: https://github.com/hleejr/makescraper/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/hleejr/makescraper.svg?style=for-the-badge
[forks-url]: https://github.com/hleejr/makescraper/network/members
[stars-shield]: https://img.shields.io/github/stars/hleejr/makescraper.svg?style=for-the-badge
[stars-url]: https://github.com/hleejr/makescraper/stargazers
[issues-shield]: https://img.shields.io/github/issues/hleejr/makescraper.svg?style=for-the-badge
[issues-url]: https://github.com/hleejr/makescraper/issues
[license-shield]: https://img.shields.io/github/license/hleejr/makescraper.svg?style=for-the-badge
[license-url]: https://github.com/hleejr/makescraper/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/henry-bowe-jr-31498916a/
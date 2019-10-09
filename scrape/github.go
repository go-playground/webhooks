// Scrape the github webhook documentation for example payloads
// The data might need some manual massaging because some example values
// are set to null.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type ExamplePayload struct {
	name string
	json string
}

func getName(selection *goquery.Selection) *goquery.Selection {
	if match, _ := regexp.MatchString("^(\\w+_)*\\w+$", selection.Text()); match {
		// if match, _ := regexp.MatchString("^\\S$", selection.Text()); match {
		return selection
	}
	return getName(selection.Next())
}

func getExample(selection *goquery.Selection) *goquery.Selection {
	if selection.Is("pre") {
		return selection
	}
	return getExample(selection.Next())
}
func collect() []ExamplePayload {
	c := colly.NewCollector()
	collected := make([]ExamplePayload, 0)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("h3", func(e *colly.HTMLElement) {

		if strings.Contains(e.Text, "Webhook event name") {
			name := getName(e.DOM.Next())
			exampleHeader := name.Next()
			if !strings.Contains(exampleHeader.Text(), "Webhook payload example") {
				fmt.Printf(exampleHeader.Text())
				return
			}
			example := getExample(exampleHeader.Next())
			collected = append(collected, ExamplePayload{
				name: name.Text(),
				json: example.Text(),
			})
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit("https://developer.github.com/v3/activity/events/types/")

	return collected
}

// Installation id does not seem to be present in most examples, but they do appear for many
// Payloads, at least if they are for a Github app.
var installationSnippet = `,
  "installation": {
    "id": 2311213,
    "node_id": "MDIzOkludGVncmF0aW9uSW5zdGFsbGF0aW9uMjMxMTIxMw=="
  }
}`
var dontAddInstallationFor = map[string]bool{
	"installation":              true,
	"installation_repositories": true,
	"content_reference":         true,
	"repository_dispatch":       true,
}

func addInstallationId(example *ExamplePayload) {
	if _, ok := dontAddInstallationFor[example.name]; ok {
		return
	}
	r, _ := regexp.Compile("\n}[ \n\r]+$")
	// res := r.FindString(example.json)
	// g	fmt.Printf("found %s", res)
	example.json = r.ReplaceAllStringFunc(example.json, func(str string) string {
		return strings.Replace(str, "\n}", installationSnippet, 1)
	})
}
func write(folder string, example ExamplePayload) error {
	filename := fmt.Sprintf("%s.json", strings.ReplaceAll(example.name, "_", "-"))
	fullPath := path.Join(folder, filename)

	fmt.Printf("Writing %s\n", fullPath)
	err := ioutil.WriteFile(fullPath, []byte(example.json), 0777)
	if err != nil {
		fmt.Printf("Error writing file %s: %s\n", fullPath, err)
	}

	return err
}
func main() {
	collected := collect()
	fmt.Printf("Found %s examle payloads. \n", len(collected))
	root, _ := os.Getwd()
	folder := path.Join(root, "testdata/github")
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0777)
	}
	for _, ex := range collected {
		addInstallationId(&ex)
		_ = write(folder, ex)
	}
}

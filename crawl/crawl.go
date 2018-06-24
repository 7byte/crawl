package crawl

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type visitNode func(n *html.Node) (links []string, urls []string)

func extract(domain, url string, visits []visitNode) (links []string, urls []string, err error) {
	if !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
		if domain == "" {
			return nil, nil, fmt.Errorf("domain is empty")
		}
		url = domain + "/" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	links, urls = forEachNode(doc, visits)
	return
}

func forEachNode(n *html.Node, visits []visitNode) (links []string, urls []string) {
	for _, v := range visits {
		l, u := v(n)
		links = append(links, l...)
		urls = append(urls, u...)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		l, u := forEachNode(c, visits)
		links = append(links, l...)
		urls = append(urls, u...)
	}
	return
}

// Crawl get links
func Crawl(domain, url string) []string {
	list, _, err := extract(domain, url, []visitNode{inputImg()})
	if err != nil {
		log.Print(err)
	}
	return list
}

package crawl

import (
	"log"

	"github.com/7byte/crawl/downloader"
	"golang.org/x/net/html"
)

func hrefImg() (v visitNode) {
	return func(n *html.Node) (links []string, urls []string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				links = append(links, a.Val)
			}
			for s := n.FirstChild; s != nil; s = s.NextSibling {
				if s.Type == html.ElementNode && s.Data == "img" {
					for _, a := range s.Attr {
						if a.Key != "src" {
							continue
						}
						log.Print(a.Val)
						urls = append(links, a.Val)
						downloader.Download("", []string{a.Val}, "/data/image")
					}
				}
			}
		}
		return
	}
}

func inputImg() (v visitNode) {
	return func(n *html.Node) (links []string, urls []string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				links = append(links, a.Val)
			}
		}
		if n.Type == html.ElementNode && n.Data == "input" {
			var dataType, src string
			for _, a := range n.Attr {
				if a.Key == "data-src" {
					src = a.Val
				}
				if a.Key == "type" {
					dataType = a.Val
				}
			}
			if dataType == "image" {
				log.Print(src)
				urls = append(links, src)
				downloader.Download("", []string{src}, "/data/image")
			}
		}
		return
	}
}

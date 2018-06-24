package downloader

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// Download 下载
func Download(domain string, urls []string, root string) (err error) {
	for _, addr := range urls {
		u, err := url.Parse(addr)
		if err != nil {
			continue
		}
		str := strings.Split(u.Path, "/")
		filename := str[len(str)-1]

		resp, err := http.Get(addr)
		if err != nil {
			continue
		}
		if resp.StatusCode != http.StatusOK {
			continue
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		image, err := os.Create(path.Join(root, filename))
		if err != nil {
			log.Print(err)
			resp.Body.Close()
			continue
		}
		image.Write(data)
		image.Close()
	}
	return
}

package readability

import (
	nurl "net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

// indexOf returns the position of the first occurrence of a
// specified  value in a string array. Returns -1 if the
// value to search for never occurs.
func indexOf(array []string, key string) int {
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return i
		}
	}
	return -1
}

// wordCount returns number of word in str.
func wordCount(str string) int {
	return len(strings.Fields(str))
}

// toAbsoluteURI convert uri to absolute path based on base.
// However, if uri is prefixed with hash (#), the uri won't be changed.
func toAbsoluteURI(uri string, base *nurl.URL) string {
	if uri == "" || base == nil {
		return ""
	}

	// If it is hash tag, return as it is
	if uri[:1] == "#" {
		return uri
	}

	// If it is already an absolute URL, return as it is
	tmp, err := nurl.ParseRequestURI(uri)
	if err == nil && tmp.Scheme != "" && tmp.Hostname() != "" {
		return uri
	}

	// Otherwise, resolve against base URI.
	tmp, err = nurl.Parse(uri)
	if err != nil {
		return uri
	}

	return base.ResolveReference(tmp).String()
}

// renderToFile ender an element and save it to file.
// It will panic if it fails to create destination file.
func renderToFile(element *html.Node, filename string) {
	dstFile, err := os.Create(filename)
	if err != nil {
		logrus.Fatalln("failed to create file:", err)
	}
	defer dstFile.Close()
	html.Render(dstFile, element)
}

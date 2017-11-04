// Package parser provides methods to grab all links from markdown files.
package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// DownloadURL returns Body response from the URL.
func DownloadURL(URL string) (string, error) {
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.String(), nil
}

// startsWithSlash checks if string starts with a slash.
func startsWithSlash(line string) bool {
	return strings.HasPrefix(line, "#")
}

// GetAllLinks returns all links and their names from a given markdown file.
func GetAllLinks(markdown string) map[string]string {
	// holds all the links and their corresponding values
	m := make(map[string]string)

	// regex to extract link and text attached to link
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	// scans line by line
	for scanner.Scan() {
		// make regex
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)

		// only apply regex if there are links
		if matches != nil {
			m[matches[0][1]] = matches[0][2]
		}
	}
	return m
}

// ParseMarkdownURL parses an URL and grabs all markdown links from it.
func ParseMarkdownURL(URL string) (map[string]string, error) {
	file, err := DownloadURL(URL)
	if err != nil {
		return make(map[string]string), err
	}
	return GetAllLinks(file), nil
}

// readFile returns contents of the file.
func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

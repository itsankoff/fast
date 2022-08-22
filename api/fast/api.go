package fast

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/itsankoff/fast/api"
)

// BaseDomain defines the speed test domain
const BaseDomain = "fast.com"

// API implementation static check
var _ api.API = (*FastAPI)(nil)

// FastAPI implements the api.API interface using fast.com as speed test
type FastAPI struct {
	domain string
	secure bool
}

// New returns a new instance of FastAPI
func New(secure bool) api.API {
	return &FastAPI{
		domain: BaseDomain,
		secure: secure,
	}
}

// GetDownloadURLs returns URLs for the download of data during speed test
// measurement
func (f *FastAPI) GetDownloadURLs(count int, secure bool) ([]string, error) {
	token, err := f.Token(secure)
	if err != nil {
		return nil, err
	}

	protocol := "http"
	if secure {
		protocol += "s"
	}

	url := fmt.Sprintf("%s://api.fast.com/netflix/speedtest?https=%t&token=%s&urlCount=%d",
		protocol, secure, token, count)

	content, err := api.Get(url)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile("(?U)\"url\":\"(.*)\"")
	reUrls := re.FindAllStringSubmatch(content, -1)

	var outputURLs []string
	for _, arr := range reUrls {
		outputURLs = append(outputURLs, arr[1])
	}

	return outputURLs, nil
}

// Token returns a fast.com token for requesting download URLs
func (f *FastAPI) Token(secure bool) (string, error) {
	protocol := "http"
	if secure {
		protocol += "s"
	}

	baseURL := fmt.Sprintf("%s://%s", protocol, f.domain)
	content, err := api.Get(baseURL)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("app-.*\\.js")
	scriptNames := re.FindAllString(content, 1)
	scriptURL := fmt.Sprintf("%s/%s", baseURL, scriptNames[0])

	scriptContent, err := api.Get(scriptURL)
	if err != nil {
		return "", err
	}

	// Extract the token
	re = regexp.MustCompile("token:\"[[:alpha:]]*\"")
	tokens := re.FindAllString(scriptContent, 1)
	if len(tokens) > 0 {
		return tokens[0][7 : len(tokens[0])-1], nil
	}

	return "", errors.New("token not found")
}

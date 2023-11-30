package vuta

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

type App struct {
	*http.Client
	url      string
	shutdown chan<- error
}

func NewApp(url string) *App {
	return &App{
		Client: http.DefaultClient,
		url:    url,
	}
}

// Download uses single thread
// downloads the file and saves it
// in the file name provided
func (app *App) Download() error {
	fn, err := os.Create(path.Base(app.url))
	if err != nil {
		return err
	}

	resp, err := app.Get(app.url)
	if err != nil {
		return err
	}

	// defer resp.Body.Close() // No need, the server closes this.

	if _, err = io.Copy(fn, resp.Body); err != nil {
		return err
	}

	return nil
}

// Header returns header details of the file
// returns the Content-Length, ETag,acceptRange, error
// error if the header request fails
func Header(url string) (int64, string, bool, error) {
	var acceptRanges bool

	// Make a head request to get file details
	resp, err := http.DefaultClient.Head(url)
	if err != nil {
		return -1, "", acceptRanges, err
	}

	contentLength := resp.ContentLength
	eTag := resp.Header.Get("Etag")

	if resp.Header.Get("Accept-Ranges") == "bytes" {
		acceptRanges = true
	}

	return contentLength, eTag, acceptRanges, nil
}

// DownloadChunk gets chunks of bytes and assembles it.
// to the file.
func (app *App) DownloadChunks(f string, wg *sync.WaitGroup, start, end int) error {
	defer wg.Done()

	// Adds a Range Header to the request
	r := fmt.Sprintf("bytes=%d-%d", start, end)

	req, err := http.NewRequest(http.MethodGet, app.url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", r)

	resp, err := app.Client.Do(req)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(f, data, os.FileMode(os.O_APPEND)|os.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}

	return nil
}

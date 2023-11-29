package vuta

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
func (app *App) Download(fn *os.File) error {
	// fn, err := os.Create(path.Base(app.url))
	// if err != nil {
	// 	return err
	// }

	resp, err := app.Get(app.url)
	if err != nil {
		return err
	}

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

	fmt.Println(resp.Header)

	return contentLength, eTag, acceptRanges, nil
}

// DownloadChunk gets chunks of bytes and assembles it.
// to the file.
func (app *App) DownloadChunk(fn *os.File, offset, size byte) error {
	return nil

}

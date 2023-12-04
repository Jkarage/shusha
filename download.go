package shusha

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

type Download struct {
	*http.Client
	url      string
	shutdown chan<- error
}

func NewDownload(url string) *Download {
	return &Download{
		Client: http.DefaultClient,
		url:    url,
	}
}

// Download uses single thread
// downloads the file and saves it
// in the file name provided
func (d *Download) Download() error {
	fn, err := os.Create(path.Base(d.url))
	if err != nil {
		return err
	}

	resp, err := d.Get(d.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(fn, resp.Body); err != nil {
		return err
	}

	return nil
}

// Header returns header details of the file
// returns the Content-Length, ETag,acceptRange, error
// error if the header request fails
func (d *Download) Header() (int64, string, bool, error) {
	var acceptRanges bool

	// Make a head request to get file details
	resp, err := http.DefaultClient.Head(d.url)
	if err != nil {
		return -1, "", acceptRanges, err
	}

	contentLength := resp.ContentLength
	fileName := path.Base(d.url)

	if resp.Header.Get("Accept-Ranges") == "bytes" {
		acceptRanges = true
	}

	return contentLength, fileName, acceptRanges, nil
}

// DownloadChunk gets chunks of bytes and assembles it.
// to the file.
func (d *Download) DownloadChunks(f string, wg *sync.WaitGroup, start, end int) error {
	defer wg.Done()

	// Adds a Range Header to the request
	r := fmt.Sprintf("bytes=%d-%d", start, end)

	req, err := http.NewRequest(http.MethodGet, d.url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Range", r)

	resp, err := d.Client.Do(req)
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

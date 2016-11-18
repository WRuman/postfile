package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type BodyDescriptor struct {
	body *bytes.Buffer
	contentType string
}

func getRequestBody(filename string, fieldname string) (bd *BodyDescriptor, err error) {
	var fw io.Writer
	var f *os.File
	bd = new(BodyDescriptor)
	bd.body = new(bytes.Buffer)
	w := multipart.NewWriter(bd.body)
	defer w.Close()
	// Try opening the file
	if f, err = os.Open(filename); err != nil {
		return
	}
	defer f.Close()
	// Create the header field and get a destination in the buffer
	if fw, err = w.CreateFormFile(fieldname, filename); err != nil {
		return
	}
	// Copy the file into the dest buffer
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	bd.contentType = w.FormDataContentType()
	return
}

func post(body *bytes.Buffer, url string, contentType string) (status int, err error) {
	client := new(http.Client)
	res, err := client.Post(url, contentType, body)
	if err != nil {
		return
	}
	status = res.StatusCode
	res.Body.Close()
	return
}

func main() {
	var filename string
	var fieldname string
	var url string
	flag.StringVar(&filename, "filename", "", "File to upload to the target url")
	flag.StringVar(&fieldname, "fieldname", "", "Fieldname to use in the POSTed form")
	flag.StringVar(&url, "url", "", "Url to POST to")
	flag.Parse()
	if len(filename) < 1 || len(fieldname) < 1 || len(url) < 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	fmt.Printf("Uploading '%s' to '%s' in form field  '%s'... \n", filename, url, fieldname)
	bd, ioerr := getRequestBody(filename, fieldname)
	if ioerr != nil {
		fmt.Printf("Could not read file for upload: %s \n", ioerr.Error())
		os.Exit(2)
	}
	if status, err := post(bd.body, url, bd.contentType); err != nil {
		fmt.Printf("Request failed: %s \n", err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("Remote responded with status code %d \n", status)
	}
}

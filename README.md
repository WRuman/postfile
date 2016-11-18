# Postfile
=======
Simple utility for POSTing files as `multipart/form-data`, written in golang

## Building
You will need a go build environment, which you can learn how to set up [in the golang documentation](https://golang.org/doc/install)

Once you have your environment set up, creating an executable should be as simple as running `go get`
and `go install`

## Usage
When running the program, it takes the following flags:
1. `-filename` The name of the file you want to upload
2. `-fieldname` The field name of the uploaded file in the simulated form that is sent to the url
3. `-url` The url to make a POST request to
4. `-help` Show acceptable argument list with descriptions and exit

All the flags except `-help` are required to make a complete request.

When run, the program will provide the status code of the response and exit.

## Improvements
This is a very basic program that I threw together to help me at work. Potential improvements include:
* Better error reporting and the ability to retry failed requests
* Ability to batch requests based on some input file
* Ability make POSTs with content types besides `multipart/form-data`
* Avoiding buffering the whole upload into memory before sending it


# Frontend Server for JavaScript SPA with Go

Server for a JavaScript SPA Frontend to server assets and proxy requests to the api backend.

Built with [Go](), [Gorilla Router](https://github.com/gorilla/mux) and [NYTimes Gziphandler](https://github.com/NYTimes/gziphandler).

## Step 01
Serve static assets with Go standard library. First we need to download the Gorilla Mux router with:

	go get github.com/gorilla/mux

Now we can run the server for our asets with:

	# Server runs at http://localhost:3000
	go run step-01/main.go


## Step 02
Added proxy for requests to the api backend.

	# Server runs at http://localhost:3000
	go run step-02/main.go

## Step 03
Added gzip filter to compress assets. We will use the [NYTimes Gziphandler](https://github.com/NYTimes/gziphandler) so 
we need to download it first:

	go get github.com/NYTimes/gziphandler

Now we can run the server:

	# Server runs at http://localhost:3000
	go run step-03/main.go

## Step 04
Added health check and ability to configure api url.

	# Server runs at http://localhost:3000
	go run step-04/main.go

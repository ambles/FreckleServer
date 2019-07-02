# Basic Web Server with GO

## Installing golang and git

`sudo apt-get install golang git`

 

## Setting up the directories

`cd ~`

`mkdir -p work/src`

`mkdir work/bin`

`cd work/src`

`mkdir FreckleServer `

`cd FreckleServer/`

 

## Getting Ambles’s sample code

`git init`

`git pull git@github.com:ambles/FreckleServer.git`

 

## Building the server

`export GOPATH=$HOME/work`

`export PATH=$PATH:$GOPATH/bin`

`cd ../`(you should be in \~/work/src now)

`go get github.com/gorilla/mux`

`go get github.com/lib/pq`

`cd ../bin`

`go build FreckleServer`

 

## Running the server

`./FreckleServer`

 

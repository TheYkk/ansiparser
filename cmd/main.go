package main

import "github.com/theykk/ansiparser"

func main() {
	data := `[0K[36;1mGetting source from Git repository[0;m
[0;m[32;1mFetching changes...[0;m
Initialized empty Git repository in /builds/TheYkk/docker-nginx-static/.git/
[32;1mCreated fresh repository.[0;m
[32;1mChecking out 160213f6 as refs/heads/v3...[0;m`
	ansiparser.Parser([]byte(data))
}

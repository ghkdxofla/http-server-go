package main

import "flag"

var directory string

func init() {
	if flag.Lookup("directory") == nil {
		flag.StringVar(&directory, "directory", "", "for directory")
	}
}

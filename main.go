package main

import (
	"errors"
	"fmt"

	"github.com/snowmerak/error-tree/etree"
)

func main() {
	Timeout := etree.New(errors.New("timeout"), nil)
	Badrequest := etree.New(errors.New("bad request"), nil)
	Notfound := etree.New(errors.New("not found"), nil)
	Facebook := etree.New(errors.New("facebook"), nil)
	Google := etree.New(errors.New("google"), nil)

	FacebookTimeout := etree.New(errors.New("facebook timeout"), Timeout, Facebook)
	FacebookBadrequest := etree.New(errors.New("facebook bad request"), Badrequest, Facebook)
	FacebookNotfound := etree.New(errors.New("facebook not found"), Notfound, Facebook)
	GoogleTimeout := etree.New(errors.New("google timeout"), Timeout, Google)
	GoogleBadrequest := etree.New(errors.New("google bad request"), Badrequest, Google)
	GoogleNotfound := etree.New(errors.New("google not found"), Notfound, Google)

	fmt.Println(
		etree.Cover(FacebookTimeout, Timeout), " ",
		etree.Cover(FacebookBadrequest, Facebook), " ",
		etree.Cover(FacebookNotfound, Facebook), " ",
		etree.Cover(GoogleTimeout, Timeout), " ",
		etree.Cover(GoogleBadrequest, Google), " ",
		etree.Cover(GoogleNotfound, Notfound), " ",
		etree.Cover(GoogleNotfound, Google), " ",
		etree.Cover(GoogleBadrequest, Facebook), " ",
		etree.Cover(FacebookTimeout, Google), " ",
	)
}

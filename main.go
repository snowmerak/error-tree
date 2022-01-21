package main

import (
	"fmt"

	"github.com/snowmerak/error-tree/etree"
)

func main() {
	Timeout := etree.New("timeout", nil)
	Badrequest := etree.New("bad request", nil)
	Notfound := etree.New("not found", nil)
	Facebook := etree.New("facebook", nil)
	Google := etree.New("google", nil)

	FacebookTimeout := etree.New("facebook timeout", Timeout, Facebook)
	FacebookBadrequest := etree.New("facebook bad request", Badrequest, Facebook)
	FacebookNotfound := etree.New("facebook not found", Notfound, Facebook)
	GoogleTimeout := etree.New("google timeout", Timeout, Google)
	GoogleBadrequest := etree.New("google bad request", Badrequest, Google)
	GoogleNotfound := etree.New("google not found", Notfound, Google)

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

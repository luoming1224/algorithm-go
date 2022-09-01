package main

type S int

func main() {
	a := make([]*int, 1)
	b := 12
	a[0] = &b

	c := make(map[string]*int)
	d := 14
	c["aaa"] = &d

	e := make(chan *int, 1)
	f := 15
	e <- &f
}

package main

import (
	"fmt"

	bw "github.com/immesys/bw2bind"
	. "github.com/immesys/bw2bind/expr"
)

const NS = "scratch.ns"

func main() {
	// This will connect to local router on localhost:28589
	cl := bw.ConnectOrExit("")
	cl.SetEntityFromEnvironOrExit()

	//Create a view that contains all resources on
	//the NS namespace.
	// M is shorthand for map[string]interface{}
	// A is shorthand for []string
	v1, err := cl.CreateView(M{
		"ns": A{NS}, "meta": M{"inview": "true"},
	})

	if err != nil {
		panic(err)
	}

	res, err := v1.List()
	if err != nil {
		panic(err)
	}
	for _, i := range res {
		fmt.Println(i)
	}
	//Subscribe to all i.helloworld interfaces
	//in the view:
	v1.SubSignalFOrExit("i.echo", "msg", func(sm *bw.SimpleMessage) {
		sm.Dump()
	})

	for {
	}
}

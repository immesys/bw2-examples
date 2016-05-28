// This application is a little bit like an echo server, but also keeps
// state. You can publish to the msg slot on i.echo to set what the message
// is, and that is the message that will be published on msg signal every
// five seconds.

package main

import (
	"fmt"
	"time"

	"github.com/immesys/spawnpoint/spawnable"

	bw "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	//Connect
	cl := bw.ConnectOrExit("")
	cl.SetEntityFromEnvironOrExit()
	params := spawnable.GetParamsOrExit()

	//This registers a service rooted at the uri specified by spawnpoint
	//or params.yml
	uri := params.MustString("base_uri")
	svc := cl.RegisterService(uri, "s.helloworld")

	//This sets a metadata key on the service
	svc.SetMetadata("demoapp", "helloworld")

	//You can have multiple interfaces per service. The second parameter
	//is the interface type, the first is the name of that instance of the
	//interface. We only have one interface, so underscore is a placeholder
	iface := svc.RegisterInterface("_", "i.echo")

	//Now let us implement the interface:
	msg := "unset echo message"

	//People can set what the message should be
	iface.SubscribeSlot("msg", func(m *bw.SimpleMessage) {
		if newmsg := m.GetOnePODF(bw.PODFString); newmsg != nil {
			msg = newmsg.(bw.TextPayloadObject).Value()
		}
	})

	//Also, every five seconds, we publish the message
	for {
		po := bw.CreateTextPayloadObject(bw.PONumString, msg)
		err := iface.PublishSignal("msg", po)
		fmt.Println("Published, error was", err)
		time.Sleep(5 * time.Second)
	}
}

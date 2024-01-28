package snmps

import (
	"net"
	"sync"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/ruraomsk/ag-server/logger"
)

var mutex sync.Mutex
var CommandSnmp chan gosnmp.SnmpPDU

func Start(address string) {

	CommandSnmp = make(chan gosnmp.SnmpPDU)

	gosnmp.Default.Target = "192.168.88.1"
	gosnmp.Default.Community = "private"
	for {
		err := gosnmp.Default.Connect()
		if err != nil {
			logger.Error.Printf("Connect() err: %v", err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	logger.Info.Printf("Starting SNMP TRAP Server on: %s\n", address)
	go requestor()
	go commander()
	tl := gosnmp.NewTrapListener()
	tl.OnNewTrap = myTrapHandlerTCP
	tl.Params = gosnmp.Default
	tl.Params.Community = "private"

	err := tl.Listen(address)
	if err != nil {
		time.Sleep(1 * time.Second)
		logger.Error.Fatalf("Error in TRAP listen: %s\n", err)
	}
}
func commander() {
	for {
		cmd := <-CommandSnmp
		cmds := make([]gosnmp.SnmpPDU, 0)
		cmds = append(cmds, cmd)
		mutex.Lock()
		t := gosnmp.SnmpTrap{
			AgentAddress: gosnmp.Default.Target,
			Variables:    cmds,
			Timestamp:    uint(time.Now().Unix()),
		}
		_, err := gosnmp.Default.SendTrap(t)
		mutex.Unlock()
		if err != nil {
			logger.Error.Printf("send trap %s", err.Error())
		}
	}
}
func requestor() {
	for {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		t := gosnmp.SnmpTrap{
			AgentAddress: gosnmp.Default.Target,
			Variables:    makereq(),
			Timestamp:    uint(time.Now().Unix()),
		}
		_, err := gosnmp.Default.SendTrap(t)
		mutex.Unlock()
		if err != nil {
			logger.Error.Printf("send trap %s", err.Error())
		}
	}
}
func myTrapHandlerTCP(packet *gosnmp.SnmpPacket, addr *net.UDPAddr) {

	logger.Debug.Printf("SNMP trap received from: %s:%d. Community:%s, SnmpVersion:%s \n",
		addr.IP, addr.Port, packet.Community, packet.Version)
	for _, variable := range packet.Variables {
		// var val string
		// switch variable.Type {
		// case gosnmp.OctetString:
		// 	val = string(variable.Value.([]byte))
		// case gosnmp.ObjectIdentifier:
		// 	val = fmt.Sprintf("%s", variable.Value)
		// case gosnmp.TimeTicks:
		// 	a := gosnmp.ToBigInt(variable.Value)
		// 	val = fmt.Sprintf("%d", (*a).Int64())
		// case gosnmp.Null:
		// 	val = ""
		// default:
		// 	// ... or often you're just interested in numeric values.
		// 	// ToBigInt() will return the Value as a BigInt, for plugging
		// 	// into your calculations.
		// 	a := gosnmp.ToBigInt(variable.Value)
		// 	val = fmt.Sprintf("%d", (*a).Int64())
		// }
		// logger.Debug.Printf("trap - oid[%d]: %s (%s) = %v \n", i, variable.Name, variable.Type, val)
		updater(variable)
	}
}

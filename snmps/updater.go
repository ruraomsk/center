package snmps

import (
	"fmt"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/scenter/state"
)

func getInteger(variable gosnmp.SnmpPDU) int {
	switch variable.Type {
	case gosnmp.OctetString:
		return 0
	case gosnmp.ObjectIdentifier:
		return 0
	case gosnmp.TimeTicks:
		return 0
	case gosnmp.Null:
		return 0
	}
	return int(gosnmp.ToBigInt(variable.Value).Uint64())
}
func getString(variable gosnmp.SnmpPDU) string {
	switch variable.Type {
	case gosnmp.OctetString:
		return string(variable.Value.([]byte))
	case gosnmp.ObjectIdentifier:
		return fmt.Sprintf("%s", variable.Value)
	case gosnmp.TimeTicks:
		a := gosnmp.ToBigInt(variable.Value)
		return fmt.Sprintf("%d", (*a).Int64())
	case gosnmp.Null:
		return ""
	default:
		a := gosnmp.ToBigInt(variable.Value)
		return fmt.Sprintf("%d", (*a).Int64())
	}
}
func getTime(variable gosnmp.SnmpPDU) time.Time {
	switch variable.Type {
	case gosnmp.OctetString:
		return time.Unix(0, 0)
	case gosnmp.ObjectIdentifier:
		return time.Unix(0, 0)
	case gosnmp.TimeTicks:
		a := gosnmp.ToBigInt(variable.Value)
		return time.UnixMicro((*a).Int64())
	case gosnmp.Null:
		return time.Unix(0, 0)
	default:
		a := gosnmp.ToBigInt(variable.Value)
		return time.UnixMicro((*a).Int64())
	}
}
func updater(variable gosnmp.SnmpPDU) {
	var oid = gosnmp.SnmpPDU{Name: variable.Name}
	switch variable.Name[1:] {
	case "1.3.6.1.4.1.1618.3.6.2.1.2.0":
		//indicates the major state of a traffic controller
		state.RemoteDevice.SetStatus(getInteger(variable))
	case "1.3.6.1.4.1.1618.3.6.2.2.2.0":
		//Режим работы
		state.RemoteDevice.SetState(getInteger(variable))
	case "1.3.6.1.4.1.1618.3.5.2.1.7.0":
		//Номер плана
		state.RemoteDevice.SetPlan(getInteger(variable))
	case "1.3.6.1.4.1.1618.3.7.2.1.3.0":
		//Источник плана
		state.RemoteDevice.SetSource(getInteger(variable))
	case "1.3.6.1.4.1.1618.3.7.2.11.2.0":
		//Номер фазы
		state.RemoteDevice.SetPhase(getInteger(variable))
	case "1.3.6.1.4.1.1618.3.5.2.1.6.0":
		//Состояние сигнальных групп
		state.RemoteDevice.SetSignals([]byte(getString(variable)))
	case "1.3.6.1.4.1.1618.3.1.2.2.2.0":
		//Тревоги
		state.RemoteDevice.SetAlarm(getInteger(variable), 0)
	case "1.3.6.1.4.1.1618.3.1.2.2.3.0":
		//Тревоги
		state.RemoteDevice.SetAlarm(getInteger(variable), 1)
	case "1.3.6.1.4.1.1618.3.1.2.2.4.0":
		//Тревоги
		state.RemoteDevice.SetMessage(getString(variable))
	default:
		logger.Error.Printf("Not updater from  %s \n", oid.Name)
	}
}

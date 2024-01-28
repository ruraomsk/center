package snmps

import "github.com/gosnmp/gosnmp"

func SetPlan(value int) gosnmp.SnmpPDU {
	// "Вызов плана":
	return gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.1618.3.7.2.2.1.0",
		Type:  gosnmp.Integer,
		Value: value,
	}
}
func SetPhase(value int) gosnmp.SnmpPDU {
	// "Вызов фазы":
	return gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.1618.3.7.2.11.1.0",
		Type:  gosnmp.Integer,
		Value: value,
	}
}
func SetFlashing(value int) gosnmp.SnmpPDU {
	// "Вызов ЖМ":
	return gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.1618.3.2.2.1.1.0",
		Type:  gosnmp.Integer,
		Value: value,
	}
}
func SetDark(value int) gosnmp.SnmpPDU {
	// "Вызов ОС":
	return gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.1618.3.2.2.2.1.0",
		Type:  gosnmp.Integer,
		Value: value,
	}
}
func SetAllRed(value int) gosnmp.SnmpPDU {
	// "Вызов КК":
	return gosnmp.SnmpPDU{
		Name:  "1.3.6.1.4.1.1618.3.2.2.4.1.0",
		Type:  gosnmp.Integer,
		Value: value,
	}
}

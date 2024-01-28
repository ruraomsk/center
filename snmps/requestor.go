package snmps

import "github.com/gosnmp/gosnmp"

func makereq() []gosnmp.SnmpPDU {
	result := make([]gosnmp.SnmpPDU, 0)
	result = append(result, gosnmp.SnmpPDU{
		//indicates the major state of a traffic controller
		Name: "1.3.6.1.4.1.1618.3.6.2.1.2.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Режим работы
		Name: "1.3.6.1.4.1.1618.3.6.2.2.2.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Номер плана
		Name: "1.3.6.1.4.1.1618.3.5.2.1.7.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Источник плана
		Name: "1.3.6.1.4.1.1618.3.7.2.1.3.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Номер фазы
		Name: "1.3.6.1.4.1.1618.3.7.2.11.2.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Состояние сигнальных групп
		Name: "1.3.6.1.4.1.1618.3.5.2.1.6.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Тревоги
		Name: "1.3.6.1.4.1.1618.3.1.2.2.2.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Тревоги
		Name: "1.3.6.1.4.1.1618.3.1.2.2.3.0",
		Type: gosnmp.Null,
	})
	result = append(result, gosnmp.SnmpPDU{
		//Тревоги
		Name: "1.3.6.1.4.1.1618.3.1.2.2.4.0",
		Type: gosnmp.Null,
	})
	return result
}

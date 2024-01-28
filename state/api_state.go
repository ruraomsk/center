package state

import "fmt"

func CopyDevice() Device {
	mutex.Lock()
	defer mutex.Unlock()
	return RemoteDevice
}

func (d *Device) PlanToString() string {
	if d.Plan > 0 && d.Plan <= 13 {
		return fmt.Sprintf("Локальный %d", d.Plan)
	}
	switch d.Plan {
	case 14:
		return "Управление Spot/Utopia"
	case 15:
		return "ВПУ"
	case 16:
		return "Центр"
	case 17:
		return "ЖМ"
	case 18:
		return "ОС"
	case 19:
		return "КК"
	}
	return fmt.Sprintf("НП %d", d.Plan)
}
func (d *Device) SourceToString() string {
	switch d.Source {
	case 1:
		return "trafficActuatedPlanSelectionCommand(1)"
	case 2:
		return "currentTrafficSituationCentral(2)"
	case 3:
		return "controlBlockOrInput(3)"
	case 4:
		return "manuallyFromWorkstation(4)"
	case 5:
		return "emergencyRoute(5)"
	case 6:
		return "currentTrafficSituation(6)"
	case 7:
		return "calendarClock(7)"
	case 8:
		return "controlBlockInLocal(8)"
	case 9:
		return "forcedByParameterBP40(9)"
	case 10:
		return "startUpPlan(10)"
	case 11:
		return "localPlan(11)"
	case 12:
		return "manualControlPlan(12)"
	}
	return fmt.Sprintf("НИ %d", d.Plan)
}
func (d *Device) StateToString() string {
	switch d.State {
	case 0:
		return "noInfo(0)"
	case 1:
		return "operatesInPriorityA4(1),"
	case 2:
		return "operatesInPriorityA3(2)"
	case 3:
		return "operatesInPriorityA2(3)"
	case 4:
		return "operatesInPriorityA1(4)"
	case 5:
		return "operatesInEmergencyMode(5)"
	case 6:
		return "operatesInManualControl(6)"
	case 7:
		return "operatesInLinkedMode(7)"
	case 8:
		return "operatesInIsolatedMode(8)"
	case 9:
		return "operatesInCoordinatedMode(9)"
	case 10:
		return "operatesInSPOTManualMode(10)"
	case 11:
		return "operatesInSPOTCentralizedMode(11)"
	case 12:
		return "operatesInSPOTLocalMode(12)"
	}
	return fmt.Sprintf("НИ %d", d.Plan)
}
func (d *Device) StatusToString() string {
	switch d.Status {
	case 0:
		return "noInformation(0)"
	case 1:
		return "workingProperly(1)"
	case 2:
		return "powerUp(2)"
	case 3:
		return "dark(3),"
	case 4:
		return "flash(4)"
	case 5:
		return "partialFlash(5)"
	case 6:
		return "allRed(6)"
	}
	return fmt.Sprintf("НC %d", d.Plan)
}
func (d *Device) AlarmToString() string {
	return fmt.Sprintf("[%d,%d] <%s>", d.Alarm[0], d.Alarm[1], d.Message)
}

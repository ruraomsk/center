package state

import (
	"sync"
	"time"
)

var mutex sync.Mutex

var RemoteDevice = Device{Alarm: [2]int{0, 0},
	Message:     "",
	SignalGroup: make([]byte, 0),
}

type Device struct {
	LastOperation time.Time
	Plan          int
	Source        int //Источник плана
	Phase         int
	Status        int //Статус
	State         int //Режим работы
	SignalGroup   []byte
	Alarm         [2]int
	Message       string
}

func (d *Device) SetPlan(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	d.LastOperation = time.Now()
	d.Plan = value
}
func (d *Device) SetPhase(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	d.LastOperation = time.Now()
	d.Phase = value
}
func (d *Device) SetSource(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	d.LastOperation = time.Now()
	d.Source = value
}
func (d *Device) SetStatus(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	d.LastOperation = time.Now()
	d.Status = value
}
func (d *Device) SetState(value int) {
	mutex.Lock()
	defer mutex.Unlock()
	d.LastOperation = time.Now()
	d.State = value
}
func (d *Device) SetSignals(value []byte) {
	mutex.Lock()
	defer mutex.Unlock()
	d.LastOperation = time.Now()
	d.SignalGroup = value
}

func (d *Device) SetAlarm(value int, pos int) {
	mutex.Lock()
	defer mutex.Unlock()
	if pos < len(d.Alarm) {
		d.LastOperation = time.Now()
		d.Alarm[pos] = value
	}
}
func (d *Device) SetMessage(value string) {
	mutex.Lock()
	defer mutex.Unlock()
	d.LastOperation = time.Now()
	d.Message = value
}

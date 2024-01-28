package web

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anoshenko/rui"
	"github.com/ruraomsk/ag-server/logger"
	"github.com/ruraomsk/scenter/snmps"
	"github.com/ruraomsk/scenter/state"
)

const mainText = `
ListLayout {
	style = showPage,
	orientation = vertical,
	content = [
		GridLayout {
			content = [
				ListLayout {
					orientation = vertical,
					row = 0, column = 0,
					content = [
						TextView {
							id=idLine1,text = "Line1",text-size="18px",
						},
						TextView {
							id=idLine2,text = "Line2",text-size="18px",
						},
						TextView {
							id=idLine3,text = "Line3",text-size="18px",
						},
						TableView {
							style=table,
							cell-horizontal-align = left,  title = "Направления", id=idNaps,
						},
					]
				},
				GridLayout {
					row = 0, column = 1,
					content = [
						TextView {
							row = 0, column = 0,
							id="Flashing",text-align = center,
							text = "ЖМ",background-color=gray,
						},
						TextView {
							row = 0, column = 1,
							id="AllRed",text-align = center,
							text = "КК",background-color=gray,
						},
						TextView {
							row = 0, column = 2,
							id="Dark",text-align = center,
							text = "ОС",background-color=gray,
						},
						Button {
							row = 2, column = 1,
							id=setAllRedOn,content="КК вкл"
						},
						Button {
							row = 2, column = 0,
							id=setFlashingOn,content="ЖМ вкл"
						},
						Button {
							row = 2, column = 2,
							id=setDarkOn,content="ОС вкл"
						},
						Button {
							row = 3, column = 1,
							id=setAllRedOff,content="КК выкл"
						},
						Button {
							row = 3, column = 0,
							id=setFlashingOff,content="ЖМ выкл"
						},
						Button {
							row = 3, column = 2,
							id=setDarkOff,content="ОС выкл"
						},
						Button {
							row = 6, column = 0:1,
							id=setPlan,content="Установить План"
						},
						NumberPicker {
							row = 6, column = 2,
							id=idPlan,type=editor,min=0,max=32,value=0
						},
						Button {
							row = 7, column = 0:1,
							id=setPhase,content="Установить Фазу"
						},
						NumberPicker {
							row = 7, column = 2,
							id=idPhase,type=editor,min=0,max=32,value=0
						},

					]
				},
			]
		},
	]
}
`

func getInteger(a any) (result int) {
	s, ok := a.(string)
	if ok {
		result, _ = strconv.Atoi(s)
	}
	f, ok := a.(float64)
	if ok {
		result = int(f)
	}
	return
}

func makeButtonOnScreen(view rui.View) {
	rui.Set(view, "setAllRedOn", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetAllRed(1)
		logger.Info.Printf("Оператор установил КК")
	})
	rui.Set(view, "setAllRedOff", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetAllRed(0)
		logger.Info.Printf("Оператор отменил КК")
	})
	rui.Set(view, "setFlashingOn", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetFlashing(1)
		logger.Info.Printf("Оператор установил ЖМ")
	})
	rui.Set(view, "setFlashingOff", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetFlashing(0)
		logger.Info.Printf("Оператор отменил ЖМ")
	})
	rui.Set(view, "setDarkOn", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetDark(1)
		logger.Info.Printf("Оператор установил ОС")
	})
	rui.Set(view, "setDarkOff", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetDark(0)
		logger.Info.Printf("Оператор отменил ОС")
	})
	rui.Set(view, "setPlan", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetPlan(getInteger(rui.Get(view, "idPlan", "value")))
		logger.Info.Printf("Оператор вызвал план %d", getInteger(rui.Get(view, "idPlan", "value")))
	})
	rui.Set(view, "setPhase", rui.ClickEvent, func(rui.View) {
		snmps.CommandSnmp <- snmps.SetPhase(getInteger(rui.Get(view, "idPhase", "value")))
		logger.Info.Printf("Оператор вызвал фазу %d", getInteger(rui.Get(view, "idPhase", "value")))
	})
}
func updatePartKDM(view rui.View) {
	hs := state.CopyDevice()
	rui.Set(view, "idLine1", "text", fmt.Sprintf("<b>План</b> %s <b>Фаза</b> %d",
		hs.PlanToString(), hs.Phase))
	rui.Set(view, "idLine2", "text", fmt.Sprintf("<b>Источник</b> %s <b>Состояние</b> %s %s",
		hs.SourceToString(), hs.StatusToString(), hs.StateToString()))
	rui.Set(view, "idLine3", "text", fmt.Sprintf("<b>Alarms</b> %s ",
		hs.AlarmToString()))

	if hs.Status == 6 {
		rui.Set(view, "AllRed", "background-color", "green")
	} else {
		rui.Set(view, "AllRed", "background-color", "gray")
	}
	if hs.Status == 5 {
		rui.Set(view, "Flashing", "background-color", "green")
	} else {
		rui.Set(view, "Flashing", "background-color", "gray")
	}
	if hs.Status == 3 {
		rui.Set(view, "Dark", "background-color", "green")
	} else {
		rui.Set(view, "Dark", "background-color", "gray")
	}

	var content [][]any
	content = append(content, []any{"Нап", "Состояние"})
	count := 1
	for i := 0; i < len(hs.SignalGroup); i++ {
		ds := ""
		switch hs.SignalGroup[i] {
		case 0:
			ds = "все сигналы выключены"
		case 1:
			ds = "направление перешло в неактивное состояние, желтый после зеленого"
		case 2:
			ds = "направление перешло в неактивное состояние, красный"
		case 3:
			ds = "направление перешло в активное состояние, красный"
		case 4:
			ds = "направление перешло в активное состояние, красный c желтым"
		case 5:
			ds = "направление перешло в активное состояние, зеленый"
		case 6:
			ds = "направление не меняло свое состояние, зеленый"
		case 7:
			ds = "направление не меняло свое состояние, красный"
		case 8:
			ds = "зеленый мигающий сигнал"
		case 9:
			ds = "желтый мигающий в режиме ЖМ"
		case 10:
			ds = "сигналы выключены в режиме ОС"
		case 11:
			ds = "неиспользуемое направление"
		default:
			ds = "error code"
		}
		content = append(content, []any{i, ds})
		count++
	}
	rui.SetParams(view, "idNaps", rui.Params{
		rui.Content:             content,
		rui.HeadHeight:          1,
		rui.CellPadding:         "1px",
		rui.CellHorizontalAlign: "left",
	})

}
func makeMainScreen(view rui.View) {
	mutex.Lock()
	defer mutex.Unlock()
	updatePartKDM(view)
}
func updaterScreen(view rui.View, session rui.Session) {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		if view == nil {
			return
		}
		w, ok := SessionStatus[session.ID()]
		if !ok {
			continue
		}

		if !w {
			continue
		}
		makeMainScreen(view)
	}
}

func mainScreen(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, mainText)
	if view == nil {
		return nil
	}
	makeMainScreen(view)
	makeButtonOnScreen(view)
	go updaterScreen(view, session)
	return view
}

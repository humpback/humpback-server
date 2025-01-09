package controller

type ControllerInter interface {
}

var Controller ControllerInter = &controller{}

type controller struct{}

func Start(stopCh <-chan struct{}) {
	go SessionGCInterval(stopCh)
}

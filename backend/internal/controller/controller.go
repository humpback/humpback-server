package controller

type ControllerInter interface {
}

var Controller ControllerInter = &controller{}

type controller struct{}

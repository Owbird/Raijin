package app

import (
	"log"
	"reflect"
)

type Action struct {
	Obj     any
	Pkg     string
	Methods []string
}

type App struct {
	Actions []Action
}

func NewApp() *App { return &App{} }

func (a *App) Bind(act any) {
	t := reflect.TypeOf(act)

	pkg := t.String()

	action := Action{Obj: act, Pkg: pkg, Methods: []string{}}

	for i := 0; i < t.NumMethod(); i++ {

		method := t.Method(i)

		action.Methods = append(action.Methods, method.Name)
	}

	log.Println(act)

	a.Actions = append(a.Actions, action)
}

func (a *App) Run() {
	for _, act := range a.Actions {
		log.Println(act.Obj, act.Pkg, act.Methods)

		v := reflect.ValueOf(act.Obj)

		for _, m := range act.Methods {

			method := v.MethodByName(m)

			method.Call(nil)

		}
	}
}

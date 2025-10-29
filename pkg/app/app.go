package app

import (
	"log"
	"os"
	"path/filepath"
	"raijin/internal/config"
	"raijin/internal/generator"
	"reflect"
	"strings"
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
	appDirs := config.GetAppDirs(nil)

	for _, act := range a.Actions {
		log.Println(act.Obj, act.Pkg, act.Methods)

		structName := strings.Split(act.Pkg, ".")[1]

		actionDir := filepath.Join(appDirs.ActionsDir, structName)

		os.MkdirAll(actionDir, config.FileMode)

		actionFile := filepath.Join(actionDir, "index.ts")

		actionMethods := []string{}

		v := reflect.ValueOf(act.Obj)
		for _, m := range act.Methods {
			params := []string{}
			returnValues := []string{}

			log.Println("*****************")

			log.Printf("Method Name: %v\n", m)

			// os.WriteFile(appDirs.ActionsDir, []byte(), config.FileMode)

			method := v.MethodByName(m)

			log.Printf("Method Signature: %v\n", method.Type())
			log.Printf("Number of params: %v\n", method.Type().NumIn())
			log.Printf("Number of return values: %v\n", method.Type().NumOut())

			for i := 0; i < method.Type().NumIn(); i++ {
				params = append(params, method.Type().In(i).Name())
			}

			totalReturnValues := method.Type().NumOut()
			if totalReturnValues == 0 {
				returnValues = append(returnValues, "void")
			} else {
				for i := range totalReturnValues {
					if tsType, ok := generator.TypeMap[method.Type().Out(i).Name()]; ok {
						returnValues = append(returnValues, tsType)
					} else {
						returnValues = append(returnValues, "any")
					}
				}
			}

			log.Println("*****************")

			actionMethods = append(actionMethods, generator.GenerateActionsMethod(generator.MethodMeta{
				Name:         m,
				ParamTypes:       params,
				ReturnValues: returnValues,
			}))

		}

		actionsFileContent := strings.Join(actionMethods, "\n")

		os.WriteFile(actionFile, []byte(actionsFileContent), config.FileMode)

	}
}

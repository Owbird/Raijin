package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
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
				ParamTypes:   params,
				ReturnValues: returnValues,
			}))

		}

		actionsFileContent := strings.Join(actionMethods, "\n")

		os.WriteFile(actionFile, []byte(actionsFileContent), config.FileMode)

	}

	buildCmd := exec.Command("pnpm", "-C", appDirs.FrontendSrc, "build")
	log.Println(buildCmd.String())
	res, err := buildCmd.Output()
	if err != nil {
		log.Fatalln(err)

	}

	log.Println(string(res))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, appDirs.FrontendDist+"/index.html")
	})

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {

		frontendDir := appDirs.FrontendDist

		path := r.URL.Path
		data, err := os.ReadFile(filepath.Join(frontendDir, path))
		if err != nil {
			fmt.Print(err)
			http.NotFound(w, r)
			return
		}
		if strings.HasSuffix(path, ".js") {
			w.Header().Set("Content-Type", "text/javascript")
		} else if strings.HasSuffix(path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		_, err = w.Write(data)
		if err != nil {
			fmt.Print(err)
		}

	})

	http.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("a")

		defer r.Body.Close()

		var data map[string]any
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}

		params := []reflect.Value{}
		for _, value := range data {
			params = append(params, reflect.ValueOf(value))
		}

		for _, action := range a.Actions {
			for _, methodName := range action.Methods {
				if methodName == query {
					v := reflect.ValueOf(action.Obj)
					method := v.MethodByName(query)
					if !method.IsValid() {
						http.Error(w, "method not found", http.StatusNotFound)
						return
					}

					returnValues := method.Call(params)

					var results []any
					for _, rv := range returnValues {
						results = append(results, rv.Interface())
					}

					w.Header().Set("Content-Type", "application/json")
					if err := json.NewEncoder(w).Encode(results); err != nil {
						http.Error(w, "failed to encode response", http.StatusInternalServerError)
					}
					return
				}
			}
		}

		http.Error(w, "action not found", http.StatusNotFound)
	})

	
	log.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

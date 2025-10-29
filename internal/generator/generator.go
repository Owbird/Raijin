package generator

import (
	"fmt"
	"strings"
)

type MethodMeta struct {
	Name         string
	ParamTypes   []string
	ReturnValues []string
}

var TypeMap = map[string]string{
	"string":      "string",
	"bool":        "boolean",
	"int":         "number",
	"int8":        "number",
	"int16":       "number",
	"int32":       "number",
	"int64":       "number",
	"uint":        "number",
	"uint8":       "number",
	"uint16":      "number",
	"uint32":      "number",
	"uint64":      "number",
	"float32":     "number",
	"float64":     "number",
	"interface{}": "any",
	"error":       "string",
}

func GenerateExampleEntryFile() string {
	return `package main

import "raijin/pkg/app"

type AuthActions struct{}

func NewAuthActions() *AuthActions {
	return &AuthActions{}
}

func (aa *AuthActions) Login(email, password string) bool { 
	println("Logging in") 

	return true
}

func (aa *AuthActions) Logout(username string) bool { 
	println("Logging out") 

	return true
}

func (aa *AuthActions) Auth(username, password, confirmPassword, email string) map[string]any { 
	println("Authing") 

	return map[string]
}

func main() {
	a := app.NewApp()

	a.Bind(NewAuthActions())

	a.Run()
}`
}

func GenerateActionsMethod(metadata MethodMeta) string {
	paramsWithTypes := []string{}
	paramsOnly := []string{}

	for paramIndex, paramType := range metadata.ParamTypes {
		paramsWithTypes = append(paramsWithTypes, fmt.Sprintf("param%v: %v", paramIndex+1, paramType))

		paramsOnly = append(paramsOnly, fmt.Sprintf("param%v", paramIndex+1))
	}

	content := fmt.Sprintf(`export const %v = async (%v): Promise<%v> => {

		const res = await fetch("", {method:"POST", body: JSON.stringify({%v})})

		const data = (await res.json())

		return data

		}`, metadata.Name, strings.Join(paramsWithTypes, ", "), strings.Join(metadata.ReturnValues, "|"),
		strings.Join(paramsOnly, ","),
	)

	return content
}

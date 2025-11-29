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

import "github.com/owbird/raijin/pkg/app"

type CounterActions struct {
	Count float64
}

func NewCounterAction() *CounterActions {
	return &CounterActions{Count: 0}
}

func (aa *CounterActions) Add(value float64) bool {
	aa.Count += value
	return true
}

func (aa *CounterActions) Subtract(value float64) bool {
	aa.Count -= value
	return true
}

func (aa *CounterActions) Multiply(value float64) bool {
	aa.Count *= value
	return true
}

func (aa *CounterActions) Divide(value float64) bool {
	aa.Count /= value
	return true
}

func (aa *CounterActions) Reset() bool {
	aa.Count = 0
	return true
}

func (aa *CounterActions) GetCount() float64 {
	return aa.Count
}

func main() {
	a := app.NewApp()

	a.Bind(NewCounterAction())

	a.Run()
}
`
}

func GenerateActionsMethod(metadata MethodMeta) string {
	paramsWithTypes := []string{}
	paramsOnly := []string{}

	for paramIndex, paramType := range metadata.ParamTypes {

		paramsWithTypes = append(paramsWithTypes, fmt.Sprintf("param%v: %v", paramIndex+1, TypeMap[paramType]))

		paramsOnly = append(paramsOnly, fmt.Sprintf("param%v", paramIndex+1))
	}

	content := fmt.Sprintf(`export const %v = async (%v): Promise<%v> => {

		const res = await fetch("/action?a=%v", {method:"POST", body: JSON.stringify({%v})})

		const data = (await res.json())

		return data

		}`, metadata.Name, strings.Join(paramsWithTypes, ", "), strings.Join(metadata.ReturnValues, "|"), metadata.Name,
		strings.Join(paramsOnly, ","),
	)

	return content
}

package main

import (
	"raijin/pkg/app"
)

type AuthActions struct{}

func NewAuthAction() *AuthActions {
	return &AuthActions{}
}

func (aa *AuthActions) Login()      { println("Logging In") }
func (aa *AuthActions) Logout()     { println("Logging Out") }
func (aa *AuthActions) UpsertUser() { println("Upsert User") }

func main() {
	a := app.NewApp()

	a.Bind(NewAuthAction())

	a.Run()
}

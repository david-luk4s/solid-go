package main

import "fmt"

type Task interface {
	StartTask()
	PauseTask()
	FinallyTask()
}

type TaskHuman interface {
	Task
	CoffeeBreak()
}

type Human struct{ TaskHuman }

type TaskBot interface {
	Task
	BateryBreak()
}

type Bot struct{ TaskBot }

func main() {
	var ihm TaskHuman
	ihm = NewHuman()
	ihm.StartTask()
	ihm.PauseTask()
	ihm.CoffeeBreak()
	ihm.FinallyTask()

	var ibot TaskBot
	ibot = NewBot()
	ibot.StartTask()
	ibot.PauseTask()
	ibot.BateryBreak()
	ibot.FinallyTask()
}

func NewHuman() *Human        { return &Human{} }
func (h *Human) StartTask()   { fmt.Println("> start task human") }
func (h *Human) PauseTask()   { fmt.Println("> pause task human") }
func (h *Human) CoffeeBreak() { fmt.Println("> coffe break") }
func (h *Human) FinallyTask() { fmt.Println("> finally task human") }

func NewBot() *Bot          { return &Bot{} }
func (b *Bot) StartTask()   { fmt.Println("> start task bot") }
func (b *Bot) PauseTask()   { fmt.Println("> pause task bot") }
func (b *Bot) BateryBreak() { fmt.Println("> break for batery") }
func (b *Bot) FinallyTask() { fmt.Println("> finally task bot") }

package main

import "fmt"

// Command interface
type Command interface {
    Execute()
}

// Receiver
type Light struct {
    isOn bool
}

func NewLight() *Light {
    return &Light{isOn: false}
}

func (l *Light) TurnOn() {
    l.isOn = true
    fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
    l.isOn = false
    fmt.Println("Light is off")
}

// Concrete command
type TurnOnCommand struct {
    light *Light
}

func NewTurnOnCommand(light *Light) *TurnOnCommand {
    return &TurnOnCommand{light: light}
}

func (c *TurnOnCommand) Execute() {
    c.light.TurnOn()
}

// Concrete command
type TurnOffCommand struct {
    light *Light
}

func NewTurnOffCommand(light *Light) *TurnOffCommand {
    return &TurnOffCommand{light: light}
}

func (c *TurnOffCommand) Execute() {
    c.light.TurnOff()
}

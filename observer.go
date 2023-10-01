package main

import "fmt"

type Observer interface {
    CS2Update(string)
}

type Subject interface {
    RegisterObserver(Observer)
    DeregisterObserver(Observer)
    NotifyObservers(string)
}

type ConcreteObserver struct {
    Name string
}

func (co *ConcreteObserver) CS2Update(message string) {
    fmt.Printf("[%s] Получено обновление: %s\n", co.Name, message)
}

type ConcreteSubject struct {
    observers []Observer
}

func (cs *ConcreteSubject) RegisterObserver(observer Observer) {
    cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) DeregisterObserver(observer Observer) {
    for i, o := range cs.observers {
        if o == observer {
            cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
            break
        }
    }
}

func (cs *ConcreteSubject) NotifyObservers(message string) {
    for _, observer := range cs.observers {
        observer.CS2Update(message)
    }
}



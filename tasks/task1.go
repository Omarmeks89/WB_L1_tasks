package main

import (
    "fmt"
)

// Описываем структуру Human
type Human struct {
    Name string
    Age int
}

func (h *Human) GetName() string {
    return h.Name
}

func (h *Human) GetAge() int {
    return h.Age
}

func (h *Human) GrowOld() {
    h.Age++
}

// Now we create Action struct, that will "inherit" from Human
type Action struct {
    Action string
    // встраиваем структуру Human по имени типа
    // что позволяет обращаться к полям и методам
    // Human напрямую. Поиск полей и методов осуществляется
    // от начального адреса объемлющей структуры и далее ищет
    // нужные поля во вложенных структурах.
    Human
}

// Создаем метод у нового типа, в котором обращаемся
// к методу GetName() Human.
func (a *Action) DoAction() {
    fmt.Printf("My name is %s, i`m %s now!\n", a.GetName(), a.Action)
}

func main() {
    // Создаем структуру Human
    h := Human{Name: "Alex", Age: 25}
    fmt.Println(h.GetAge())
    // Создаем Action и встраиваем Human
    a := Action{"reading", h}
    // внутри DoAction будет вызван GetName (Human)
    a.DoAction()
    // Также можно вынести метод DoAction наружу
    // и пользоваться им для вызова DoAction
    // новых Action.
    doAction := (*Action).DoAction
    // Создаем новое действие для Human
    new_action := &Action{"walking", h}
    doAction(new_action)
    doAction(&a)
}

package main

import (
    "fmt"
)

//
// As recommended here:
// https://github-com.translate.goog/golang/go/wiki/CodeReviewComments?_x_tr_sl=auto&_x_tr_tl=ru&_x_tr_hl=ru#receiver-type
//
// i`ll use pointer as receiver in methods, bsc later
// we can increase mutable atts inside a structs.
// I wouldn`t separate methods on get/set.
//

//
// describe Human struct
//
type Human struct {
    Name string
    //
    // i know that here we can use uint8
    // bcs age can`t be bigger as ... 1xx,
    // but int is most practice i think.
    //
    Age int
}

func (h *Human) GetName() string {
    return h.Name
}

func (h *Human) GetAge() int {
    return h.Age
}

//
// we will increment age by one.
//
func (h *Human) GrowOld() {
    h.Age++
    //
    // or we can use (*h).Age like in "C" language
    // but that not needed here.
}

//
// Now we create Action struct, that will "inherit" from Human
//
type Action struct {
    //
    // we embed struct Human by name
    // TODO describe how it works...
    //
    Human
    Action string
}

//
// Now we describe method for Action
// that will show (print / send to stdout) which action do a Human now.
//
func (a *Action) DoAction() {
    //
    // here we can use getter-method GetName from Human
    // bcs TODO -->>>...
    //
    // i prefer using methods bcs attr name can be changed
    // we embed Human here and not sure about exact attr name, but
    // we`re sure about method name and signature so if it will
    // be changed to, it will be much more easier to find and debug
    // this change.
    //
    fmt.Printf("My name is %s, i`m %s now!\n", a.GetName(), a.Action)
}

// so, now we can check how it works:
//
func main() {
    h := Human{Name: "Alex", Age: 25}
    //
    // so, Alex, how old are you??
    //
    fmt.Println(h.GetAge())
    //
    // below we can create an Action like:
    // a := Action{Human{Name: "Name", Age: 10}, Action: "playing"}
    //
    // it`s time to use action:
    //
    a := Action{h, "reading"}
    //
    // So, Alex, what are you doing now??
    //
    a.DoAction()
    // 
    // but we can sepatate this method <DoAction> and
    // use it as a function for other Actions:
    //
    doAction := (*Action).DoAction
    //
    // but we should use pointer, bcs we put Action
    // into method as a ptr.
    //
    new_action := &Action{h, "walking"}
    //
    // so, Alex, and what are you doing now?
    //
    doAction(new_action)
    doAction(&a)
}

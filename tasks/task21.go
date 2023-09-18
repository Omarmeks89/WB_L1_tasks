package main

import (
    "fmt"
)

// определяем общий интерфейс для всех
// устройств
type USB interface {
    Connect()
}

type User struct {
}

// создаем пользователя и его метод
// для подключения устройств
func (u *User) connectDevice(d USB) {
    fmt.Printf("%T connect device %T\n", *u, d)
    d.Connect()
}

// создаем устройства, реализующие
// интернфейс Connect()
type Flash struct {
}

func (f *Flash) Connect() {
    fmt.Println("Flash connected...")
}

type SSD struct {
}

func (s *SSD) Connect() {
    fmt.Println("SSD connected...")
}

// И устройство, релизующее иной интерфейс
type OldFashionKeyboard struct {
}

func (ok *OldFashionKeyboard) PCI() {
    fmt.Println("OldFashionKeyboard connected by PCI...")
}

// для устройства с другим интерфейсом создадим адаптер - 
// тип, который будет раелизовывать общий интерфейс COnnect()
// и инкапсулировать устройство в своем атрибуте, вызывая
// его метод в методе Connect()
type KeyboardAdapter struct {
    k *OldFashionKeyboard
}

func (ka *KeyboardAdapter) Connect() {
    ka.k.PCI()
    fmt.Printf("%T connected via %T...\n", *ka.k, *ka)
}

func main() {
    // Благодаря адаптеру, пользователь
    // может подключить все устройства, пользуясь
    // привычным ему интерфейсом.
    user := User{}
    flash := Flash{}
    ssdDisk := SSD{}

    oldKeyboard := OldFashionKeyboard{}
    kAdapter := KeyboardAdapter{k: &oldKeyboard}

    user.connectDevice(&flash)
    user.connectDevice(&ssdDisk)
    user.connectDevice(&kAdapter)
}

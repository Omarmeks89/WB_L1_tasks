package main

import (
    "fmt"
)

// system interface that use User
type USB interface {
    Connect()
}

type User struct {
}

// User`s connect method
func (u *User) connectDevice(d USB) {
    fmt.Printf("%T connect device %T\n", *u, d)
    d.Connect()
}

type Flash struct {
}

// kind of device, that implement USB interface
func (f *Flash) Connect() {
    fmt.Println("Flash connected...")
}

type SSD struct {
}

// next device, that implement USB to
func (s *SSD) Connect() {
    fmt.Println("SSD connected...")
}

type OldFashionKeyboard struct {
}

// ups... what`s a PCI?
func (ok *OldFashionKeyboard) PCI() {
    fmt.Println("OldFashionKeyboard connected by PCI...")
}

type KeyboardAdapter struct {
    k *OldFashionKeyboard
}

// here we go with USB adapter for PCI
func (ka *KeyboardAdapter) Connect() {
    ka.k.PCI()
    fmt.Printf("%T connected via %T...\n", *ka.k, *ka)
}

// so, let`s build it all together!
func main() {
    user := User{}
    flash := Flash{}
    ssdDisk := SSD{}

    oldKeyboard := OldFashionKeyboard{}
    kAdapter := KeyboardAdapter{k: &oldKeyboard}

    user.connectDevice(&flash)
    user.connectDevice(&ssdDisk)
    user.connectDevice(&kAdapter)
}

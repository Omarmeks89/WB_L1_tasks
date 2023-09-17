package main

import (
    "fmt"
    "log"
    "os"
    "math/big"
    "strings"
)

const (
    add string = "+"
    sub string = "-"
    mul string = "x"
    div string = "/"
    dot byte = '.'
    base int = 10
    floatType string = "float"
    intType string = "int"
)

type args []string

type Operation interface {
    Add() string
    Mul() string
    Div() string
    Sub() string
}

type IntOperation struct {
    op *big.Int
    left *big.Int
    right *big.Int
    base int
}

func (o *IntOperation) Add() string {
    return o.op.Add(o.left, o.right).Text(o.base)
}

func (o *IntOperation) Mul() string {
    return o.op.Mul(o.left, o.right).Text(o.base)
}

func (o *IntOperation) Div() string {
    return o.op.Div(o.left, o.right).Text(o.base)
}

func (o *IntOperation) Sub() string {
    return o.op.Sub(o.left, o.right).Text(o.base)
}

func NewIntOpr(a, b string, base int) *IntOperation {
    op := new(big.Int)
    op1 := new(big.Int)
    op_l, _ := op.SetString(a, base)
    op_r, _ := op1.SetString(b, base)
    return &IntOperation{
        base:               base,
        op:                 op,
        left:               op_l,
        right:              op_r,
    }
}

type FloatOperation struct {
    op *big.Float
    left *big.Float
    right *big.Float
    prec int
    // data repres mode
    mode byte
}

func (o *FloatOperation) Add() string {
    return o.op.Add(o.left, o.right).Text(o.mode, o.prec)
}

func (o *FloatOperation) Mul() string {
    return o.op.Mul(o.left, o.right).Text(o.mode, o.prec)
}

func (o *FloatOperation) Div() string {
    return o.op.Quo(o.left, o.right).Text(o.mode, o.prec)
}

func (o *FloatOperation) Sub() string {
    return o.op.Sub(o.left, o.right).Text(o.mode, o.prec)
}

func NewFloatOpr(a, b string) *FloatOperation {
    op := new(big.Float)
    op1 := new(big.Float)
    op_l, _ := op.SetString(a)
    op_r, _ := op1.SetString(b)
    return &FloatOperation{
        op:                 op,
        left:               op_l,
        right:              op_r,
        prec:               int(10),
        mode:               'e',
    }
}

func isValidDig(arg string) bool {
    for i := 0; i < len(arg); i++ {
        if arg[i] < 48 || arg[i] > 57 {
            return false
        }
    }
    return true
}

func isFloat(arg string) bool {
    parts := strings.Split(arg, string(dot))
    if len(parts) > 1 {
        num := strings.Join(parts, "")
        return isValidDig(num)
    }
    return false
}

func isInt(arg string) bool {
    idx := strings.IndexByte(arg, dot)
    if idx < 0 {
        return isValidDig(arg)
    }
    return false
}

func ReadCLI() args {
    var a args
    if len(os.Args) < 4 {
        log.Fatal("Not enough args...")
    }
    return append(a, os.Args[1:]...)
}

func Exec(code string, op Operation) (string, error) {
    var result string
    var err error
    switch code {
    case add:
        result = op.Add()
    case div:
        result = op.Div()
    case sub:
        result = op.Sub()
    case mul:
        result = op.Mul()
    default:
        result = ""
        err = fmt.Errorf("Unexpected opcode: %s", code)
    }
    return result, err
}

func main() {
    var op Operation
    a := ReadCLI()
    if isFloat(a[0]) && isFloat(a[2]) {
        op = NewFloatOpr(a[0], a[2])
    }
    if isInt(a[0]) && isInt(a[2]) {
        op = NewIntOpr(a[0], a[2], base)
    }
    if op == nil {
        fmt.Printf("Invalid operands: %s | %s\n", a[0], a[2])
        os.Exit(1)
    }
    res, err := Exec(a[1], op)
    if err != nil {
        fmt.Printf("ERROR: %s\n", err.Error())
        os.Exit(1)
    }
    fmt.Println(res)
}

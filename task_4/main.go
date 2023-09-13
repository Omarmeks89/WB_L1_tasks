package main

import (
    "os"
    "os/signal"
    "context"
    "fmt"
    "math/rand"
    "time"
    "strings"
    "strconv"
    "log"
)

const (
    symbs string = "zaqxswcdevfrbgtnhymjukilop"
)

type (
    Poolsize int
    Token int
)

func RandString(strlen int) string {
    rand.Seed(time.Now().UTC().UnixNano())
    var randStr string
    res := make([]string, strlen)
    for i := 0; i < strlen; i++ {
        s := symbs[rand.Intn(len(symbs) - 1)]
        res[i] = string(s)
    }
    return strings.Join(res, randStr)
}

func ReadCLI() Poolsize {
    var cnt int
    var err error
    arg := os.Args[1]
    if cnt, err = strconv.Atoi(arg); err != nil {
        log.Fatal("Invalid arg")
    }
    return Poolsize(cnt)
}

func NewPool(size Poolsize) (func(c *context.Context, ch <-chan string)) {
    tokens := make(chan Token, size)
    for i := 0; i < int(size); i++ {
        tokens<- Token(i)
    }

    worker := func(ctx *context.Context, t Token, d string) {
        tok := tokens
        select {
        case <-(*ctx).Done():
            return
        default:
            fmt.Printf("Worker [%d] data: [%s]\n", t, d)
            select {
            case tok<- t:
            default:
                return
            }
        }
    }

    return func(ctx *context.Context, c <-chan string) {
        for {
            select {
            case data := <-c:
                select {
                case t := <- tokens:
                    go worker(ctx, t, data)
                case <-(*ctx).Done():
                    return
                }
            case <-(*ctx).Done():
                return
            }
        }
    }
}

func on_shutdown(f func()) {
    f()
    fmt.Println("Done...")
}

func main() {

    dataChan := make(chan string)
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
    defer on_shutdown(stop)

    PoolSize := ReadCLI()
    Run := NewPool(PoolSize)
    go Run(&ctx, dataChan)

    fmt.Println("Run...")
    for {
        data := RandString(6)
        select {
        case <-ctx.Done():
            fmt.Println("SIGINT received...")
            return
        case dataChan<- data:
        }
    }
}

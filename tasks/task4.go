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

// функция чтения данных из командной строки
func ReadCLI() Poolsize {
    var cnt int
    var err error
    arg := os.Args[1]
    if cnt, err = strconv.Atoi(arg); err != nil {
        log.Fatal("Invalid arg")
    }
    return Poolsize(cnt)
}

// В качестве способа завершения работы всех воркеров
// был выбран context как инструмент специально созданный
// для таких случаев, расширяемый при необходимости (отмена по 
// таймауту, дедлайну), а также (для данной ситуации) поддержива
// ющий перехват сигналов.
func NewPool(size Poolsize) (func(c *context.Context, ch <-chan string)) {
    // Используя свойство буферизованного канала
    // создаем пул горутин нужного размера.
    // Пока в пуле есть токены (т.е. канал не пустой)
    // горутина может взять токен и начать работу.
    // По завершении работы горутина возвращает токен
    // в пул, чтобы ожидающие горутины (если их больше
    // чем размер пула), могли начать работу.
    // При попытке чтения из пустого канала, горутина
    // паркуется и ждет появления свободного токена.
    tokens := make(chan Token, size)
    for i := 0; i < int(size); i++ {
        tokens<- Token(i)
    }

    worker := func(ctx *context.Context, t Token, d string) {
        tok := tokens
        // Работаем с каналами с помощью 
        // select -> он читает состояние
        // каналов и выполняет незаблокированную ветвь.
        // В данном случае было выбрано неблокирующее
        // поведение с веткой default
        select {
        case <-(*ctx).Done():
            return
        default:
            fmt.Printf("Worker [%d] data: [%s]\n", t, d)
            select {
            // Воркер возвращает выданный токен в пул
            case tok<- t:
            default:
                return
            }
        }
    }

    // горутина запускающая воркеров и выдающая им токены
    return func(ctx *context.Context, c <-chan string) {
        for {
            select {
            case data := <-c:
                // используем вложенный select
                // чтобы имеь возможность обработать
                // ситуацию завершения внутри
                // объемлющего блока select.
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
    // Создаем отложенный вызов 
    // функции cancel для отмены горутин 
    // чтобы была возможность завершения
    // если SIGINT не будет получен
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

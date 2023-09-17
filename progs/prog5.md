## Что выведет программа:

```go
func main() {
    slice := []string{"a", "a"}
    func(slice []string) {
        slice = append(slice, "a")
        slice[0] = "b"
        slice[1] = "b"
        fmt.Print(slice)
    }(slice)
    fmt.Print(slice)
}
```

Результат:

```bash
[b, b, a][a, a]
```

В строке `slice = append(slice, "a")` мы создаем новый слайс
внутри анонимной функции, так как добавление нового элемента
превышает размеры текущего слайса.(создается новый массив и слайс с указателем на новый массив). Последующие операции
изменяют значение нового локального слайса анонимной функции.
Исходный слайс не меняется.

Чтобы менялся исходный слайс, нужно передать его по указателю,
либо вернуть из функции.

```go
func main() {
    slice := []string{"a", "a"}
    func(slice *[]string) {
        *slice = append(*slice, "a")
        (*slice)[0] = "b"
        (*slice)[1] = "b"
        fmt.Print(&slice)
    }(&slice)
    fmt.Print(slice)
}
```

Результат:

```bash
[b, b, a][b, b, a]
```

## Что выведет программа:

```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```

Будет дедлок, так как при передаче WaitGroup по значению
она копируется и при копировании обнуляет свое состояние (счетчик).
Программа выведет:

```bash
1
0
3
4
2
```

после чего упадет с дедлоком.
WaitGroup нужно передавать по указателю, чтобы
не копировать его:

```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```

Программа отработает корректно и выведет:
(числа от 0 до 4 в рандомном порядке)
```bash
2
4
0
3
1
exit
```



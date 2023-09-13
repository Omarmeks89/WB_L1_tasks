## Задание №7

Для конкурентной записи был создан буферизованный
канал на 4 токена: A, B, C, D.

Было создано 8 горутин (по 2 на каждый токен) и запущены параллельно.
Каждая из горутин брала токен из канала, записывала по нему сообщение в map,
в формате: <номер воркера> <имя токена>.
После чего возвращала токен обратно. Каждая горутина должна была выполнить 10
операций записи в map.

Общее количество операций = 80.

Ниже прилагаю вывод программы:

### Вывод программы:

```bash
WRK [7] done [10] operations...
WRK [3] done [10] operations...
WRK [2] done [10] operations...
WRK [1] done [10] operations...
WRK [0] done [10] operations...
WRK [6] done [10] operations...
WRK [4] done [10] operations...
WRK [5] done [10] operations...
For KEY [A] found:
	> MSG [1]: [WRK [7] KEY [A]]
	> MSG [2]: [WRK [7] KEY [A]]
	> MSG [3]: [WRK [7] KEY [A]]
	> MSG [4]: [WRK [3] KEY [A]]
	> MSG [5]: [WRK [3] KEY [A]]
	> MSG [6]: [WRK [3] KEY [A]]
	> MSG [7]: [WRK [0] KEY [A]]
	> MSG [8]: [WRK [4] KEY [A]]
	> MSG [9]: [WRK [0] KEY [A]]
	> MSG [10]: [WRK [4] KEY [A]]
	> MSG [11]: [WRK [1] KEY [A]]
	> MSG [12]: [WRK [4] KEY [A]]
	> MSG [13]: [WRK [4] KEY [A]]
	> MSG [14]: [WRK [5] KEY [A]]

For KEY [B] found:
	> MSG [1]: [WRK [7] KEY [B]]
	> MSG [2]: [WRK [7] KEY [B]]
	> MSG [3]: [WRK [7] KEY [B]]
	> MSG [4]: [WRK [3] KEY [B]]
	> MSG [5]: [WRK [3] KEY [B]]
	> MSG [6]: [WRK [3] KEY [B]]
	> MSG [7]: [WRK [4] KEY [B]]
	> MSG [8]: [WRK [4] KEY [B]]
	> MSG [9]: [WRK [0] KEY [B]]
	> MSG [10]: [WRK [0] KEY [B]]
	> MSG [11]: [WRK [0] KEY [B]]
	> MSG [12]: [WRK [0] KEY [B]]
	> MSG [13]: [WRK [0] KEY [B]]
	> MSG [14]: [WRK [0] KEY [B]]
	> MSG [15]: [WRK [0] KEY [B]]
	> MSG [16]: [WRK [5] KEY [B]]
	> MSG [17]: [WRK [5] KEY [B]]
	> MSG [18]: [WRK [5] KEY [B]]
	> MSG [19]: [WRK [4] KEY [B]]
	> MSG [20]: [WRK [5] KEY [B]]

For KEY [C] found:
	> MSG [1]: [WRK [7] KEY [C]]
	> MSG [2]: [WRK [7] KEY [C]]
	> MSG [3]: [WRK [2] KEY [C]]
	> MSG [4]: [WRK [6] KEY [C]]
	> MSG [5]: [WRK [1] KEY [C]]
	> MSG [6]: [WRK [2] KEY [C]]
	> MSG [7]: [WRK [6] KEY [C]]
	> MSG [8]: [WRK [1] KEY [C]]
	> MSG [9]: [WRK [2] KEY [C]]
	> MSG [10]: [WRK [6] KEY [C]]
	> MSG [11]: [WRK [1] KEY [C]]
	> MSG [12]: [WRK [2] KEY [C]]
	> MSG [13]: [WRK [6] KEY [C]]
	> MSG [14]: [WRK [1] KEY [C]]
	> MSG [15]: [WRK [2] KEY [C]]
	> MSG [16]: [WRK [6] KEY [C]]
	> MSG [17]: [WRK [1] KEY [C]]
	> MSG [18]: [WRK [2] KEY [C]]
	> MSG [19]: [WRK [6] KEY [C]]
	> MSG [20]: [WRK [1] KEY [C]]
	> MSG [21]: [WRK [2] KEY [C]]
	> MSG [22]: [WRK [6] KEY [C]]
	> MSG [23]: [WRK [1] KEY [C]]
	> MSG [24]: [WRK [2] KEY [C]]
	> MSG [25]: [WRK [6] KEY [C]]
	> MSG [26]: [WRK [1] KEY [C]]
	> MSG [27]: [WRK [2] KEY [C]]
	> MSG [28]: [WRK [6] KEY [C]]
	> MSG [29]: [WRK [1] KEY [C]]
	> MSG [30]: [WRK [2] KEY [C]]
	> MSG [31]: [WRK [6] KEY [C]]
	> MSG [32]: [WRK [5] KEY [C]]
	> MSG [33]: [WRK [4] KEY [C]]
	> MSG [34]: [WRK [4] KEY [C]]

For KEY [D] found:
	> MSG [1]: [WRK [7] KEY [D]]
	> MSG [2]: [WRK [7] KEY [D]]
	> MSG [3]: [WRK [3] KEY [D]]
	> MSG [4]: [WRK [3] KEY [D]]
	> MSG [5]: [WRK [3] KEY [D]]
	> MSG [6]: [WRK [3] KEY [D]]
	> MSG [7]: [WRK [0] KEY [D]]
	> MSG [8]: [WRK [4] KEY [D]]
	> MSG [9]: [WRK [5] KEY [D]]
	> MSG [10]: [WRK [5] KEY [D]]
	> MSG [11]: [WRK [5] KEY [D]]
	> MSG [12]: [WRK [5] KEY [D]]
```

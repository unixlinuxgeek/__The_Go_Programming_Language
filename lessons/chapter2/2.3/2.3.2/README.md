### 2.3.2 Указатели

> Переменная это участок памяти, содержащее значение.

> & это оператор получения адреса.

> Указатель это местоположение в памяти, где хранится значение.

Значение указателя это адрес переменной.
Все переменные имеют адрес.

С помощью указателя можно считывать или изменять значение переменной
косвенно, не используя ее имя, если оно у нее есть.

Если переменная объявлена как ```var x int```, выражение &x (адрес x)
дает указатель на целочисленную переменную,
т.е значение типа *int, который произносится как "указатель на int".

Если это значение называется ```p```, мы говорим ```p указывает на x``` 
или  ```p содержит адрес x```.

Переменная на которую указывает ```p``` пишется как *p.
Выражение ```*p```  дает значение этой переменной, int.

Пример:
```go
x := 1          
p := &x         // p имеет тип *int и указывает на x
fmt.Println(*p) // "1"
*p = 2          // Эквивалентно присваиванию x = 2 
fmt.Println(x)  // "2"
```

> Каждый компонент переменной составного типа, 
поле структуры или элемент массива также является переменной, 
а значит имеет свой адрес.  

> Нулевое значение указателя любого типа равно ```nil```. 

> Проверка ```p != nil ``` true если p указывает на переменную.

> Указатели можно сравнивать:
Два указателя равны когда они указывают на одну и туже переменную,
или когда они оба равны nil.
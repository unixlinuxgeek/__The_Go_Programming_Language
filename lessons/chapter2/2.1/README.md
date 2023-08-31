### 2.1 Имена

Go имеет 25 ключевых слов и они не могут использоваться в качестве имен.

```
break     default          func      interface     select
case      defer            go        map           struct
chan      else             goto      package       switch
const     fallthrough      if        range         type
continue  for              import    return        var
```

Предопределенные имена:
Они не являются зарезервированными.

Константы ```true false iota nil```

Типы:     ```int int8 int16 int32 int64```
          ```uint uint8 uint16 uint32 uint64 uintptr```
          ```float32 float64 complex128 complex64```
          ```bool byte rune string error```

Функции:  ```make len cap new append copy close delete```
          ```complex real imag panic recover```

Если сущьность объявленпа вне функции, она видна для всех файлов пакета,
к которому она пренадлежит.

Регистр первой буквы имени пакета определяет его видимость, за пределами пакета.
Есл она начинается с прописной буквы оно экспортируемое,
она видна за пределами собственного пакета, и к нему могут обращаться другие части программы,
как например функция Printf из пакета fmt.

Соглашения Go:
Чем больше область видимости переменной, тем длинее должно быть имя переменной.
И чем меньше область видимости переменной, тем короче имя переменной.



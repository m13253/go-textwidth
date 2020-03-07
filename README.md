# go-textwidth

Go library to determine the required columns and rows for a Unicode string to
display in a terminal

## Usage

```go
package main

import (
    "fmt"
    textwidth "github.com/m13253/go-textwidth"
)

func main() {
    const s = "Go（又称 Golang）是 Google 开发的一种静态强类型、编译型、并发"+
        "型，并具有垃圾回收功能的编程语言。Go 的主要特色在于易于使用的并行设"+
        "计，叫做 Goroutine，透过 Goroutine 能够让程序以异步的方式运行，而不需"+
        "要担心一个函数导致程序中断，因此 Go 也非常地适合网络服务。\t————维基"+
        "百科 Wikipedia"
    row, col := textwidth.GetTextOffset(s, 0, 80)
    total := textwidth.GetTextWidth(s, 0, 80)

    fmt.Printf("rows = %d, columns = %d, total = %d", row, col, total)
}
```

The paragraph takes 3 lines and 46 more columns. 3×80+44 = 286.
```
<--------10--------20--------30--------40--------50--------60--------70-------->
Go（又称 Golang）是 Google 开发的一种静态强类型、编译型、并发型，并具有垃圾回收
功能的编程语言。Go 的主要特色在于易于使用的并行设计，叫做 Goroutine，透过 Gorout
ine 能够让程序以异步的方式运行，而不需要担心一个函数导致程序中断，因此 Go 也非常
地适合网络服务。        ————维基百科 Wikipedia
```

Therefore, the result is:
```
rows = 3, columns = 44, total = 284
```

## License

go-textwidth is released under MIT license. Please see [LICENSE](LICENSE) for
more imformation.

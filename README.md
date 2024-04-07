# EarthMC-Go
Unofficial wrapper for the **EarthMC** Dynmap and Official APIs in Golang.<br>

# Installation
Enter the following line into your project's terminal.

```console
go get github.com/earthmc-toolkit/emcgo
```

# Usage
```go
import (
    emc "github.com/earthmc-toolkit/emcgo"
)

func main() {
    towns, err := emc.Aurora.Towns.All()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(towns)
}
```

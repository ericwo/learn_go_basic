package main

import "fmt"

func main()  {
    m := make(map[string]int)
    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    elem, ok := m["Answer"]
    fmt.Println("The value:", elem, "Present?", ok)

    delete(m, "Answer")

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}

package main

import (
    "container/list"
    "fmt"
)

func main() {
    var intList list.List
    intList.PushBack(11)
    intList.PushBack(23)
    intList.PushBack(34)

    for element := intList.Front(); element != nil; element = element.Next() {
        fmt.Print(element.Value.(int))
        fmt.Print(" ")
    }
}

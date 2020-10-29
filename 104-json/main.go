package main

import (
    "encoding/json"
    "fmt"
)

type Member struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Active bool   `json:"active"`
}

func main() {
    //marshal()
    unmarchal()
}

func marshal() {
    mem := Member{"Alex", 10, true}

    jsonBytes, err := json.Marshal(mem)
    if err != nil {
        panic(err)
    }

    jsonString := string(jsonBytes)

    fmt.Println(jsonString)
}

func unmarchal() {
    jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})
    
    var mem Member
    if err := json.Unmarshal(jsonBytes, &mem); err != nil {
        panic(err)
    }
    
    fmt.Println(mem.Name, mem.Age, mem.Active)
}
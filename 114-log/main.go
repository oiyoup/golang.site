package main

import (
    "log"
    "os"
)

var myLogger *log.Logger

func main() {
    log.Println("Logging")
    
    myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
    
    run()
    
    myLogger.Println("End of Program")
}

func run() {
    myLogger.Print("Test")
}

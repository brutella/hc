package main

import(
    "github.com/brutella/log"
)

func main() {    
    log.Println("[ERRO] Just an error")
    log.Printf("[ERRO] Just an error %d", 1)
    log.Println("[WARN] Just a warning")
    log.Printf("[WARN] Just a warning %d", 1)
    log.Println("[INFO] Just an info")
    log.Printf("[INFO] Just an info %d", 1)
    log.Println("[VERB] Just an info")
    log.Printf("[VERB] Just an info %d", 1)
    
    log.Error = false
    log.Warn  = false
    log.Info  = false
    log.Verbose = false
    
    log.Println("MUST be displayed")
    log.Println("[ERRO] MUST NOT be displayed")
    log.Printf("[ERRO] MUST NOT be displayed %d", 1)
    log.Println("[WARN] MUST NOT be displayed")
    log.Printf("[WARN] MUST NOT be displayed %d", 1)
    log.Println("[INFO] MUST NOT be displayed")
    log.Printf("[INFO] MUST NOT be displayed %d", 1)
    log.Println("[VERB] MUST NOT be displayed")
    log.Printf("[VERB] MUST NOT be displayed %d", 1)
}
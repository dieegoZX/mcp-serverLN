package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "MCP Server online e funcional! 🔥")
}

func main() {
    http.HandleFunc("/", handler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000" // fallback local
    }

    fmt.Println("Servidor rodando na porta:", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

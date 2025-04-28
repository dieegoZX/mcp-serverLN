package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var authorizedUsers = map[string]bool{
	"Admin": true,
}

type CommandRequest struct {
	User    string `json:"user"`
	Command string `json:"command"`
}

type CommandResponse struct {
	Response string `json:"response"`
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Access Denied. MCP accepts only POST requests.", http.StatusMethodNotAllowed)
		return
	}

	var cmdReq CommandRequest
	err := json.NewDecoder(r.Body).Decode(&cmdReq)
	if err != nil {
		http.Error(w, "Invalid request payload.", http.StatusBadRequest)
		return
	}

	if !authorizedUsers[cmdReq.User] {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(CommandResponse{Response: "Access Denied. You are not authorized by MCP."})
		log.Printf("[MCP] Unauthorized access attempt by: %s\n", cmdReq.User)
		return
	}

	response := "Command '" + cmdReq.Command + "' executed successfully by " + cmdReq.User + "."
	log.Printf("[MCP] Command executed: %s by user: %s\n", cmdReq.Command, cmdReq.User)
	json.NewEncoder(w).Encode(CommandResponse{Response: response})
}

func main() {
	http.HandleFunc("/command", commandHandler)

	log.Println("🧠 MCP Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

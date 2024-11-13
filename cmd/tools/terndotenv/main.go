package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	args := []string{
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	}
	cmd := exec.Command("tern", args...)
	
	if err := cmd.Run(); err != nil {
		log.Printf("Error running tern command: %v", err)
		os.Exit(1)
		}
	
}
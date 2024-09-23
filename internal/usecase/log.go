package usecase

import (
	"log"
	"os"
)

func logIt(msg string) {
	// Abra ou crie o arquivo de log
	file, err := os.OpenFile("RPSUtilities.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %v", err)
	}
	defer file.Close()

	// Configure o logger para escrever no arquivo
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime)

	// Exemplo de registros de log
	log.Println(msg)

	log.SetOutput(os.Stdout)
}

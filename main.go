package main

import (
	client "github.com/pedroteixeirabisognin/swapi-go/client"
	"github.com/pedroteixeirabisognin/swapi-go/database"
	"github.com/pedroteixeirabisognin/swapi-go/server"
)

//Implementar cache com redis
//Implementar banco de dados
//Criar serviço
//Fazer testes unitários
//Fazer testes integrados
//Dockerizar
//Mensageria
func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
	client.GetPlanets()
}

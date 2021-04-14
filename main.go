package main

import (
	"landlords/agentservice"
	"landlords/config"
	"landlords/initcards"
	"landlords/log"
	"landlords/manager"
	"landlords/signal"
)

func main() {
	log.InitLog()

	config.InitConfig()

	//redis.InitRedis()

	signal.InitSignal()

	initcards.InitNewCards()

	manager.InitRoomManager()

	manager.InitPvpPoolManager()

	//websocket.Run()
	agentservice.AgentRun()
}

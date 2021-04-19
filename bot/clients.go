package bot

import (
	"fmt"
	"github.com/NazarNintendo/cryptobot/model"
	"github.com/NazarNintendo/cryptobot/utils"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func getClient(user *tb.User) *model.Client {
	client, exists := activeClients[user.ID]
	if !exists {
		activeClients[user.ID] = &model.Client{User: user}
		client = activeClients[user.ID]
	}
	return client
}

func addMonitoring(user *tb.User, monitoringConfig string) {
	client := getClient(user)
	crypto := utils.ParseCryptoId(strings.Split(monitoringConfig, "|")[1])
	frequency := utils.ParseFrequencyId(strings.Split(monitoringConfig, "|")[0])
	client.AddMonitoring(crypto, frequency)
	for _, client := range activeClients {
		fmt.Println(client.String())
	}
}

func deleteMonitorings(user *tb.User) {
	client := getClient(user)
	client.Monitorings = []model.Monitoring{}
	for _, client := range activeClients {
		fmt.Println(client.String())
	}
}

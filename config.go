package main

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"io/ioutil"
)

type Config struct {
	ListenAddress    string          `json:"listen_address"`
	MaxPlayers       int             `json:"max_players"`
	Motd             string          `json:"motd"`
	Restricted       bool            `json:"restricted"`
	Logs             bool            `json:"logs"`
	JoinMessage      json.RawMessage `json:"join_message"`
	BossBar          json.RawMessage `json:"boss_bar"`
	PlayerListHeader json.RawMessage `json:"playerlist_header"`
	PlayerListFooter json.RawMessage `json:"playerlist_footer"`
}

var (
	config         Config
	join_message   PacketPlayMessage
	bossbar_create PacketBossBar
	playerlist_hf  PacketPlayerListHeaderFooter
)

func InitConfig() (err error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	if config.JoinMessage != nil {
		join_message = PacketPlayMessage{
			string(config.JoinMessage),
			CHAT_BOX,
		}
	}

	if config.BossBar != nil {
		bossbar_create = PacketBossBar{
			uuid: uuid.Must(uuid.NewV4()),
			action:   BOSSBAR_ADD,
			title:    string(config.BossBar),
			health:   1.0,
			color:    BOSSBAR_COLOR_RED,
			division: BOSSBAR_NODIVISION,
			flags:    0,
		}
	}

	playerlist_hf = PacketPlayerListHeaderFooter{}
	if config.PlayerListHeader != nil {
		msg := string(config.PlayerListHeader)
		playerlist_hf.header = &msg
	}
	if config.PlayerListFooter != nil {
		msg := string(config.PlayerListFooter)
		playerlist_hf.footer = &msg
	}
	return
}

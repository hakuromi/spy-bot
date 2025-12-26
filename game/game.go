package game

import (
	"errors"

	"github.com/hakuromi/spy-bot/models"
)

type Manage struct {
	Game models.Game
}

func (m *Manage) NewGame() { // создание игры
	m.Game = models.Game{
		Players: []models.Player{},
		Active:  false,
	}
}

func (m *Manage) AddPlayer(p models.Player) error { // добавление пользователя
	if m.Game.Active {
		return errors.New("Игра уже началась.")
	}

	for _, player := range m.Game.Players {
		if player.ID == p.ID {
			return errors.New("игрок уже в игре")
		}
	}
	m.Game.Players = append(m.Game.Players, p)

	return nil
}

func (m *Manage) CanStart() bool { // проверка количества игроков
	return len(m.Game.Players) >= 3
}

func (m *Manage) Start() error {
	if !m.CanStart() {
		return errors.New("нужно минимум 3 игрока")
	}
	m.Game.Active = true
	return nil
}

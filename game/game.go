package game

import (
	"errors"
	"math/rand"

	"github.com/hakuromi/spy-bot/models"
)

type Manager struct {
	Game models.Game
}

func (m *Manager) NewGame() { // создание игры
	m.Game = models.Game{
		Players: []models.Player{},
		Active:  false,
	}
}

func (m *Manager) AddPlayer(p models.Player) error { // добавление пользователя
	if m.Game.Active {
		return errors.New("Игра уже началась.")
	}

	for _, player := range m.Game.Players {
		if player.ID == p.ID {
			return errors.New("Игрок уже в игре.")
		}
	}
	m.Game.Players = append(m.Game.Players, p)

	return nil
}

func (m *Manager) CanStart() bool { // проверка количества игроков
	return len(m.Game.Players) >= 3
}

func (m *Manager) chooseSpy() {
	num := rand.Intn(len(m.Game.Players))
	m.Game.SpyID = m.Game.Players[num].ID
}

func (m *Manager) chooseHero() {
	num := rand.Intn(len(models.Heroes))
	m.Game.Hero = models.Heroes[num]
}

func (m *Manager) Start() error {
	if !m.CanStart() {
		return errors.New("нужно минимум 3 игрока")
	}
	m.chooseSpy()
	m.chooseHero()
	m.Game.Active = true
	return nil
}

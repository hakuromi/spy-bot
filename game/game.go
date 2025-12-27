package game

import (
	"errors"
	"fmt"
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
		return errors.New("игра уже началась.")
	}

	for _, player := range m.Game.Players {
		if player.ID == p.ID {
			return errors.New("игрок уже в игре.")
		}
	}
	m.Game.Players = append(m.Game.Players, p)

	return nil
}

func (m *Manager) CanStart() bool { // проверка количества игроков
	return len(m.Game.Players) >= 3
}

func (m *Manager) chooseSpy() { // выбор шпиона рандомно
	num := rand.Intn(len(m.Game.Players))
	m.Game.SpyID = m.Game.Players[num].ID
}

func (m *Manager) chooseHero() { // выбор героя рандомно
	if m.Game.Mode == "clash" {
		num := rand.Intn(len(models.Royale))
		m.Game.Hero = models.Royale[num]
	} else if m.Game.Mode == "dota" {
		num := rand.Intn(len(models.Heroes))
		m.Game.Hero = models.Heroes[num]
	}
}

func (m *Manager) Start() error { //
	if !m.CanStart() {
		return errors.New("нужно минимум 3 игрока")
	}
	m.chooseSpy()
	m.chooseHero()
	m.Game.Active = true
	return nil
}

func (m *Manager) GetRoles() map[int64]string {
	roles := make(map[int64]string)
	for _, player := range m.Game.Players {
		if player.ID == m.Game.SpyID {
			roles[player.ID] = "вы шпион"
		} else {
			roles[player.ID] = "вы не шпион. персонаж: " + m.Game.Hero
		}
	}
	return roles
}

func (m *Manager) End() {
	m.Game.Players = []models.Player{} // удаляем всех игроков
	m.Game.Active = false              // игра неактивна
	m.Game.Mode = ""                   // режим не выбран
	m.Game.SpyID = 0                   // шпион очищен
	m.Game.Hero = ""                   //герой очищен
}

func PlayerList(players []models.Player) string {
	if len(players) == 0 {
		return "никто еще не присоединился."
	}
	list := "игроки:\n"
	for _, p := range players {
		list += fmt.Sprintf("@%s\n", p.Name)
	}
	return list
}

package manager

import "click-game/internal/storage"

type Manager struct {
	storage storage.Storage
}

func NewManager(s storage.Storage) *Manager {
	return &Manager{
		storage: s,
	}
}

func (m *Manager) Run() {
	// zapusk bota
}

func (m *Manager) handleBotUpdates() {
	for {
		//m.storage.GetUser()
	}
}

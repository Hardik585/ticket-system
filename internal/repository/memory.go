package repository

import (
	"errors"
	"sync"
	"ticket-system/internal/models"
	"time"
)

type MemoryStore struct {
	mu           sync.RWMutex
	users        map[string]models.User 
	tickets      map[int]models.Ticket   
	nextUserID   int
	nextTicketID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:        make(map[string]models.User),
		tickets:      make(map[int]models.Ticket),
		nextUserID:   1,
		nextTicketID: 1,
	}
}

func (s *MemoryStore) CreateUser(username, passwordHash string) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[username]; exists {
		return models.User{}, errors.New("username already exists")
	}
	user := models.User{
		ID:           s.nextUserID,
		Username:     username,
		PasswordHash: passwordHash,
	}
	s.users[username] = user
	s.nextUserID++
	return user, nil
}

func (s *MemoryStore) GetUserByUsername(username string) (models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, exists := s.users[username]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *MemoryStore) CreateTicket(title, description string, ownerID int) models.Ticket {
	s.mu.Lock()
	defer s.mu.Unlock()
	ticket := models.Ticket{
		ID:          s.nextTicketID,
		Title:       title,
		Description: description,
		Status:      "open", 
		OwnerID:     ownerID,
		CreatedAt:   time.Now(),
	}
	s.tickets[s.nextTicketID] = ticket
	s.nextTicketID++
	return ticket
}

func (s *MemoryStore) GetTicketsByOwner(ownerID int) []models.Ticket {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var userTickets []models.Ticket
	for _, ticket := range s.tickets {
		if ticket.OwnerID == ownerID {
			userTickets = append(userTickets, ticket)
		}
	}
	return userTickets
}

func (s *MemoryStore) GetTicketByID(id int) (models.Ticket, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ticket, exists := s.tickets[id]
	if !exists {
		return models.Ticket{}, errors.New("ticket not found")
	}
	return ticket, nil
}

func (s *MemoryStore) UpdateTicketStatus(id int, status string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	ticket, exists := s.tickets[id]
	if !exists {
		return errors.New("ticket not found")
	}
	ticket.Status = status
	s.tickets[id] = ticket
	return nil
}
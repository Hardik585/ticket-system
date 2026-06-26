package ticket

import (
	"errors"
	"ticket-system/internal/models"
	"ticket-system/internal/repository"
)

type Service struct {
	store *repository.MemoryStore
}

func NewService(store *repository.MemoryStore) *Service {
	return &Service{store: store}
}

func (s *Service) Create(title, description string, ownerID int) models.Ticket {
	return s.store.CreateTicket(title, description, ownerID)
}

func (s *Service) ListByOwner(ownerID int) []models.Ticket {
	return s.store.GetTicketsByOwner(ownerID)
}

func (s *Service) GetForOwner(id, ownerID int) (models.Ticket, error) {
	ticket, err := s.store.GetTicketByID(id)
	if err != nil {
		return models.Ticket{}, errors.New("ticket not found")
	}
	if ticket.OwnerID != ownerID {
		return models.Ticket{}, errors.New("unauthorized access to ticket")
	}
	return ticket, nil
}

func (s *Service) UpdateStatus(id, ownerID int, newStatus string) (models.Ticket, error) {
	if newStatus != "open" && newStatus != "in_progress" && newStatus != "closed" {
		return models.Ticket{}, errors.New("invalid status value")
	}
	ticket, err := s.store.GetTicketByID(id)
	if err != nil {
		return models.Ticket{}, errors.New("ticket not found")
	}
	if ticket.OwnerID != ownerID {
		return models.Ticket{}, errors.New("unauthorized action")
	}
	if ticket.Status == "closed" {
		return models.Ticket{}, errors.New("cannot update a closed ticket")
	}
	err = s.store.UpdateTicketStatus(id, newStatus)
	if err != nil {
		return models.Ticket{}, err
	}
	ticket.Status = newStatus
	return ticket, nil
}
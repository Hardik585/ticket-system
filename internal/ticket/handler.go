package ticket

import (
	"net/http"
	"strconv"
    "ticket-system/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

type createRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type updateStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// POST /tickets
func (h *Handler) Create(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket := h.service.Create(req.Title, req.Description, userID.(int))
	c.JSON(http.StatusCreated, ticket)
}

// GET /tickets
func (h *Handler) List(c *gin.Context) {
	userID, _ := c.Get("user_id")
	tickets := h.service.ListByOwner(userID.(int))
	if tickets == nil {
		tickets = []models.Ticket{}
	}
	c.JSON(http.StatusOK, tickets)
}

// GET /tickets/{id}
func (h *Handler) GetByID(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ticket id"})
		return
	}
	ticket, err := h.service.GetForOwner(ticketID, userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// PATCH /tickets/{id}/status
func (h *Handler) UpdateStatus(c *gin.Context) {
	userID, _ := c.Get("user_id")
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ticket id"})
		return
	}
	var req updateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket, err := h.service.UpdateStatus(ticketID, userID.(int), req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}
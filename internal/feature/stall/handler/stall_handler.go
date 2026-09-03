package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/stall/dto"
	"lapakita-backend/internal/feature/stall/usecase"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StallHandler struct {
	usecase *usecase.StallUsecase
}

func NewStallHandler(usecase *usecase.StallUsecase) *StallHandler {
	return &StallHandler{usecase: usecase}
}

// GET /api/v1/stalls (Search with Filters)
func (h *StallHandler) Search(c *gin.Context) {
	var req dto.SearchStallRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	stalls, meta, err := h.usecase.Search(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyStallFailedToGet))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyStallSearchSuccess), stalls, meta)
}

// GET /api/v1/stalls/:id
func (h *StallHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	stall, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyStallNotFound))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyStallGetSuccess), stall)
}

// POST /api/v1/stalls (Protected)
func (h *StallHandler) Create(c *gin.Context) {
	ownerID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.CreateStallRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	stall, err := h.usecase.Create(c.Request.Context(), ownerID, req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyStallFailedToCreate))
		return
	}

	api.Success(c, http.StatusCreated, i18n.T(c, i18n.KeyStallCreateSuccess), stall)
}

// PUT /api/v1/stalls/:id (Protected)
func (h *StallHandler) Update(c *gin.Context) {
	ownerID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.UpdateStallRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	req.ID = c.Param("id")

	stall, err := h.usecase.Update(c.Request.Context(), ownerID, req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyStallFailedToUpdate))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyStallUpdateSuccess), stall)
}

// DELETE /api/v1/stalls/:id (Protected)
func (h *StallHandler) Delete(c *gin.Context) {
	ownerID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	if err := h.usecase.Delete(c.Request.Context(), ownerID, id); err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyStallFailedToDelete))
		return
	}

	api.Success[any](c, http.StatusOK, i18n.T(c, i18n.KeyStallDeleteSuccess))
}

// GET /api/v1/stalls/my-stalls
func (h *StallHandler) GetByOwner(c *gin.Context) {
	ownerID, ok := middleware.GetUserIDFromContext(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.GetOwnerStallsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	// Auto-fill owner_id dari JWT jika tidak dikirimkan di query param
	if req.OwnerID == "" {
		req.OwnerID = ownerID.String()
	}

	stalls, meta, err := h.usecase.GetByOwnerID(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyStallFailedToGet))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyStallSearchSuccess), stalls, meta)
}

// GET /api/v1/stalls/:id/similar
func (h *StallHandler) GetSimilar(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	req := dto.GetSimilarStallsRequest{ID: id}
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	stalls, meta, err := h.usecase.GetSimilar(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyStallNotFound))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyStallGetSimilarSuccess), stalls, meta)
}

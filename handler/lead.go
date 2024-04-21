package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	oteldemo "github.com/phbpx/otel-demo"
	"github.com/uptrace/opentelemetry-go-extra/otelutil"
	"go.uber.org/zap"
	"net/http"
)

type LeadHandler struct {
	service oteldemo.LeadService
	log     *zap.SugaredLogger
}

func NewLeadHandler(service oteldemo.LeadService, log *zap.SugaredLogger) *LeadHandler {
	return &LeadHandler{
		service: service,
		log:     log,
	}
}

func (lh LeadHandler) Create(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request oteldemo.LeadRequest

	if err := decode(r, &request); err != nil {
		lh.log.Errorw("GetByID", "error", err.Error())
		respondErr(ctx, rw, http.StatusBadRequest, err)
		return
	}

	lead := oteldemo.CreateLeadByParam(request)

	if err := lh.service.Create(r.Context(), lead); err != nil {
		lh.log.Errorw("Create", "error", err.Error())
		switch {
		case errors.Is(err, oteldemo.ErrDuplicatedLead):
			respondErr(ctx, rw, http.StatusConflict, err)
		default:
			respondErr(ctx, rw, http.StatusInternalServerError, err)
		}
		return
	}
	attribute := otelutil.Attribute("groupId", lead.Group)
	respond(ctx, http.StatusCreated, attribute)
	respondWriter(rw, http.StatusCreated, &lead)
}

func (lh LeadHandler) GetByID(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		lh.log.Errorw("GetByID", "error", err.Error())
		respondErr(ctx, rw, http.StatusBadRequest, errors.New("ID is not in its proper form"))
		return
	}

	lead, err := lh.service.GetByID(ctx, id.String())
	if err != nil {
		lh.log.Errorw("GetByID", "error", err.Error())
		switch err {
		case oteldemo.ErrLeadNotFound:
			respondErr(ctx, rw, http.StatusNotFound, err)
		default:
			respondErr(ctx, rw, http.StatusInternalServerError, err)
		}
		return
	}

	attribute := otelutil.Attribute("groupId", lead.Group)
	respond(ctx, http.StatusOK, attribute)
	respondWriter(rw, http.StatusOK, &lead)
}

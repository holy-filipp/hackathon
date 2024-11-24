package bymonth

import (
	"calendar-backend/internal/lib/logger/api/response"
	"calendar-backend/internal/lib/logger/sl"
	"calendar-backend/internal/storage/postgresql/ent"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Month int `json:"month" validate:"gte=1,lte=12"`
	Year  int `json:"year" validate:"gte=2024,lte=2024"`
}

type Response struct {
	response.Response
	Events []*ent.Events
}

type EventsFetcher interface {
	GetEventsByMonth(year int, month int) ([]*ent.Events, error)
}

func New(log *slog.Logger, fetcher EventsFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.events.bymonth.New"

		log = log.With(slog.String("op", op), slog.String("request_id", middleware.GetReqID(r.Context())))

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, response.Error("failed to decode request"))

			return
		}

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, response.ValidationError(validateErr))

			return
		}

		res, err := fetcher.GetEventsByMonth(req.Year, req.Month)
		if err != nil {
			log.Error("failed to fetch from db", sl.Err(err))

			render.JSON(w, r, response.Error("failed to fetch from db"))

			return
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
			Events:   res,
		})
	}
}

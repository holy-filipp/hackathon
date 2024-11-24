package parse

import (
	"calendar-backend/internal/config"
	"calendar-backend/internal/lib/logger/api/response"
	"calendar-backend/internal/lib/logger/sl"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ledongthuc/pdf"
)

type Request struct {
	URL string `json:"url" validate:"required,url"`
}

type Response struct {
	response.Response
	Count int `json:"count"`
}

type Event struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	SexAge       string    `json:"sex_age"`
	Discipline   string    `json:"discipline"`
	TimeStart    time.Time `json:"time_start"`
	TimeEnd      time.Time `json:"time_end"`
	Venue        string    `json:"venue"`
	Participants int       `json:"participants"`
	SportID      string    `json:"sport_id"`
}

type Parser interface {
	SaveEvents(events []Event) (int, error)
	SaveSport(sport string) (string, error)
}

func New(log *slog.Logger, cfg config.Config, parser Parser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.parser.parse.New"

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

		fileName := fmt.Sprintf("%s/%s.%s", cfg.PDFStorage, uuid.New().String(), "pdf")
		// fileName := fmt.Sprintf("%s/%s.%s", cfg.PDFStorage, "cb064543-412d-4dcd-b64c-8859111574f0", "pdf")

		out, err := os.Create(fileName)
		if err != nil {
			log.Error("failed to create pdf file", sl.Err(err))

			render.JSON(w, r, response.Error("failed to create pdf file"))

			return
		}
		defer out.Close()

		resp, err := http.Get(req.URL)
		if err != nil {
			log.Error("failed to download pdf file", sl.Err(err))

			render.JSON(w, r, response.Error("failed to download pdf file"))

			return
		}
		defer resp.Body.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Error("failed to copy pdf file", sl.Err(err))

			render.JSON(w, r, response.Error("failed to copy pdf file"))

			return
		}

		file, reader, err := pdf.Open(fileName)
		if err != nil {
			log.Error("failed to open pdf file", sl.Err(err))

			render.JSON(w, r, response.Error("failed to open pdf file"))

			return
		}
		defer file.Close()

		log.Info("file downloaded, parsing...")

		// Получаем количество страниц
		totalPages := reader.NumPage()

		eventRegex := regexp.MustCompile(`([А-Яа-я-\s]+)?(\d{16})\s+(.*?)\s(\d{2}\.\d{2}\.\d{4})\s(\d{2}\.\d{2}\.\d{4})+(.*?)\s+(\d+)`)
		nameRegex := regexp.MustCompile(`([А-Я\s-\d()."«»№#]+)\s([а-я,\s\d-]+)\s?([А-ЯAA-Z][А-ЯA-Z\s-\d,а-я()]+)?`)
		sportRegex := regexp.MustCompile(`(?:^|\s)([А-ЯЁ]+(?:-[А-ЯЁ]+)*(?:\s[А-ЯЁ]+(?:-[А-ЯЁ]+)*)*)(=?\s|$)`)

		var events []Event
		lastSport := ""

		// Извлекаем текст
		for pageNum := 1; pageNum <= totalPages; pageNum++ {
			page := reader.Page(pageNum)
			if page.V.IsNull() {
				continue
			}

			// Извлекаем текст
			text, err := extractTextFromPage(page, cfg)
			if err != nil {
				log.Error("error while extracting text from pdf file", sl.Err(err))

				render.JSON(w, r, response.Error("error while extracting text from pdf file"))

				return
			}

			matches := eventRegex.FindAllStringSubmatch(text, -1)

			for _, match := range matches {
				format := "02.01.2006"

				if strings.Trim(match[1], " ") != "" {
					sport := sportRegex.FindAllStringSubmatch(strings.Trim(match[1], " "), -1)

					if len(sport) == 1 {
						name := sport[0][1]

						id, err := parser.SaveSport(name)
						if err != nil {
							log.Error("error while saving sport", sl.Err(err))

							render.JSON(w, r, response.Error("error while saving sport"))

							return
						}

						lastSport = id
					}
				}

				timeStart, err := time.Parse(format, match[4])
				if err != nil {
					log.Error("error while converting date to timestamp", sl.Err(err))

					render.JSON(w, r, response.Error("error while converting date to timestamp"))

					return
				}

				timeEnd, err := time.Parse(format, match[5])
				if err != nil {
					log.Error("error while converting date to timestamp", sl.Err(err))

					render.JSON(w, r, response.Error("error while converting date to timestamp"))

					return
				}

				parsedName := nameRegex.FindAllStringSubmatch(match[3], -1)

				id, err := strconv.Atoi(match[2])
				if err != nil {
					log.Error("error while strconv.Atoi", sl.Err(err))

					render.JSON(w, r, response.Error("error while strconv.Atoi"))

					return
				}

				participants, err := strconv.Atoi(match[7])
				if err != nil {
					log.Error("error while strconv.Atoi", sl.Err(err))

					render.JSON(w, r, response.Error("error while strconv.Atoi"))

					return
				}

				if len(parsedName) == 0 {
					log.Error("error while parsing name", slog.String("name", match[3]))

					continue
				}

				discipline := ""
				if len(parsedName[0]) == 4 {
					discipline = parsedName[0][3]
				}

				events = append(events, Event{
					ID:           id,
					Name:         strings.Trim(parsedName[0][1], " "),
					SexAge:       strings.Trim(parsedName[0][2], " "),
					Discipline:   strings.Trim(discipline, " "),
					TimeStart:    timeStart,
					TimeEnd:      timeEnd,
					Venue:        strings.Trim(match[6], " "),
					Participants: participants,
					SportID:      lastSport,
				})
			}
		}

		totalCount := 0

		for offset := 1; offset <= (len(events)/4000)+1; offset++ {
			log.Info("saving events", slog.Int("part", offset))

			sliceStart := (offset - 1) * 4000
			sliceEnd := offset * 4000

			if sliceEnd > len(events) {
				sliceEnd = len(events)
			}

			count, err := parser.SaveEvents(events[sliceStart:sliceEnd])
			if err != nil {
				log.Error("error while saving events to database", sl.Err(err))

				render.JSON(w, r, response.Error("error while saving events to database"))

				return
			}

			totalCount += count
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
			Count:    totalCount,
		})
	}
}

func extractTextFromPage(page pdf.Page, cfg config.Config) (string, error) {
	var builder strings.Builder

	// Переменные для отслеживания последней вертикальной и горизонтальной позиции
	var lastY, lastX float64

	// Перебираем содержимое страницы
	for _, text := range page.Content().Text {
		if text.S != "" {
			// Если вертикальная позиция изменяется на большее расстояние, чем порог, то начинаем новую строку
			if lastY > 0 && text.Y < lastY-cfg.Parser.VerticalThreshold {
				builder.WriteRune(' ') // Добавляем символ новой строки
			}

			// Если горизонтальная позиция изменяется на большее расстояние, чем порог, то считаем, что это новая колонка
			if lastX > 0 && text.X > lastX+cfg.Parser.HorizontalThreshold {
				builder.WriteRune(' ') // Добавляем пробел для новой колонки
			}

			// Добавляем текст
			builder.WriteString(text.S)

			// Обновляем последние позиции
			lastY = text.Y
			lastX = text.X
		}
	}

	return strings.TrimSpace(builder.String()), nil
}

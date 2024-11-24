package postgresql

import (
	"calendar-backend/internal/http-server/handlers/parser/parse"
	"calendar-backend/internal/storage/postgresql/ent"
	"calendar-backend/internal/storage/postgresql/ent/events"
	"context"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	Client ent.Client
	Ctx    context.Context
}

func New(host string, port int, user string, dbName string, password string, sslMode string) (*Storage, error) { // Создание хранилища
	const op = "storage.postgresql.New"

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbName, password, sslMode))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{Client: *client, Ctx: ctx}, nil
}

func (s *Storage) SaveEvents(events []parse.Event) (int, error) {
	const op = "storage.postgresql.SaveEvents"

	res, err := s.Client.Events.MapCreateBulk(events, func(c *ent.EventsCreate, i int) {
		event := events[i]

		c.SetID(event.ID).SetName(event.Name).
			SetSexAge(event.SexAge).
			SetDiscipline(event.Discipline).
			SetTimeStart(event.TimeStart).
			SetTimeEnd(event.TimeEnd).
			SetVenue(event.Venue).
			SetParticipants(event.Participants).
			SetSportID(event.SportID)
	}).Save(s.Ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return len(res), nil
}

func (s *Storage) SaveSport(sport string) (string, error) {
	const op = "storage.postgresql.SaveSport"

	res, err := s.Client.Sports.Create().SetName(sport).Save(s.Ctx)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return res.ID.String(), nil
}

func (s *Storage) GetEventsByMonth(year int, month int) ([]*ent.Events, error) {
	const op = "storage.postgresql.GetEventsByMonth"

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	fmt.Println(startDate)
	fmt.Println(endDate)

	res, err := s.Client.Events.Query().Where(events.TimeStartGTE(startDate), events.TimeStartLT(endDate)).All(s.Ctx)
	if err != nil {
		return []*ent.Events{}, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}

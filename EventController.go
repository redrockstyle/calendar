package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

type EventController struct {
	eventRepository *EventRepository
}

func NewEventController() (*EventController, error) {
	repository, err := NewEventRepository()

	if err != nil {
		return nil, err
	}

	return &EventController{
		eventRepository: repository,
	}, nil
}

func (c *EventController) Filter(ctx *fasthttp.RequestCtx) {
	start := ctx.Request.URI().QueryArgs().Peek("start")
	end := ctx.Request.URI().QueryArgs().Peek("end")
	if start == nil {
		responseJSON(ctx, &map[string]string{"error": "Не указана дата начала события"}, fasthttp.StatusBadRequest)
		return
	}

	if end == nil {
		responseJSON(ctx, &map[string]string{"error": "Не указана дата окончания события"}, fasthttp.StatusBadRequest)
		return
	}

	startTime, err := time.Parse("2006-01-02 15:04", string(start))

	if  err != nil {
		responseJSON(ctx, &map[string]string{"error": "Неправильный формат даты начала события"}, fasthttp.StatusBadRequest)
		return
	}

	endTime, err := time.Parse("2006-01-02 15:04", string(end))

	if  err != nil {
		responseJSON(ctx, &map[string]string{"error": "Неправильный формат даты окончания события"}, fasthttp.StatusBadRequest)
		return
	}

	events, err := c.eventRepository.Filter(startTime, endTime)

	if err != nil {
		responseJSON(ctx, &map[string]string{"error": "Ошибка при получении событий"}, fasthttp.StatusInternalServerError)
		return
	}

	responseJSON(ctx, events, fasthttp.StatusOK)
}

func (c *EventController) Create(ctx *fasthttp.RequestCtx) {
	var events []Event
	err := json.Unmarshal(ctx.Request.Body(), &events)

	if err != nil {
		responseJSON(ctx, &map[string]string{"error": "Неправильный json"}, fasthttp.StatusBadRequest)
		return
	}

	for _, event := range events {
		startTime, err := time.Parse("2006-01-02 15:04", string(event.Start))

		if  err != nil {
			responseJSON(ctx, &map[string]string{"error": "Неправильный формат даты начала события"}, fasthttp.StatusBadRequest)
			return
		}

		endTime, err := time.Parse("2006-01-02 15:04", string(event.End))

		if  err != nil {
			responseJSON(ctx, &map[string]string{"error": "Неправильный формат даты окончания события"}, fasthttp.StatusBadRequest)
			return
		}
		event.Duration = endTime.Sub(startTime).String()
		c.eventRepository.Create(event)
	}
}

func (c *EventController) Update(ctx *fasthttp.RequestCtx) {
	var events []Event
	err := json.Unmarshal(ctx.Request.Body(), &events)

	if err != nil {
		responseJSON(ctx, &map[string]string{"error": "Неправильный json"}, fasthttp.StatusBadRequest)
		return
	}

	for _, event := range events {
		startTime, err := time.Parse("2006-01-02 15:04", string(event.Start))

		if  err != nil {
			responseJSON(ctx, &map[string]string{"error": "Неправильный формат даты начала события"}, fasthttp.StatusBadRequest)
			return
		}

		endTime, err := time.Parse("2006-01-02 15:04", string(event.End))

		if  err != nil {
			responseJSON(ctx, &map[string]string{"error": "Неправильный формат даты окончания события"}, fasthttp.StatusBadRequest)
			return
		}
		event.Duration = endTime.Sub(startTime).String()
		c.eventRepository.Update(event)
	}
}

func (c *EventController) Delete(ctx *fasthttp.RequestCtx) {
	var ids []int
	json.Unmarshal(ctx.Request.Body(), &ids)

	if ids == nil {
		responseJSON(ctx, &map[string]string{"error": "Неправильный формат данных"}, fasthttp.StatusBadRequest)
		return
	}

	var errorsList []string
	for _, id := range ids {
		err := c.eventRepository.Delete(id)
		if err != nil {
			errorsList = append(errorsList, fmt.Sprintf("Ошибка при удалении события с id = %d", id))
		}
	}

	if errorsList != nil {
		responseJSON(ctx, &map[string]string{"error": strings.Join(errorsList, "\n")}, fasthttp.StatusInternalServerError)
		return
	}

}
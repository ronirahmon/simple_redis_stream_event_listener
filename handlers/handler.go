package handlers

import (
	"context"
	"simple_redis_stream_event_listener/event"
	"simple_redis_stream_event_listener/models"
	"simple_redis_stream_event_listener/services"
)

type Handler struct {
}

func (h *Handler) ChangeStatusEvent(ctx context.Context) (err error) {
	data := &models.Status{}
	if err = event.BindContexData(ctx, data); err != nil {
		return
	}
	err = services.PrintStatus(data)
	return
}

package amqprpc

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase"
	"github.com/dimk00z/GophKeeper/pkg/rabbitmq/rmq_rpc/server"
)

type GophKeeperRoutes struct {
	GophKeeperUseCase usecase.GophKeeper
}

func newGophKeeperRoutes(routes map[string]server.CallHandler, t usecase.GophKeeper) {
	r := &GophKeeperRoutes{t}
	{
		routes["getHistory"] = r.getHistory()
	}
}

type historyResponse struct {
	History []entity.GophKeeper `json:"history"`
}

func (r *GophKeeperRoutes) getHistory() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		GophKeepers, err := r.GophKeeperUseCase.History(context.Background())
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - GophKeeperRoutes - getHistory - r.GophKeeperUseCase.History: %w", err)
		}

		response := historyResponse{GophKeepers}

		return response, nil
	}
}

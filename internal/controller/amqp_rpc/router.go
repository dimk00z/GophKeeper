package amqprpc

import (
	"github.com/dimk00z/GophKeeper/internal/usecase"
	"github.com/dimk00z/GophKeeper/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.GophKeeper) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newGophKeeperRoutes(routes, t)
	}

	return routes
}

package application

import (
	ordersInfra "go-echo-template/internal/infrastructure/orders"
	productsInfra "go-echo-template/internal/infrastructure/products"
	usersInfra "go-echo-template/internal/infrastructure/users"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
	HTTPServer   *echo.Echo
	UsersRepo    *usersInfra.InMemoryRepo
	OrdersRepo   *ordersInfra.InMemoryRepo
	ProductsRepo *productsInfra.InMemoryRepo
}

func (s *ServerSuite) SetupTest() {
	s.UsersRepo = usersInfra.NewInMemoryRepo()
	s.OrdersRepo = ordersInfra.NewInMemoryRepo()
	s.ProductsRepo = productsInfra.NewInMemoryRepo()
	s.HTTPServer = SetupHTTPServer(s.UsersRepo, s.OrdersRepo, s.ProductsRepo)
}

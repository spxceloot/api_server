package service

import (
	"context"
	"github/luqxus/spxce/database"
	"github/luqxus/spxce/types"
)

type Service interface {
	GetParkingSpaces(ctx context.Context, geoLocation *types.GeoLocation) ([]*types.ParkingSpace, error)
	GetParkingSpace(ctx context.Context) (*types.ParkingSpace, error)
	CreateUser(ctx context.Context, data *types.CreateUserRequest) (string, error)
	Login(ctx context.Context, data *types.LoginRequest) (string, error)
}

type DefaultService struct {
	AuthService
	ParkingService
}

func New(datastore database.Database) Service {
	return &DefaultService{
		AuthService: AuthService{
			datastore: datastore,
		},
		ParkingService: ParkingService{
			datastore: datastore,
		},
	}
}

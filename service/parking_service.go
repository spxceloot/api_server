package service

import (
	"context"
	"github/luqxus/spxce/database"
	"github/luqxus/spxce/types"
)

type ParkingService struct {
	datastore database.Database
}

func NewParkingService(datastore database.Database) *ParkingService {
	return &ParkingService{
		datastore: datastore,
	}
}

func (s *ParkingService) GetParkingSpaces(ctx context.Context, geoLocation *types.GeoLocation) ([]*types.ParkingSpace, error) {
	return nil, nil
}

func (s *ParkingService) GetParkingSpace(ctx context.Context) (*types.ParkingSpace, error) {
	return nil, nil
}

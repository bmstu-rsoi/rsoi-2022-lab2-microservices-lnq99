package service

import (
	"app/flight/model"
	"app/flight/repository"
	"context"
)

type Service interface {
	ListFlights(ctx context.Context, page, size int32) model.PaginationResponse
}

type ServiceImpl struct {
	repo repository.Repo
}

var airportMap = make(map[int32]model.Airport)

func NewService(repo repository.Repo) Service {
	return &ServiceImpl{repo: repo}
}

func (s *ServiceImpl) ListFlights(ctx context.Context, page, size int32) model.PaginationResponse {
	//flight := s.repo.SelectFlightsWithOffsetLimit((page-1)*size, size)

	res := model.PaginationResponse{
		TotalElements: 0,
		Page:          page,
		PageSize:      size,
		Items:         []model.FlightResponse{},
	}

	flights, err := s.repo.ListFlightsWithOffsetLimit(ctx,
		repository.ListFlightsWithOffsetLimitParams{size, (page - 1) * size})

	if err != nil {
		return res
	}

	res.TotalElements = int32(len(flights))

	for _, f := range flights {
		fromAirport := s.loadAirport(ctx, f.FromAirportID)
		toAirport := s.loadAirport(ctx, f.ToAirportID)
		if fromAirport == nil || toAirport == nil {
			continue
		}
		res.Items = append(res.Items, model.FlightResponse{
			FlightNumber: f.FlightNumber,
			FromAirport:  fromAirport.Name + " " + fromAirport.City,
			ToAirport:    toAirport.Name + " " + fromAirport.City,
			Date:         f.Datetime.String(),
			Price:        f.Price,
		})
	}

	return res
}

func (s *ServiceImpl) loadAirport(ctx context.Context, id int32) *model.Airport {
	airport, ok := airportMap[id]
	if !ok {
		newAirport, err := s.repo.GetAirport(ctx, id)
		if err != nil {
			return nil
		}
		airport = model.Airport(newAirport)
		airportMap[id] = airport
	}
	return &airport
}

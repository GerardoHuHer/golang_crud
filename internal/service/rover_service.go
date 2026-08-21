package service

import (
	"context"
	"errors"

	"github.com/GerardoHuHer/go_crud/internal/domain"
	"github.com/GerardoHuHer/go_crud/internal/repository"
)

var ErrIdTaken = errors.New("Id already taken")

type RoverService struct {
	repo repository.RoverRepository
}

func NewRoverService(repo repository.RoverRepository) *RoverService {
	return &RoverService{repo: repo}
}

func (r *RoverService) CreateRover(ctx context.Context, pos_x, pos_y int) (*domain.Rover, error) {
	rover := &domain.Rover{Pos_x: pos_x, Pos_y: pos_y}
	if err := r.repo.Create(ctx, rover); err != nil {
		return nil, err
	}
	return rover, nil
}

func (r *RoverService) GetRoverById(ctx context.Context, id uint) (*domain.Rover, error) {
	return r.repo.FindById(ctx, id)
}

func (r *RoverService) UpdateRoverById(ctx context.Context, id uint, pos_x, pos_y int) (*domain.Rover, error) {
	rover, err := r.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	rover.Pos_x = pos_x
	rover.Pos_y = pos_y

	if err := r.repo.Update(ctx, rover); err != nil {
		return nil, err
	}
	return rover, nil
}

func (r *RoverService) DeleteRoverById(ctx context.Context, id uint) error {
	return r.repo.Delete(ctx, id)
}

func (r *RoverService) GetAll(ctx context.Context) (*[]domain.Rover, error) {
	return r.repo.GetAll(ctx)
}

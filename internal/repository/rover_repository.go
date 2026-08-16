package repository

import (
	"context"

	"github.com/GerardoHuHer/go_crud/internal/domain"
)

type RoverRepository interface {
	FindById(ctx context.Context, id uint) (*domain.Rover, error)
	Create(ctx context.Context, r *domain.Rover) error
	Update(ctx context.Context, r *domain.Rover) error
	Delete(ctx context.Context, id uint) error
}

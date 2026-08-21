package postgres

import (
	"context"

	"github.com/GerardoHuHer/go_crud/internal/domain"
	"gorm.io/gorm"
)

type roverRepository struct {
	db *gorm.DB
}

func NewRoverRepository(db *gorm.DB) *roverRepository {
	return &roverRepository{db: db}
}

// READ
func (r *roverRepository) FindById(ctx context.Context, id uint) (*domain.Rover, error) {
	var rover domain.Rover
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&rover).Error; err != nil {
		return nil, err
	}

	return &rover, nil
}

// CREATE
func (r *roverRepository) Create(ctx context.Context, rover *domain.Rover) error {
	return r.db.WithContext(ctx).Create(rover).Error
}

// UPDATE
func (r *roverRepository) Update(ctx context.Context, rover *domain.Rover) error {
	return r.db.WithContext(ctx).Save(rover).Error
}

// DELETE
func (r *roverRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Rover{}, id).Error
}

func (r *roverRepository) GetAll(ctx context.Context) (*[]domain.Rover, error) {
	var rovers *[]domain.Rover
	if err := r.db.WithContext(ctx).Find(&rovers).Error; err != nil {
		return nil, err
	}
	return rovers, nil
}

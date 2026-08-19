package service

import (
	"context"
	"errors"
	"testing"

	"github.com/GerardoHuHer/go_crud/internal/domain"

	"github.com/stretchr/testify/assert"
)

type fakeRoverRepo struct {
	rovers map[uint]*domain.Rover
	nextID uint
}

package repository

import "github.com/sharifulb07/goku/internal/models"

type ResourceRepo struct {
}

func NewResourceRepo() *ResourceRepo {
	return &ResourceRepo{}
}

func (r *ResourceRepo) Create(resource *models.Resource)error{
	
}
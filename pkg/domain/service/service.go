package service

import (
	"context"
	"github.com/mdapathy/imageuploader/pkg/domain"
	command "github.com/mdapathy/imageuploader/pkg/domain/cmd"
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"github.com/mdapathy/imageuploader/pkg/domain/query"
)

type Service struct {
	repo domain.ImageRepository

	filterer query.Filterer
}

func New(r domain.ImageRepository, f query.Filterer) *Service {
	return &Service{
		repo:     r,
		filterer: f,
	}
}

func (svc *Service) List(ctx context.Context, q *query.List) (*query.Result, error) {
	return svc.repo.List(ctx, q)
}

func (svc *Service) Details(ctx context.Context, q *query.Detail) (*model.Image, error) {
	return svc.repo.Details(ctx, q)
}

func (svc *Service) Create(ctx context.Context, cmd *command.Create) error {
	if err := command.Process(cmd); err != nil {
		return err
	}
	img, err := model.NewImage(cmd.Content, cmd.UserID)
	if err != nil {
		return err
	}

	return svc.repo.Create(ctx, img)
}

func (svc *Service) Delete(ctx context.Context, cmd *command.Delete) error {
	if err := command.Process(cmd); err != nil {
		return err
	}
	return svc.repo.Delete(ctx, cmd.ID, cmd.UserID)
}

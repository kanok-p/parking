package service

import "parking/repository"

type implementation struct {
	repo repository.Repo
}

func New(repo repository.Repo) (service Service) {
	return &implementation{repo}
}

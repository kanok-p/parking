package service

import (
	"parking/repository"
)

type Service interface {
	Stdin(input CommandInput) string
	OpenFile(path string) (command []CommandInput, err error)
	List() (items []repository.Park, err error)
	Create(input int) (err error)
	Parking(input *repository.Park) (ID string, err error)
	Delete(input int) (err error)
	RunCommand(command []CommandInput)
}

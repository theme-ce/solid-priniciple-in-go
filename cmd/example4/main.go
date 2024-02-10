package main

import "fmt"

type Repository interface {
	GetData() string
}

type respository struct {
	someField string
}

func NewRepository(someField string) Repository {
	return &respository{
		someField: someField,
	}
}

func (r *respository) GetData() string {
	return r.someField
}

type Service interface {
	Process()
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Process() {
	fmt.Println(s.repo.GetData())
}

func main() {
	repo := NewRepository("some value")
	s := NewService(repo)

	s.Process()
}

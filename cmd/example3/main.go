package main

type Repository interface {
	GetDataFromA() string
	GetDataFromB() string
}

type repository struct {
	someFieldA string
	someFieldB string
}

func NewRepository(someFieldA, someFieldB string) Repository {
	return &repository{
		someFieldA: someFieldA,
		someFieldB: someFieldB,
	}
}

func (r *repository) GetDataFromA() string {
	return r.someFieldA
}

func (r *repository) GetDataFromB() string {
	return r.someFieldB
}

type RepositoryA interface {
	GetDataFromA() string
}

type repositoryA struct {
	someFieldA string
}

func NewRepositoryA(someFieldA string) RepositoryA {
	return &repositoryA{
		someFieldA: someFieldA,
	}
}

func (r *repositoryA) GetDataFromA() string {
	return r.someFieldA
}

type RepositoryB interface {
	GetDataFromB() string
}

type repositoryB struct {
	someFieldB string
}

func NewRepositoryB(someFieldB string) RepositoryB {
	return &repositoryB{
		someFieldB: someFieldB,
	}
}

func (r *repositoryB) GetDataFromB() string {
	return r.someFieldB
}

func main() {
	repo := NewRepository("some value A", "some value B")
	_ = repo.GetDataFromA()
	_ = repo.GetDataFromB()

	repoA := NewRepositoryA("some value")
	repoB := NewRepositoryB("some value")
	_ = repoA.GetDataFromA()
	_ = repoB.GetDataFromB()
}

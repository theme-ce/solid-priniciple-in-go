package main

/*
Example.
Retrieve data from a database and cache data
*/

/*
Bad practice.
This function is not following the Interface Segregation Principle.
It is doing two things: getting data from a database and caching data to redis.
*/

type Repository interface {
	GetData() string
	CacheData(data string)
}

type repository struct {
	// some database connection
	// some redis connection
}

func (r *repository) GetData() string {
	// get data from database
	return "data"
}

func (r *repository) CacheData(data string) {
	// cache data to redis
}

func NewRepository() Repository {
	return &repository{}
}

/*
Good practice.
This function is following the Interface Segregation Principle.
It is doing one thing: getting data from a database.
*/

type RepositoryDb interface {
	GetData() string
}

type repositoryDb struct {
	// some database connection
}

func (r *repositoryDb) GetData() string {
	// get data from database
	return "data"
}

func NewRepositoryDb() RepositoryDb {
	return &repositoryDb{}
}

/*
Good practice.
This function is following the Interface Segregation Principle.
It is doing one thing: caching data to redis.
*/

type RepositoryRedis interface {
	CacheData(data string)
}

type repositoryRedis struct {
	// some redis connection
}

func (r *repositoryRedis) CacheData(data string) {
	// cache data to redis
}

func NewRepositoryRedis() RepositoryRedis {
	return &repositoryRedis{}
}

func main() {
	repo := NewRepository()
	data := repo.GetData()
	repo.CacheData(data)

	repoDb := NewRepositoryDb()
	dataDb := repoDb.GetData()
	repoRedis := NewRepositoryRedis()
	repoRedis.CacheData(dataDb)
}

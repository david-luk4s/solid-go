package main

import (
	"fmt"
	"log"
	"math/rand"
)

var (
	Storage map[int64]*User
)

type User struct {
	id       int64
	username string
}

type ServiceUserRepos interface {
	InsertUser(username string) (*User, error)
	GetUserById(user_id int64) (*User, error)
}

type UserRepos struct {
	repo ServiceUserRepos
}

func (r *UserRepos) ServiceInsertUser(username string) (*User, error) {
	user, err := r.repo.InsertUser(username)
	if err != nil {
		log.Panic(err)
	}
	return user, nil
}

func (r *UserRepos) ServiceGetUserById(user_id int64) (*User, error) {
	user, err := r.repo.GetUserById(user_id)
	if err != nil {
		log.Panic(err)
	}
	return user, nil
}

type RepoPostgresql struct{}

func (rpg *RepoPostgresql) InsertUser(username string) (*User, error) {
	fmt.Println("#driver postgresql")
	id := rand.Intn(10000)
	user := &User{int64(id), username}

	Storage[user.id] = user
	return user, nil
}

func (rpg *RepoPostgresql) GetUserById(user_id int64) (*User, error) {
	return Storage[user_id], nil
}

type RepoMongoDB struct{}

func (rpg *RepoMongoDB) InsertUser(username string) (*User, error) {
	fmt.Println("#driver mongodb")
	id := rand.Intn(10000)
	user := &User{int64(id), username}

	Storage[user.id] = user
	return user, nil
}

func (rpg *RepoMongoDB) GetUserById(user_id int64) (*User, error) {
	return Storage[user_id], nil
}

func init() {
	Storage = make(map[int64]*User)
}

func main() {
	//any driver
	repoMG := &RepoMongoDB{}

	//abstract for UserRepos
	repoUser := &UserRepos{repoMG}
	repoUser.ServiceInsertUser("david_luk4s")

	for _, v := range Storage {
		fmt.Println(v)
	}
}

package container

import (
	"fmt"
	"math/rand"
)

type Container struct {
	ID string
}

func NewContainer() (*Container, error) {
	id := generateID()
	return &Container{ID: id}, nil
}

func generateID() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
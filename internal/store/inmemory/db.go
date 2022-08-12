package inmemory

import (
	"animeshop/internal/models"
	"animeshop/internal/store"
	"context"
	"fmt"

	"sync"
)

// DB saves information about laptops
type DB struct {
	data map[int]*models.Manga
	mu   *sync.RWMutex
}

// NewDB is the function for creating basic Database
func NewDB() store.Store {
	return &DB{
		data: make(map[int]*models.Manga),
		mu:   new(sync.RWMutex),
	}
}

// Create for creating new element in DB
func (db *DB) Create(ctx context.Context, manga *models.Manga) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[manga.ID] = manga

	return nil
}

// All is used for reading all elements in DB
func (db *DB) All(ctx context.Context) ([]*models.Manga, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	mangas := make([]*models.Manga, 0, len(db.data))
	for _, laptop := range db.data {
		mangas = append(mangas, laptop)
	}

	return mangas, nil
}

// ByID is used for reading elements by id in DB
func (db *DB) ByID(ctx context.Context, id int) (*models.Manga, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	manga, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No manga eith id: %d", id)
	}

	return manga, nil
}

// Update is used for updating elements in DB
func (db *DB) Update(ctx context.Context, manga *models.Manga) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[manga.ID] = manga
	return nil
}

// Delete is used for deleting elements by id in DB
func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, id)
	return nil
}

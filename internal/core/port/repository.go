// internal/core/port/repository.go
package port

import (
	"context"

	"github.com/TubagusAldiMY/Go-React-ComicReader-Be/internal/core/domain" // Sesuaikan path
)

// GenreRepository mendefinisikan operasi yang bisa dilakukan pada data genre.
type GenreRepository interface {
	List(ctx context.Context) ([]domain.Genre, error)
	Create(ctx context.Context, genre *domain.Genre) error
	GetBySlug(ctx context.Context, slug string) (*domain.Genre, error)
	Update(ctx context.Context, genre *domain.Genre) error
	DeleteBySlug(ctx context.Context, slug string) error
}

package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	ucase "github.com/bxcodec/go-clean-arch/movie/usecase"
)

func TestFetchMovies(t *testing.T) {
	mockMovieRepo := new(mocks.MovieRepository)
	mockMovie := domain.Movie{
		Title:   "Hello",
		Content: "Content",
	}

	mockListArtilce := make([]domain.Movie, 0)
	mockListArtilce = append(mockListArtilce, mockMovie)

	t.Run("success", func(t *testing.T) {
		mockMovieRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListArtilce, "next-cursor", nil).Once()
		mockLogmovie := domain.Logmovie{
			ID:   1,
			Name: "Iman Tumorang",
		}
		mockLogmovierepo := new(mocks.LogmovieRepository)
		mockLogmovierepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockLogmovie, nil)
		u := ucase.NewMovieUsecase(mockMovieRepo, mockLogmovierepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
		cursorExpected := "next-cursor"
		assert.Equal(t, cursorExpected, nextCursor)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListArtilce))

		mockMovieRepo.AssertExpectations(t)
		mockLogmovierepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockMovieRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(nil, "", errors.New("Unexpexted Error")).Once()

		mockLogmovierepo := new(mocks.LogmovieRepository)
		u := ucase.NewMovieUsecase(mockMovieRepo, mockLogmovierepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)

		assert.Empty(t, nextCursor)
		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockMovieRepo.AssertExpectations(t)
		mockLogmovierepo.AssertExpectations(t)
	})

}

func TestGetMovieByID(t *testing.T) {
	mockMovieRepo := new(mocks.MovieRepository)
	mockMovie := domain.Movie{
		Title:   "Hello",
		Content: "Content",
	}
	mockLogmovie := domain.Logmovie{
		ID:   1,
		Name: "Iman Tumorang",
	}

	t.Run("success", func(t *testing.T) {
		mockMovieRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockMovie, nil).Once()
		mockLogmovierepo := new(mocks.LogmovieRepository)
		mockLogmovierepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockLogmovie, nil)
		u := ucase.NewMovieUsecase(mockMovieRepo, mockLogmovierepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockMovie.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockMovieRepo.AssertExpectations(t)
		mockLogmovierepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockMovieRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Movie{}, errors.New("Unexpected")).Once()

		mockLogmovierepo := new(mocks.LogmovieRepository)
		u := ucase.NewMovieUsecase(mockMovieRepo, mockLogmovierepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockMovie.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.Movie{}, a)

		mockMovieRepo.AssertExpectations(t)
		mockLogmovierepo.AssertExpectations(t)
	})

}

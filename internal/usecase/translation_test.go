package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/usecase"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func GophKeeper(t *testing.T) (*usecase.GophKeeperUseCase, *MockGophKeeperRepo, *MockGophKeeperWebAPI) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockGophKeeperRepo(mockCtl)
	webAPI := NewMockGophKeeperWebAPI(mockCtl)

	GophKeeper := usecase.New(repo, webAPI)

	return GophKeeper, repo, webAPI
}

func TestHistory(t *testing.T) {
	t.Parallel()

	GophKeeper, repo, _ := GophKeeper(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background()).Return(nil, nil)
			},
			res: []entity.GophKeeper(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background()).Return(nil, errInternalServErr)
			},
			res: []entity.GophKeeper(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := GophKeeper.History(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestTranslate(t *testing.T) {
	t.Parallel()

	GophKeeper, repo, webAPI := GophKeeper(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				webAPI.EXPECT().Translate(entity.GophKeeper{}).Return(entity.GophKeeper{}, nil)
				repo.EXPECT().Store(context.Background(), entity.GophKeeper{}).Return(nil)
			},
			res: entity.GophKeeper{},
			err: nil,
		},
		{
			name: "web API error",
			mock: func() {
				webAPI.EXPECT().Translate(entity.GophKeeper{}).Return(entity.GophKeeper{}, errInternalServErr)
			},
			res: entity.GophKeeper{},
			err: errInternalServErr,
		},
		{
			name: "repo error",
			mock: func() {
				webAPI.EXPECT().Translate(entity.GophKeeper{}).Return(entity.GophKeeper{}, nil)
				repo.EXPECT().Store(context.Background(), entity.GophKeeper{}).Return(errInternalServErr)
			},
			res: entity.GophKeeper{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := GophKeeper.Translate(context.Background(), entity.GophKeeper{})

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

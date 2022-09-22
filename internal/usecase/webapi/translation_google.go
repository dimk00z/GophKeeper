package webapi

import (
	"fmt"

	translator "github.com/Conight/go-googletrans"

	"github.com/dimk00z/GophKeeper/internal/entity"
)

// GophKeeperWebAPI -.
type GophKeeperWebAPI struct {
	conf translator.Config
}

// New -.
func New() *GophKeeperWebAPI {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	return &GophKeeperWebAPI{
		conf: conf,
	}
}

// Translate -.
func (t *GophKeeperWebAPI) Translate(GophKeeper entity.GophKeeper) (entity.GophKeeper, error) {
	trans := translator.New(t.conf)

	result, err := trans.Translate(GophKeeper.Original, GophKeeper.Source, GophKeeper.Destination)
	if err != nil {
		return entity.GophKeeper{}, fmt.Errorf("GophKeeperWebAPI - Translate - trans.Translate: %w", err)
	}

	GophKeeper.GophKeeper = result.Text

	return GophKeeper, nil
}

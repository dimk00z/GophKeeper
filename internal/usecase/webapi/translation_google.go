package webapi

// GophKeeperWebAPI -.
type GophKeeperWebAPI struct{}

// New -.
func New() *GophKeeperWebAPI {
	return &GophKeeperWebAPI{}
}

// // Translate -.
// func (t *GophKeeperWebAPI) Translate(GophKeeper entity.GophKeeper) (entity.GophKeeper, error) {
// 	trans := translator.New(t.conf)

// 	result, err := trans.Translate(GophKeeper.Original, GophKeeper.Source, GophKeeper.Destination)
// 	if err != nil {
// 		return entity.GophKeeper{}, fmt.Errorf("GophKeeperWebAPI - Translate - trans.Translate: %w", err)
// 	}

// 	GophKeeper.GophKeeper = result.Text

// 	return GophKeeper, nil
// }

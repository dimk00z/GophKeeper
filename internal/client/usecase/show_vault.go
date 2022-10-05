package usecase

import (
	"fmt"

	"github.com/dimk00z/GophKeeper/internal/client/usecase/viewsets"
	"github.com/fatih/color"
)

func (uc *GophKeeperClientUseCase) ShowVault(userPassword, showVaultOption string) {
	if !uc.verifyPassword(userPassword) {
		return
	}

	switch showVaultOption {
	case "a":
		uc.showCards(uc.repo.LoadCards())
		uc.showLogins(uc.repo.LoadLogins())
		uc.showNotes(uc.repo.LoadNotes())
	case "c":
		uc.showCards(uc.repo.LoadCards())
	case "l":
		uc.showLogins(uc.repo.LoadLogins())
	case "n":
		uc.showNotes(uc.repo.LoadNotes())
	}
}

func (uc *GophKeeperClientUseCase) showCards(cards []viewsets.CardForList) {
	color.Yellow("Users cards:")
	yellow := color.New(color.FgYellow).SprintFunc()
	for _, card := range cards {
		fmt.Printf("ID: %s name:%s brand: %s\n", //nolint:forbidigo // cli printing
			yellow(card.ID),
			yellow(card.Name),
			yellow(card.Brand))
	}
	fmt.Printf("Total %s cards\n", yellow(len(cards))) //nolint:forbidigo // cli printing
}

func (uc *GophKeeperClientUseCase) showLogins(logins []viewsets.LoginForList) {
	color.Yellow("Users logins:")
	yellow := color.New(color.FgYellow).SprintFunc()
	for _, login := range logins {
		fmt.Printf("ID: %s name:%s uri: %s\n", //nolint:forbidigo // cli printing
			yellow(login.ID),
			yellow(login.Name),
			yellow(login.URI))
	}
	fmt.Printf("Total %s logins\n", yellow(len(logins))) //nolint:forbidigo // cli printing
}

func (uc *GophKeeperClientUseCase) showNotes(notes []viewsets.NoteForList) {
	color.Yellow("Users notes:")
	yellow := color.New(color.FgYellow).SprintFunc()
	for _, note := range notes {
		fmt.Printf("ID: %s name:%s\n", //nolint:forbidigo // cli printing
			yellow(note.ID),
			yellow(note.Name))
	}
	fmt.Printf("Total %s notes\n", yellow(len(notes))) //nolint:forbidigo // cli printing
}

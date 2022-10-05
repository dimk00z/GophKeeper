package usecase

import (
	"fmt"
	"log"

	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/dimk00z/GophKeeper/internal/utils"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func (uc *GophKeeperClientUseCase) loadNotes(accessToken string) {
	notes, err := uc.clientAPI.GetNotes(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = uc.repo.SaveNotes(notes); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v notes", len(notes))
}

func (uc *GophKeeperClientUseCase) AddNote(userPassword string, note *entity.SecretNote) {
	accessToken, err := uc.authorisationCheck(userPassword)
	if err != nil {
		log.Fatalf("GophKeeperClientUseCase - AddNote - %v", err)
	}
	uc.encryptNote(userPassword, note)

	if err = uc.clientAPI.AddNote(accessToken, note); err != nil {
		log.Fatalf("GophKeeperClientUseCase - AddNote - %v", err)
	}

	uc.repo.AddNote(note)

	color.Green("Note %q added, id: %v", note.Name, note.ID)
}

func (uc *GophKeeperClientUseCase) ShowNote(userPassword, noteID string) {
	if !uc.verifyPassword(userPassword) {
		return
	}
	noteUUID, err := uuid.Parse(noteID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	note, err := uc.repo.GetNoteByID(noteUUID)
	if err != nil {
		color.Red(err.Error())

		return
	}

	uc.decryptNote(userPassword, &note)
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("ID: %s\nname:%s\nNote:%s\n", //nolint:forbidigo // cli printing
		yellow(note.ID),
		yellow(note.Name),
		yellow(note.Note),
	)
}

func (uc *GophKeeperClientUseCase) encryptNote(userPassword string, note *entity.SecretNote) {
	note.Note = utils.Encrypt(userPassword, note.Note)
}

func (uc *GophKeeperClientUseCase) decryptNote(userPassword string, note *entity.SecretNote) {
	note.Note = utils.Decrypt(userPassword, note.Note)
}

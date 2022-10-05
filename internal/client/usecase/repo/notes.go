package repo

import (
	"errors"

	"github.com/dimk00z/GophKeeper/internal/client/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/internal/client/usecase/viewsets"
	"github.com/dimk00z/GophKeeper/internal/entity"
	"github.com/google/uuid"
)

var errNoteNotFound = errors.New("note not found")

func (r *GophKeeperRepo) AddNote(note *entity.SecretNote) {
	noteForSaving := models.Note{
		ID:     note.ID,
		Name:   note.Name,
		Note:   note.Note,
		UserID: r.getUserID(),
	}
	r.db.Save(&noteForSaving)
}

func (r *GophKeeperRepo) LoadNotes() []viewsets.NoteForList {
	userID := r.getUserID()
	var notes []models.Note
	r.db.Where("user_id", userID).Find(&notes)
	if len(notes) == 0 {
		return nil
	}

	notesViewSet := make([]viewsets.NoteForList, len(notes))

	for index := range notes {
		notesViewSet[index].ID = notes[index].ID
		notesViewSet[index].Name = notes[index].Name
	}

	return notesViewSet
}

func (r *GophKeeperRepo) SaveNotes(notes []entity.SecretNote) error {
	if len(notes) == 0 {
		return nil
	}
	userID := r.getUserID()
	notesForDB := make([]models.Note, len(notes))
	for index := range notes {
		notesForDB[index].ID = notes[index].ID
		notesForDB[index].Name = notes[index].Name
		notesForDB[index].Note = notes[index].Note
		notesForDB[index].UserID = userID
	}

	return r.db.Save(notesForDB).Error
}

func (r *GophKeeperRepo) GetNoteByID(noteID uuid.UUID) (note entity.SecretNote, err error) {
	var noteFromDB models.Note
	if err = r.db.Find(&noteFromDB, noteID).Error; noteFromDB.ID == uuid.Nil || err != nil {
		return note, errNoteNotFound
	}

	note.ID = noteFromDB.ID
	note.Note = noteFromDB.Note
	note.Name = noteFromDB.Name

	return
}

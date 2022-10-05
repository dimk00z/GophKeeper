package repo

import (
	"github.com/dimk00z/GophKeeper/internal/client/usecase/repo/models"
	"github.com/dimk00z/GophKeeper/internal/client/usecase/viewsets"
)

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

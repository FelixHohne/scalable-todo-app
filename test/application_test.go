package test

import (
	"testing"
	"time"
	"todo-app/src"
)

func TestNoteStoreCreate(t *testing.T) {
	noteStore := src.CreateNotStore()
	// [] is not legal syntax in go.
	noteId := noteStore.CreateNote("", []string{})

	note, err := noteStore.GetNote(0)
	if err != nil {
		t.Fatal(err)
	}

	if noteId != 0 {
		t.Fail()
	}

	if note.Content != "" {
		t.Fail()
	}

	if noteStore.NextId != 1 {
		t.Fail()
	}

}

func TestNoteStoreInvalidGet(t *testing.T) {

	noteStore := src.CreateNotStore()

	_, err := noteStore.GetNote(0)

	if err == nil {
		t.Fail()
	}

	_, err2 := noteStore.GetNote(-1)

	if err2 == nil {
		t.Fail()
	}

}

func TestNoteUpdate(t *testing.T) {

	noteStore := src.CreateNotStore()
	noteId := noteStore.CreateNote("", []string{})
	note, _ := noteStore.GetNote(noteId)

	if note.Content != "" {
		t.Fail()
	}
	time.Sleep(100 * time.Millisecond)
	err := noteStore.UpdateNote(noteId, "Hello World", []string{})
	if err != nil {
		t.Fatal(err)
	}

	updatedNote, _ := noteStore.GetNote(noteId)

	if updatedNote.Content != "Hello World" {
		t.Fail()
	}

	isAfter := updatedNote.LastUpdatedAt.Nanosecond() > updatedNote.Created.Nanosecond()

	if !isAfter {
		t.Fail()
	}

}

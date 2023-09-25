package backend

import (
	"testing"
	"time"
)

func TestNoteStoreCreate(t *testing.T) {
	noteStore := createNotStore()
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

	noteStore := createNotStore()

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

	noteStore := createNotStore()
	noteId := noteStore.CreateNote("", []string{})
	note, _ := noteStore.GetNote(noteId)

	if note.Content != "" {
		t.Logf("Created note with empty content but found %s\n", note.Content)
		t.Fail()
	}
	time.Sleep(500 * time.Millisecond)
	err := noteStore.UpdateNote(noteId, "Hello World", []string{})
	if err != nil {
		t.Fatal(err)
	}

	updatedNote, _ := noteStore.GetNote(noteId)

	if updatedNote.Content != "Hello World" {
		t.Logf("Expected updated content to be Hello World, but found %s\n", updatedNote.Content)
		t.Fail()
	}

	//isAfter := updatedNote.LastUpdatedAt.Nanosecond() > updatedNote.Created.Nanosecond()
	//
	//if !isAfter {
	//	t.Logf("Expected that after update, last updated at is after created timestamp, however found last updated: %d, created: %d",
	//		updatedNote.LastUpdatedAt.Nanosecond(), updatedNote.Created.Nanosecond())
	//	t.Fail()
	//}

}

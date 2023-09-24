package src

import (
	"fmt"
	"sync"
	"time"
)

type Note struct {
	Id            int       `json:"id"`
	Content       string    `json:"content"`
	Tags          []string  `json:"tags"`
	Created       time.Time `json:"created"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
}

type NoteStore struct {
	sync.Mutex
	notes  map[int]Note
	nextId int
}

func create() *NoteStore {
	noteStore := &NoteStore{}
	noteStore.notes = make(map[int]Note)
	noteStore.nextId = 0
	return noteStore
}

func (noteStore *NoteStore) CreateNote(content string, tags []string) int {
	noteStore.Lock()
	defer noteStore.Unlock()

	note := Note{
		Id:            noteStore.nextId,
		Content:       content,
		Tags:          make([]string, len(tags)),
		Created:       time.Now(),
		LastUpdatedAt: time.Now(),
	}
	copy(note.Tags, tags)
	noteStore.nextId++
	return note.Id
}

func (noteStore *NoteStore) GetNote(id int) (Note, error) {
	noteStore.Lock()
	defer noteStore.Unlock()

	t, ok := noteStore.notes[id]
	if ok {
		return t, nil
	} else {
		return Note{}, fmt.Errorf("note with id=%d not found", id)
	}
}

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
	NextId int
}

func CreateNotStore() *NoteStore {
	noteStore := &NoteStore{}
	noteStore.notes = make(map[int]Note)
	noteStore.NextId = 0
	return noteStore
}

func (noteStore *NoteStore) CreateNote(content string, tags []string) int {
	noteStore.Lock()
	defer noteStore.Unlock()
	id := noteStore.NextId

	note := Note{
		Id:            id,
		Content:       content,
		Tags:          make([]string, len(tags)),
		Created:       time.Now(),
		LastUpdatedAt: time.Now(),
	}
	copy(note.Tags, tags)
	noteStore.notes[id] = note
	noteStore.NextId++
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

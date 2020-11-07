package storage

import (
	"encoding/json"
	"github.com/nanobox-io/golang-scribble"
	"log"
	"sync"

	"github.com/hackformissions/discipulador/model"
)

const (
	PERSON_STORE = "person"
)

type Store struct {
	Mu sync.RWMutex
	db *scribble.Driver

	discipulados map[string]*model.Person
}

func Init(path string) (*Store, error) {
	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New(path, nil)
	if err != nil {
		log.Printf("Unable to initialize storage. err=%s", err)
		return nil, err
	}

	// load existing data
	records, err := db.ReadAll(PERSON_STORE)
	d := make(map[string]*model.Person)
	if err != nil {
		log.Printf("No previous data. err=%s", err)
	} else {
		for _, p := range records {
			person := &model.Person{}
			if err := json.Unmarshal([]byte(p), person); err != nil {
				log.Printf("Unable to unmarshal data. err=%s", err)
			}
			d[person.ID] = person
		}
	}

	return &Store{
		db:           db,
		discipulados: d,
	}, nil
}

func (s *Store) Write(collection string, resource string, v interface{}) error {
	err := s.db.Write(collection, resource, v)
	if err != nil {
		log.Printf("Unable to write to Store. err=%s", err)
	} else {
		switch v.(type) {
		case *model.Person:
			s.Mu.Lock()
			defer s.Mu.Unlock()
			person := v.(*model.Person)
			s.discipulados[person.ID] = person
		}
	}
	return err
}

func (s *Store) UnsafeReadAllPersons() map[string]*model.Person {
	return s.discipulados
}

func (s *Store) Delete(collection, resource string) error {
	err := s.db.Delete(collection, resource)
	if err != nil {
		log.Printf("Unable to delete from Store. err=%s", err)
	} else {
		s.Mu.Lock()
		defer s.Mu.Unlock()
		switch collection {
		case PERSON_STORE:
			delete(s.discipulados, resource)
		}
	}
	return err
}

package storage

import (
	"github.com/hackformissions/discipulador/model"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	// storage
	db, err := Init("./test_data")
	if err != nil {
		t.Fatalf("Failure. err=%s", err)
	}

	p := &model.Person{
		ID:        time.Now().String(),
		Firstname: "T O",
	}
	db.Write(PERSON_STORE, p.ID, p)

	pp := db.UnsafeReadAllPersons()
	db.Mu.RLock()
	if len(pp) != 1 {
		t.FailNow()
	}
	db.Mu.RUnlock()

	db.Delete(PERSON_STORE, p.ID)
	pp = db.UnsafeReadAllPersons()
	if len(pp) != 0 {
		t.FailNow()
	}
}

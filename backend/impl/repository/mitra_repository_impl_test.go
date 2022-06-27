package repository

import (
	"FinalProject/utility"
	"testing"
	"os"
)

func TestMitraRepository_Login_MitraExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantEmail := "rezky@email.com"
	wantRole := "MITRA"

	mitraRepositoryImpl := NewMitraRepositoryImpl(db)

	record, err := mitraRepositoryImpl.Login("rezky@email.com", "123456")
	if err != nil {
		t.Fatal("Err:", err)
	}

	if record.Email != wantEmail {
		t.Fatalf("got %q want %q", record.Email, wantEmail)
	}

	if record.KategoriUser != wantRole {
		t.Fatalf("got %q want %q", record.KategoriUser, wantRole)
	}

	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}

func TestMitraRepository_Login_MitraNotExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantMessageError := "sql: no rows in result set"

	mitraRepositoryImpl := NewMitraRepositoryImpl(db)

	_, err := mitraRepositoryImpl.Login("rezky@email.com", "123456")
	if err != nil {
		if err.Error() != wantMessageError {
			t.Fatalf("got %q want %q", err.Error(), wantMessageError)
		}
	}

	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}

func TestMitraRepository_Login_WrongPaswordMitra(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}
	
	wantMessageError := "ERR_UNAUTHORIZED"

	mitraRepositoryImpl := NewMitraRepositoryImpl(db)
	
	_, err := mitraRepositoryImpl.Login("rezky@email.com", "12345sjds6")
	if err != nil {
		if err.Error() != wantMessageError {
			t.Fatalf("got %q want %q", err.Error(), wantMessageError)
		}
	}

	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}

package service

import (
	"FinalProject/impl/repository"
	"FinalProject/payload"
	"FinalProject/utility"
	"os"
	"testing"
)


func TestLoginService_MitraExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Error("Err:", err)
	}

	wantEmail := "rezky@email.com"
	wantRole := "MITRA"

	mitraServiceImpl := NewMitraServiceImpl(repository.NewMitraRepositoryImpl(db))

	response, err := mitraServiceImpl.Login(payload.LoginRequest{
		Email: "rezky@email.com",
		Password: "123456",
	})
	if err != nil {
		t.Fatal("Err:", err)
	}

	if response.Email != wantEmail {
		t.Fatalf("got %q want %q", response.Email, wantEmail)
	}

	if response.Role != wantRole {
		t.Fatalf("got %q want %q", response.Role, wantRole)
	}

	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}

func TestLoginService_MitraNotExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantMessageError := "ERR_NO_DATA_FOUND"

	mitraServiceImpl := NewMitraServiceImpl(repository.NewMitraRepositoryImpl(db))

	_, err := mitraServiceImpl.Login(payload.LoginRequest{
		Email: "rezkyyyyy@email.com",
		Password: "123456",
	})
	if err != nil {
		if err.Error() != wantMessageError {
			t.Fatalf("got %q want %q", err.Error(), wantMessageError)
		}
	}

	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}

func TestLoginService_WrongPasswordMitra(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantMessageError := "ERR_UNAUTHORIZED"

	mitraServiceImpl := NewMitraServiceImpl(repository.NewMitraRepositoryImpl(db))

	_, err := mitraServiceImpl.Login(payload.LoginRequest{
		Email:    "rezky@email.com",
		Password: "sdsjdskjdksjd",
	})
	if err != nil {
		if err.Error() != wantMessageError {
			t.Fatalf("got %q want %q", err.Error(), wantMessageError)
		}
	}
	
	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}
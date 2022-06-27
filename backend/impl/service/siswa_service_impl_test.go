package service

import (
	"FinalProject/impl/repository"
	"FinalProject/payload"
	"FinalProject/utility"
	"os"
	"testing"
)

func TestSiswaService_Login_SiswaExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantEmail := "denny@email.com"
	wantRole := "SISWA"

	siswaServiceImpl := NewSiswaServiceImpl(
		repository.NewSiswaRepositoryImpl(db),
		repository.NewBeasiswaSiswaRepositoryImpl(db),
	)

	response, err := siswaServiceImpl.Login(payload.LoginRequest{
		Email:    "denny@email.com",
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

func TestSiswaService_Login_SiswaNotExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantMessageError := "ERR_NO_DATA_FOUND"

	siswaServiceImpl := NewSiswaServiceImpl(
		repository.NewSiswaRepositoryImpl(db),
		repository.NewBeasiswaSiswaRepositoryImpl(db),
	)

	_, err := siswaServiceImpl.Login(payload.LoginRequest{
		Email:    "denny1@email.com",
		Password: "13456",
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

func TestSiswaService_Login_WrongPasswordSiswa(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantMessageError := "ERR_UNAUTHORIZED"

	siswaServiceImpl := NewSiswaServiceImpl(
		repository.NewSiswaRepositoryImpl(db),
		repository.NewBeasiswaSiswaRepositoryImpl(db))

	_, err := siswaServiceImpl.Login(payload.LoginRequest{
		Email:    "denny@email.com",
		Password: "123456-",
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

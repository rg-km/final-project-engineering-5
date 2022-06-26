package service

import (
	"FinalProject/impl/repository"
	"FinalProject/payload"
	"FinalProject/utility"
	"testing"
)

func TestLogin_SiswaExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Error("Err:", err)
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
		t.Error("Err:", err)
	}

	if response.Email != wantEmail {
		t.Errorf("got %q want %q", response.Email, wantEmail)
	}

	if response.Role != wantRole {
		t.Errorf("got %q want %q", response.Role, wantRole)
	}
}

func TestLogin_SiswaNotExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Error("Err:", err)
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
			t.Errorf("got %q want %q", err.Error(), wantMessageError)
		}
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Error("Err:", err)
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
			t.Helper()
			t.Errorf("got %q want %q", err.Error(), wantMessageError)
		}
	}
}

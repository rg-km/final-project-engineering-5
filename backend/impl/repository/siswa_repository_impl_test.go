package repository

import (
	"testing"
	"os"

	"FinalProject/utility"
)

func TestSiswaRepository_Login_SiswaExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantEmail := "sinulingga@email.com"
	wantRole := "SISWA"

	siswaRepositoryImpl := NewSiswaRepositoryImpl(db)

	record, err := siswaRepositoryImpl.Login("sinulingga@email.com", "123456")
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

func TestSiswaRepository_Login_SiswaNotExists(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}

	wantMessageError := "sql: no rows in result set"

	siswaRepositoryImpl := NewSiswaRepositoryImpl(db)

	_, err := siswaRepositoryImpl.Login("sinulinggaaa@email.com", "123456")
	if err != nil {
		if err.Error() != wantMessageError {
			t.Fatalf("got %q want %q", err.Error(), wantMessageError)
		}
	}

	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}

func TestSiswaRepository_Login_WrongPasswordSiswa(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		t.Fatal("Err:", err)
	}
	
	wantMessageError := "ERR_UNAUTHORIZED"

	siswaRepositoryImpl := NewSiswaRepositoryImpl(db)
	
	_, err := siswaRepositoryImpl.Login("sinulingga@email.com", "skjdkdks")
	if err != nil {
		if err.Error() != wantMessageError {
			t.Fatalf("got %q want %q", err.Error(), wantMessageError)
		}
	}

	if err := os.Remove("beasiswa.db"); err != nil {
		t.Fatal("Err:", err)
	}
}

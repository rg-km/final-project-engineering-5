## Final Project Engineering 5 - Backend

## Schema Database
https://dbdesigner.page.link/ZFAEXxamCFztcUUB7

## API Specification
https://www.postman.com/final-project-engineering-5/workspace/final-project-engineering-5

## Run Backend
1. Install Golang, lihat panduan: https://go.dev/doc/install
2. Clone repo ini.
3. Masuk ke folder <b>backend</b> (```cd backend/```)
4. Jalankan command: ```go mod tidy```
5. Jalankan server: ```go run main.go```
6. Lihat endpoint pada log.
7. Done.

## Run Test
Pada command line jalankan <b>test.sh</b>
```$ ./test.sh```
. Jika semua test berjalan dengan baik outputnya adalah:
```
=== RUN   TestMitraService_Login_MitraExists
--- PASS: TestMitraService_Login_MitraExists (0.14s)
=== RUN   TestMitraService_Login_MitraNotExists
--- PASS: TestMitraService_Login_MitraNotExists (0.11s)
=== RUN   TestMitraService_Login_WrongPasswordMitra
--- PASS: TestMitraService_Login_WrongPasswordMitra (0.10s)
=== RUN   TestSiswaService_Login_SiswaExists
--- PASS: TestSiswaService_Login_SiswaExists (0.10s)
=== RUN   TestSiswaService_Login_SiswaNotExists
--- PASS: TestSiswaService_Login_SiswaNotExists (0.05s)
=== RUN   TestSiswaService_Login_WrongPasswordSiswa
--- PASS: TestSiswaService_Login_WrongPasswordSiswa (0.15s)
PASS
ok      FinalProject/impl/service       0.651s
=== RUN   TestMitraRepository_Login_MitraExists
--- PASS: TestMitraRepository_Login_MitraExists (0.14s)
=== RUN   TestMitraRepository_Login_MitraNotExists
--- PASS: TestMitraRepository_Login_MitraNotExists (0.16s)
=== RUN   TestMitraRepository_Login_WrongPaswordMitra
--- PASS: TestMitraRepository_Login_WrongPaswordMitra (0.10s)
=== RUN   TestSiswaRepository_Login_SiswaExists
--- PASS: TestSiswaRepository_Login_SiswaExists (0.10s)
=== RUN   TestSiswaRepository_Login_SiswaNotExists
--- PASS: TestSiswaRepository_Login_SiswaNotExists (0.10s)
=== RUN   TestSiswaRepository_Login_WrongPasswordSiswa
--- PASS: TestSiswaRepository_Login_WrongPasswordSiswa (0.10s)
PASS
ok      FinalProject/impl/repository    0.708s
```

Untuk saat ini test yang dilakukan untuk method / API Login pada package <b>repository</b> dan <b>service</b>.

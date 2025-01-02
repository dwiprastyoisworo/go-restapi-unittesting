# Makefile untuk proyek Golang

# Variabel default
GO := go
MAIN_FILE := main.go

# Target default: jalankan aplikasi Go
run:
	$(GO) run $(MAIN_FILE)

# Menjalankan go mod tidy untuk membersihkan dependensi
tidy:
	$(GO) mod tidy

# Menjalankan unit tests dan menghasilkan laporan coverage
test:
	$(GO) test -cover ./... -coverprofile=coverage.out
	$(GO) tool cover -html=coverage.out -o coverage.html

# Membersihkan file-file yang dihasilkan oleh test coverage
clean:
	rm -f coverage.out coverage.html
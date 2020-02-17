package alltransaksi

// Repository :
type Repository interface {
	AddAllTransaksi() error
}

// Usecase :
type Usecase interface {
	AddAllTransaksi() error
}

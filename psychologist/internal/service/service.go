package service

// Service presents psychologist service
type Service interface {
	Clients()
	ClientAppoints()
	ClientTransferActivity()
}

func New(db *postgres.DB, dtb DTB) *Psychologist {
	return &Psychologist{db: db, dtb: dtb}
}

// Psychologist ...
type Psychologist struct {
	db  *postgres.DB
	dtb DTB
}

//DTB presents database repository
type DTB interface {
	GetClients()
	SetClientAppoints()
	SetClientTrasferActivity()
}

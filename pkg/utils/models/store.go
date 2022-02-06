package models

const (
	MasterURL = "postgres://mtuser:mt123@localhost:5432/test?sslmode=disable&connect_timeout=10"
)

type SqlSettings struct {
	MasterURL   string
	ReplicaURLS []string
	Seed        bool
}

func (s *SqlSettings) SetDefaults() {
	s.MasterURL = MasterURL
	s.ReplicaURLS = nil
}

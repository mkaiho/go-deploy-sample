package server

type Server interface {
	Init() error
	Run() error
}

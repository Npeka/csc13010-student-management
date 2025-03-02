package server

import "fmt"

func (s *Server) StartWorker() {
	kurl := fmt.Sprintf("%s:%d", s.cfg.Kafka.Host, s.cfg.Kafka.Port)
	go s.w.authWorker.Start(kurl)
	go s.w.stdWorker.Start(kurl)
}

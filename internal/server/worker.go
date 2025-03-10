package server

import "fmt"

func (s *Server) StartWorker() {
	kurl := fmt.Sprintf("%s:%d", s.cf.Kafka.Host, s.cf.Kafka.Port)
	go s.w.authWorker.Start(kurl)
	go s.w.stdWorker.Start(kurl)
	go s.w.notiWorker.Start(kurl)
	go s.w.auditWorker.Start(kurl)
}

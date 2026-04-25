package servers

import "sync"

type f func()

type Manager struct {
	wg   *sync.WaitGroup
	errs chan error
	fs   []f
}

func newManager() *Manager {
	return &Manager{
		wg:   &sync.WaitGroup{},
		errs: make(chan error),
	}
}

func (m *Manager) Add(f f) {
	m.fs = append(m.fs, f)
}

func (m *Manager) run(f f) {
	m.wg.Add(1)
	wf := func() {
		defer m.wg.Done()
		f()
	}
	go wf()
}

func (m *Manager) Run() {
	for _, f := range m.fs {
		m.run(f)
	}
	m.wg.Wait()
	close(m.errs)
}

func (s *FiberServer) startFiberServer(errs chan error) f {
	return func() {
		errs <- s.Run()
	}
}

func (s *FiberServer) exposeViaNgrok(errs chan error) f {
	return func() {
		if s.UseNgrok != "true" {
			return
		}
		ngrok := s.NewNgrokServer()
		errs <- ngrok.Run()
	}
}

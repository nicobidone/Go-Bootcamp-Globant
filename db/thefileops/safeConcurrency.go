package thefileops

import "sync"

// SafeFile needs a mutex and the file location to secure
type SafeFile struct {
	mu   sync.Mutex
	path string
}

// NewSafeFile inics a SafeFile struct with an incoming path
func (s *SafeFile) NewSafeFile(path string) {
	s.path = path
}

func (s *SafeFile) SafeCreateFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	CreateFile(s.path)
}

func (s *SafeFile) SafeDeleteFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	DeleteFile(s.path)
}

func (s *SafeFile) SafeReadFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	ReadFile(s.path)
}

func (s *SafeFile) SafeWriteFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	WriteFile(s.path)
}

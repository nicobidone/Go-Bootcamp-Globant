package thefileops

import "sync"

// SafeFile needs a mutex and the file location to secure
type SafeFile struct {
	mu   sync.Mutex
	path string
}

// NewSafeFile inics a SafeFile struct with an incoming path
func NewSafeFile(path string) *SafeFile {
	s := new(SafeFile)
	s.path = path
	return s
}

// SafeCreateFile applies mutex on a specific file
func (s *SafeFile) SafeCreateFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	CreateFile(s.path)
}

// SafeDeleteFile applies mutex on a specific file
func (s *SafeFile) SafeDeleteFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	DeleteFile(s.path)
}

// SafeReadFile applies mutex on a specific file
func (s *SafeFile) SafeReadFile() {
	s.mu.Lock()
	defer s.mu.Unlock()
	ReadFile(s.path)
}

// SafeWriteFile applies mutex on a specific file
func (s *SafeFile) SafeWriteFile(text string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	WriteFile(s.path, text)
}

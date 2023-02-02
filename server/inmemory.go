package server

type InmemoryStore struct {
	redirects map[string]string
}

func NewInmemoryStore(redirects map[string]string) Store {
	return &InmemoryStore{redirects}
}

func (s *InmemoryStore) GetByKey(short string) (url string, found bool) {
	url, found = s.redirects[short]
	return
}

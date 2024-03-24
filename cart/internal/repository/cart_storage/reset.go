package cartstorage

func (s *storage) Reset() {
	s.Lock()
	defer s.Unlock()

	cartStorage = map[int64]*Cart{}
}

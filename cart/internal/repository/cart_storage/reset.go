package cartstorage

func (cs *cartStorage) Reset() {
	cs.Lock()
	defer cs.Unlock()

	storage = map[int64]*Cart{}
}

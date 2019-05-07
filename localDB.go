package kurz

type localDB map[string]string

func (l localDB) Get(key []byte) []byte {

	formattedKey := string(key)
	target, exist := l[formattedKey]

	if !exist {
		return nil
	}

	return []byte(target)

}

func (l *localDB) Put(key, value []byte) {
	(*l)[string(key)] = string(value)
}

func newLocalDB() *localDB {
	regularMap := make(map[string]string)
	loc := localDB(regularMap)
	return &loc
}

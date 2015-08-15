package myjudge

//TODO implement Cartesian tree for caching

type Cacher struct {
	cachedPages map[string]WebPage
	cachedUsers map[string]User
}

var singleton *Cacher = nil

func getCacher() *Cacher {
	if singleton != nil {
		return singleton
	}
	singleton = &Cacher{}
	return singleton
}

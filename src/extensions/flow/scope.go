package flow

type Scope struct {
	parent *Scope
	values map[string]*ScopeEntry
}

func NewScope() *Scope {
	return &Scope{values: map[string]*ScopeEntry{}}
}

func (self *Scope) Create() *Scope {
	child := NewScope()
	child.parent = self
	return child
}

func (self Scope) HasLocal(key string) bool {
	_, exists := self.values[key]
	return exists
}

func (self Scope) Has(key string) bool {
	if self.HasLocal(key) {
		return true
	}

	if self.parent != nil {
		return self.parent.Has(key)
	}

	return false
}

func (self Scope) GetLocal(key string) *ScopeEntry {
	entry, _ := self.values[key]
	return entry
}

func (self Scope) Get(key string) *ScopeEntry {
	if !self.Has(key) {
		return nil
	}

	if self.HasLocal(key) {
		return self.GetLocal(key)
	}

	return self.parent.Get(key)
}

func (self *Scope) SetLocal(key string, value *ScopeEntry) {
	self.values[key] = value
}

func (self *Scope) Set(key string, value *ScopeEntry) {
	if self.HasLocal(key) {
		self.values[key] = value
		return
	}

	self.parent.Set(key, value)
}

func (self *Scope) Del(key string) {
	if !self.Has(key) {
		return
	}

	if self.HasLocal(key) {
		delete(self.values, key)
		return
	}

	self.parent.Del(key)
}

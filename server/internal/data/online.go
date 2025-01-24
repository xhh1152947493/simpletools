package data

import "sync"

type OnlineUser struct {
	Platform string
	Username string
	Password string
}

type OnlineUserMgr struct {
	mu    sync.RWMutex
	users map[string]*OnlineUser
}

func NewOnlineUserMgr() *OnlineUserMgr {
	return &OnlineUserMgr{
		users: make(map[string]*OnlineUser),
	}
}

func hash(username, platform string) string {
	return username + "|" + platform
}

func (m *OnlineUserMgr) Add(user *OnlineUser) {
	if user == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.users[hash(user.Username, user.Platform)] = user
}

func (m *OnlineUserMgr) Del(user *OnlineUser) {
	if user == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.users, hash(user.Username, user.Platform))
}

func (m *OnlineUserMgr) Get(username, platform string) *OnlineUser {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.users[hash(username, platform)]
}

func (m *OnlineUserMgr) GetAll() []*OnlineUser {
	m.mu.RLock()
	defer m.mu.RUnlock()
	users := make([]*OnlineUser, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users
}

func (m *OnlineUserMgr) GetByPlatform(platform string) []*OnlineUser {
	m.mu.RLock()
	defer m.mu.RUnlock()
	users := make([]*OnlineUser, 0, len(m.users))
	for _, user := range m.users {
		if user.Platform == platform {
			users = append(users, user)
		}
	}
	return users
}

func (m *OnlineUserMgr) GetByUser(user string) []*OnlineUser {
	m.mu.RLock()
	defer m.mu.RUnlock()
	users := make([]*OnlineUser, 0, len(m.users))
	for _, u := range m.users {
		if u.Username == user {
			users = append(users, u)
		}
	}
	return users
}

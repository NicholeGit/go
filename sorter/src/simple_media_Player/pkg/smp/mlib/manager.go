package library

import "errors"

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, a := range m.musics {
		if a.Name == name {
			return &a
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	// Remove the found item from the slice.
	if index < len(m.musics)-1 { // 中间元素 Element between first and last
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { //删除仅有一个元素
		m.musics = make([]MusicEntry, 0)
	} else { //删除最后一个元素
		m.musics = m.musics[:index-1]
	}

	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for i, v := range m.musics {
		if v.Name == name {
			return m.Remove(i)
		}
	}
	return nil
}

package main

type GuestBookEntry struct {
	Id      int
	Email   string
	Title   string
	Content string
}

type GuestBook struct {
	guestBookData []*GuestBookEntry
}

func NewGuestBook() *GuestBook

func (g *GuestBook) GetEntry(id int) (*GuestBookEntry, error)

func (g *GuestBook) GetAllEntries() []*GuestBookEntry

func (g *GuestBook) RemoveAllEntries()

package server

/*
Typy wiadomości przychodzącej
	0 - startgame
	1 - Pozycja pada
	2 - Pozycja piłki
*/
type Message struct {
	MessageType int
	Data        map[string]interface{}
}

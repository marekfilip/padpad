package message

const (
	STARTGAME_TYPE     int = 0
	PAD_POSITION_TYPE  int = 1
	BALL_POSITION_TYPE int = 2
)

type Message struct {
	MessageType int                `json:"t"`
	Data        map[string]float32 `json:"d"`
}

package message

const (
	STARTGAME_TYPE             int = 1
	INC_PAD_POSITION_TYPE      int = 2
	BALL_POSITION_TYPE         int = 3
	PLAYER_PAD_POSITION_TYPE   int = 4
	OPPONENT_PAD_POSITION_TYPE int = 5
)

type Message struct {
	MessageType int                `json:"t"`
	Data        map[string]float32 `json:"d"`
}

package chess

var Symbols = map[string]map[string]string{
	"white": map[string]string{
		"king":   "♔",
		"queen":  "♕",
		"rook":   "♖",
		"bishop": "♗",
		"knight": "♘",
		"pawn":   "♙",
	},
	"black": map[string]string{
		"king":   "♚",
		"queen":  "♛",
		"rook":   "♜",
		"bishop": "♝",
		"knight": "♞",
		"pawn":   "♟",
	},
}

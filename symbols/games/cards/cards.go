package cards

//  "heart": map[string]string{
//    "ace": "🂱",
//     🂲 🂳 🂴 🂵 🂶 🂷 🂸 🂹 🂺
//  🂱 🂲 🂳 🂴 🂵 🂶 🂷 🂸 🂹 🂺
// 🂻 🂼 🂽 🂾
// 🂡 🂢 🂣 🂤 🂥 🂦 🂧 🂨 🂩 🂪
// 🂫 🂬 🂭 🂮
// 🃁 🃂 🃃 🃄 🃅 🃆 🃇 🃈 🃉 🃊
// 🃋 🃌 🃍 🃎
// 🃑 🃒 🃓 🃔 🃕 🃖 🃗 🃘 🃙 🃚
// 🃛 🃜 🃝 🃞
// 🂠 🃏 🃟
var Symbols = map[string]map[string]string{
	"spade": map[string]string{
		"empty":  "♣",
		"filled": "♠",
	},
	"heart": map[string]string{
		"empty":  "♡",
		"filled": "♡",
	},
	"diamond": map[string]string{
		"empty":  "♢",
		"filled": "",
	},
	"club": map[string]string{
		"empty":  "♣",
		"filled": "♣",
	},
}

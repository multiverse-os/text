package cli

// A collection of UTF-8 symbols that work by default in Gnome terminal on
// Debian AND that are specifically useful for UI design.

// https://en.wikipedia.org/wiki/Miscellaneous_Symbols
// https://www.w3schools.com/charsets/ref_html_utf8.asp

// TODO: This will be moved into its own package and called in optionally
var Settings = map[string]string{
	"brightness": "☼",
}

// https://github.com/golang/text/blob/master/unicode/norm/data12.0.0_test.go
var Symbols = map[string]string{
	"copyright":             "© ",
	"registered":            "® ",
	"check":                 "✔",
	"micro":                 "µ",
	"emphasized_x":          "✘",
	"light_x":               "✕",
	"star":                  "★",
	"exclamation_mark":      "❗",
	"question_mark":         "❓",
	"quote_open":            "❝",
	"quote_close":           "❞",
	"square":                "▇",
	"circle_double":         "◎",
	"circle_circle":         "ⓞ",
	"circled_x":             "ⓧ",
	"circled_pipe":          "Ⓘ",
	"circled_question_mark": "?⃝",
	"underscore":            "─",
	"ellipsis":              "…",
	"pointer":               "❯",
	"pointer_small":         "›",
	"hamburger":             "☰",
	"smirk":                 "㋡",
	"mustache":              "෴",
	"heart":                 "♥",
	"section":               "§",
	"letter":                "✉",
	"paragraph":             "¶",
	"tape_drive":            "✇",
	"double_bang":           "‼",
	"crescent_moon":         "☾",
	"arrow_outline_left":    "◁",
	"arrow_outline_right":   "▷",
	"space":                 "␣",
	"bold":                  "␢",
	"missing_symbol":        "�",
	"sissors":               "✂",
	"sissors_outline":       "✄",
}

// ⋈ ⧓

// Could be useful for indicating corners
//  ◰ ◱ ◲ ◳

// Stacked Lines  ▤  𝌆

// Editing  ⁀ ⎁ ⎂ ⎃

// White space representation
//tab ↹ ⇄ ⇤ ⇥ ↤ ↦
//space ·  ˽

// Various
// ^ ⌃
// ✲
//  ⎇ ⌥ ✦ ✧   ⌤
// ⎋ ⌫  ⌦ ⎀
// ⌧  ⇞ ⇟   ⎉ ⎊ ⍰
// ⌂

var ReloadSymbols = []string{"↶", "↷", "⟲", "⟳", "↺", "↻"}

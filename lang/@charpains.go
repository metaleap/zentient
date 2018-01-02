package ℤ

// not imported anywhere, just for live-testing potential unicode pains (none exist as of right now, for strs that DO render on MY system =)

type _tmpTestFoo struct {
	A, B, C string
}

func init() {
	i1, i2, i3 := 1, 2, 1+2
	println(i1 + i2 + i3)
	foo := append([]string{"hello"}, "world")
	for _, f := range foo {
		vs := _tmpTestFoo{A: "⟨ℤ⟩" + f + "\t"}
		if f != "" && vs.B != "▶▶▶" {
			print("❬" + f + "❭\n" + f + "➜")
			// something i dont even grok 世界世界世界 moar of same:  世界世界世界
			vs.C = "❗" + vs.A + "▶▶▶" + vs.B
		}
	}
}

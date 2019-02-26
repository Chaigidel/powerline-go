package main

func segmentRoot(p *powerline) {
	var foreground, background uint8
	var content string
	if *p.args.PrevError == 0 {
		foreground = p.theme.CmdPassedFg
		background = p.theme.CmdPassedBg
		content    = `(っ･ω･)っ`
	} else {
		foreground = p.theme.CmdFailedFg
		background = 3 // p.theme.CmdFailedBg
		content    = `(｡•́︿•̀｡) `
	}

	p.appendSegment("root", segment{
		content:       content,
		foreground:    foreground,
		background:    background,
		noEndingSpace: true,
	})
}

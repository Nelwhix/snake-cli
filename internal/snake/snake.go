package snake

type Body struct {
	Parts  []Part
	Xspeed int
	Yspeed int
}

func NewBody(parts []Part, xSpeed, ySpeed int) *Body {
	return &Body{
		Parts:  parts,
		Xspeed: xSpeed,
		Yspeed: ySpeed,
	}
}

func (sb *Body) ChangeDir(vertical, horizontal int) {
	sb.Xspeed = horizontal
	sb.Yspeed = vertical
}

func (sb *Body) Update(width, height int) {
	sb.Parts = append(sb.Parts, sb.Parts[len(sb.Parts)-1].GetUpdatedPart(sb, width, height))
	sb.Parts = sb.Parts[1:]
}

type Part struct {
	X int
	Y int
}

func (p *Part) GetUpdatedPart(sb *Body, width, height int) Part {
	newPart := *p
	newPart.X = (newPart.X + sb.Xspeed) % width
	if newPart.X < 0 {
		newPart.X += width
	}

	newPart.Y = (newPart.Y + sb.Yspeed) % height
	if newPart.Y < 0 {
		newPart.Y += height
	}

	return newPart
}

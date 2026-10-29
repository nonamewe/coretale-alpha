package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type gamestate int

const (
	statestart gamestate = iota
	statelobby
	statemove
	statefight
)
const (
	Titlehotland0 = 0
	Titlehotland1 = 1
	Titlehotland2 = 2
	Titlehotland3 = 3
)

type room struct {
	roomNAME string
	Tiles    [][]int
}

var openRoom = room{
	roomNAME: "startroom",
	Tiles: [][]int{
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{3, 3, 3, 3, 3, 3, 3, 3},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{1, 1, 1, 1, 1, 1, 1, 1},
	},
}

type Game struct {
	State       gamestate
	Selected    int
	Xplayer     int
	Yplayer     int
	player      *ebiten.Image
	sansheadd   *ebiten.Image
	HPpalyer    int
	LVLplayer   int
	KILLSplayer int
	NAMEplayer  string
	hotland0    *ebiten.Image
	hotland1    *ebiten.Image
	hotland2    *ebiten.Image
	hotland3    *ebiten.Image
}

func loadImage(path string) *ebiten.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
func loadImageSans(path string) *ebiten.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sansheadd, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(sansheadd)
}

func (g *Game) Updatestart() {
	for _, char := range ebiten.AppendInputChars(nil) {
		if len([]rune(g.NAMEplayer)) < 10 {
			g.NAMEplayer += string(char)

		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {

		runes := []rune(g.NAMEplayer)

		if len(runes) > 0 {
			g.NAMEplayer = string(runes[:len(runes)-1])
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		runes := []rune(g.NAMEplayer)

		if len(runes) > 0 {
			g.State = statemove
		}
	}

}
func (g *Game) Updatemovement() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Xplayer += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Xplayer -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.Yplayer -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Yplayer += 3
	}

	return nil
}

func (g *Game) Update() error {
	switch g.State {
	case statestart:
		g.Updatestart()
	case statemove:
		return g.Updatemovement()
	}
	return nil
}

// explain
func (g *Game) DrawRoom(screen *ebiten.Image) {
	tiles := []*ebiten.Image{
		g.hotland0,
		g.hotland1,
		g.hotland2,
		g.hotland3,
	}

	const tileSize = 32

	for y, row := range openRoom.Tiles {
		for x, tile := range row {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64(x*tileSize),
				float64(y*tileSize),
			)

			screen.DrawImage(tiles[tile], op)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	//	xwidth := screen.Bounds().Dx()
	//	yheight := screen.Bounds().Dy()
	switch g.State {

	case statestart:
		if g.NAMEplayer == "sans" {
			sansop := &ebiten.DrawImageOptions{}
			sansop.GeoM.Scale(2, 2)
			w := g.sansheadd.Bounds().Dx() * 2
			h := g.sansheadd.Bounds().Dy() * 2

			sansop.GeoM.Translate(
				float64((640-w)/2),
				float64((480-h)/2),
			)
			screen.DrawImage(g.sansheadd, sansop)

		}
		ebitenutil.DebugPrintAt(screen, "NAME: "+g.NAMEplayer, 20, 20)
	case statemove:
		g.DrawRoom(screen)

		playerop := &ebiten.DrawImageOptions{}
		playerop.GeoM.Scale(0.03, 0.03)
		playerop.GeoM.Translate(
			float64(g.Xplayer),
			float64(g.Yplayer),
		)
		screen.DrawImage(g.player, playerop)
	}

}

func (g *Game) Layout(
	outsideWidth, outsideHeight int,
) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowTitle("coretale")

	game := &Game{
		State:     statestart,
		player:    loadImage("assets/goferimage.png"),
		sansheadd: loadImage("assets/sanshead.png"),
		Xplayer:   100,
		Yplayer:   100,
		hotland0:  loadImage("assets/0roadmap(hotland).png"),
		hotland1:  loadImage("assets/1roadmap(hotland).png"),
		hotland2:  loadImage("assets/2roadmap(hotland).png"),
		hotland3:  loadImage("assets/3roadmap(hotland).png"),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct{}

var tilt int

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if tilt > -3 {
			tilt--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if tilt < 3 {
			tilt++
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	bgImage := ebiten.NewImage(640, 480)
	bgImage.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	var eimg *ebiten.Image
	switch tilt {
	case -3:
		// load image
		img, _, err := ebitenutil.NewImageFromFile("./left3.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case -2:
		// load image
		img, _, err := ebitenutil.NewImageFromFile("./left2.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case -1:
		// load image
		img, _, err := ebitenutil.NewImageFromFile("./left1.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case 0:
		// load image
		img, _, err := ebitenutil.NewImageFromFile("./straight (1).png")
		if err != nil {
			log.Fatal("what", err)
		}
		eimg = img
	case 1:
		// load image
		img, _, err := ebitenutil.NewImageFromFile("./right1.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case 2:
		// load image
		img, _, err := ebitenutil.NewImageFromFile("./right2.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case 3:
		// load image
		img, _, err := ebitenutil.NewImageFromFile("./right3.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	}

	op := &ebiten.DrawImageOptions{}

	//dRAW IMAGE
	fmt.Println(tilt)
	screen.DrawImage(bgImage, op)
	if eimg != nil {
		screen.DrawImage(eimg, op)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

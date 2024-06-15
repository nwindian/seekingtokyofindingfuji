package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game implements ebiten.Game interface.
type Game struct {
	counter int
}

var tilt int
var roadCount = 1

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if tilt > -4 {
			tilt--
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if tilt < 4 {
			tilt++
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	var eimg *ebiten.Image
	switch tilt {
	case -4:
		img, _, err := ebitenutil.NewImageFromFile("./left_crash.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case 4:
		img, _, err := ebitenutil.NewImageFromFile("./right_crash.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case -3:
		img, _, err := ebitenutil.NewImageFromFile("./left3.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case -2:
		img, _, err := ebitenutil.NewImageFromFile("./left2.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case -1:
		img, _, err := ebitenutil.NewImageFromFile("./left1.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case 0:
		img, _, err := ebitenutil.NewImageFromFile("./straight.png")
		if err != nil {
			log.Fatal("what", err)
		}
		eimg = img
	case 1:
		img, _, err := ebitenutil.NewImageFromFile("./right1.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case 2:
		img, _, err := ebitenutil.NewImageFromFile("./right2.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	case 3:

		img, _, err := ebitenutil.NewImageFromFile("./right3.png")
		if err != nil {
			log.Fatal(err)
		}
		eimg = img
	}

	//dRAW IMAGE
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(0.5, 1)
	opts.GeoM.Translate(0, 175)

	var road string
	var sky string
	g.counter++
	if g.counter%30 == 0 {
		if roadCount == 3 {
			roadCount = 1
		} else {
			roadCount++
		}
		road = fmt.Sprintf("./theRoad%d.png", roadCount)
		sky = fmt.Sprintf("./theSky%d.png", roadCount)
	} else {
		road = fmt.Sprintf("./theRoad%d.png", roadCount)
		sky = fmt.Sprintf("./theSky%d.png", roadCount)
	}

	img, _, err := ebitenutil.NewImageFromFile(sky)
	if err != nil {
		log.Fatal(err)
	}

	skyopts := &ebiten.DrawImageOptions{}
	skyopts.GeoM.Scale(.5, 1)
	screen.DrawImage(img, skyopts)

	img, _, err = ebitenutil.NewImageFromFile(road)
	if err != nil {
		log.Fatal(err)
	}
	screen.DrawImage(img, opts)

	if eimg != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 130)
		screen.DrawImage(eimg, op)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 450
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 900)
	ebiten.Monitor().Size()
	ebiten.SetWindowTitle("Seeking Tokyo Finding Fuji")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

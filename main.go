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

const (
	screenWidth  = 640
	screenHeight = 900
)

var tilt int
var roadCount = 1
var skyCount = 1
var crashed = false
var crashedFade = 0
var started = false

// Used for distance
const FUJI_DISTANCE = 5000

var currentDistance = 0.0

// Used for determining acceleration and speed
const PIXELS_PER_BANANA = 10
const TICKS_PER_SECOND = 60
const ACCELERATION_DECAY = 0.01
const SPEED_DECAY = 0.01

var acceleration = 0.0
var speed = 0.0
var ticksSinceLastPress = 0
var bananasPerSecond = 0.0
var bananasPerSecond2 = 0.0

func UpdateDistance() {
	bananasPerTick := speed / PIXELS_PER_BANANA
	currentDistance += bananasPerTick
}

func UpdateAcceleration() {
	rate := float64(TICKS_PER_SECOND / ticksSinceLastPress)
	acceleration = rate * 0.001
}

func UpdateSpeed() {
	speed += acceleration
	speed = max(0, speed)
}

func Decelerate() {
	if acceleration > 0 {
		acceleration -= ACCELERATION_DECAY
	} else {
		acceleration = 0
	}

	if speed > 0 {
		speed -= SPEED_DECAY
	} else {
		speed = 0
	}
}

func UpdateBPS() {
	bananasPerTick := speed / PIXELS_PER_BANANA
	bananasPerSecond = bananasPerTick * TICKS_PER_SECOND
}

func UpdateBPS2() {
	bananasPerTick2 := acceleration / PIXELS_PER_BANANA
	bananasPerSecond2 = bananasPerTick2 * TICKS_PER_SECOND * TICKS_PER_SECOND
}

func UpdateBalance() {
	tiltRate := 60 - min(55, int(bananasPerSecond)/3)

	if ticksSinceLastPress%tiltRate == 0 {
		if tilt < 0 {
			tilt -= 1
		} else if tilt > 0 {
			tilt += 1
		}
	}
}

func ShowReset() {

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	if !crashed && started {
		leftPressed := inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft)
		rightPressed := inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight)
		straightPressed := inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp)

		if leftPressed {
			if tilt == 1 {
				tilt = -1
			} else {
				tilt--
			}
		}
		if rightPressed {
			if tilt == -1 {
				tilt = 1
			} else {
				tilt++
			}
		}
		if straightPressed {
			if tilt == -1 || tilt == 1 {
				tilt = 0
			}
		}

		if leftPressed && rightPressed {
			Crash()
		} else if leftPressed != rightPressed {
			UpdateAcceleration()
			ticksSinceLastPress = 0

		} else {
			UpdateBalance()
		}

		if tilt == -4 || tilt == 4 {
			Crash()
		}

		UpdateSpeed()
		ticksSinceLastPress++
		if ticksSinceLastPress > 60 {
			Decelerate()
		}

		UpdateBPS()
		UpdateBPS2()
		UpdateDistance()
	} else if started {
		crashedFade++
		if inpututil.IsKeyJustPressed(ebiten.KeyR) && crashedFade >= 60 {
			reset()
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		started = true
	}

	return nil
}

func reset() {
	currentDistance = 0
	tilt = 0
	crashedFade = 0
	crashed = false
}

func Crash() {
	speed = 0
	acceleration = 0
	crashed = true
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	var eimg *ebiten.Image
	var middleText *ebiten.Image
	var titleScreen *ebiten.Image

	if crashed {
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
		default:
			img, _, err := ebitenutil.NewImageFromFile("./straight_crash.png")
			if err != nil {
				log.Fatal(err)
			}
			eimg = img
		}
		if crashedFade >= 60 {
			eimg = nil
			if g.counter%60 < 30 {
				img, _, err := ebitenutil.NewImageFromFile("./restart.png")
				if err != nil {
					log.Fatal(err)
				}
				middleText = img
			}
		}
	} else if started {
		switch tilt {
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
	} else {
		img, _, err := ebitenutil.NewImageFromFile("./title.png")
		if err != nil {
			log.Fatal(err)
		}
		titleScreen = img
	}

	//dRAW IMAGE
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(0.5, 1)
	opts.GeoM.Translate(0, 175)

	var road string
	var sky string

	roadRate := 60 - min(59, int(bananasPerSecond)/2)

	g.counter++
	if speed != 0 && g.counter%roadRate == 0 {
		if roadCount == 3 {
			roadCount = 1
		} else {
			roadCount++
		}
	}

	if g.counter%30 == 0 {
		if skyCount == 3 {
			skyCount = 1
		} else {
			skyCount++
		}
	}

	road = fmt.Sprintf("./theRoad%d.png", roadCount)
	sky = fmt.Sprintf("./theSky%d.png", skyCount)

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

	if middleText != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 150)
		op.GeoM.Scale(.25, .25)
		screen.DrawImage(middleText, op)
	}

	if titleScreen != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 50)
		op.GeoM.Scale(.25, .25)
		screen.DrawImage(titleScreen, op)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Speed: %.2f Bananas / Sec\nAcceleration: %.2f Bananas / Sec^2\nDistance: %.2f Bananas\nCompleted: %.1f%%\n", bananasPerSecond, bananasPerSecond2, currentDistance, (currentDistance*100)/FUJI_DISTANCE))
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 450
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.Monitor().Size()
	ebiten.SetWindowTitle("Seeking Tokyo Finding Fuji")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

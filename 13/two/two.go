package two

import (
	"fmt"

	"../computer"
)

const empty = 0
const wall = 1
const block = 2
const paddle = 3
const ball = 4

const left = -1
const neutral = 0
const right = 1

var glyphs map[int]string

func getJoystick(toMove int) int {
	if toMove < 0 {
		return left
	}
	if toMove > 0 {
		return right
	}
	return neutral
}

func printGame(tiles map[pos]int) {
	for y := 0; y < 20; y++ {
		for x := 0; x < 44; x++ {
			fmt.Print(glyphs[tiles[[2]int{x, y}]])
		}
		fmt.Println(" ", y)
	}
	for x := 0; x < 44; x++ {
		fmt.Print(x / 10)
	}
	fmt.Println()
	for x := 0; x < 44; x++ {
		fmt.Print(x % 10)
	}
	fmt.Println()
}

type velocity [2]int
type pos [2]int

func getNextPos(p pos, v velocity) pos {
	return pos{p[0] + v[0], p[1] + v[1]}
}

func copyTiles(tiles map[pos]int) map[pos]int {
	newtiles := make(map[pos]int)
	for k, v := range tiles {
		newtiles[k] = v
	}
	return newtiles
}

func getNextVelocity(tiles map[pos]int, ballPos pos, ballVel velocity) velocity {
	vx, vy := ballVel[0], ballVel[1]
	nextPos := getNextPos(ballPos, ballVel)
	nextx := pos{ballPos[0] + vx, ballPos[1]}
	if tiles[nextx] != empty && tiles[nextx] != ball {
		vx = -vx
	}
	nexty := pos{ballPos[0], ballPos[1] + vy}
	if tiles[nexty] != empty && tiles[nexty] != ball {
		vy = -vy
	}
	if ballVel[0] == vx && ballVel[1] == vy && tiles[nextPos] != empty && tiles[nextPos] != ball {
		vx = -vx
		vy = -vy
	}
	if tiles[nextx] == empty && tiles[nexty] == empty && tiles[nextPos] == block {
		tiles[nextPos] = empty
	} else {
		if tiles[nextx] == block {
			tiles[nextx] = empty
		}
		if tiles[nexty] == block {
			tiles[nexty] = empty
		}
	}
	return velocity{vx, vy}
}

func findNextPaddleSpot(tiles map[pos]int, ballPos pos, ballVel velocity) int {
	for {
		if ballPos[1] > 17 || (ballPos[1] == 17 && ballVel[1] == 1) {
			break
		}
		ballVel = getNextVelocity(tiles, ballPos, ballVel)
		ballPos = getNextPos(ballPos, ballVel)
	}
	return ballPos[0]
}

func calcVelocity(posA pos, posB pos) velocity {
	v := velocity{0, 0}
	if posA[0] > posB[0] {
		v[0] = -1
	} else if posA[0] < posB[0] {
		v[0] = 1
	}
	if posA[1] > posB[1] {
		v[1] = -1
	} else if posA[1] < posB[1] {
		v[1] = 1
	}
	return v
}

func arcade(instructions chan int, move chan int, done chan bool) {
	var x, y, tid int
	var ok bool
	tiles := make(map[pos]int)
	blocks := 0
	score := 0
	paddlex := 0
	ballPos := pos{0, 0}
	ballVelocity := velocity{1, 1}
	joystick := 0
	goalx := 22
	toMove := 0
	playing := false
	for {
		select {
		case x, ok = <-instructions:
			if !ok {
				fmt.Println("Final Score:", score, "blocks", blocks)
				close(done)
				return
			}
			y = <-instructions
			tid = <-instructions
		default:
			x = -2
		}
		if x == -2 {

		} else if x == -1 && y == 0 {
			score = tid
			playing = true
		} else {
			p := pos{x, y}
			if tid == block {
				blocks++
			} else if tiles[p] == block {
				blocks--
			}
			tiles[p] = tid
			if tid == paddle {
				paddlex = x
				toMove = goalx - paddlex
				//printGame(tiles)
			}
			if tid == ball {
				ballVelocity = calcVelocity(ballPos, p)
				ballPos = p
				goalx = findNextPaddleSpot(copyTiles(tiles), p, ballVelocity)
				if paddlex > 0 {
					toMove = goalx - paddlex
				}
				//printGame(tiles)
			}
		}
		if playing {
			joystick = getJoystick(toMove)
			select {
			case move <- joystick:
				toMove -= joystick
			default:
			}
		}
	}
}

// Run is the entry point for this solution.
func Run() {
	fmt.Println("Part Two")
	glyphs = make(map[int]string)
	glyphs[wall] = "#"
	glyphs[block] = "*"
	glyphs[ball] = "o"
	glyphs[paddle] = "-"
	glyphs[empty] = " "
	game := computer.LoadProgram()
	game[0] = 2 // Free play!
	io := make(chan int)
	joy := make(chan int)
	done := make(chan bool)
	go game.Execute(joy, io)
	go arcade(io, joy, done)
	<-done
}

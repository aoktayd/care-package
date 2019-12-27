package main

import (
	"strconv"

	"github.com/gdamore/tcell"
)

const (
	empty int = iota
	wall
	block
	paddle
	ball
)

var tiles = map[int]rune{
	empty:  ' ',
	wall:   '█',
	block:  '█',
	paddle: '▀',
	ball:   '●',
}

var background = tcell.StyleDefault.Background(tcell.ColorAntiqueWhite)
var styles = map[int]tcell.Style{
	empty:  background,
	wall:   tcell.StyleDefault.Foreground(tcell.ColorDimGrey),
	block:  tcell.StyleDefault.Foreground(tcell.ColorDodgerBlue),
	paddle: background.Foreground(tcell.ColorCornflowerBlue),
	ball:   background.Foreground(tcell.ColorCornflowerBlue),
}

type tile struct {
	x, y  int
	tile  rune
	style tcell.Style
}

const width, height = 45, 23

type game struct {
	grid             [height][width]tile
	score            int
	paddleDir        int
	ballX, ballY     int
	paddleX, paddleY int
}

func updateGame(game *game, data []int) {
	for i := 0; i < len(data); i += 3 {
		x := data[i]
		y := data[i+1]
		z := data[i+2]

		if x == -1 && y == 0 {
			game.score = z
			continue
		}

		if z == ball {
			game.ballX, game.ballY = x, y
		} else if z == paddle {
			game.paddleX, game.paddleY = x, y
		}

		game.grid[y][x] = tile{
			x:     x,
			y:     y,
			tile:  tiles[z],
			style: styles[z],
		}
	}

	sc := strconv.Itoa(game.score)
	x := width/2 - len(sc)/2
	for _, v := range sc {
		game.grid[game.paddleY+1][x] = tile{
			x:     x + 1,
			y:     game.paddleY + 1,
			tile:  v,
			style: styles[paddle],
		}
		x++
	}
}

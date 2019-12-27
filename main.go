package main

import (
	"fmt"
	"time"

	"github.com/aoktayd/intgode"
	"github.com/gdamore/tcell"
)

func main() {
	intcode := intcodeStdin()
	intcode[0] = 2
	program := intgode.NewIntcodeProgram(intcode)

	go program.Exec()
	game := &game{}
	updateGame(game, <-program.Output())

	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	} else if err = s.Init(); err != nil {
		panic(err)
	}

	quit := make(chan string)
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					quit <- "User has quit the game."
				}
			}
		}
	}()

	start := make(chan struct{})
	go func() {
		for y, row := range game.grid {
			for x, cell := range row {
				time.Sleep(time.Microsecond * 555)
				s.SetContent(x, y, cell.tile, nil, cell.style)
				s.Show()
			}
		}
		for x := 0; x < width; x++ {
			s.SetContent(x, game.paddleY+2, tiles[wall], nil, styles[wall])
			s.Show()
		}
		close(start)
	}()

	go func() {
		<-start

		for {
			time.Sleep(time.Second / 60)

			if program.Halted() {
				quit <- "Game over!"
				return
			}

			switch {
			case game.ballX == game.paddleX:
				game.paddleDir = 0
			case game.ballX > game.paddleX:
				game.paddleDir = 1
			case game.ballX < game.paddleX:
				game.paddleDir = -1
			}

			program.Input() <- game.paddleDir
			updateGame(game, <-program.Output())

			s.Clear()
			for y, row := range game.grid {
				for x, cell := range row {
					s.SetContent(x, y, cell.tile, nil, cell.style)
				}
			}
			for x := 0; x < width; x++ {
				s.SetContent(x, game.paddleY+2, tiles[wall], nil, styles[wall])
			}
			s.Show()
		}
	}()

	msg := <-quit
	s.Fini()
	if len(msg) > 0 {
		fmt.Println(msg)
	}
	fmt.Printf("Score: %d\n", game.score)
}

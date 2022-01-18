package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func constructcube() (newcube [3][3][3][6]string) {
	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {
					newcube[d][h][w][f] = " "
					switch f {
					case 0:
						if d == 2 {
							newcube[d][h][w][f] = "m"
						}
					case 1:
						if w == 0 {
							newcube[d][h][w][f] = "y"
						}
					case 2:
						if d == 0 {
							newcube[d][h][w][f] = "r"
						}
					case 3:
						if w == 2 {
							newcube[d][h][w][f] = "c"
						}
					case 4:
						if h == 0 {
							newcube[d][h][w][f] = "b"
						}
					case 5:
						if h == 2 {
							newcube[d][h][w][f] = "g"
						}
					}
				}
			}
		}
	}
	return newcube
}

func printcolor(currcolor string) {

	b := color.New(color.FgBlue)

	r := color.New(color.FgRed)

	g := color.New(color.FgGreen)

	y := color.New(color.FgYellow)

	m := color.New(color.FgMagenta)

	c := color.New(color.FgCyan)

	switch currcolor {
	case "b":
		b.Print(currcolor)
	case "g":
		g.Print(currcolor)
	case "r":
		r.Print(currcolor)
	case "m":
		m.Print(currcolor)
	case "y":
		y.Print(currcolor)
	case "c":
		c.Print(currcolor)
	}

}

func printspace(count int) {

	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}

}

func printx(count int) {

	for i := 0; i < count; i++ {
		fmt.Print("x")
	}

}

func printface(count int, color string) {

	for i := 0; i < count; i++ {
		printcolor(color)
	}

}

func drawcube(cube [3][3][3][6]string) {

	printspace(22)

	printface(4, cube[2][0][2][4])

	fmt.Println()

	printspace(15)

	printface(1, cube[2][0][1][4])

	printspace(1)

	printx(2)

	printface(10, cube[2][0][2][4])

	printx(2)

	fmt.Println()

	printspace(8)

	printx(1)

	printface(13, cube[2][0][1][4])

	printx(3)

	printface(12, cube[1][0][2][4])

	fmt.Println()

	printspace(2)

	printface(10, cube[2][0][0][4])

	printx(2)

	printface(2, cube[2][0][1][4])

	printx(2)

	printface(10, cube[1][0][1][4])

	printx(2)

	printface(3, cube[1][0][2][4])

	printx(2)

	printface(8, cube[0][0][2][4])

	fmt.Println()

	printface(1, cube[2][0][0][1])

	printx(3)

	printface(5, cube[2][0][0][4])

	printx(2)

	printface(7, cube[1][0][0][4])

	printx(2)

	printface(6, cube[1][0][1][4])

	printx(2)

	printface(5, cube[0][0][1][4])

	printx(2)

	printface(8, cube[0][0][2][4])

	printx(4)

	fmt.Println()

	printface(6, cube[2][0][0][1])

	printx(3)

	printface(10, cube[1][0][0][4])

	printx(3)

	printface(1, cube[0][0][0][4])

	printx(2)

	printface(12, cube[0][0][1][4])

	printx(3)

	printface(7, cube[0][0][2][2])

	fmt.Println()

	printface(6, cube[2][0][0][1])

	printx(1)

	printface(5, cube[1][0][0][1])

	printx(3)

	printface(14, cube[0][0][0][4])

	printx(5)

	printface(4, cube[0][0][1][2])

	printx(2)

	printface(7, cube[0][0][2][2])

	fmt.Println()

	printface(6, cube[2][0][0][1])

	printx(1)

	printface(6, cube[1][0][0][1])

	printx(1)

	printface(3, cube[0][0][0][1])

	printx(3)

	printface(4, cube[0][0][0][4])

	printx(3)

	printface(3, cube[0][0][0][2])

	printx(1)

	printface(7, cube[0][0][1][2])

	printx(2)

	printface(7, cube[0][0][2][2])

	fmt.Println()

	printface(1, cube[2][1][0][1])

	printx(3)

	printface(2, cube[2][0][0][1])

	printx(1)

	printface(6, cube[1][0][0][1])

	printx(1)

	printface(7, cube[0][0][0][1])

	printx(1)

	printface(8, cube[0][0][0][2])

	printx(1)

	printface(7, cube[0][0][1][2])

	printx(2)

	printface(4, cube[0][0][2][2])

	printx(3)

	fmt.Println()

	printface(6, cube[2][1][0][1])

	printx(3)

	printface(4, cube[1][0][0][1])

	printx(1)

	printface(7, cube[0][0][0][1])

	printx(1)

	printface(8, cube[0][0][0][2])

	printx(1)

	printface(7, cube[0][0][1][2])

	printx(4)

	printface(4, cube[0][1][2][2])

	printx(1)

	fmt.Println()

	printface(6, cube[2][1][0][1])

	printx(1)

	printface(4, cube[1][1][0][1])

	printx(3)

	printface(7, cube[0][0][0][1])

	printx(1)

	printface(8, cube[0][0][0][2])

	printx(5)

	printface(3, cube[0][1][1][2])

	printx(2)

	printface(6, cube[0][1][2][2])

	printx(1)

	fmt.Println()

	printface(6, cube[2][1][0][1])

	printx(1)

	printface(6, cube[1][1][0][1])

	printx(1)

	printface(1, cube[0][1][0][1])

	printx(3)

	printface(3, cube[0][0][0][1])

	printx(1)

	printface(3, cube[0][0][0][2])

	printx(3)

	printface(2, cube[0][1][0][2])

	printx(1)

	printface(7, cube[0][1][1][2])

	printx(2)

	printface(6, cube[0][1][2][2])

	fmt.Println()

	printx(3)

	printface(3, cube[2][1][0][1])

	printx(1)

	printface(6, cube[1][1][0][1])

	printx(1)

	printface(6, cube[0][1][0][1])

	printx(2)

	printface(8, cube[0][1][0][2])

	printx(1)

	printface(7, cube[0][1][1][2])

	printx(2)

	printface(4, cube[0][1][2][2])

	printx(2)

	fmt.Println()

	printface(5, cube[2][2][0][1])

	printx(3)

	printface(5, cube[1][1][0][1])

	printx(1)

	printface(7, cube[0][1][0][1])

	printx(1)

	printface(8, cube[0][1][0][2])

	printx(1)

	printface(7, cube[0][1][1][2])

	printx(4)

	printface(4, cube[0][2][2][2])

	fmt.Println()

	printface(6, cube[2][2][0][1])

	printx(1)

	printface(3, cube[1][2][0][1])

	printx(4)

	printface(7, cube[0][1][0][1])

	printx(1)

	printface(8, cube[0][1][0][2])

	printx(1)

	printface(2, cube[0][1][1][2])

	printx(3)

	printface(2, cube[0][2][1][2])

	printx(1)

	printface(7, cube[0][2][2][2])

	fmt.Println()

	printface(6, cube[2][2][0][1])

	printx(1)

	printface(6, cube[1][2][0][1])

	printx(4)

	printface(4, cube[0][1][0][1])

	printx(1)

	printface(4, cube[0][1][0][2])

	printx(5)

	printface(7, cube[0][2][1][2])

	printx(1)

	printface(7, cube[0][2][2][2])

	fmt.Println()

	printspace(2)

	printx(1)

	printface(3, cube[2][2][0][1])

	printx(1)

	printface(6, cube[1][2][0][1])

	printx(2)

	printface(4, cube[0][2][0][1])

	printx(5)

	printface(6, cube[0][2][0][2])

	printx(1)

	printface(7, cube[0][2][1][2])

	printx(1)

	printface(5, cube[0][2][2][2])

	printx(1)

	fmt.Println()

	printspace(7)

	printx(1)

	printface(6, cube[1][2][0][1])

	printx(1)

	printface(6, cube[0][2][0][1])

	printx(1)

	printface(8, cube[0][2][0][2])

	printx(1)

	printface(7, cube[0][2][1][2])

	printx(2)

	fmt.Println()

	printspace(12)

	printx(1)

	printface(1, cube[1][2][0][1])

	printx(1)

	printface(6, cube[0][2][0][1])

	printx(1)

	printface(8, cube[0][2][0][2])

	printx(1)

	printface(2, cube[0][2][1][2])

	printx(2)

	fmt.Println()

	printspace(15)

	printx(2)

	printface(4, cube[0][2][0][1])

	printx(2)

	printface(4, cube[0][2][0][2])

	printx(2)

	fmt.Println()

	printspace(20)

	printx(4)

	fmt.Println()

	time.Sleep(1 * time.Second)

}

func turncuberight(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					pd := d
					pw := w
					pf := 0

					switch f {
					case 0:
						pf = 1
					case 1:
						pf = 2
					case 2:
						pf = 3
					case 3:
						pf = 0
					case 4:
						pf = 4
					case 5:
						pf = 5

					}

					newcube[d][h][w][f] = cube[pw][h][2-pd][pf]

				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func turncubeleft(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					pd := d
					pw := w
					pf := 0

					switch f {
					case 0:
						pf = 3
					case 1:
						pf = 0
					case 2:
						pf = 1
					case 3:
						pf = 2
					case 4:
						pf = 4
					case 5:
						pf = 5

					}

					newcube[d][h][w][f] = cube[2-pw][h][pd][pf]

				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func turncubeup(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					pd := d
					ph := h
					pf := 0

					switch f {
					case 0:
						pf = 4
					case 1:
						pf = 1
					case 2:
						pf = 5
					case 3:
						pf = 3
					case 4:
						pf = 2
					case 5:
						pf = 0
					}

					newcube[d][h][w][f] = cube[ph][2-pd][w][pf]

				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func turncubedown(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					pd := d
					ph := h
					pf := 0

					switch f {
					case 0:
						pf = 5
					case 1:
						pf = 1
					case 2:
						pf = 4
					case 3:
						pf = 3
					case 4:
						pf = 0
					case 5:
						pf = 2
					}

					newcube[d][h][w][f] = cube[2-ph][pd][w][pf]

				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotaterightcw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {
					if w == 2 {
						pd := d
						ph := h
						pf := 0

						switch f {
						case 0:
							pf = 4
						case 1:
							pf = 1
						case 2:
							pf = 5
						case 3:
							pf = 3
						case 4:
							pf = 2
						case 5:
							pf = 0
						}

						newcube[d][h][w][f] = cube[ph][2-pd][w][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotaterightccw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					if w == 2 {

						pd := d
						ph := h
						pf := 0

						switch f {
						case 0:
							pf = 5
						case 1:
							pf = 1
						case 2:
							pf = 4
						case 3:
							pf = 3
						case 4:
							pf = 0
						case 5:
							pf = 2
						}

						newcube[d][h][w][f] = cube[2-ph][pd][w][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotateleftcw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					if w == 0 {

						pd := d
						ph := h
						pf := 0

						switch f {
						case 0:
							pf = 5
						case 1:
							pf = 1
						case 2:
							pf = 4
						case 3:
							pf = 3
						case 4:
							pf = 0
						case 5:
							pf = 2
						}

						newcube[d][h][w][f] = cube[2-ph][pd][w][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotateleftccw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {
					if w == 0 {
						pd := d
						ph := h
						pf := 0

						switch f {
						case 0:
							pf = 4
						case 1:
							pf = 1
						case 2:
							pf = 5
						case 3:
							pf = 3
						case 4:
							pf = 2
						case 5:
							pf = 0
						}

						newcube[d][h][w][f] = cube[ph][2-pd][w][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotatetopcw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {
					if h == 0 {
						pd := d
						pw := w
						pf := 0

						switch f {
						case 0:
							pf = 1
						case 1:
							pf = 2
						case 2:
							pf = 3
						case 3:
							pf = 0
						case 4:
							pf = 4
						case 5:
							pf = 5

						}

						newcube[d][h][w][f] = cube[pw][h][2-pd][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotatetopccw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					if h == 0 {
						pd := d
						pw := w
						pf := 0

						switch f {
						case 0:
							pf = 3
						case 1:
							pf = 0
						case 2:
							pf = 1
						case 3:
							pf = 2
						case 4:
							pf = 4
						case 5:
							pf = 5

						}

						newcube[d][h][w][f] = cube[2-pw][h][pd][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotatebottomcw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {

					if h == 2 {
						pd := d
						pw := w
						pf := 0

						switch f {
						case 0:
							pf = 3
						case 1:
							pf = 0
						case 2:
							pf = 1
						case 3:
							pf = 2
						case 4:
							pf = 4
						case 5:
							pf = 5

						}

						newcube[d][h][w][f] = cube[2-pw][h][pd][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotatebottomccw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {
					if h == 2 {
						pd := d
						pw := w
						pf := 0

						switch f {
						case 0:
							pf = 1
						case 1:
							pf = 2
						case 2:
							pf = 3
						case 3:
							pf = 0
						case 4:
							pf = 4
						case 5:
							pf = 5

						}

						newcube[d][h][w][f] = cube[pw][h][2-pd][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotatefrontcw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {
					if d == 0 {
						pw := w
						ph := h
						pf := 0

						switch f {
						case 0:
							pf = 0
						case 1:
							pf = 5
						case 2:
							pf = 2
						case 3:
							pf = 4
						case 4:
							pf = 1
						case 5:
							pf = 3
						}

						newcube[d][h][w][f] = cube[d][2-pw][ph][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func rotatefrontccw(cube [3][3][3][6]string) (newcube [3][3][3][6]string) {

	newcube = cube

	for d := 0; d < 3; d++ {
		for h := 0; h < 3; h++ {
			for w := 0; w < 3; w++ {
				for f := 0; f < 6; f++ {
					if d == 0 {
						pw := w
						ph := h
						pf := 0

						switch f {
						case 0:
							pf = 0
						case 1:
							pf = 4
						case 2:
							pf = 2
						case 3:
							pf = 5
						case 4:
							pf = 3
						case 5:
							pf = 1
						}

						newcube[d][h][w][f] = cube[d][pw][2-ph][pf]
					}
				}
			}
		}
	}

	drawcube(newcube)

	return newcube

}

func scramblecube(cube [3][3][3][6]string) (scrambledcube [3][3][3][6]string) {

	rand.Seed(time.Now().UnixNano())

	for x := 0; x < 12; x++ {
		switch rand.Intn(4) {
		case 1:
			cube = rotaterightcw(cube)
		case 2:
			cube = rotateleftccw(cube)
		case 3:
			cube = rotatebottomcw(cube)
		case 4:
			cube = rotatefrontcw(cube)
		}
	}

	return cube
}

func main() {

	cube := constructcube()

	fmt.Println("Welcome to the rubix cube!")

	time.Sleep(1 * time.Second)

	drawcube(cube)

	time.Sleep(1 * time.Second)

	options := make(map[string]string)

	fmt.Print("\nTime to use the cube! Add a 'c' at the end of any command to reverse it.\n\n")

	options["tr"] = "Rotate the cube right."

	options["tl"] = "Rotate the cube left."

	options["r"] = "Rotate the right segment of the cube."

	options["l"] = "Rotate the left segment of the cube."

	options["f"] = "Rotate the front segment of the cube."

	options["s"] = "Scramble cube."

	options["x"] = "Exit."

	scanner := bufio.NewScanner(os.Stdin)

optionsbegin:

	fmt.Print("What do you want to do?\n\n")

	for i, y := range options {
		if len(i) == 2 {
			fmt.Println(i + ":    " + y)
			continue
		}
		fmt.Println(i + ":     " + y)
	}

	fmt.Println()

	scanner.Scan()

	switch strings.ToLower(scanner.Text()) {
	case "tr":
		drawcube(cube)
		cube = turncuberight(cube)
		goto optionsbegin
	case "trc":
		drawcube(cube)
		cube = turncubeleft(cube)
		goto optionsbegin
	case "tl":
		drawcube(cube)
		cube = turncubeleft(cube)
		goto optionsbegin
	case "tlc":
		drawcube(cube)
		cube = turncuberight(cube)
		goto optionsbegin
	case "r":
		drawcube(cube)
		cube = rotaterightcw(cube)
		goto optionsbegin
	case "rc":
		drawcube(cube)
		cube = rotaterightccw(cube)
		goto optionsbegin
	case "l":
		drawcube(cube)
		cube = rotateleftccw(cube)
		goto optionsbegin
	case "lc":
		drawcube(cube)
		cube = rotateleftcw(cube)
		goto optionsbegin
	case "f":
		drawcube(cube)
		cube = rotatefrontcw(cube)
		goto optionsbegin
	case "fc":
		drawcube(cube)
		cube = rotatefrontccw(cube)
		goto optionsbegin
	case "s":
		drawcube(cube)
		cube = scramblecube(cube)
		goto optionsbegin
	case "t":
		drawcube(cube)
		cube = rotatetopcw(cube)
		goto optionsbegin
	case "tc":
		drawcube(cube)
		cube = rotatetopccw(cube)
		goto optionsbegin
	case "b":
		drawcube(cube)
		cube = rotatebottomcw(cube)
		goto optionsbegin
	case "bc":
		drawcube(cube)
		cube = rotatebottomccw(cube)
		goto optionsbegin
	case "x":
		return
	}

}

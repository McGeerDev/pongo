package game

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
  
	fmt.Println("hello World")
  screen, err := tcell.NewScreen()

  if err != nil {
    log.Fatalf("%+v",err)
  }
  if err := screen.Init(); err != nil {
    log.Fatalf("%+v",err)
  }

}

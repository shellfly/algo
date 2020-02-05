/*
Execution: go run cmd/degrees-of-separation/main.go filename delimiter source
Data files:
   https://algs4.cs.princeton.edu / 41graph / routes.txt
   https://algs4.cs.princeton.edu / 41graph / movies.txt

% go run cmd/degrees-of-separation/main.go routes.txt " " "JFK"
LAS
   JFK
   ORD
   DEN
   LAS
DFW
   JFK
   ORD
   DFW
EWR
   Not in database.

% go run cmd/degrees-of-separation/main.go movies.txt "/" "Bacon, Kevin"
Kidman, Nicole
   Bacon, Kevin
   Woodsman, The(2004)
   Grier, David Alan
   Bewitched(2005)
   Kidman, Nicole
Grant, Cary
   Bacon, Kevin
   Planes, Trains & Automobiles(1987)
   Martin, Steve(I)
   Dead Men Don't Wear Plaid(1982)
   Grant, Cary

% go run cmd/degrees-of-separation/main.go movies.txt "/" "Animal House (1978)"
Titanic(1997)
   Animal House(1978)
   Allen, Karen(I)
   Raiders of the Lost Ark(1981)
   Taylor, Rocky(I)
   Titanic(1997)
To Catch a Thief(1955)
   Animal House(1978)
   Vernon, John(I)
   Topaz(1969)
   Hitchcock, Alfred(I)
   To Catch a Thief(1955)
*/

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shellfly/algo/algs4"
	"github.com/shellfly/algo/stdin"
)

func main() {
	sg := algs4.NewSymbolGraph(os.Args[1], os.Args[2])
	g := sg.G()
	source := os.Args[3]
	if !sg.Contains(source) {
		fmt.Println(source, "not in database")
		return
	}
	s := sg.Index(source)
	bfs := algs4.NewBreadFirstPaths(g, s)
	stdin := stdin.NewStdIn()
	for !stdin.IsEmpty() {
		sink := strings.Trim(stdin.ReadString(), " ")
		if sg.Contains(sink) {
			t := sg.Index(sink)
			if bfs.HasPathTo(t) {
				for _, v := range bfs.PathTo(t) {
					fmt.Println("  ", sg.Name(v))
				}
			} else {
				fmt.Println("Not connected")
			}
		} else {
			fmt.Println("Not in database")
		}
	}
}

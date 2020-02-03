/******************************************************************************
 *  Execution:    go run cmd/kosaraju-scc/main.go filename.txt
 *  Dependencies: Digraph.java TransitiveClosure.java StdOut.java In.java
 *  Data files:   https://algs4.cs.princeton.edu/42digraph/tinyDG.txt
 *                https://algs4.cs.princeton.edu/42digraph/mediumDG.txt
 *                https://algs4.cs.princeton.edu/42digraph/largeDG.txt
 *
 *  Compute the strongly-connected components of a digraph using the
 *  Kosaraju-Sharir algorithm.
 *
 *  Runs in O(E + V) time.
 *
 *  % go run cmd/kosaraju-scc/main.go tinyDG.txt
 *  5 strong components
 *  1
 *  0 2 3 4 5
 *  9 10 11 12
 *  6 8
 *  7
 *
 *  % go run cmd/kosaraju-scc/main.go mediumDG.txt
 *  10 strong components
 *  21
 *  2 5 6 8 9 11 12 13 15 16 18 19 22 23 25 26 28 29 30 31 32 33 34 35 37 38 39 40 42 43 44 46 47 48 49
 *  14
 *  3 4 17 20 24 27 36
 *  41
 *  7
 *  45
 *  1
 *  0
 *  10
 *
 *  % java -Xss50m KosarajuSharirSCC mediumDG.txt
 *  25 strong components
 *  7 11 32 36 61 84 95 116 121 128 230   ...
 *  28 73 80 104 115 143 149 164 184 185  ...
 *  38 40 200 201 207 218 286 387 418 422 ...
 *  12 14 56 78 87 103 216 269 271 272    ...
 *  42 48 112 135 160 217 243 246 273 346 ...
 *  46 76 96 97 224 237 297 303 308 309   ...
 *  9 15 21 22 27 90 167 214 220 225 227  ...
 *  74 99 133 146 161 166 202 205 245 262 ...
 *  43 83 94 120 125 183 195 206 244 254  ...
 *  1 13 54 91 92 93 106 140 156 194 208  ...
 *  10 39 67 69 131 144 145 154 168 258   ...
 *  6 52 66 113 118 122 139 147 212 213   ...
 *  8 127 150 182 203 204 249 367 400 432 ...
 *  63 65 101 107 108 136 169 170 171 173 ...
 *  55 71 102 155 159 198 228 252 325 419 ...
 *  4 25 34 58 70 152 172 196 199 210 226 ...
 *  2 44 50 88 109 138 141 178 197 211    ...
 *  57 89 129 162 174 179 188 209 238 276 ...
 *  33 41 49 119 126 132 148 181 215 221  ...
 *  3 18 23 26 35 64 105 124 157 186 251  ...
 *  5 16 17 20 31 47 81 98 158 180 187    ...
 *  24 29 51 59 75 82 100 114 117 134 151 ...
 *  30 45 53 60 72 85 111 130 137 142 163 ...
 *  19 37 62 77 79 110 153 352 353 361    ...
 *  0 68 86 123 165 176 193 239 289 336   ...
 *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo"
	"github.com/shellfly/algo/stdin"
)

func main() {
	g := algo.NewDigraph(stdin.NewIn(os.Args[1]))
	scc := algo.NewKosarajuSCC(g)
	m := scc.Count()
	fmt.Println(m, "strong components")
	components := make([]*algo.Queue, m)
	for i := 0; i < m; i++ {
		components[i] = algo.NewQueue()
	}
	for v := 0; v < g.V(); v++ {
		components[scc.ID(v)].Enqueue(v)
	}
	for i := 0; i < m; i++ {
		for _, v := range components[i].Slice() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

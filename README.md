## Overview

This repository contains the Go source code for the algorithms in the textbook
<a href = "http://amzn.to/13VNJi7">Algorithms, 4th Edition</a> by Robert Sedgewick and Kevin Wayne.

The official Java source code is <a href="https://github.com/kevin-wayne/algs4">here</a>.

## Goals

Make a Golang implementation of the library so that a Go programmer can learn this book easily.

Try to keep the interface and variable name consistent with the original book while writing idiomatic go code.

## Note

Unlike Java or Python where you can put `main` function in a file directly. In `go` the `main` function has to be in the `main` package and in the **cmd** directory by convention.
To test each algorithm(data structure), instead of running the file directly, you have to run the file under **cmd** directory. Example: [cmd/bag/main.go](cmd/bag/main.go)

## Index

* 1 FUNDAMENTALS

  * [Bag](bag.go)
  * [Queue](queue.go)
  * [Stack](stack.go)
  * [UnionFind](uf.go)

* 2 SORTING

  * [Selection](selection.go)
  * [Insertion](insertion.go)
  * [Shell](shell.go)
  * [Merge](merge.go)
  * [Quick](quick.go)
  * [Quick3Way](quick_3way.go)
  * [MaxPQ](max_pq.go)
  * [TopM](cmd/topm/main.go)
  * [IndexMinPQ](index_min_pq.go)
  * [Multiway](cmd/multiway/main.go)
  * [Heap](heap.go)

* 3 SEARCHING

  * [FrequencyCounter](cmd/frequency-counter/main.go)
  * [SequentialSearchST](sequential_search.go)
  * [BinarySearchST](binary_search_st.go)
  * [BST](bst.go)
  * [RedBlackBST](red_black_bst.go)
  * [SeparateChainingHashST](separate_chaining_hash_st.go)
  * [LinearProbingHashST](linear_probing_hash_st.go)

* 4 GRAPHS

  * [Graph](graph.go)
  * [DepthFirstSearch](depth_first_search.go)
  * [DepthFirstPaths](depth_first_paths.go)
  * [BreadthFirstPaths](breadth_first_paths.go)
  * [CC](cc.go)
  * [SymbolGraph](symbol_graph.go)
  * [DegreesOfSeparation](cmd/degrees-of-separation/main.go)
  * [Digraph](digraph.go)
  * [DirectedDFS](directed_dfs.go)
  * [DirectedCycle](directed_cycle.go)

* 5 STRING

  * [LSD](lsd.go)

## License

This code is released under MIT.

## Contribute to this repository

Issue reports and code fixes are welcome. please follow the same style as the code in the repository and add test for your code.
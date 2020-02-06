## Overview

This repository contains the Go source code for the algorithms in the textbook
<a href = "http://amzn.to/13VNJi7">Algorithms, 4th Edition</a> by Robert Sedgewick and Kevin Wayne.

The official Java source code is <a href="https://github.com/kevin-wayne/algs4">here</a>.

## Goals

Make a Go implementation of the library so that a Go programmer can follow this book easily or prefer to demonstrate the algorithms using Go.

Try to keep the interface and variable name consistent with the original book while writing idiomatic go code.

## Note

Unlike Java or Python where you can put `main` function in a file directly. In `go` the `main` function has to be in the `main` package and in the **cmd** directory by convention.
To test each algorithm(data structure), instead of running the file directly, you have to run the file under **cmd** directory. Example: [cmd/bag/main.go](algs4/cmd/bag/main.go)

## Index

* 1 FUNDAMENTALS
  * [Bag](algs4/bag.go)
  * [Queue](algs4/queue.go)
  * [Stack](algs4/stack.go)
  * [UnionFind](algs4/uf.go)
* 2 SORTING
  * [Selection](algs4/selection.go)
  * [Insertion](algs4/insertion.go)
  * [Shell](algs4/shell.go)
  * [Merge](algs4/merge.go)
  * [Quick](algs4/quick.go)
  * [Quick3Way](algs4/quick_3way.go)
  * [MaxPQ](algs4/max_pq.go)
  * [TopM](algs4/cmd/topm/main.go)
  * [IndexMinPQ](algs4/index_min_pq.go)
  * [Multiway](algs4/cmd/multiway/main.go)
  * [Heap](algs4/heap.go)
* 3 SEARCHING
  * [FrequencyCounter](algs4/cmd/frequency-counter/main.go)
  * [SequentialSearchST](algs4/sequential_search.go)
  * [BinarySearchST](algs4/binary_search_st.go)
  * [BST](algs4/bst.go)
  * [RedBlackBST](algs4/red_black_bst.go)
  * [SeparateChainingHashST](algs4/separate_chaining_hash_st.go)
  * [LinearProbingHashST](algs4/linear_probing_hash_st.go)
* 4 GRAPHS
  * Graph
    * [Graph](algs4/graph.go)
    * [DepthFirstSearch](algs4/depth_first_search.go)
    * [DepthFirstPaths](algs4/depth_first_paths.go)
    * [BreadthFirstPaths](algs4/breadth_first_paths.go)
    * [CC](algs4/cc.go)
    * [Cycle](algs4/cycle.go)
    * [SymbolGraph](algs4/symbol_graph.go)
    * [DegreesOfSeparation](algs4/cmd/degrees-of-separation/main.go)
  * Digraph
    * [Digraph](algs4/digraph.go)
    * [SymbolDigraph](algs4/symbol_digraph.go)
    * [DirectedDFS](algs4/directed_dfs.go)
    * [DirectedCycle](algs4/directed_cycle.go)
    * [DepthFirstOrder](algs4/depth_first_order.go)
    * [Topological](algs4/topological.go)
    * [KosarajuSCC](algs4/kosaraju_scc.go)
  * MST
    * [EdgeWeightedGraph](algs4/edge_weighted_graph.go)
    * [LazyPrimMST](algs4/lazy_prim_mst.go)
    * [PrimMST](algs4/prim_mst.go)
    * [KruskalMST](algs4/kruskal_mst.go)
  * Shortest Paths
    * [EdgeWeightedDigraph](algs4/edge_weighted_digraph.go)
    * [DijkstraSP](algs4/dijkstra_sp.go)
    * [AcyclicSP](algs4/acyclic_sp.go)
    * [BellmanFordSP](algs4/bellman_ford_sp.go)
* 5 STRING
  * [LSD](algs4/lsd.go)
  * [MSD](algs4/msd.go)
  * [Quick3string](algs4/quick3_string.go)
  * [TrieST](trie_st.go)


## License

This code is released under MIT.

## Contribute to this repository

Issue reports and code fixes are welcome. please follow the same style as the code in the repository and add test for your code.
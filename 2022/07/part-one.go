package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
	tree := TreeNode{
		Type:       TreeNodeTypeDirectory,
		ChildNodes: make([]TreeNode, 0),
	}

	var currentNode *TreeNode
	currentNode = &tree

	scanner.Scan() // Skip the first `cd` line to root

	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Split(line, " ")

		if lineParts[0] == "$" { // Issuing new command
			if lineParts[1] == "cd" {
				if lineParts[2] == ".." {
					currentNode = currentNode.ParentNode

					continue
				}

				var childNodeIndex int

				for i, childNode := range currentNode.ChildNodes {
					if childNode.Name == lineParts[2] {
						childNodeIndex = i

						break
					}
				}

				currentNode = &currentNode.ChildNodes[childNodeIndex]
			}

			// We issued an `ls` command if not `cd` so we're good for next loop iteration

			continue
		}

		// We must be in an ls otherwise
		childNode := TreeNode{
			Name:       lineParts[1],
			ParentNode: currentNode,
		}

		if lineParts[0] == "dir" { // Directory item
			childNode.Type = TreeNodeTypeDirectory
			childNode.ChildNodes = make([]TreeNode, 0)
		} else {
			childNode.Type = TreeNodeTypeFile
			childNode.Size, _ = strconv.ParseInt(lineParts[0], 10, 64)
		}

		currentNode.ChildNodes = append(currentNode.ChildNodes, childNode)
	}

	tree.calculateSize()
	tree.print(0)

	fmt.Printf("\nDirectories sum is %d\n", tree.sumDirectoriesAtMostSize(100000))
}

func (node *TreeNode) sumDirectoriesAtMostSize(directorySizeHighBound int64) int64 {
	var sum int64

	if node.Size <= directorySizeHighBound {
		sum += node.Size
	}

	for _, childNode := range node.ChildNodes {
		if childNode.Type == TreeNodeTypeDirectory {
			sum += childNode.sumDirectoriesAtMostSize(directorySizeHighBound)
		}
	}

	return sum
}

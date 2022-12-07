package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) {
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

	fmt.Printf("Closest directory size is %d\n", tree.getMinimumSizeOver(30000000-(70000000-tree.Size)))
}

func (node *TreeNode) getMinimumSizeOver(size int64) int64 {
	var minimumSize int64

	if node.Size >= size {
		minimumSize = node.Size
	}

	for _, childNode := range node.ChildNodes {
		if childNode.Type == TreeNodeTypeDirectory {
			childClosestSize := childNode.getMinimumSizeOver(size)

			if childClosestSize > 0 && childClosestSize < minimumSize {
				minimumSize = childClosestSize
			}
		}
	}

	return minimumSize
}

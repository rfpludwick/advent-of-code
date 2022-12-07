package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	flagShowHelp  bool
	flagTestMode  bool
	flagInputFile string
)

func init() {
	flag.BoolVar(&flagShowHelp, "help", false, "Show this help")
	flag.BoolVar(&flagTestMode, "test", false, "Enable test mode")
	flag.StringVar(&flagInputFile, "input", "./input.txt", "Input file to use")
}

type TreeNodeType byte

const (
	TreeNodeTypeFile      TreeNodeType = 0
	TreeNodeTypeDirectory TreeNodeType = 1
)

type TreeNode struct {
	Name       string
	Type       TreeNodeType
	Size       int64 // For directories, this is a value populated by calculation later
	ParentNode *TreeNode
	ChildNodes []TreeNode
}

func main() {
	flag.Parse()

	if flagShowHelp {
		flag.Usage()

		os.Exit(0)
	}

	if flagTestMode {
		flagInputFile = "./test-input.txt"
	}

	file, err := os.Open(flagInputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (node *TreeNode) calculateSize() int64 {
	for i, childNode := range node.ChildNodes {
		if childNode.Type == TreeNodeTypeFile {
			node.Size += childNode.Size
		} else {
			node.Size += childNode.calculateSize()
		}

		node.ChildNodes[i] = childNode
	}

	return node.Size
}

func (node *TreeNode) print(offset int) {
	var nodeName string
	var nodeType string

	if len(node.Name) == 0 {
		nodeName = "/"
	} else {
		nodeName = node.Name
	}

	if node.Type == TreeNodeTypeDirectory {
		nodeType = "dir"
	} else {
		nodeType = "file"
	}

	fmt.Printf("%"+fmt.Sprintf("%d", offset)+"s %s (%s, size=%d)\n", "-", nodeName, nodeType, node.Size)

	for _, childNode := range node.ChildNodes {
		childNode.print(offset + 2)
	}
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

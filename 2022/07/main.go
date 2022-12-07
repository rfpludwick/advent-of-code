package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagShowHelp  bool
	flagTestMode  bool
	flagInputFile string
	flagRunPart   int
)

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

func init() {
	flag.BoolVar(&flagShowHelp, "help", false, "Show this help")
	flag.BoolVar(&flagTestMode, "test", false, "Enable test mode")
	flag.StringVar(&flagInputFile, "input", "./input.txt", "Input file to use")
	flag.IntVar(&flagRunPart, "run-part", 1, "The part to run")

	flag.Parse()

	if flagShowHelp {
		flag.Usage()

		os.Exit(0)
	}

	if flagTestMode {
		flagInputFile = "./test-input.txt"
	}
}

func main() {
	file, err := os.Open(flagInputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch flagRunPart {
	case 1:
		partOne(scanner)
	case 2:
		partTwo(scanner)
	default:
		log.Fatalf("Error: Part number %d is not supported", flagRunPart)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

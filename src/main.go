package main

import "tree"
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
	reader := bufio.NewReader(os.Stdin)
	tree := tree.NewBinaryTreeArrary(10)
	for {
		fmt.Print("Enter CMD: ")
		rawLine, _, _ := reader.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		numA, _ := strconv.Atoi(string(tokens[1]))
		numB, _ := strconv.Atoi(string(tokens[2]))
		if tokens[0] == "update" {
			tree.Update(numA, numB)
			tree.Print()
		} else if tokens[0] == "query" {
			sumLeft := tree.Sum(numA)
			sumRight := tree.Sum(numB)
			fmt.Println(sumLeft, sumRight)
			tree.Print()
			fmt.Println("Sum[", tokens[1], ":", tokens[2], "]=", sumRight-sumLeft)
		} else {
			fmt.Println("Invalid CMD")
		}

		tree.Print()
	}
}

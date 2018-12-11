package day8

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

var example bool

type node struct {
	nchild    int
	nmetadata int
	children  []*node
	metadata  []int
}

func (n *node) parseRoot(s *bufio.Scanner) error {

	if s.Scan() {
		t, err := strconv.Atoi(s.Text())
		if err != nil {
			return err
		}
		n.nchild = t
		if s.Scan() {
			m, err := strconv.Atoi(s.Text())
			if err != nil {
				return err
			}
			n.nmetadata = m
		} else {
			return fmt.Errorf("could read nchildren but no n metadata")
		}
	} else {
		return fmt.Errorf("EOF")
	}
	err := n.readChildren(s)
	if err != nil {
		return err
	}
	return nil

}

func (n *node) readMetadata(s *bufio.Scanner) error {
	if s.Scan() {
		t, err := strconv.Atoi(s.Text())
		if err != nil {
			return err
		}
		n.metadata = append(n.metadata, t)
	} else {
		return fmt.Errorf("couln't read all metadata")
	}
	return nil
}

func (n *node) readChildren(s *bufio.Scanner) error {
	for i := 0; i < n.nchild; i++ {
		c := new(node)
		n.children = append(n.children, c)
		if s.Scan() {
			t, err := strconv.Atoi(s.Text())
			if err != nil {
				return err
			}
			c.nchild = t
			if s.Scan() {
				t, err := strconv.Atoi(s.Text())
				if err != nil {
					return err
				}
				c.nmetadata = t
			} else {
				return fmt.Errorf("couldnt read amount of metadata")
			}
		} else {
			fmt.Errorf("EOF")
		}
		err := c.readChildren(s)
		if err != nil {
			return err
		}
	}
	for j := 0; j < n.nmetadata; j++ {
		err := n.readMetadata(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *node) readTree() int {
	result := 0
	result = result + acc(n.metadata)
	for _, n := range n.children {
		result = result + n.readTree()
	}
	return result
}

func (n *node) readTree2() int {
	result := 0
	if n.nchild == 0 {
		result = acc(n.metadata)
	} else {
		for _, c := range n.metadata {
			if c == 0 {
				continue
			}
			if c > n.nchild {
				continue
			}
			result = result + n.children[c-1].readTree2()
		}
	}

	return result
}

func acc(a []int) int {
	result := 0
	for _, e := range a {
		result = result + e
	}
	return result
}

func Run(input string) (string, string, error) {

	example = false

	if example {
		input = "day8/example.txt"
	}

	seq, err := ioutil.ReadFile(input)
	if err != nil {
		return "", "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(seq))
	scanner.Split(bufio.ScanWords)

	n := new(node)
	err = n.parseRoot(scanner)
	if err != nil {
		return "", "", err
	}

	result1 := strconv.Itoa(n.readTree())
	result2 := strconv.Itoa(n.readTree2())

	return result1, result2, nil
}

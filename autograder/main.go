package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type entry struct {
	val int
	key string
}

type entries []entry

const ShellToUse = "bash"

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func (s entries) Len() int           { return len(s) }
func (s entries) Less(i, j int) bool { return s[i].val < s[j].val }
func (s entries) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	err, out, errout := Shellout("ls Recruitment\\ Task/")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println("--- stdout ---")
	temp := strings.Fields(out)
	length := len(temp)
	m := make(map[string]int)

	for i := 0; i < length; i++ {
		fmt.Println(temp[i])
		element := temp[i]
		words := strings.Split(element, "-")
		name := words[len(words)-1]
		m[name] = m[name] + 0
		strExec := "cd Recruitment\\ Task/" + element + " && go test"
		err, out, errout = Shellout(strExec)
		badRoute := 0
		if err != nil {
			log.Printf("error: %v\n", err)
			badRoute = 1
		}
		if errout != "" {
			fmt.Println(errout)
			badRoute = 1
		}
		if badRoute > 0 {
			continue
		}

		if out[0:4] == "PASS" {
			//fmt.Println("done")
			m[name] = m[name] + 1
		} else {
			m[name] = m[name] + 0
		}
	}
	var scores entries
	for k, v := range m {
		scores = append(scores, entry{v, k})
	}
	sort.Sort(sort.Reverse(scores))
	b, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	_, err = b.WriteString("RESULTS\n")
	for _, v := range scores {
		if _, err = b.WriteString(v.key + "\t" + fmt.Sprintf("%d", v.val) + "\n"); err != nil {
			fmt.Println(err)
		}
	}
	if err := b.Close(); err != nil {
		fmt.Println(err)
	}
	if errout != "" {
		fmt.Println(errout)
	}
}

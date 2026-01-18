package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	mode := flag.String("mode", "export", "export (s4 share SLIP-39) or import (SLIP-39 to s4 share)")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("read stdin: %v", err)
	}

	switch *mode {
	case "export":
		mnemonics, err := slip39adapter.SharesToSlip39Mnemonics(lines)
		if err != nil {
			log.Fatalf("convert shares: %v", err)
		}
		for _, mn := range mnemonics {
			fmt.Println(mn)
		}
	case "import":
		shares, err := slip39adapter.Slip39MnemonicsToShares(lines)
		if err != nil {
			log.Fatalf("convert mnemonics: %v", err)
		}
		for _, sh := range shares {
			fmt.Println(sh)
		}
	default:
		log.Fatalf("unknown mode %q (expected export|import)", *mode)
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const version = "0.3.0"

func main() {
	var (
		force       bool
		src         string
		dest        string
		whitelist   string
		help        bool
		showVersion bool
	)

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while parsing args: %v\n", err)
		return
	}

	flag.StringVar(&src, "i", "", "a input yaml file (required)")
	flag.StringVar(&dest, "o", filepath.Join(wd, "surfaces.txt"), "an output file")
	flag.BoolVar(&force, "f", false, "skip overwriting confirmation")
	flag.StringVar(&whitelist, "w", "", "a whitelist surfaces separated by comma")
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if showVersion {
		fmt.Println(version)
		return
	}

	whitelistSurfaces := strings.Split(whitelist, ",")

	if src == "" {
		flag.Usage()
		return
	}

	fileInfo, _ := os.Stat(dest)
	if fileInfo != nil && fileInfo.IsDir() {
		fmt.Fprintln(os.Stderr, "error while parsing args:", dest, "is directory")
		return
	}
	if _, err := os.Stat(filepath.Dir(dest)); err != nil {
		fmt.Fprintf(os.Stderr, "error while parsing args: %v\n", err)
		return
	}

	data, err := loadYaml(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while loading yaml: %v\n", err)
		return
	}

	result := renderRaw(data.Raw)
	max := 0
	offset := 0
	offsetOrigin := 0
	for i, character := range data.Characters {
		surfaces := generateSurfaces(character.Parts)

		surfaceList := classifySurfaces(character.Parts, surfaces)

		if i > 0 {
			if i == 1 {
				digit := countDigit(max)
				offset = 1
				for j := 0; j < digit; j++ {
					offset *= 10
				}
				offsetOrigin = offset
			} else {
				offset += offsetOrigin
			}
		}
		var r string
		r, max = formatSurfaces(&character, surfaces, surfaceList, whitelistSurfaces, offset)
		result += r + "\n\n"
	}
	result = strings.TrimRight(result, "\n")

	if !force && isFileExists(dest) {
		fmt.Println("overwrite", dest+"? (y/n)")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "n") {
				return
			} else if strings.Contains(scanner.Text(), "y") {
				break
			}
		}
	}

	if err := os.WriteFile(dest, []byte(result), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "error while writing surfaces.txt: %v\n", err)
		return
	}

	fmt.Println("saved to", dest)
}

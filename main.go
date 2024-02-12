package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		force bool
		src   string
		dest  string
    whitelist string
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
	flag.Parse()

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

	surfaces := generateSurfaces(data.Parts)

	surfaceList := classifySurfaces(data.Parts, surfaces)

	result := formatSurfaces(data, surfaces, surfaceList, whitelistSurfaces)

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

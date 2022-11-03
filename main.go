package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	data, err := loadYaml(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while loading yaml: %v\n", err)
		return
	}

	surfaces := generateSurfaces(data.Parts, 0, []string{})

	surfaceList := classifySurfaces(data.Parts, surfaces)

	result := formatSurfaces(data, surfaces, surfaceList)

	execPath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while writing surfaces.txt: %v\n", err)
	}

	dest := filepath.Join(filepath.Dir(execPath), "surfaces.txt")
	if err := os.WriteFile(dest, []byte(result), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "error while writing surfaces.txt: %v\n", err)
		return
	}

	fmt.Println("saved to", dest)
}

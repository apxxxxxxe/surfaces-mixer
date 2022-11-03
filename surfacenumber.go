package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SurfaceList = [][][][]string

func formatSurfaces(data *Root, surfaces [][]string, surfaceList SurfaceList) string {
	const indentCount = 2

	res := "charset,UTF-8\n\ndescript\n{\n  Version,1\n}\n\n"

	// Partsの定義
	for i, sp := range surfaceList {
		res += fmt.Sprintf("// %s\n\n", data.Parts[i].Group)
		for j, n := range sp {
			num := ""
			for _, s := range n {
				num += combineNum(s) + ","
			}
			num = strings.TrimSuffix(num, ",")
			res += fmt.Sprintf("surface%s\n{\n  // %s\n%s}\n\n", num, data.Parts[i].Details[j].Name, addIndents(data.Parts[i].Details[j].Text, indentCount))
		}
	}

	min := combineNum(surfaces[0])
	max := combineNum(surfaces[len(surfaces)-1])

	// Baseの定義
	res += fmt.Sprintf("\n\nsurface.append%s-%s\n{\n%s}\n", min, max, addIndents(strings.TrimSpace(data.Base), indentCount))

	return res
}

// 各パーツを必要とするサーフェスの分類を行い、配列として返す
func classifySurfaces(parts []PartGroup, surfaces [][]string) SurfaceList {
	res := make(SurfaceList, len(parts))
	for i := range res {
		res[i] = make([][][]string, len(parts[i].Details))
	}

	for _, s := range surfaces {
		for i, p := range parts {
			for j := 1; j < len(p.Details)+1; j++ {
				n, _ := strconv.Atoi(s[i])
				if n == j {
					res[i][j-1] = append(res[i][j-1], s)
				}
			}
		}
	}
	return res
}

// パーツの種類と数から取りうるサーフェス番号を列挙する
func generateSurfaces(parts []PartGroup, i int, tmp []string) [][]string {
	if i == len(parts) {
		return [][]string{tmp}
	}
	res := [][]string{}
	for j := 1; j < len(parts[i].Details)+1; j++ {
		res = append(res, generateSurfaces(parts, i+1, append(tmp, printPartNum(len(parts[i].Details), j)))...)
	}
	return res
}

func combineNum(ary []string) string {
	s := ""
	for _, a := range ary {
		s += a
	}
	return s
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// パーツグループ(ex: 腕)
type Group = []Pose

// パーツ(ex: 後手)
type Pose = []SurfaceNumber

type SurfaceNumber = []SurfacePart

type SurfacePart struct {
	Number int
	Digit  int
}

func includeWhitelist(whitelist []string, num string) bool {
	if len(whitelist) <= 1 {
		return true
	}
	for _, w := range whitelist {
		if num == w {
			return true
		}
	}
	return false
}

func formatSurfaces(data *Root, surfaces []SurfaceNumber, surfaceList []Group, whitelist []string) string {
	const indentCount = 2

	res := "charset,UTF-8\n\ndescript\n{\n  Version,1\n}\n\n"

	// partsの定義
	for i, group := range surfaceList {
		res += fmt.Sprintf("// %s\n\n", data.Parts[i].Name)
		for j, pose := range group {
			num := ""
			numHistory := []SurfaceNumber{}
			for _, number := range pose {
				if !isIncludeSurfaceNumber(numHistory, number) &&
					includeWhitelist(whitelist, combineNum(number)) {
					numHistory = append(numHistory, number)
					num += combineNum(number) + ","
				}
			}
			if num != "" {
				num = strings.TrimSuffix(num, ",")
				res += fmt.Sprintf("surface%s\n{\n  // %s\n%s}\n\n", num, data.Parts[i].Poses[j].Name, addIndents(data.Parts[i].Poses[j].Text, indentCount))
			}
		}
	}

	min := combineNum(surfaces[0])
	max := combineNum(surfaces[len(surfaces)-1])

	// baseの定義
	res += fmt.Sprintf("\n\nsurface.append%s-%s\n{\n%s}\n", min, max, addIndents(strings.TrimSpace(data.Base), indentCount))

	return res
}

// 各パーツを必要とするサーフェスの分類を行い、配列として返す
func classifySurfaces(data []GroupData, surfaces []SurfaceNumber) []Group {
	res := make([]Group, len(data))
	for i := range res {
		res[i] = make([]Pose, len(data[i].Poses))
	}

	for i, group := range data {
		for j := range group.Poses {
			for _, number := range surfaces {
				if number[i].Number == j+1 {
					res[i][j] = append(res[i][j], number)
				}
			}
		}
	}
	return res
}

// パーツの種類と数から取りうるサーフェス番号を列挙する
func generateSurfaces(groupDatas []GroupData) []SurfaceNumber {
	const min = 1

	partCounts := []int{}
	digitCounts := []int{}
	for _, g := range groupDatas {
		l := len(g.Poses)
		partCounts = append(partCounts, l)
		digitCounts = append(digitCounts, countDigit(l))
	}

	// 作業用変数の初期化
	tmp := make([]SurfacePart, len(groupDatas))
	for i := range groupDatas {
		tmp[i] = SurfacePart{Number: min, Digit: digitCounts[i]}
	}

	sum := 1
	for _, c := range partCounts {
		sum *= c
	}

	surfaceNumbers := make([]SurfaceNumber, sum)

	for k := 0; k < len(surfaceNumbers); k++ {
		// tmpはスライスなのでcopyで値をコピー
		surfaceNumbers[k] = make([]SurfacePart, len(groupDatas))
		copy(surfaceNumbers[k], tmp)

		for i := len(groupDatas) - 1; i >= 0; i-- {

			tmp[i].Number += 1
			if tmp[i].Number > partCounts[i] {
				// 繰り上げて続行
				tmp[i].Number = min
			} else {
				// 繰り上がりがないので次ループへ
				break
			}
		}
	}

	return surfaceNumbers
}

func combineNum(sn SurfaceNumber) string {
	s := ""
	for _, sp := range sn {
		s += fmt.Sprintf("%0"+strconv.Itoa(sp.Digit)+"d", sp.Number)
	}
	return s
}

func isIncludeSurfaceNumber(ary []SurfaceNumber, n SurfaceNumber) bool {
	isExist := false
	for _, s := range ary {
		if isEqualSurfaceNumber(s, n) {
			isExist = true
			break
		}
	}
	return isExist
}

func isEqualSurfaceNumber(a, b SurfaceNumber) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].Number != b[i].Number || a[i].Digit != b[i].Digit {
			return false
		}
	}

	return true
}

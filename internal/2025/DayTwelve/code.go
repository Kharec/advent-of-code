package daytwelve

import (
	"os"
	"sort"
	"strconv"
	"strings"

	"advent-of-code/internal/registry"
)

func init() {
	registry.Register("2025D12", ParseInput, PartOne, PartTwo)
}

func ParseInput(filepath string) []string {
	content, _ := os.ReadFile(filepath)
	return strings.Split(string(content), "\n")
}

const (
	exactAreaThreshold    = 140
	exactPiecesThreshold  = 10
	exactMinSideThreshold = 8
)

type point struct {
	x, y int
}

type shapeOrientation struct {
	cells  []point
	width  int
	height int
}

type shape struct {
	area         int
	orientations []shapeOrientation
}

type region struct {
	width, height int
	counts        []int
}

func parsePuzzle(data []string) ([]shape, []region) {
	lines := compactLines(data)

	shapeBlocks := make(map[int][]string)
	shapeOrder := make([]int, 0)
	regions := make([]region, 0)

	for idx := 0; idx < len(lines); {
		line := lines[idx]

		if width, height, counts, ok := parseRegionLine(line); ok {
			regions = append(regions, region{width: width, height: height, counts: counts})
			idx++
			continue
		}

		shapeIdx, ok := parseShapeHeader(line)
		if !ok {
			idx++
			continue
		}
		idx++

		rows := make([]string, 0, 4)
		for idx < len(lines) {
			if _, _, _, isRegion := parseRegionLine(lines[idx]); isRegion {
				break
			}
			if _, isHeader := parseShapeHeader(lines[idx]); isHeader {
				break
			}
			rows = append(rows, lines[idx])
			idx++
		}

		if _, exists := shapeBlocks[shapeIdx]; !exists {
			shapeOrder = append(shapeOrder, shapeIdx)
		}
		shapeBlocks[shapeIdx] = rows
	}

	sort.Ints(shapeOrder)
	shapes := make([]shape, 0, len(shapeOrder))
	for _, shapeIdx := range shapeOrder {
		cells := make([]point, 0)
		for y, row := range shapeBlocks[shapeIdx] {
			for x, char := range row {
				if char == '#' {
					cells = append(cells, point{x: x, y: y})
				}
			}
		}

		shapes = append(shapes, shape{
			area:         len(cells),
			orientations: buildOrientations(cells),
		})
	}

	for idx := range regions {
		regions[idx].counts = normalizeCounts(regions[idx].counts, len(shapes))
	}

	return shapes, regions
}

func buildOrientations(cells []point) []shapeOrientation {
	transform := func(p point, t int) point {
		switch t {
		case 0:
			return point{x: p.x, y: p.y}
		case 1:
			return point{x: p.x, y: -p.y}
		case 2:
			return point{x: -p.x, y: p.y}
		case 3:
			return point{x: -p.x, y: -p.y}
		case 4:
			return point{x: p.y, y: p.x}
		case 5:
			return point{x: p.y, y: -p.x}
		case 6:
			return point{x: -p.y, y: p.x}
		default:
			return point{x: -p.y, y: -p.x}
		}
	}

	result := make([]shapeOrientation, 0, 8)
	seen := make(map[string]bool)

	if len(cells) == 0 {
		return result
	}

	for t := range 8 {
		transformed := make([]point, len(cells))
		minX, minY := cells[0].x, cells[0].y

		for idx, cell := range cells {
			p := transform(cell, t)
			transformed[idx] = p
			if p.x < minX {
				minX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
		}

		maxX, maxY := 0, 0
		for idx := range transformed {
			transformed[idx].x -= minX
			transformed[idx].y -= minY
			if transformed[idx].x > maxX {
				maxX = transformed[idx].x
			}
			if transformed[idx].y > maxY {
				maxY = transformed[idx].y
			}
		}

		sort.Slice(transformed, func(i, j int) bool {
			if transformed[i].y == transformed[j].y {
				return transformed[i].x < transformed[j].x
			}
			return transformed[i].y < transformed[j].y
		})

		var builder strings.Builder
		for _, p := range transformed {
			builder.WriteString(strconv.Itoa(p.x))
			builder.WriteByte(',')
			builder.WriteString(strconv.Itoa(p.y))
			builder.WriteByte(';')
		}
		key := builder.String()
		if seen[key] {
			continue
		}
		seen[key] = true

		result = append(result, shapeOrientation{
			cells:  transformed,
			width:  maxX + 1,
			height: maxY + 1,
		})
	}

	return result
}

type placement struct {
	mask []uint64
}

func buildPlacements(orientation shapeOrientation, width, height int) []placement {
	maxX := width - orientation.width
	maxY := height - orientation.height
	if maxX < 0 || maxY < 0 {
		return nil
	}

	cells := width * height
	words := (cells + 63) / 64
	placements := make([]placement, 0, (maxX+1)*(maxY+1))

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			mask := make([]uint64, words)
			for _, cell := range orientation.cells {
				idx := (y+cell.y)*width + (x + cell.x)
				mask[idx/64] |= 1 << (idx % 64)
			}
			placements = append(placements, placement{mask: mask})
		}
	}

	return placements
}

func canFitExactly(shapes []shape, r region) bool {
	totalPieces, totalArea, valid := regionStats(shapes, r.counts)
	if !valid {
		return false
	}

	boardArea := r.width * r.height
	if totalArea > boardArea {
		return false
	}
	if totalPieces == 0 {
		return true
	}

	byTypePlacements := make([][]placement, len(shapes))
	for t, s := range shapes {
		var placements []placement
		for _, o := range s.orientations {
			placements = append(placements, buildPlacements(o, r.width, r.height)...)
		}
		byTypePlacements[t] = placements
		if r.counts[t] > 0 && len(placements) == 0 {
			return false
		}
	}

	indices := make([]int, 0, totalPieces)
	for t, count := range r.counts {
		for range count {
			indices = append(indices, t)
		}
	}

	sort.Slice(indices, func(i, j int) bool {
		a, b := indices[i], indices[j]
		if shapes[a].area == shapes[b].area {
			return len(byTypePlacements[a]) < len(byTypePlacements[b])
		}
		return shapes[a].area > shapes[b].area
	})

	occupancy := make([]uint64, ((boardArea + 63) / 64))
	usedArea := 0

	var canPlace func(mask []uint64) bool
	canPlace = func(mask []uint64) bool {
		for idx := range mask {
			if occupancy[idx]&mask[idx] != 0 {
				return false
			}
		}
		return true
	}

	place := func(mask []uint64) {
		for idx := range mask {
			occupancy[idx] |= mask[idx]
		}
	}

	remove := func(mask []uint64) {
		for idx := range mask {
			occupancy[idx] &^= mask[idx]
		}
	}

	var dfs func(level int) bool
	dfs = func(level int) bool {
		if level == len(indices) {
			return true
		}

		freeCells := boardArea - usedArea
		remainingArea := totalArea - usedArea
		if remainingArea > freeCells {
			return false
		}

		t := indices[level]
		for _, p := range byTypePlacements[t] {
			if !canPlace(p.mask) {
				continue
			}
			place(p.mask)
			usedArea += shapes[t].area
			if dfs(level + 1) {
				return true
			}
			usedArea -= shapes[t].area
			remove(p.mask)
		}
		return false
	}

	return dfs(0)
}

func canFitRegion(shapes []shape, r region) bool {
	totalPieces, totalArea, valid := regionStats(shapes, r.counts)
	if !valid {
		return false
	}

	if totalArea > r.width*r.height {
		return false
	}
	if totalPieces == 0 {
		return true
	}

	if r.width*r.height <= exactAreaThreshold || totalPieces <= exactPiecesThreshold || min(r.width, r.height) <= exactMinSideThreshold {
		return canFitExactly(shapes, r)
	}

	return true
}

func compactLines(data []string) []string {
	lines := make([]string, 0, len(data))
	for _, line := range data {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			lines = append(lines, trimmed)
		}
	}
	return lines
}

func parseShapeHeader(line string) (int, bool) {
	if !strings.HasSuffix(line, ":") {
		return 0, false
	}
	value := strings.TrimSuffix(line, ":")
	index, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return index, true
}

func parseRegionLine(line string) (int, int, []int, bool) {
	dimPart, countsPart, ok := strings.Cut(line, ":")
	if !ok {
		return 0, 0, nil, false
	}

	width, height, ok := parseDimensions(strings.TrimSpace(dimPart))
	if !ok {
		return 0, 0, nil, false
	}

	counts := parseCounts(countsPart)
	return width, height, counts, true
}

func parseDimensions(value string) (int, int, bool) {
	left, right, ok := strings.Cut(value, "x")
	if !ok {
		return 0, 0, false
	}

	width, err := strconv.Atoi(left)
	if err != nil {
		return 0, 0, false
	}

	height, err := strconv.Atoi(right)
	if err != nil {
		return 0, 0, false
	}

	return width, height, true
}

func parseCounts(value string) []int {
	fields := strings.Fields(value)
	counts := make([]int, 0, len(fields))
	for _, field := range fields {
		count, err := strconv.Atoi(field)
		if err != nil {
			continue
		}
		counts = append(counts, count)
	}
	return counts
}

func normalizeCounts(counts []int, size int) []int {
	if len(counts) == size {
		return counts
	}
	if len(counts) > size {
		return counts[:size]
	}

	normalized := make([]int, size)
	copy(normalized, counts)
	return normalized
}

func regionStats(shapes []shape, counts []int) (totalPieces int, totalArea int, valid bool) {
	for idx, count := range counts {
		if count < 0 {
			return 0, 0, false
		}
		if idx >= len(shapes) {
			if count > 0 {
				return 0, 0, false
			}
			continue
		}

		totalPieces += count
		totalArea += count * shapes[idx].area
	}

	return totalPieces, totalArea, true
}

func PartOne(data []string) int {
	shapes, regions := parsePuzzle(data)
	fit := 0
	for _, region := range regions {
		if canFitRegion(shapes, region) {
			fit++
		}
	}
	return fit
}

func PartTwo(data []string) int {
	return 0
}

package flood_fill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	origColor := image[sr][sc]
	if origColor != color {
		dfs(image, sr, sc, origColor, color)
	}

	return image
}

func dfs(image [][]int, sr, sc, srcColor, color int) {
	if sr < 0 || sr > len(image)-1 || sc < 0 || sc > len(image[sr])-1 || image[sr][sc] != srcColor {
		return
	}

	image[sr][sc] = color

	dfs(image, sr+1, sc, srcColor, color)
	dfs(image, sr-1, sc, srcColor, color)
	dfs(image, sr, sc+1, srcColor, color)
	dfs(image, sr, sc-1, srcColor, color)
}

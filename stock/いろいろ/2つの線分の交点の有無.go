// 2つの線分に交点があるかどうかを判定
func judgeIentersected(ax, ay, bx, by, cx, cy, dx, dy int) bool {
	ta := (cx-dx)*(ay-cy) + (cy-dy)*(cx-ax)
	tb := (cx-dx)*(by-cy) + (cy-dy)*(cx-bx)
	tc := (ax-bx)*(cy-ay) + (ay-by)*(ax-cx)
	td := (ax-bx)*(dy-ay) + (ay-by)*(ax-dx)

	return tc*td < 0 && ta*tb < 0
	// return tc * td <= 0 && ta * tb <= 0; // 端点を含む場合
}
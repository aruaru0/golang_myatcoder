package main

// 丸め計算
func round(f float64, digit int) float64 {

	// 残す桁数
	pow := 1.0
	for i := 0; i < digit; i++ {
		pow *= 10
	}

	return float64(int(f*pow+0.5)) / pow
}

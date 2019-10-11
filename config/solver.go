package config

// Get : 任意の試合IDで使うソルバイメージを参照する
func (s Solver) Get(battleID int) string {
	val, exists := s.manualSet[battleID]
	if exists {
		return val
	}
	return s.Image
}

// Set : 任意の試合IDで使用するソルバイメージを変更する
func (s Solver) Set(battleID int, image string) {
	s.manualSet[battleID] = image
}

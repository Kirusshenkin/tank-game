package game

// Base представляет базу игрока или бота
type Base struct {
	Health int
	X      int
	Y      int
}

// NewBase создает новую базу с дефолтным здоровьем
func NewBase() *Base {
	return &Base{
		Health: 100, // Например, базовое здоровье базы
		X:      10,  // Примерное положение базы на карте
		Y:      10,
	}
}

// AttackBase уменьшает здоровье базы на указанную величину
func (b *Base) AttackBase(damage int) {
	if damage > b.Health {
		b.Health = 0
	} else {
		b.Health -= damage
	}
}

// IsDestroyed возвращает true, если база разрушена
func (b *Base) IsDestroyed() bool {
	return b.Health <= 0
}

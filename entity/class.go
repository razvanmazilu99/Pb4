package entity

type Class struct {
	ID          string     `gorm:"type: varchar(10); not null; primary key"`
	Title       string     `gorm:"type: varchar(100); not null"`
	Description string     `gorm:"type: varchar(255); not null"`
	Students    []*Student `gorm:"many2many:studentclasses"`
}

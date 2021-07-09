package entity

type Student struct {
	ID        string   `gorm:"type: varchar(10); not null; primary key"`
	FirstName string   `gorm:"type: varchar(100); not null" json:"first_name"`
	LastName  string   `gorm:"type: varchar(100); not null"`
	Age       int      `gorm:"type: int; not null"`
	Classes   []*Class `gorm:"many2many:studentclasses"`
}

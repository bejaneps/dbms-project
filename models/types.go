package models

// User struct represents teacher, student and faculty coordinator data
type User struct {
	ID       int    `gorm:"type: INT; PRIMARY_KEY; AUTO_INCREMENT"`
	Name     string `gorm:"type: TEXT: NOT NULL"`
	Password string `gorm:"type: TEXT; NOT NULL"`
	Type     string `gorm:"type: TEXT: NOT NULL"`
}

type Student struct {
	ID       int    `gorm:"type: INT; PRIMARY_KEY; AUTO_INCREMENT; NOT NULL"`
	Name     string `gorm:"type: TEXT"`
	DeptName string `gorm:"type: TEXT"`
	TotCred  int    `gorm:"type: INT"`
}

type Instructor struct {
	ID       int    `gorm:"type: INT; PRIMARY_KEY; AUTO_INCREMENT; NOT NULL"`
	Name     string `gorm:"type: TEXT"`
	DeptName string `gorm:"type: TEXT"`
	Salary   string `gorm:"type: TEXT"`
}

type Department struct {
	DeptName string `gorm:"type: TEXT; PRIMARY_KEY; NOT NULL"`
	Building string `gorm:"type: TEXT"`
	Budget   string `gorm:"type: TEXT"`
}

type Course struct {
	CourseID string `gorm:"type: TEXT; PRIMARY_KEY"`
	Title    string `gorm:"type: TEXT"`
	DeptName string `gorm:"type: TEXT"`
	Credits  int    `gorm:"type: INT"`
}

type Advisor struct {
	SID int `gorm:"type: INT; PRIMARY_KEY"`
	IID int `gorm:"type; INT"`
}

type Teaches struct {
	IID      int    `gorm:"type: INT; PRIMARY_KEY"`
	CourseID string `gorm:"type: TEXT; PRIMARY_KEY"`
}

type Takes struct {
	SID      int    `gorm:"type: INT; PRIMARY_KEY"`
	CourseID string `gorm:"type: TEXT; PRIMARY_KEY"`
}

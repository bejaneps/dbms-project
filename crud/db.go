package crud

import (
	"errors"

	"github.com/bejaneps/dbms-project/auth"
	"github.com/bejaneps/dbms-project/models"
)

var errNoUser = errors.New("no such user in db, please register")

// AddUser adds a teacher, student or coordinator to DB. Returns id of a newly registered user
func AddUser(name, password, uType string) int {
	db := auth.GetDB()
	defer db.Close()

	user := models.User{
		Name:     name,
		Password: password,
		Type:     uType,
	}

	db.Table("users").Save(&user)

	return user.ID
}

// CheckUser checks if a user exists in a db, if yes returns a struct of user, else error.
func CheckUser(id int, password string) (*models.User, error) {
	db := auth.GetDB()
	defer db.Close()

	user := models.User{}

	if password == "" {
		db.Table("users").First(&user, "id = ?", id)
	} else {
		db.Table("users").First(&user, "id = ? AND password = ?", id, password)
	}

	if user.Name != "" {
		return &user, nil
	}

	return nil, errNoUser
}

// AddCourse adds a course to a 'course' table, if successful returns nil, else error
func AddCourse(course models.Course) error {
	db := auth.GetDB()
	defer db.Close()

	err := db.Table("course").Save(&course).Error
	if err != nil {
		return err
	}

	return nil
}

// AddTeacher adds an instructor to 'instructor' table, if successful returns nil, else error
func AddTeacher(teacher models.Instructor) error {
	db := auth.GetDB()
	defer db.Close()

	err := db.Table("instructor").Save(&teacher).Error
	if err != nil {
		return err
	}

	return nil
}

// AddStudent adds a student to 'student' table, if successful returns nil, else error
func AddStudent(student models.Student) error {
	db := auth.GetDB()
	defer db.Close()

	err := db.Table("student").Save(&student).Error
	if err != nil {
		return err
	}

	return nil
}

// AssignCourse adds a teaches data to 'teaches' table, if successful returns nil, else error
func AssignCourse(teaches models.Teaches) error {
	db := auth.GetDB()
	defer db.Close()

	err := db.Table("teaches").Save(&teaches).Error
	if err != nil {
		return err
	}

	return nil
}

// AssignInstructor adds an advisor data to 'advisor' table, if successful returns nil, else error
func AssignInstructor(advisor models.Advisor) error {
	db := auth.GetDB()
	defer db.Close()

	err := db.Table("advisor").Save(&advisor).Error
	if err != nil {
		return err
	}

	return nil
}

// ListTcrCourses gets a list of courses that are assigned to instructor, if successful returns list of courses & nil, else nil & error
func ListTcrCourses(instructorID int) ([]string, error) {
	var courses []string

	db := auth.GetDB()
	defer db.Close()

	usr := models.User{}
	err := db.Table("users").First(&usr, "id = ?", instructorID).Error
	if err != nil {
		return nil, err
	}

	if usr.Name != "" {
		tcr := models.Instructor{}
		r := db.Table("instructor").Where("name = ?", usr.Name).Row()
		r.Scan(&tcr.ID, &tcr.Name, &tcr.DeptName, &tcr.Salary)

		if tcr.ID == 0 {
			return nil, nil
		}

		teaches := []models.Teaches{}
		err = db.Table("teaches").Find(&teaches, "i_id = ?", tcr.ID).Error
		if err != nil {
			return nil, err
		}

		for _, val := range teaches {
			courses = append(courses, val.CourseID)
		}
	}

	return courses, nil
}

// ListStdCourses gets a list of courses that are available for student, if successful returns list of courses & nil, else nil & error
func ListStdCourses(studentID int) ([]string, error) {
	var courses []string

	db := auth.GetDB()
	defer db.Close()

	usr := models.User{}
	err := db.Table("users").First(&usr, "id = ?", studentID).Error
	if err != nil {
		return nil, err
	}

	if usr.Name != "" {
		std := models.Student{}
		r := db.Table("student").Where("name = ?", usr.Name).Row()
		r.Scan(&std.ID, &std.Name, &std.DeptName, &std.TotCred)

		if std.ID == 0 {
			return nil, nil
		}

		takes := []models.Takes{}
		err = db.Table("takes").Find(&takes, "s_id = ?", std.ID).Error
		if err != nil {
			return nil, err
		}

		for _, val := range takes {
			courses = append(courses, val.CourseID)
		}
	}

	return courses, nil
}

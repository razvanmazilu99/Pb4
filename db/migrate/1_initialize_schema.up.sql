BEGIN;

CREATE TABLE Students (
    id varchar(10) NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    age int NOT NULL,
    CONSTRAINT student_key PRIMARY KEY (id)
);

CREATE TABLE Classes ( 
    id varchar(10) NOT NULL,
    title varchar(100) NOT NULL,
    description varchar(255) NOT NULL,
    CONSTRAINT class_key PRIMARY KEY (id)
);

CREATE TABLE StudentClasses (
    StudentID varchar(10) NOT NULL,
    ClassID varchar(10) NOT NULL,
    CONSTRAINT StudentClasses_PK PRIMARY KEY (
        StudentID,
        ClassID
    ), 
    CONSTRAINT Student_FK FOREIGN KEY (StudentID) REFERENCES Students(id),
    CONSTRAINT Class_FK FOREIGN KEY (ClassID) REFERENCES Classes(id)
);

COMMIT;
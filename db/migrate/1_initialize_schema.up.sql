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
    student_id varchar(10) NOT NULL,
    class_id varchar(10) NOT NULL,
    CONSTRAINT Student_Classes_PK PRIMARY KEY (student_id, class_id), 
    CONSTRAINT Student_FK FOREIGN KEY (student_id) REFERENCES Students(id),
    CONSTRAINT Class_FK FOREIGN KEY (class_id) REFERENCES Classes(id)
);

COMMIT;
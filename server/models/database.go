package models

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MsgPayload(payload interface{}) gin.H {
	return gin.H{
		"msg": payload,
	}
}

type Database struct {
	ready  bool
	client *mongo.Client
	db     *mongo.Database
}

const uri = "mongodb+srv://scheduler:Schedulethejoy@eclipsecluster.mvengjn.mongodb.net/?retryWrites=true&w=majority"

func NewDatabase() *Database {
	db := &Database{}

	if err := db.Connect(uri); err != nil {
		log.Fatal(err)
	}

	return db
}

const StudentsCollection = "Students"
const TeachersCollection = "TeachersCollection"

func (dba *Database) Connect(uri string) error {
	ctx := context.Background()

	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri)); err != nil {
		return err
	} else {
		dba.client = client
		dba.db = client.Database("Eclipse-Academy")
		dba.ready = true
		return err
	}
}

func (dba *Database) AddStudent(student *Student) error {
	studentCollection := dba.db.Collection(StudentsCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	_, err := studentCollection.InsertOne(ctx, *student)
	defer cancel()
	return err
}

func (dba *Database) AddTeacher(teacher *Teacher) error {
	teacherCollection := dba.db.Collection(TeachersCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	teacher.Nickname = strings.ToLower(teacher.Nickname)
	_, err := teacherCollection.InsertOne(ctx, *teacher)
	defer cancel()
	return err
}

func (dba *Database) AddSubjects(subjectPayload SubjectPayload) error {
	student, err := dba.GetStudentById(subjectPayload.StudentId)

	if err != nil {
		return errors.New("invalid body, could not find student by id")
	}

	studentSubjects := student.Subjects

	for _, subject := range subjectPayload.Subjects {
		studentSubjects[subject.Name] = SessionData{PricePerSession: subject.PricePerSession, SessionLengthInMinutes: subject.SessionLengthInMinutes}
	}

	studentCollection := dba.db.Collection(StudentsCollection)

	objectId, err := primitive.ObjectIDFromHex(subjectPayload.StudentId)

	if err != nil {
		return err
	}

	var lowercaseSubjects = make(map[string]SessionData)

	for k, v := range studentSubjects {
		computedKey := ComputeSubjectName(k)
		lowercaseSubjects[computedKey] = v
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.M{"$set": bson.D{{Key: "subjects", Value: lowercaseSubjects}}}

	result, err := studentCollection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("failed to modify student `subjects` field")
	}

	return nil
}

func (dba *Database) GetStudentById(studentId string) (*Student, error) {
	studentCollection := dba.db.Collection(StudentsCollection)

	objectId, err := primitive.ObjectIDFromHex(studentId)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objectId}}

	foundStudent := NewStudent()

	err = studentCollection.FindOne(context.Background(), filter).Decode(foundStudent)

	if err != nil {
		return nil, err
	}

	return foundStudent, nil
}

func (dba *Database) GetTeacherById(teacherId string) (*Teacher, error) {
	teacherCollection := dba.db.Collection(TeachersCollection)

	objectId, err := primitive.ObjectIDFromHex(teacherId)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objectId}}

	foundTeacher := NewTeacher()

	err = teacherCollection.FindOne(context.Background(), filter).Decode(foundTeacher)

	if err != nil {
		return nil, err
	}

	return foundTeacher, nil
}

func (dba *Database) GetTeacherByNickname(nickname string) (*Teacher, error) {
	teacherCollection := dba.db.Collection(TeachersCollection)

	filter := bson.D{{Key: "nickname", Value: strings.ToLower(nickname)}}

	foundTeacher := NewTeacher()

	err := teacherCollection.FindOne(context.Background(), filter).Decode(foundTeacher)

	if err != nil {
		return nil, err
	}

	return foundTeacher, nil
}

func (dba *Database) GetStudents() (*[]Student, error) {
	studentCollection := dba.db.Collection(StudentsCollection)

	students := NewStudents()
	ctx := context.Background()

	cursor, err := studentCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	for cursor.Next(ctx) {
		currentStudent := NewStudent()

		if err := cursor.Decode(currentStudent); err != nil {
			return nil, err
		}

		*students = append(*students, *currentStudent)
	}

	return students, err
}

func (dba *Database) GetTeachers() (*[]Teacher, error) {
	teachersCollection := dba.db.Collection(TeachersCollection)

	teachers := NewTeachers()
	ctx := context.Background()

	cursor, err := teachersCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	for cursor.Next(ctx) {
		currentTeacher := NewTeacher()

		if err := cursor.Decode(currentTeacher); err != nil {
			return nil, err
		}

		*teachers = append(*teachers, *currentTeacher)
	}

	return teachers, err
}

func ComputeSubjectName(input string) string {
	splitInput := strings.Split(strings.ToLower(input), " ")
	return strings.Join(splitInput, "_")
}

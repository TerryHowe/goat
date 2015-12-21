package model

import (
	"database/sql"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestModel struct {
	id   int64
	name sql.NullString
}

func (model *TestModel) scan(iterator *ListIterator) bool {
	return iterator.scan(&model.id, &model.name)
}

func (model *TestModel) create(d *Database) bool {
	fields := []string{"name"}
	values := []sql.NullString{model.name}
	model.id = d.create("clusters", fields, values)
	return (model.id != 0)
}

func (model *TestModel) delete(d *Database) bool {
	return d.delete("clusters", model.id)
}

func TestDatabaseOpen(t *testing.T) {
	var sut Database = Database{}
	var iterator ListIterator
	Convey("The result should be true", t, func() {
		So(sut.open(), ShouldEqual, true)
	})
	Convey("Insert a record record should be true", t, func() {
		var testmodel TestModel
		testmodel.id = 0
		testmodel.name.String = "foo"
		testmodel.name.Valid = true
		testmodel.create(&sut)
		So(testmodel.id, ShouldNotEqual, 0)
		Convey("The list should return something", func() {
			iterator = sut.list("clusters")
			So(iterator.rows, ShouldNotEqual, nil)
			Convey("Scan should return record", func() {
				var listmodel TestModel
				var found bool
				for listmodel.scan(&iterator) {
					if listmodel.id == testmodel.id {
						So(listmodel.name.String, ShouldEqual, "foo")
						So(listmodel.name.Valid, ShouldEqual, true)
						found = true
					}
				}
				So(found, ShouldEqual, true)
			})
			Convey("The close should succeed", func() {
				So(iterator.close(), ShouldEqual, true)
			})
		})
		So(testmodel.delete(&sut), ShouldEqual, true)
	})
	Convey("The close should succeed", t, func() {
		So(sut.close(), ShouldEqual, true)
	})
}

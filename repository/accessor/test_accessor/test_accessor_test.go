package accessor

// import (
// 	"context"
// 	"testing"

// 	"github.com/takuya-okada-01/heart-note/repository/database"
// )

// func TestTestAccessor(t *testing.T) {

// 	db := database.Connect()
// 	insertTest := Test{
// 		// Name: "test",
// 	}

// 	testAccessor := NewTestAccessor(db)
// 	err := db.ResetModel(context.Background(), (*Test)(nil))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	id, err := testAccessor.InsertTest(&insertTest)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	want := insertTest.ID

// 	test, err := testAccessor.SelectTest(id)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if test.ID != want {
// 		t.Errorf("InsertTest == %s, want %s", test.ID, want)
// 	}
// }

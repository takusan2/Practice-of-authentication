package accessor

// import (
// 	"context"
// 	"fmt"

// 	"github.com/google/uuid"
// 	"github.com/uptrace/bun"
// )

// type Test struct {
// 	bun.BaseModel `bun:"table:tests,alias:t"`
// 	ID            string `bun:"id,pk"`
// 	// Name          string    `bun:"name,notnull"`
// }

// type TestAccessor interface {
// 	InsertTest(test *Test) (string, error)
// 	SelectTest(id string) (Test, error)
// 	SelectLastTest(id string) (Test, error)
// }

// type testAccessor struct {
// 	db *bun.DB
// }

// func NewTestAccessor(db *bun.DB) TestAccessor {
// 	return &testAccessor{db: db}
// }

// func (t *testAccessor) InsertTest(test *Test) (string, error) {
// 	test.ID = uuid.New().String()
// 	fmt.Print(test.ID)
// 	_, err := t.db.NewInsert().Model(test).Exec(context.Background())
// 	if err != nil {
// 		return "", err
// 	}
// 	return test.ID, nil
// }

// func (t *testAccessor) SelectTest(id string) (Test, error) {
// 	var test Test
// 	err := t.db.NewSelect().Model(&test).Where("id = ?", id).Scan(context.Background())
// 	return test, err
// }

// func (t *testAccessor) SelectLastTest(id string) (Test, error) {
// 	var test Test
// 	err := t.db.NewSelect().Model(&test).Where("id = ?", id).Scan(context.Background())
// 	return test, err
// }

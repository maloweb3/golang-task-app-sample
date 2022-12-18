// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTasks(t *testing.T) {
	t.Parallel()

	query := Tasks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTasksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Tasks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TaskExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Task exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaskExists to return true, but got false.")
	}
}

func testTasksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	taskFound, err := FindTask(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if taskFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTasksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Tasks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTasksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Tasks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTasksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taskOne := &Task{}
	taskTwo := &Task{}
	if err = randomize.Struct(seed, taskOne, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTwo, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tasks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTasksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taskOne := &Task{}
	taskTwo := &Task{}
	if err = randomize.Struct(seed, taskOne, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTwo, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func taskBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func testTasksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Task{}
	o := &Task{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Task object: %s", err)
	}

	AddTaskHook(boil.BeforeInsertHook, taskBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taskBeforeInsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterInsertHook, taskAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taskAfterInsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterSelectHook, taskAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taskAfterSelectHooks = []TaskHook{}

	AddTaskHook(boil.BeforeUpdateHook, taskBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taskBeforeUpdateHooks = []TaskHook{}

	AddTaskHook(boil.AfterUpdateHook, taskAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taskAfterUpdateHooks = []TaskHook{}

	AddTaskHook(boil.BeforeDeleteHook, taskBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taskBeforeDeleteHooks = []TaskHook{}

	AddTaskHook(boil.AfterDeleteHook, taskAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taskAfterDeleteHooks = []TaskHook{}

	AddTaskHook(boil.BeforeUpsertHook, taskBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taskBeforeUpsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterUpsertHook, taskAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taskAfterUpsertHooks = []TaskHook{}
}

func testTasksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTasksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(taskColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTasksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTasksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTasksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tasks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taskDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `Description`: `character varying`, `EndDate`: `date`, `Status`: `character varying`, `Priority`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testTasksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(taskAllColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskDBTypes, true, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTasksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taskAllColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskDBTypes, true, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taskAllColumns, taskPrimaryKeyColumns) {
		fields = taskAllColumns
	} else {
		fields = strmangle.SetComplement(
			taskAllColumns,
			taskPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := TaskSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTasksUpsert(t *testing.T) {
	t.Parallel()

	if len(taskAllColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Task{}
	if err = randomize.Struct(seed, &o, taskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Task: %s", err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, taskDBTypes, false, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Task: %s", err)
	}

	count, err = Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

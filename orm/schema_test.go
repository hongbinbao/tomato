package orm

import (
	"reflect"
	"testing"
	"time"

	"github.com/lfq7413/tomato/errs"
	"github.com/lfq7413/tomato/types"
)

func Test_AddClassIfNotExists(t *testing.T) {
	// validateNewClass
	// TODO
}

func Test_UpdateClass(t *testing.T) {
	// GetOneSchema
	// validateSchemaData
	// deleteField
	// reloadData
	// enforceFieldExists
	// setPermissions
	// TODO
}

func Test_deleteField(t *testing.T) {
	// GetOneSchema
	// TODO
}

func Test_validateObject(t *testing.T) {
	// EnforceClassExists
	// enforceFieldExists
	// thenValidateRequiredColumns
	// TODO
}

func Test_testBaseCLP(t *testing.T) {
	// TODO
}

func Test_validatePermission(t *testing.T) {
	// testBaseCLP
	// TODO
}

func Test_EnforceClassExists(t *testing.T) {
	// AddClassIfNotExists
	// TODO
}

func Test_validateNewClass(t *testing.T) {
	// validateSchemaData
	// TODO
}

func Test_validateSchemaData(t *testing.T) {
	// TODO
}

func Test_validateRequiredColumns(t *testing.T) {
	// TODO
}

func Test_enforceFieldExists(t *testing.T) {
	// getExpectedType
	// TODO
}

func Test_setPermissions(t *testing.T) {
	// reloadData
	// TODO
}

func Test_HasClass(t *testing.T) {
	// reloadData
	// TODO
}

func Test_getExpectedType(t *testing.T) {
	// TODO
}

func Test_reloadData(t *testing.T) {
	// GetAllClasses
	// TODO
}

func Test_GetAllClasses(t *testing.T) {
	// TODO
}

func Test_GetOneSchema(t *testing.T) {
	// TODO
}

////////////////////////////////////////////////////////////

func Test_thenValidateRequiredColumns(t *testing.T) {
	// validateRequiredColumns
	// TODO
}

func Test_getType(t *testing.T) {
	var object interface{}
	var result types.M
	var err error
	var expect interface{}
	/************************************************************/
	object = true
	result, err = getType(object)
	expect = types.M{"type": "Boolean"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = "hello"
	result, err = getType(object)
	expect = types.M{"type": "String"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = 1024
	result, err = getType(object)
	expect = types.M{"type": "Number"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = 10.24
	result, err = getType(object)
	expect = types.M{"type": "Number"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__type": "Date",
		"iso":    "abc",
	}
	result, err = getType(object)
	expect = types.M{"type": "Date"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = map[string]interface{}{
		"__type": "File",
		"name":   "abc",
	}
	result, err = getType(object)
	expect = types.M{"type": "File"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.S{1, 2, 3}
	result, err = getType(object)
	expect = types.M{"type": "Array"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = []interface{}{1, 2, 3}
	result, err = getType(object)
	expect = types.M{"type": "Array"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = time.Now()
	result, err = getType(object)
	expect = errs.E(errs.IncorrectType, "bad obj. can not get type")
	if err == nil || reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
}

func Test_getObjectType(t *testing.T) {
	var object interface{}
	var result types.M
	var err error
	var expect interface{}
	/************************************************************/
	object = []interface{}{1, 2, 3}
	result, err = getObjectType(object)
	expect = types.M{"type": "Array"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__type":    "Pointer",
		"className": "abc",
	}
	result, err = getObjectType(object)
	expect = types.M{
		"type":        "Pointer",
		"targetClass": "abc",
	}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__type": "File",
		"name":   "abc",
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "File"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__type": "Date",
		"iso":    "abc",
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "Date"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__type":    "GeoPoint",
		"latitude":  10,
		"longitude": 10,
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "GeoPoint"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__type": "Bytes",
		"base64": "abc",
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "Bytes"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__type": "Other",
	}
	result, err = getObjectType(object)
	expect = errs.E(errs.IncorrectType, "This is not a valid Other")
	if err == nil || reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"$ne": types.M{
			"__type":    "Pointer",
			"className": "abc",
		},
	}
	result, err = getObjectType(object)
	expect = types.M{
		"type":        "Pointer",
		"targetClass": "abc",
	}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "Increment",
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "Number"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "Delete",
	}
	result, err = getObjectType(object)
	expect = nil
	if err != nil || result != nil {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "Add",
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "Array"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "AddUnique",
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "Array"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "Remove",
	}
	result, err = getObjectType(object)
	expect = types.M{"type": "Array"}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "AddRelation",
		"objects": types.S{
			types.M{
				"className": "abc",
			},
		},
	}
	result, err = getObjectType(object)
	expect = types.M{
		"type":        "Relation",
		"targetClass": "abc",
	}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "RemoveRelation",
		"objects": types.S{
			types.M{
				"className": "abc",
			},
		},
	}
	result, err = getObjectType(object)
	expect = types.M{
		"type":        "Relation",
		"targetClass": "abc",
	}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{
		"__op": "Batch",
		"ops": types.S{
			types.M{
				"__type": "File",
				"name":   "abc",
			},
		},
	}
	result, err = getObjectType(object)
	expect = types.M{
		"type": "File",
	}
	if err != nil || reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{"__op": "Other"}
	result, err = getObjectType(object)
	expect = errs.E(errs.IncorrectType, "unexpected op: Other")
	if err == nil || reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
	/************************************************************/
	object = types.M{"key": "value"}
	result, err = getObjectType(object)
	expect = types.M{"type": "object"}
	if err != nil || reflect.DeepEqual(expect, expect) == false {
		t.Error("expect:", expect, "result:", result, err)
	}
}

func Test_ClassNameIsValid(t *testing.T) {
	var className string
	var ok bool
	var expect bool
	/************************************************************/
	className = "_User"
	ok = ClassNameIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "_Installation"
	ok = ClassNameIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "_Role"
	ok = ClassNameIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "_Session"
	ok = ClassNameIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "_Join:abc:123"
	ok = ClassNameIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "abc"
	ok = ClassNameIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
}

func Test_InvalidClassNameMessage(t *testing.T) {
	var className string
	var result string
	var expect string
	/************************************************************/
	className = "abc"
	result = InvalidClassNameMessage(className)
	expect = "Invalid classname: abc, classnames can only have alphanumeric characters and _, and must start with an alpha character "
	if result != expect {
		t.Error("expect:", expect, "result:", result)
	}
}

func Test_joinClassIsValid(t *testing.T) {
	var className string
	var ok bool
	var expect bool
	/************************************************************/
	className = "_Join:abc:def"
	ok = joinClassIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "_Join:abc123:def123"
	ok = joinClassIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "_Join:_abc123:def_123"
	ok = joinClassIsValid(className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "abc"
	ok = joinClassIsValid(className)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	className = "_Join:@123:!def"
	ok = joinClassIsValid(className)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
}

func Test_fieldNameIsValid(t *testing.T) {
	var fieldName string
	var ok bool
	var expect bool
	/************************************************************/
	fieldName = "abc_123"
	ok = fieldNameIsValid(fieldName)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "abc123"
	ok = fieldNameIsValid(fieldName)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "123abc"
	ok = fieldNameIsValid(fieldName)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "*abc"
	ok = fieldNameIsValid(fieldName)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "abc@123"
	ok = fieldNameIsValid(fieldName)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
}

func Test_fieldNameIsValidForClass(t *testing.T) {
	var fieldName string
	var className string
	var ok bool
	var expect bool
	/************************************************************/
	fieldName = ""
	className = ""
	ok = fieldNameIsValidForClass(fieldName, className)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "abc"
	className = ""
	ok = fieldNameIsValidForClass(fieldName, className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "objectId"
	className = ""
	ok = fieldNameIsValidForClass(fieldName, className)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "abc"
	className = "_User"
	ok = fieldNameIsValidForClass(fieldName, className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "username"
	className = "_User"
	ok = fieldNameIsValidForClass(fieldName, className)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	fieldName = "key"
	className = "class"
	ok = fieldNameIsValidForClass(fieldName, className)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
}

func Test_fieldTypeIsInvalid(t *testing.T) {
	var tp types.M
	var err error
	var expect error
	/************************************************************/
	tp = nil
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.InvalidJSON, "invalid JSON")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.InvalidJSON, "invalid JSON")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{"type": 1024}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.InvalidJSON, "invalid JSON")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{"type": "Pointer"}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.MissingRequiredFieldError, "type Pointer needs a class name")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type":        "Pointer",
		"targetClass": 1024,
	}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.InvalidJSON, "invalid JSON")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type":        "Pointer",
		"targetClass": "@abc",
	}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.InvalidClassName, "Invalid classname: @abc, classnames can only have alphanumeric characters and _, and must start with an alpha character ")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type":        "Pointer",
		"targetClass": "abc",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{"type": "Relation"}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.MissingRequiredFieldError, "type Relation needs a class name")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type":        "Relation",
		"targetClass": 1024,
	}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.InvalidJSON, "invalid JSON")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type":        "Relation",
		"targetClass": "@abc",
	}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.InvalidClassName, "Invalid classname: @abc, classnames can only have alphanumeric characters and _, and must start with an alpha character ")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type":        "Relation",
		"targetClass": "abc",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "Number",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "String",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "Boolean",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "Date",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "Object",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "Array",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "GeoPoint",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "File",
	}
	err = fieldTypeIsInvalid(tp)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	tp = types.M{
		"type": "Other",
	}
	err = fieldTypeIsInvalid(tp)
	expect = errs.E(errs.IncorrectType, "invalid field type: Other")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
}

func Test_validateCLP(t *testing.T) {
	var perms types.M
	var fields types.M
	var err error
	var expect error
	/************************************************************/
	perms = nil
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{}
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"get": types.M{"012345678901234567890123": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"find": types.M{"012345678901234567890123": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"create": types.M{"012345678901234567890123": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"update": types.M{"012345678901234567890123": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"delete": types.M{"012345678901234567890123": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"addField": types.M{"012345678901234567890123": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"other": types.M{"012345678901234567890123": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "other is not a valid operation for class level permissions")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"readUserFields": types.S{"key1", "key2"},
	}
	fields = types.M{
		"key1": types.M{
			"type":        "Pointer",
			"targetClass": "_User",
		},
		"key2": types.M{
			"type":        "Pointer",
			"targetClass": "_User",
		},
	}
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"writeUserFields": types.S{"key1", "key2"},
	}
	fields = types.M{
		"key1": types.M{
			"type":        "Pointer",
			"targetClass": "_User",
		},
		"key2": types.M{
			"type":        "Pointer",
			"targetClass": "_User",
		},
	}
	err = validateCLP(perms, fields)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"readUserFields": "hello",
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "this perms[operation] is not a valid value for class level permissions readUserFields")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"readUserFields": types.S{"key1", "key2"},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "key1 is not a valid column for class level pointer permissions readUserFields")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"readUserFields": types.S{"key1", "key2"},
	}
	fields = types.M{}
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "key1 is not a valid column for class level pointer permissions readUserFields")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"readUserFields": types.S{"key1", "key2"},
	}
	fields = types.M{"key1": 1024}
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "key1 is not a valid column for class level pointer permissions readUserFields")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"readUserFields": types.S{"key1", "key2"},
	}
	fields = types.M{
		"key1": types.M{
			"type": "Other",
		},
	}
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "key1 is not a valid column for class level pointer permissions readUserFields")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"get": types.M{"abc": true},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "abc is not a valid key for class level permissions")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"get": types.M{"role:abc": false},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "false is not a valid value for class level permissions get:role:abc:false")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	perms = types.M{
		"get": types.M{"role:abc": "hello"},
	}
	fields = nil
	err = validateCLP(perms, fields)
	expect = errs.E(errs.InvalidJSON, "this perm is not a valid value for class level permissions get:role:abc:perm")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
}

func Test_verifyPermissionKey(t *testing.T) {
	var key string
	var err error
	var expect error
	/************************************************************/
	key = "0123456789abcdefghij0123"
	err = verifyPermissionKey(key)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	key = "role:1024"
	err = verifyPermissionKey(key)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	key = "*"
	err = verifyPermissionKey(key)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	key = "abcd"
	err = verifyPermissionKey(key)
	expect = errs.E(errs.InvalidJSON, key+" is not a valid key for class level permissions")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	key = "*abc"
	err = verifyPermissionKey(key)
	expect = errs.E(errs.InvalidJSON, key+" is not a valid key for class level permissions")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	key = "role:*abc"
	err = verifyPermissionKey(key)
	expect = nil
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
	/************************************************************/
	key = "@mail"
	err = verifyPermissionKey(key)
	expect = errs.E(errs.InvalidJSON, key+" is not a valid key for class level permissions")
	if reflect.DeepEqual(expect, err) == false {
		t.Error("expect:", expect, "result:", err)
	}
}

func Test_buildMergedSchemaObject(t *testing.T) {
	var existingFields types.M
	var putRequest types.M
	var result types.M
	var expect types.M
	/************************************************************/
	existingFields = nil
	putRequest = nil
	result = buildMergedSchemaObject(existingFields, putRequest)
	expect = types.M{}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	existingFields = types.M{}
	putRequest = types.M{}
	result = buildMergedSchemaObject(existingFields, putRequest)
	expect = types.M{}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	existingFields = types.M{
		"_id":           "_User",
		"objectId":      types.M{"type": "String"},
		"createdAt":     types.M{"type": "Date"},
		"updatedAt":     types.M{"type": "Date"},
		"ACL":           types.M{"type": "ACL"},
		"username":      types.M{"type": "String"},
		"password":      types.M{"type": "String"},
		"email":         types.M{"type": "String"},
		"emailVerified": types.M{"type": "Boolean"},
		"name":          types.M{"type": "String"},
		"skill":         types.M{"type": "Array"},
	}
	putRequest = types.M{
		"age":   types.M{"type": "Number"},
		"skill": types.M{"__op": "Delete"},
	}
	result = buildMergedSchemaObject(existingFields, putRequest)
	expect = types.M{
		"name": types.M{"type": "String"},
		"age":  types.M{"type": "Number"},
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	existingFields = types.M{
		"_id":           "user",
		"objectId":      types.M{"type": "String"},
		"createdAt":     types.M{"type": "Date"},
		"updatedAt":     types.M{"type": "Date"},
		"ACL":           types.M{"type": "ACL"},
		"username":      types.M{"type": "String"},
		"password":      types.M{"type": "String"},
		"email":         types.M{"type": "String"},
		"emailVerified": types.M{"type": "Boolean"},
		"name":          types.M{"type": "String"},
		"skill":         types.M{"type": "Array"},
	}
	putRequest = types.M{
		"age":   types.M{"type": "Number"},
		"skill": types.M{"__op": "Delete"},
	}
	result = buildMergedSchemaObject(existingFields, putRequest)
	expect = types.M{
		"username":      types.M{"type": "String"},
		"password":      types.M{"type": "String"},
		"email":         types.M{"type": "String"},
		"emailVerified": types.M{"type": "Boolean"},
		"name":          types.M{"type": "String"},
		"age":           types.M{"type": "Number"},
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
}

func Test_injectDefaultSchema(t *testing.T) {
	var schema types.M
	var result types.M
	var expect types.M
	/************************************************************/
	schema = nil
	result = injectDefaultSchema(schema)
	expect = nil
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "user",
	}
	result = injectDefaultSchema(schema)
	expect = types.M{
		"className": "user",
		"fields": types.M{
			"objectId":  types.M{"type": "String"},
			"createdAt": types.M{"type": "Date"},
			"updatedAt": types.M{"type": "Date"},
			"ACL":       types.M{"type": "ACL"},
		},
		"classLevelPermissions": nil,
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "user",
		"fields": types.M{
			"key": types.M{"type": "String"},
		},
	}
	result = injectDefaultSchema(schema)
	expect = types.M{
		"className": "user",
		"fields": types.M{
			"objectId":  types.M{"type": "String"},
			"createdAt": types.M{"type": "Date"},
			"updatedAt": types.M{"type": "Date"},
			"ACL":       types.M{"type": "ACL"},
			"key":       types.M{"type": "String"},
		},
		"classLevelPermissions": nil,
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "_User",
		"fields": types.M{
			"key": types.M{"type": "String"},
		},
	}
	result = injectDefaultSchema(schema)
	expect = types.M{
		"className": "_User",
		"fields": types.M{
			"objectId":      types.M{"type": "String"},
			"createdAt":     types.M{"type": "Date"},
			"updatedAt":     types.M{"type": "Date"},
			"ACL":           types.M{"type": "ACL"},
			"key":           types.M{"type": "String"},
			"username":      types.M{"type": "String"},
			"password":      types.M{"type": "String"},
			"email":         types.M{"type": "String"},
			"emailVerified": types.M{"type": "Boolean"},
		},
		"classLevelPermissions": nil,
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "_User",
		"fields": types.M{
			"key": types.M{"type": "String"},
		},
		"classLevelPermissions": types.M{
			"find": types.M{"*": true},
		},
	}
	result = injectDefaultSchema(schema)
	expect = types.M{
		"className": "_User",
		"fields": types.M{
			"objectId":      types.M{"type": "String"},
			"createdAt":     types.M{"type": "Date"},
			"updatedAt":     types.M{"type": "Date"},
			"ACL":           types.M{"type": "ACL"},
			"key":           types.M{"type": "String"},
			"username":      types.M{"type": "String"},
			"password":      types.M{"type": "String"},
			"email":         types.M{"type": "String"},
			"emailVerified": types.M{"type": "Boolean"},
		},
		"classLevelPermissions": types.M{
			"find": types.M{"*": true},
		},
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
}

func Test_convertSchemaToAdapterSchema(t *testing.T) {
	var schema types.M
	var result types.M
	var expect types.M
	/************************************************************/
	schema = nil
	result = convertSchemaToAdapterSchema(schema)
	expect = nil
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "user",
	}
	result = convertSchemaToAdapterSchema(schema)
	expect = types.M{
		"className": "user",
		"fields": types.M{
			"objectId":  types.M{"type": "String"},
			"createdAt": types.M{"type": "Date"},
			"updatedAt": types.M{"type": "Date"},
			"_rperm":    types.M{"type": "Array"},
			"_wperm":    types.M{"type": "Array"},
		},
		"classLevelPermissions": nil,
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "_User",
		"fields": types.M{
			"key": types.M{"type": "String"},
		},
	}
	result = convertSchemaToAdapterSchema(schema)
	expect = types.M{
		"className": "_User",
		"fields": types.M{
			"objectId":         types.M{"type": "String"},
			"createdAt":        types.M{"type": "Date"},
			"updatedAt":        types.M{"type": "Date"},
			"key":              types.M{"type": "String"},
			"username":         types.M{"type": "String"},
			"_hashed_password": types.M{"type": "String"},
			"email":            types.M{"type": "String"},
			"emailVerified":    types.M{"type": "Boolean"},
			"_rperm":           types.M{"type": "Array"},
			"_wperm":           types.M{"type": "Array"},
		},
		"classLevelPermissions": nil,
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
}

func Test_convertAdapterSchemaToParseSchema(t *testing.T) {
	var schema types.M
	var result types.M
	var expect types.M
	/************************************************************/
	schema = nil
	result = convertAdapterSchemaToParseSchema(schema)
	expect = nil
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{}
	result = convertAdapterSchemaToParseSchema(schema)
	expect = types.M{}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"fields": types.M{
			"_rperm": types.M{"type": "Array"},
			"_wperm": types.M{"type": "Array"},
			"key":    types.M{"type": "String"},
		},
	}
	result = convertAdapterSchemaToParseSchema(schema)
	expect = types.M{
		"fields": types.M{
			"key": types.M{"type": "String"},
			"ACL": types.M{"type": "ACL"},
		},
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "_User",
		"fields": types.M{
			"_rperm":           types.M{"type": "Array"},
			"_wperm":           types.M{"type": "Array"},
			"key":              types.M{"type": "String"},
			"authData":         types.M{"type": "String"},
			"_hashed_password": types.M{"type": "String"},
		},
	}
	result = convertAdapterSchemaToParseSchema(schema)
	expect = types.M{
		"className": "_User",
		"fields": types.M{
			"key":      types.M{"type": "String"},
			"ACL":      types.M{"type": "ACL"},
			"password": types.M{"type": "String"},
		},
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
	/************************************************************/
	schema = types.M{
		"className": "other",
		"fields": types.M{
			"_rperm":           types.M{"type": "Array"},
			"_wperm":           types.M{"type": "Array"},
			"key":              types.M{"type": "String"},
			"authData":         types.M{"type": "String"},
			"_hashed_password": types.M{"type": "String"},
		},
	}
	result = convertAdapterSchemaToParseSchema(schema)
	expect = types.M{
		"className": "other",
		"fields": types.M{
			"key":              types.M{"type": "String"},
			"ACL":              types.M{"type": "ACL"},
			"authData":         types.M{"type": "String"},
			"_hashed_password": types.M{"type": "String"},
		},
	}
	if reflect.DeepEqual(expect, result) == false {
		t.Error("expect:", expect, "result:", result)
	}
}

func Test_dbTypeMatchesObjectType(t *testing.T) {
	var dbType types.M
	var objectType types.M
	var ok bool
	var expect bool
	/************************************************************/
	dbType = nil
	objectType = nil
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	dbType = types.M{}
	objectType = nil
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	dbType = nil
	objectType = types.M{}
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	dbType = types.M{"type": "String"}
	objectType = types.M{}
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	dbType = types.M{"type": "String"}
	objectType = types.M{"type": "Date"}
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	dbType = types.M{"type": "Pointer", "targetClass": "abc"}
	objectType = types.M{"type": "Pointer", "targetClass": "def"}
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = false
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	dbType = types.M{"type": "Pointer", "targetClass": "abc"}
	objectType = types.M{"type": "Pointer", "targetClass": "abc"}
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
	/************************************************************/
	dbType = types.M{"type": "String"}
	objectType = types.M{"type": "String"}
	ok = dbTypeMatchesObjectType(dbType, objectType)
	expect = true
	if ok != expect {
		t.Error("expect:", expect, "result:", ok)
	}
}

func Test_Load(t *testing.T) {
	// reloadData
	// TODO
}

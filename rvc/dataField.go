package rvc

type FieldDataType string
type FieldName string

const (
	F64  FieldDataType = "float64"
	U8   FieldDataType = "uint8"
	BIT2 FieldDataType = "bit2"

	instance          FieldName = "instance"
	PRIORITY          FieldName = "priority"
	VOLTAGE           FieldName = "voltage"
	CURRENT           FieldName = "current"
	group             FieldName = "group"
	brightness        FieldName = "brightness"
	lock              FieldName = "lockitem"
	overCurrentStatus FieldName = "overCurrentStatus"
	enableStatus      FieldName = "enableStatus"
	delayDuration     FieldName = "delayDuration"
	lastCommand       FieldName = "lastCommand"
	interlockStatus   FieldName = "interlockStatus"
	loadStatus        FieldName = "loadStatus"
	reserved          FieldName = "reserved"
	undercurrent      FieldName = "undercurrent"
	masterMemoryValue FieldName = "master memory value"
)

type dataField struct {
	name      FieldName
	fieldType FieldDataType
}

func (f *dataField) GetName() FieldName {
	return f.name
}

func (f *dataField) Gettype() FieldDataType {
	return f.fieldType
}

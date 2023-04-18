package rvc

type FieldDataType string
type FieldName string

const (
	F64  FieldDataType = "float64"
	U8   FieldDataType = "uint8"
	BIT2 FieldDataType = "bit2"

	INSTANCE            FieldName = "instance"
	PRIORITY            FieldName = "priority"
	VOLTAGE             FieldName = "voltage"
	CURRENT             FieldName = "current"
	GROUP               FieldName = "group"
	BRIGHTNESS          FieldName = "brightness"
	LOCK                FieldName = "lockitem"
	OVER_CURRENT_STATUS FieldName = "overCurrentStatus"
	ENABLE_STATUS       FieldName = "enableStatus"
	DELAY_DURATION      FieldName = "delayDuration"
	LAST_COMMAND        FieldName = "lastCommand"
	INTERLOCK_STATUS    FieldName = "interlockStatus"
	LOAD_STATUS         FieldName = "loadStatus"
	RESERVED            FieldName = "reserved"
	UNDERCURRENT        FieldName = "undercurrent"
	MASTERMEMVAL        FieldName = "master memory value"
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

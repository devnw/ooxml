package document

import "fmt"

const _FormFieldType_name = "FormFieldTypeUnknownFormFieldTypeTextFormFieldTypeCheckBoxFormFieldTypeDropDown"

var _FormFieldType_index = [...]uint8{0, 20, 37, 58, 79}

func (i FormFieldType) String() string {
	if i >= FormFieldType(len(_FormFieldType_index)-1) {
		return fmt.Sprintf("FormFieldType(%d)", i)
	}
	return _FormFieldType_name[_FormFieldType_index[i]:_FormFieldType_index[i+1]]
}

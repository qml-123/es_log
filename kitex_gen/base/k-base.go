// Code generated by Kitex v0.5.2. DO NOT EDIT.

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *BaseData) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetLogId bool = false
	var issetCode bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetLogId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetCode = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetLogId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetCode {
		fieldId = 6
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseData[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_BaseData[fieldId]))
}

func (p *BaseData) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.LogId = v

	}
	return offset, nil
}

func (p *BaseData) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.FromService = &v

	}
	return offset, nil
}

func (p *BaseData) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.ToService = &v

	}
	return offset, nil
}

func (p *BaseData) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.ToMethod = &v

	}
	return offset, nil
}

func (p *BaseData) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Locate = &v

	}
	return offset, nil
}

func (p *BaseData) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Code = v

	}
	return offset, nil
}

func (p *BaseData) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Message = &v

	}
	return offset, nil
}

// for compatibility
func (p *BaseData) FastWrite(buf []byte) int {
	return 0
}

func (p *BaseData) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "BaseData")
	if p != nil {
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *BaseData) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("BaseData")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *BaseData) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "log_id", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.LogId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *BaseData) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetFromService() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "from_service", thrift.STRING, 2)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.FromService)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *BaseData) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetToService() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "to_service", thrift.STRING, 3)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.ToService)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *BaseData) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetToMethod() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "to_method", thrift.STRING, 4)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.ToMethod)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *BaseData) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetLocate() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "locate", thrift.STRING, 5)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.Locate)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *BaseData) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "code", thrift.I32, 6)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Code)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *BaseData) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetMessage() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "message", thrift.STRING, 7)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.Message)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *BaseData) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("log_id", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.LogId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *BaseData) field2Length() int {
	l := 0
	if p.IsSetFromService() {
		l += bthrift.Binary.FieldBeginLength("from_service", thrift.STRING, 2)
		l += bthrift.Binary.StringLengthNocopy(*p.FromService)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *BaseData) field3Length() int {
	l := 0
	if p.IsSetToService() {
		l += bthrift.Binary.FieldBeginLength("to_service", thrift.STRING, 3)
		l += bthrift.Binary.StringLengthNocopy(*p.ToService)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *BaseData) field4Length() int {
	l := 0
	if p.IsSetToMethod() {
		l += bthrift.Binary.FieldBeginLength("to_method", thrift.STRING, 4)
		l += bthrift.Binary.StringLengthNocopy(*p.ToMethod)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *BaseData) field5Length() int {
	l := 0
	if p.IsSetLocate() {
		l += bthrift.Binary.FieldBeginLength("locate", thrift.STRING, 5)
		l += bthrift.Binary.StringLengthNocopy(*p.Locate)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *BaseData) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("code", thrift.I32, 6)
	l += bthrift.Binary.I32Length(p.Code)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *BaseData) field7Length() int {
	l := 0
	if p.IsSetMessage() {
		l += bthrift.Binary.FieldBeginLength("message", thrift.STRING, 7)
		l += bthrift.Binary.StringLengthNocopy(*p.Message)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

package eth

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

func decodeTxParams(abi *abi.ABI, v map[string]interface{}, data []byte) error {
	m, err := abi.MethodById(data[:4])
	if err != nil {
		return err
	}
	if err := m.Inputs.UnpackIntoMap(v, data[4:]); err != nil {
		return err
	}
	for k, val := range v {
		v[k] = ethTypeToStringyType(val)
	}
	return nil
}

func ethTypeToStringyType(v interface{}) interface{} {
	val := reflect.Indirect(reflect.ValueOf(v))

	switch vTy := val.Interface().(type) {
	case []byte:
		return "0x" + ethcommon.Bytes2Hex(vTy)
	case [32]byte:
		return fmt.Sprintf("0x%x", vTy)
	case ethcommon.Address:
		return vTy.Hex()
	case ethcommon.Hash:
		return "0x" + vTy.Hex()
	case big.Int:
		return vTy.String()
	default:
		return handleComplexEthType(val)
	}
}

func handleComplexEthType(val reflect.Value) interface{} {
	switch val.Kind() {
	// tuple
	case reflect.Struct:
		vString := "{"
		for i := 0; i < val.NumField(); i++ {
			vString += fmt.Sprintf(" %v", val.Type().Field(i).Name)
			vString += ": "
			vString += fmt.Sprintf("%v ", ethTypeToStringyType(val.Field(i).Interface()))
		}
		vString += "}"
		return vString
	case reflect.Array:
		return handleEthSlice(val)
	case reflect.Slice:
		return handleEthSlice(val)
	default:
		return val.Interface()
	}
}

func handleEthSlice(val reflect.Value) string {
	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		return ""
	}
	vString := "["
	for i := 0; i < val.Len(); i++ {
		vString += fmt.Sprintf(" %v ", ethTypeToStringyType(val.Index(i).Interface()))
	}
	vString += "]"
	return vString
}

func RandAddress() (ethcommon.Address, error) {
	b, err := RandBytes(ethcommon.AddressLength)
	if err != nil {
		return ethcommon.BytesToAddress([]byte{}), err
	}
	return ethcommon.BytesToAddress(b), nil
}

func RandBytes(len uint) ([]byte, error) {
	data := make([]byte, len)
	_, err := rand.Read(data)

	return data, err
}

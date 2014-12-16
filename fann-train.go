package fann

/*
#cgo LDFLAGS: -lfann -lm
#include <fann.h>

static void cpFannTypeArray(fann_type* src, fann_type* dst, unsigned int n) {
	unsigned int i = 0;
	for(; i < n; i++)
		dst[i] = src[i];
}

static void get_train_input(struct fann_train_data* td, fann_type* dst, unsigned int pos, unsigned int ln) {
	cpFannTypeArray(td->input[pos], dst, ln);
}

static void get_train_output(struct fann_train_data* td, fann_type* dst, unsigned int pos, unsigned int ln) {
	cpFannTypeArray(td->output[pos], dst, ln);
}

//  Creates an empty set of training data
//  copied from fann 2.2: http://sourceforge.net/p/fann/code/ci/master/tree/src/fann_train_data.c [881e35]
struct fann_train_data * fann_create_train(unsigned int num_data, unsigned int num_input, unsigned int num_output)
{
	fann_type *data_input, *data_output;
	unsigned int i;
	struct fann_train_data *data =
		(struct fann_train_data *) malloc(sizeof(struct fann_train_data));

	if(data == NULL)
	{
		fann_error(NULL, FANN_E_CANT_ALLOCATE_MEM);
		return NULL;
	}

	fann_init_error_data((struct fann_error *) data);

	data->num_data = num_data;
	data->num_input = num_input;
	data->num_output = num_output;
	data->input = (fann_type **) calloc(num_data, sizeof(fann_type *));
	if(data->input == NULL)
	{
		fann_error(NULL, FANN_E_CANT_ALLOCATE_MEM);
		fann_destroy_train(data);
		return NULL;
	}

	data->output = (fann_type **) calloc(num_data, sizeof(fann_type *));
	if(data->output == NULL)
	{
		fann_error(NULL, FANN_E_CANT_ALLOCATE_MEM);
		fann_destroy_train(data);
		return NULL;
	}

	data_input = (fann_type *) calloc(num_input * num_data, sizeof(fann_type));
	if(data_input == NULL)
	{
		fann_error(NULL, FANN_E_CANT_ALLOCATE_MEM);
		fann_destroy_train(data);
		return NULL;
	}

	data_output = (fann_type *) calloc(num_output * num_data, sizeof(fann_type));
	if(data_output == NULL)
	{
		fann_error(NULL, FANN_E_CANT_ALLOCATE_MEM);
		fann_destroy_train(data);
		return NULL;
	}

	for(i = 0; i != num_data; i++)
	{
		data->input[i] = data_input;
		data_input += num_input;
		data->output[i] = data_output;
		data_output += num_output;
	}
	return data;
}

*/
import "C"
import (
	"reflect"
	"unsafe"
)

func ReadTrainFromData(input, output [][]FannType) *TrainData {
	num_data := len(input)
	num_input := len(input[0])
	num_output := len(output[0])
	data := C.fann_create_train(C.uint(num_data), C.uint(num_input), C.uint(num_output))
	// map input
	inputHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(*data.input)),
		Len:  num_data * num_input,
		Cap:  num_data * num_input,
	}
	inputSlice := *(*[]FannType)(unsafe.Pointer(&inputHdr))
	k := 0
	for i := 0; i < num_data; i++ {
		for j := 0; j < num_input; j++ {
			inputSlice[k] = input[i][j]
			k++
		}
	}
	// map output
	outputHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(*data.output)),
		Len:  num_data * num_output,
		Cap:  num_data * num_output,
	}
	outputSlice := *(*[]FannType)(unsafe.Pointer(&outputHdr))
	k = 0
	for i := 0; i < num_data; i++ {
		for j := 0; j < num_output; j++ {
			outputSlice[k] = output[i][j]
			k++
		}
	}
	//------
	var td TrainData
	td.object = data
	return &td
}

func ReadTrainFromFile(filename string) *TrainData {
	var td TrainData

	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))

	td.object = C.fann_read_train_from_file(cfn)

	return &td
}

func (td *TrainData) Destroy() {
	C.fann_destroy_train(td.object)
}

func (td *TrainData) GetInput(i uint32) []FannType {
	num := td.GetNumInput()
	input := make([]FannType, num)
	C.get_train_input(td.object, (*C.fann_type)(&input[0]), C.uint(i), C.uint(num))
	return input
}

func (td *TrainData) GetOutput(i uint32) []FannType {
	num := td.GetNumOutput()
	output := make([]FannType, num)
	C.get_train_output(td.object, (*C.fann_type)(&output[0]), C.uint(i), C.uint(num))
	return output
}

func (td *TrainData) Shuffle() {
	C.fann_shuffle_train_data(td.object)
}

func (ann *Ann) ScaleTrain(td *TrainData) {
	C.fann_scale_train(ann.object, td.object)
}

func (ann *Ann) DescaleTrain(td *TrainData) {
	C.fann_descale_train(ann.object, td.object)
}

func (td *TrainData) Length() uint32 {
	return uint32(C.fann_length_train_data(td.object))
}

func MergeTrainData(td1 *TrainData, td2 *TrainData) *TrainData {
	var td TrainData
	td.object = C.fann_merge_train_data(td1.object, td2.object)
	return &td
}

func (td *TrainData) Duplicate() *TrainData {
	var td_dup TrainData
	td_dup.object = C.fann_duplicate_train_data(td.object)
	return &td_dup
}

func (td *TrainData) Subset(pos uint32, length uint32) *TrainData {
	var td_sub TrainData
	td_sub.object = C.fann_subset_train_data(td.object, C.uint(pos), C.uint(length))
	return &td_sub
}

func (td *TrainData) GetNumInput() uint32 {
	return uint32(C.fann_num_input_train_data(td.object))
}

func (td *TrainData) GetNumOutput() uint32 {
	return uint32(C.fann_num_output_train_data(td.object))
}

func (td *TrainData) SaveTrain(filename string) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_save_train(td.object, cfn)
}

func (td *TrainData) SaveTrainToFixed(filename string, decimal_point uint32) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))

	C.fann_save_train_to_fixed(td.object, cfn, C.uint(decimal_point))
}

func (td *TrainData) ScaleInputTrainData(new_min FannType, new_max FannType) {
	C.fann_scale_input_train_data(td.object, C.fann_type(new_min), C.fann_type(new_max))
}

func (td *TrainData) ScaleOutputTrainData(new_min FannType, new_max FannType) {
	C.fann_scale_output_train_data(td.object, C.fann_type(new_min), C.fann_type(new_max))
}

func (td *TrainData) ScaleTrainData(new_min FannType, new_max FannType) {
	C.fann_scale_train_data(td.object, C.fann_type(new_min), C.fann_type(new_max))
}

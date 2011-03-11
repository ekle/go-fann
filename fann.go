package fann

/*
#include <fann.h>

typedef unsigned int* uint_ptr;

const unsigned int* go2uintArray(unsigned int* arr, int n) {
	return malloc(sizeof(unsigned int) * n);
}

static void cpFannTypeArray(fann_type* src, fann_type* dst, unsigned int n) {
	unsigned int i = 0;
	for(; i < n; i++)
		dst[i] = src[i];
}
*/
import "C"
import "unsafe"

type FannType C.fann_type

type TrainData struct {
	object *C.struct_fann_train_data
}

func ReadFromFile(fn string) {
}

type Ann struct {
	object *C.struct_fann
}

//Create ann functions
func CreateStandart(numLayers uint, layers []uint32) (*Ann) {
	var ann Ann
	ann.object = C.fann_create_standard_array(C.uint(numLayers), (*C.uint)(&layers[0]))
	return &ann
}

func CreateSparse(concentration float32, numLayers uint, layers []uint32) (*Ann) {
	var ann Ann
	ann.object = C.fann_create_sparse_array(C.float(concentration), C.uint(numLayers), (*C.uint)(&layers[0]))
	return &ann
}

func CreateShortcut(num_layers uint32, layers []uint32) (*Ann) {
	var ann Ann
	ann.object = C.fann_create_shortcut_array(C.uint(num_layers), (*C.uint)(&layers[0]))
	return &ann
}

func CreateFromFile(filename string) (*Ann) {
	var ann Ann
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	ann.object = C.fann_create_from_file(cfn)
	return &ann
}

//run & test functions
func (ann *Ann) Run(input []FannType) ([]FannType) {
	c_out := C.fann_run(ann.object, (*C.fann_type)(&input[0]))
	outputNum := ann.GetNumOutput()
	out := make([]FannType, outputNum)
	C.cpFannTypeArray(c_out, (*C.fann_type)(&out[0]), C.uint(outputNum))
	return out
}

//train functions
func (ann *Ann) TrainOnFile(filename string, maxEpoches uint32, epochBetweenReports uint32, desiredError float32) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_train_on_file(ann.object, cfn, C.uint(maxEpoches), C.uint(epochBetweenReports), C.float(desiredError));
}


func (ann *Ann) RandomizeWeights(min_weight FannType, max_weight FannType) ( ) {
	C.fann_randomize_weights(ann.object, C.fann_type(min_weight), C.fann_type(max_weight))
}
//print functions
func (ann *Ann) PrintConnections() ( ) {
	C.fann_print_connections(ann.object)
}

func (ann *Ann) PrintParameters() ( ) {
	C.fann_print_parameters(ann.object)
}

func (ann *Ann) SetActivationFunctionHidden(tp ActivationFunc) {
	C.fann_set_activation_function_hidden(ann.object, C.enum_fann_activationfunc_enum(tp))
}

func (ann *Ann) SetActivationFunctionOutput(tp ActivationFunc) {
	C.fann_set_activation_function_output(ann.object, C.enum_fann_activationfunc_enum(tp))
}
//save functions
func (ann *Ann) Save(filename string) {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_save(ann.object, cfn)
}

func (ann *Ann) SaveToFixed(configuration_file string) {
	cfn := C.CString(configuration_file)
	defer C.free(unsafe.Pointer(cfn))
	C.fann_save_to_fixed(ann.object, cfn)
}
//destroy function
func (ann *Ann) Destroy() {
	C.fann_destroy(ann.object)
}

//getters
func (ann *Ann) GetNumInput() (uint32) {
	return uint32(C.fann_get_num_input(ann.object))
}

func (ann *Ann) GetNumOutput() (uint32) {
	return uint32(C.fann_get_num_output(ann.object))
}

func (ann *Ann) GetTotalNeurons() (uint32) {
	return uint32(C.fann_get_total_neurons(ann.object))
}

func (ann *Ann) GetTotalConnections() (uint32) {
	return uint32(C.fann_get_total_connections(ann.object))
}

func (ann *Ann) GetConnectionRate() (float32) {
	return float32(C.fann_get_connection_rate(ann.object))
}

func (ann *Ann) GetNumLayers() (uint32) {
	return uint32(C.fann_get_num_layers(ann.object))
}
/*
func (ann *Ann) GetDecimalPoint() (uint32) {
	return uint32(C.fann_get_decimal_point(ann.object))
}

func (ann *Ann) GetMultiplier() (uint32) {
	return uint32(C.fann_get_multiplier(ann.object))
}
*/

//list all activation functions
type ActivationFunc C.enum_fann_activationfunc_enum


var FANN_LINEAR ActivationFunc = C.FANN_LINEAR
var FANN_THRESHOLD ActivationFunc = C.FANN_THRESHOLD
var FANN_THRESHOLD_SYMMETRIC ActivationFunc = C.FANN_THRESHOLD_SYMMETRIC
var FANN_SIGMOID ActivationFunc = C.FANN_SIGMOID
var FANN_SIGMOID_STEPWISE ActivationFunc = C.FANN_SIGMOID_STEPWISE
var FANN_SIGMOID_SYMMETRIC ActivationFunc = C.FANN_SIGMOID_SYMMETRIC
var FANN_GAUSSIAN ActivationFunc = C.FANN_GAUSSIAN
var FANN_GAUSSIAN_SYMMETRIC ActivationFunc = C.FANN_GAUSSIAN_SYMMETRIC
var FANN_GAUSSIAN_STEPWISE ActivationFunc = C.FANN_GAUSSIAN_STEPWISE
var FANN_ELLIOT ActivationFunc = C.FANN_ELLIOT
var FANN_ELLIOT_SYMMETRIC ActivationFunc = C.FANN_ELLIOT_SYMMETRIC
var FANN_LINEAR_PIECE ActivationFunc = C.FANN_LINEAR_PIECE
var FANN_LINEAR_PIECE_SYMMETRIC ActivationFunc = C.FANN_LINEAR_PIECE_SYMMETRIC
var FANN_SIN_SYMMETRIC ActivationFunc = C.FANN_SIN_SYMMETRIC
var FANN_COS_SYMMETRIC ActivationFunc = C.FANN_COS_SYMMETRIC
var FANN_SIN ActivationFunc = C.FANN_SIN
var FANN_COS ActivationFunc = C.FANN_COS


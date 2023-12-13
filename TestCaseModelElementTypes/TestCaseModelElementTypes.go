package TestCaseModelElementTypes

import "github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"

const (
	TestCaseModelElementType_B0          TypeAndStructs.TestCaseModelElementTypeType          = "B0_BOND"
	TestCaseModelElementDescription_B0   TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with zero connection. This is the base for any TestCase: \"B0\""
	TestCaseModelElementGrpcMappingId_B0 TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 1

	TestCaseModelElementType_B1f            TypeAndStructs.TestCaseModelElementTypeType          = "B1f_BOND_NONE_SWAPPABLE"
	TestCaseModelElementTypeDescription_B1f TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with one connection. Use for starting a TestCase and is only found in pre-created TestInstructionContainersMap"
	TestCaseModelElementGrpcMappingId_B1f   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 2

	TestCaseModelElementType_B1l            TypeAndStructs.TestCaseModelElementTypeType          = "B1l_BOND_NONE_SWAPPABLE"
	TestCaseModelElementTypeDescription_B1l TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with one connection. Use for ending a TestCase and is only found in pre-created TestInstructionContainersMap"
	TestCaseModelElementGrpcMappingId_B1l   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 3

	TestCaseModelElementType_B10            TypeAndStructs.TestCaseModelElementTypeType          = "B10_BOND"
	TestCaseModelElementTypeDescription_B10 TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with zero connections which appear as an element in a new"
	TestCaseModelElementGrpcMappingId_B10   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 4

	TestCaseModelElementType_B11f            TypeAndStructs.TestCaseModelElementTypeType          = "B11f_BOND"
	TestCaseModelElementTypeDescription_B11f TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with one connection which appear as first element inside a TestInstructionContainer"
	TestCaseModelElementGrpcMappingId_B11f   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 5

	TestCaseModelElementType_B11l            TypeAndStructs.TestCaseModelElementTypeType          = "B11l_BOND"
	TestCaseModelElementTypeDescription_B11l TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with one connection which appear as last element inside a TestInstructionContainer"
	TestCaseModelElementGrpcMappingId_B11l   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 6

	TestCaseModelElementType_B12            TypeAndStructs.TestCaseModelElementTypeType          = "B12_BOND"
	TestCaseModelElementTypeDescription_B12 TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with two connections which connects TestInstructionsMap or TestInstructionContainersMap"
	TestCaseModelElementGrpcMappingId_B12   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 7

	TestCaseModelElementType_B10oxo            TypeAndStructs.TestCaseModelElementTypeType          = "B10oxo_BOND"
	TestCaseModelElementTypeDescription_B10oxo TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x*) into TIC(B11fx-n-B11lx)."
	TestCaseModelElementGrpcMappingId_B10oxo   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 8

	TestCaseModelElementType_B10ox            TypeAndStructs.TestCaseModelElementTypeType          = "B10ox_BOND"
	TestCaseModelElementTypeDescription_B10ox TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10*x) into TIC(B11fx-n-B11l)."
	TestCaseModelElementGrpcMappingId_B10ox   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 9

	TestCaseModelElementType_B10xo            TypeAndStructs.TestCaseModelElementTypeType          = "B10xo_BOND"
	TestCaseModelElementTypeDescription_B10xo TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with zero connections. This Bond is used in Pre-created containers and is used when user swap TIC(B10x*) into TIC(B11f-n-B11lx)."
	TestCaseModelElementGrpcMappingId_B10xo   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 10

	TestCaseModelElementType_B11fx            TypeAndStructs.TestCaseModelElementTypeType          = "B11fx_BOND_NONE_SWAPPABLE"
	TestCaseModelElementTypeDescription_B11fx TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with one connection which appear as first element inside a TestInstructionContainer. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructionsMap or TestInstructionsContainers before(first) structure."
	TestCaseModelElementGrpcMappingId_B11fx   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 11

	TestCaseModelElementType_B11lx            TypeAndStructs.TestCaseModelElementTypeType          = "B11lx_BOND_NONE_SWAPPABLE"
	TestCaseModelElementTypeDescription_B11lx TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with two connections which connects TestInstructionsMap or TestInstructionContainersMap. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructionsMap or TestInstructionsContainers within a pre-created container."
	TestCaseModelElementGrpcMappingId_B11lx   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 12

	TestCaseModelElementType_B12x            TypeAndStructs.TestCaseModelElementTypeType          = "B12x_BOND_NONE_SWAPPABLE"
	TestCaseModelElementTypeDescription_B12x TypeAndStructs.TestCaseModelElementDescriptionType   = "Bond with two connections which connects TestInstructionsMap or TestInstructionContainersMap. This Bond is used in Pre-created containers and is used to stop a user to add new TestInstructionsMap or TestInstructionsContainers within a pre-created container."
	TestCaseModelElementGrpcMappingId_B12x   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 13

	TestCaseModelElementType_TI            TypeAndStructs.TestCaseModelElementTypeType          = "TI_TESTINSTRUCTION"
	TestCaseModelElementTypeDescription_TI TypeAndStructs.TestCaseModelElementDescriptionType   = "TestInstruction"
	TestCaseModelElementGrpcMappingId_TI   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 14

	TestCaseModelElementType_TIx            TypeAndStructs.TestCaseModelElementTypeType          = "TIx_TESTINSTRUCTION_NONE_REMOVABLE"
	TestCaseModelElementTypeDescription_TIx TypeAndStructs.TestCaseModelElementDescriptionType   = "A standard TestInstruction which can not be removed or swapped out by the user"
	TestCaseModelElementGrpcMappingId_TIx   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 15

	TestCaseModelElementType_TIC            TypeAndStructs.TestCaseModelElementTypeType          = "TIC_TESTINSTRUCTIONCONTAINER"
	TestCaseModelElementTypeDescription_TIC TypeAndStructs.TestCaseModelElementDescriptionType   = "TestInstructionContainer(X) where X is any valid structure. Children in TestExecutionContainer is executed in serial mode"
	TestCaseModelElementGrpcMappingId_TIC   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 16

	TestCaseModelElementType_TICx            TypeAndStructs.TestCaseModelElementTypeType          = "TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE"
	TestCaseModelElementTypeDescription_TICx TypeAndStructs.TestCaseModelElementDescriptionType   = "A standard TestInstructionContainer(X) where X is any valid structure which can not be removed or swapped out by the user. Children in TestExecutionContainer is executed in serial mode"
	TestCaseModelElementGrpcMappingId_TICx   TypeAndStructs.TestCaseModelElementGrpcMappingIdType = 17
)

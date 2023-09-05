package staticcheck

import (
	"testing"

	"honnef.co/go/tools/analysis/lint/testutil"
)

func TestAll(t *testing.T) {
	checks := map[string][]testutil.Test{
		"SA1000": {{Dir: "example.com/CheckRegexps"}},
		"SA1001": {{Dir: "example.com/CheckTemplate"}},
		"SA1002": {{Dir: "example.com/CheckTimeParse"}},
		"SA1003": {{Dir: "example.com/CheckEncodingBinary"}, {Dir: "example.com/CheckEncodingBinary_go17", Version: "1.7"}, {Dir: "example.com/CheckEncodingBinary_go18", Version: "1.8"}},
		"SA1004": {{Dir: "example.com/CheckTimeSleepConstant"}},
		"SA1005": {{Dir: "example.com/CheckExec"}},
		"SA1006": {{Dir: "example.com/CheckUnsafePrintf"}},
		"SA1007": {{Dir: "example.com/CheckURLs"}},
		"SA1008": {{Dir: "example.com/CheckCanonicalHeaderKey"}},
		"SA1010": {{Dir: "example.com/checkStdlibUsageRegexpFindAll"}},
		"SA1011": {{Dir: "example.com/checkStdlibUsageUTF8Cutset"}},
		"SA1012": {{Dir: "example.com/checkStdlibUsageNilContext"}},
		"SA1013": {{Dir: "example.com/checkStdlibUsageSeeker"}},
		"SA1014": {{Dir: "example.com/CheckUnmarshalPointer"}},
		"SA1015": {{Dir: "example.com/CheckLeakyTimeTick"}, {Dir: "example.com/CheckLeakyTimeTick-main"}},
		"SA1016": {{Dir: "example.com/CheckUntrappableSignal"}},
		"SA1017": {{Dir: "example.com/CheckUnbufferedSignalChan"}},
		"SA1018": {{Dir: "example.com/CheckStringsReplaceZero"}},
		"SA1019": {
			{Dir: "example.com/CheckDeprecated"},
			{Dir: "example.com/CheckDeprecated_go13", Version: "1.3"},
			{Dir: "example.com/CheckDeprecated_go14", Version: "1.4"},
			{Dir: "example.com/CheckDeprecated_go18", Version: "1.8"},
			{Dir: "example.com/CheckDeprecated_go119", Version: "1.19"},
		},
		"SA1020": {{Dir: "example.com/CheckListenAddress"}},
		"SA1021": {{Dir: "example.com/CheckBytesEqualIP"}},
		"SA1023": {{Dir: "example.com/CheckWriterBufferModified"}},
		"SA1024": {{Dir: "example.com/CheckNonUniqueCutset"}},
		"SA1025": {{Dir: "example.com/CheckTimerResetReturnValue"}},
		"SA1026": {{Dir: "example.com/CheckUnsupportedMarshal"}},
		"SA1027": {{Dir: "example.com/CheckAtomicAlignment"}},
		"SA1028": {{Dir: "example.com/CheckSortSlice"}},
		"SA1029": {{Dir: "example.com/CheckWithValueKey"}},
		"SA1030": {
			{Dir: "example.com/CheckStrconv"},
			{Dir: "example.com/CheckStrconv_go115", Version: "1.15"},
		},
		"SA2000": {{Dir: "example.com/CheckWaitgroupAdd"}},
		"SA2001": {{Dir: "example.com/CheckEmptyCriticalSection"}},
		"SA2002": {{Dir: "example.com/CheckConcurrentTesting"}},
		"SA2003": {{Dir: "example.com/CheckDeferLock"}},
		"SA3000": {
			{Dir: "example.com/CheckTestMainExit-1_go14", Version: "1.4"},
			{Dir: "example.com/CheckTestMainExit-2_go14", Version: "1.4"},
			{Dir: "example.com/CheckTestMainExit-3_go14", Version: "1.4"},
			{Dir: "example.com/CheckTestMainExit-4_go14", Version: "1.4"},
			{Dir: "example.com/CheckTestMainExit-5_go14", Version: "1.4"},
			{Dir: "example.com/CheckTestMainExit-1_go115", Version: "1.15"},
		},
		"SA3001": {{Dir: "example.com/CheckBenchmarkN"}},
		"SA4000": {{Dir: "example.com/CheckLhsRhsIdentical"}},
		"SA4001": {{Dir: "example.com/CheckIneffectiveCopy"}},
		"SA4003": {{Dir: "example.com/CheckExtremeComparison"}},
		"SA4004": {{Dir: "example.com/CheckIneffectiveLoop"}},
		"SA4005": {{Dir: "example.com/CheckIneffectiveFieldAssignments"}},
		"SA4006": {{Dir: "example.com/CheckUnreadVariableValues"}},
		"SA4008": {{Dir: "example.com/CheckLoopCondition"}},
		"SA4009": {{Dir: "example.com/CheckArgOverwritten"}},
		"SA4010": {{Dir: "example.com/CheckIneffectiveAppend"}},
		"SA4011": {{Dir: "example.com/CheckScopedBreak"}},
		"SA4012": {{Dir: "example.com/CheckNaNComparison"}},
		"SA4013": {{Dir: "example.com/CheckDoubleNegation"}},
		"SA4014": {{Dir: "example.com/CheckRepeatedIfElse"}},
		"SA4015": {{Dir: "example.com/CheckMathInt"}},
		"SA4016": {{Dir: "example.com/CheckSillyBitwiseOps"}, {Dir: "example.com/CheckSillyBitwiseOps_shadowedIota"}, {Dir: "example.com/CheckSillyBitwiseOps_dotImport"}},
		"SA4017": {{Dir: "example.com/CheckSideEffectFreeCalls"}},
		"SA4018": {{Dir: "example.com/CheckSelfAssignment"}},
		"SA4019": {{Dir: "example.com/CheckDuplicateBuildConstraints"}},
		"SA4020": {{Dir: "example.com/CheckUnreachableTypeCases"}},
		"SA4021": {{Dir: "example.com/CheckSingleArgAppend"}},
		"SA4022": {{Dir: "example.com/CheckAddressIsNil"}},
		"SA4023": {
			{Dir: "example.com/CheckTypedNilInterface"},
			{Dir: "example.com/CheckTypedNilInterface/i26000"},
			{Dir: "example.com/CheckTypedNilInterface/i27815"},
			{Dir: "example.com/CheckTypedNilInterface/i28241"},
			{Dir: "example.com/CheckTypedNilInterface/i31873"},
			{Dir: "example.com/CheckTypedNilInterface/i33965"},
			{Dir: "example.com/CheckTypedNilInterface/i33994"},
			{Dir: "example.com/CheckTypedNilInterface/i35217"},
		},
		"SA4024": {{Dir: "example.com/CheckBuiltinZeroComparison"}},
		"SA4025": {{Dir: "example.com/CheckIntegerDivisionEqualsZero"}},
		"SA4026": {{Dir: "example.com/CheckNegativeZeroFloat"}},
		"SA4027": {{Dir: "example.com/CheckIneffectiveURLQueryModification"}},
		"SA4028": {{Dir: "example.com/CheckModuloOne"}},
		"SA4029": {{Dir: "example.com/CheckIneffectiveSort"}},
		"SA4030": {{Dir: "example.com/CheckIneffectiveRandInt"}},
		"SA4031": {{Dir: "example.com/CheckAllocationNilCheck"}},
		"SA5000": {{Dir: "example.com/CheckNilMaps"}},
		"SA5001": {{Dir: "example.com/CheckEarlyDefer"}},
		"SA5002": {{Dir: "example.com/CheckInfiniteEmptyLoop"}},
		"SA5003": {{Dir: "example.com/CheckDeferInInfiniteLoop"}},
		"SA5004": {{Dir: "example.com/CheckLoopEmptyDefault"}},
		"SA5005": {{Dir: "example.com/CheckCyclicFinalizer"}},
		"SA5007": {{Dir: "example.com/CheckInfiniteRecursion"}},
		"SA5008": {{Dir: "example.com/CheckStructTags"}, {Dir: "example.com/CheckStructTags2"}, {Dir: "example.com/CheckStructTags3"}},
		"SA5009": {{Dir: "example.com/CheckPrintf"}},
		"SA5010": {{Dir: "example.com/CheckImpossibleTypeAssertion"}},
		"SA5011": {{Dir: "example.com/CheckMaybeNil"}},
		"SA5012": {{Dir: "example.com/CheckEvenSliceLength"}},
		"SA6000": {{Dir: "example.com/CheckRegexpMatchLoop"}},
		"SA6001": {{Dir: "example.com/CheckMapBytesKey"}},
		"SA6002": {{Dir: "example.com/CheckSyncPoolValue"}},
		"SA6003": {{Dir: "example.com/CheckRangeStringRunes"}},
		"SA6005": {{Dir: "example.com/CheckToLowerToUpperComparison"}},
		"SA6006": {{Dir: "example.com/CheckByteSliceInIOWriteString"}},
		"SA9001": {{Dir: "example.com/CheckDubiousDeferInChannelRangeLoop"}},
		"SA9002": {{Dir: "example.com/CheckNonOctalFileMode"}},
		"SA9003": {{Dir: "example.com/CheckEmptyBranch"}},
		"SA9004": {{Dir: "example.com/CheckMissingEnumTypesInDeclaration"}},
		"SA9005": {{Dir: "example.com/CheckNoopMarshal"}},
		"SA9006": {{Dir: "example.com/CheckStaticBitShift"}},
		"SA9007": {{Dir: "example.com/CheckBadRemoveAll"}},
		"SA9008": {{Dir: "example.com/CheckTypeAssertionShadowingElse"}},
	}

	testutil.Run(t, Analyzers, checks)
}

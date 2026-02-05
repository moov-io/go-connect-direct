package parser

import (
	"strings"
)

// Record IDs
// https://www.ibm.com/docs/en/connect-direct/6.3.0?topic=processes-submitting-process

// LookupRecordID will lookup a code (like CTRC) and return the RecordID for it (Copy Termination Record)
func LookupRecordID(code string) *RecordID {
	cc := strings.ToUpper(code)

	if strings.HasPrefix(cc, "QC") {
		return &TcqChange
	}

	rc, found := recordCodes[cc]
	if found {
		return &rc
	}
	return nil
}

var (
	recordCodes = map[string]RecordID{
		"CMOT":  ClientManagerCommTermination,
		"CSPA":  SecurePlus,
		"CSPE":  StrongPasswordEncryption,
		"CTRC":  CopyTerminationRecord,
		"EXFA":  ExternalIntegratedFileAgent,
		"FMSD":  FmhSent,
		"FMRV":  FmhReceived,
		"IFED":  StepEndedForIf,
		"RJED":  StepEndedForRunTask,
		"RTED":  StepEndedForRunJob,
		"SBED":  StepEndedForSubmit,
		"PSED":  StepEndedForOther,
		"LSMG":  SessionManager,
		"LSST":  LocalStepStarted,
		"RSST":  RemoteStepStarted,
		"NUIC":  ServerStartupNuic,
		"NUTR":  ServerShutdownNutr,
		"NUT1":  ServerShutdownNut1,
		"NUT2":  ServerShutdownNut2,
		"NUTC":  ServerShutdownNutc,
		"NUIS":  ServerShutdownNuis,
		"PMIP":  ProcessManagerInitializing,
		"PMST":  ProcessManagerStarted,
		"PMED":  ProcessManagerEnded,
		"PMMX":  TcqMaxAgeProcessing,
		"PSTR":  ProcessStarted,
		"PRED":  ProcessEnded,
		"PERR":  ProcessError,
		"PRIN":  ProcessInterrupted,
		"PFLS":  ProcessFlushed,
		"PSAV":  ProcessSaved,
		"QCxx":  TcqChange,
		"SCNT":  ConcurrentSessionCount,
		"SLFA":  SelectFunctionalAuthorities,
		"CHFA":  ChangeFunctionalAuthorities,
		"DLFA":  DeleteFunctionalAuthorities,
		"AUPR":  AuthorizationFileProcessing,
		"SLIP":  SelectInitparms,
		"IPPR":  UpdateInitparms,
		"RFIP":  RefreshInitparms,
		"SLNM":  SelectNetmap,
		"CHNM":  ChangeNetmap,
		"NMPR":  NetmapProcess,
		"SLPX":  SelectProxy,
		"CHPX":  ChangeProxy,
		"DLPX":  DeleteProxy,
		"SMIN":  SessionManagerInitialized,
		"SMED":  SessionManagerEnded,
		"SRSP":  SelectProcessResponse,
		"STRS":  StatisticsResponse,
		"SSTR":  SessionStarted,
		"SEND":  SessionEnded,
		"SERR":  SessionError,
		"STOP":  ShutdownCommand,
		"SUBP":  SubmitProcess,
		"CHCG":  ChangeProcess,
		"DELP":  DeleteProcess,
		"TRON":  TraceOn,
		"TROFF": TraceOff,
		"USEC":  UserSecurity,

		// warnings
		"XCPK": CheckpointingDisabled,
	}
)

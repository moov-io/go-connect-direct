package parser

// Record IDs
// https://www.ibm.com/docs/en/connect-direct/6.3.0?topic=processes-submitting-process

type RecordID struct {
	ID          string
	Category    RecordCategoy
	Description string
}

type RecordCategoy string

var (
	CategoryEvent          RecordCategoy = "Event"
	CategoryProcess        RecordCategoy = "Process"
	CategoryExternalSource RecordCategoy = "External Source"
)

var (
	ClientManagerCommTermination = RecordID{
		ID:          "CMOT",
		Category:    CategoryEvent,
		Description: "Client manager comm termination",
	}

	SecurePlus = RecordID{
		ID:          "CSPA",
		Category:    CategoryProcess,
		Description: "Secure+",
	}

	StrongPasswordEncryption = RecordID{
		ID:          "CSPE",
		Category:    CategoryEvent,
		Description: "Strong Password Encryption",
	}

	CopyTerminationRecord = RecordID{
		ID:          "CTRC",
		Category:    CategoryProcess,
		Description: "Copy Termination Record",
	}

	ExternalIntegratedFileAgent = RecordID{
		ID:          "EXFA",
		Category:    CategoryExternalSource,
		Description: "External Integrated File Agent",
	}

	FmhSent = RecordID{
		ID:          "FMSD",
		Category:    CategoryProcess,
		Description: "FMH sent",
	}

	FmhReceived = RecordID{
		ID:          "FMRV",
		Category:    CategoryProcess,
		Description: "FMH received",
	}

	StepEndedForIf = RecordID{
		ID:          "IFED",
		Category:    CategoryProcess,
		Description: "Step ended for IF",
	}

	StepEndedForRunTask = RecordID{
		ID:          "RJED",
		Category:    CategoryProcess,
		Description: "Step ended for RUN TASK",
	}

	StepEndedForRunJob = RecordID{
		ID:          "RTED",
		Category:    CategoryProcess,
		Description: "Step ended for RUN JOB",
	}

	StepEndedForSubmit = RecordID{
		ID:          "SBED",
		Category:    CategoryProcess,
		Description: "Step ended for SUBMIT",
	}

	StepEndedForOther = RecordID{
		ID:          "PSED",
		Category:    CategoryProcess,
		Description: "Step ended for other",
	}

	SessionManager = RecordID{
		ID:          "LSMG",
		Category:    CategoryProcess,
		Description: "Session Manager",
	}

	LocalStepStarted = RecordID{
		ID:          "LSST",
		Category:    CategoryProcess,
		Description: "Local step started",
	}

	RemoteStepStarted = RecordID{
		ID:          "RSST",
		Category:    CategoryProcess,
		Description: "Remote step started",
	}

	ServerStartupNuic = RecordID{
		ID:          "NUIC",
		Category:    CategoryEvent,
		Description: "Server Startup",
	}

	ServerShutdownNutr = RecordID{
		ID:          "NUTR",
		Category:    CategoryEvent,
		Description: "Server Shutdown",
	}

	ServerShutdownNut1 = RecordID{
		ID:          "NUT1",
		Category:    CategoryEvent,
		Description: "Server Shutdown",
	}

	ServerShutdownNut2 = RecordID{
		ID:          "NUT2",
		Category:    CategoryEvent,
		Description: "Server Shutdown",
	}

	ServerShutdownNutc = RecordID{
		ID:          "NUTC",
		Category:    CategoryEvent,
		Description: "Server Shutdown",
	}

	ServerShutdownNuis = RecordID{
		ID:          "NUIS",
		Category:    CategoryEvent,
		Description: "Server Shutdown",
	}

	ProcessManagerInitializing = RecordID{
		ID:          "PMIP",
		Category:    CategoryEvent,
		Description: "Process manager initializing",
	}

	ProcessManagerStarted = RecordID{
		ID:          "PMST",
		Category:    CategoryEvent,
		Description: "Process manager started",
	}

	ProcessManagerEnded = RecordID{
		ID:          "PMED",
		Category:    CategoryEvent,
		Description: "Process manager ended",
	}

	TcqMaxAgeProcessing = RecordID{
		ID:          "PMMX",
		Category:    CategoryEvent,
		Description: "TCQ max age processing",
	}

	ProcessStarted = RecordID{
		ID:          "PSTR",
		Category:    CategoryEvent,
		Description: "Process started",
	}

	ProcessEnded = RecordID{
		ID:          "PRED",
		Category:    CategoryEvent,
		Description: "Process ended",
	}

	ProcessError = RecordID{
		ID:          "PERR",
		Category:    CategoryEvent,
		Description: "Process error",
	}

	ProcessInterrupted = RecordID{
		ID:          "PRIN",
		Category:    CategoryEvent,
		Description: "Process interrupted",
	}

	ProcessFlushed = RecordID{
		ID:          "PFLS",
		Category:    CategoryEvent,
		Description: "Process flushed",
	}

	ProcessSaved = RecordID{
		ID:          "PSAV",
		Category:    CategoryEvent,
		Description: "Process saved",
	}

	TcqChange = RecordID{
		ID:          "QCxx",
		Category:    CategoryEvent,
		Description: "TCQ change (xx typically identifies the new queue)",
	}

	ConcurrentSessionCount = RecordID{
		ID:          "SCNT",
		Category:    CategoryEvent,
		Description: "Concurrent session count",
	}

	SelectFunctionalAuthorities = RecordID{
		ID:          "SLFA",
		Category:    CategoryEvent,
		Description: "Select functional authorities",
	}

	ChangeFunctionalAuthorities = RecordID{
		ID:          "CHFA",
		Category:    CategoryEvent,
		Description: "Change functional authorities",
	}

	DeleteFunctionalAuthorities = RecordID{
		ID:          "DLFA",
		Category:    CategoryEvent,
		Description: "Delete functional authorities",
	}

	AuthorizationFileProcessing = RecordID{
		ID:          "AUPR",
		Category:    CategoryEvent,
		Description: "Authorization file processing",
	}

	SelectInitparms = RecordID{
		ID:          "SLIP",
		Category:    CategoryEvent,
		Description: "Select initparms",
	}

	UpdateInitparms = RecordID{
		ID:          "IPPR",
		Category:    CategoryEvent,
		Description: "Update initparms",
	}

	RefreshInitparms = RecordID{
		ID:          "RFIP",
		Category:    CategoryEvent,
		Description: "Refresh initparms",
	}

	SelectNetmap = RecordID{
		ID:          "SLNM",
		Category:    CategoryEvent,
		Description: "Select netmap",
	}

	ChangeNetmap = RecordID{
		ID:          "CHNM",
		Category:    CategoryEvent,
		Description: "Change netmap",
	}

	NetmapProcess = RecordID{
		ID:          "NMPR",
		Category:    CategoryEvent,
		Description: "Netmap process",
	}

	SelectProxy = RecordID{
		ID:          "SLPX",
		Category:    CategoryEvent,
		Description: "Select proxy",
	}

	ChangeProxy = RecordID{
		ID:          "CHPX",
		Category:    CategoryEvent,
		Description: "Change proxy",
	}

	DeleteProxy = RecordID{
		ID:          "DLPX",
		Category:    CategoryEvent,
		Description: "Delete proxy",
	}

	SessionManagerInitialized = RecordID{
		ID:          "SMIN",
		Category:    CategoryEvent,
		Description: "Session manager initialized",
	}

	SessionManagerEnded = RecordID{
		ID:          "SMED",
		Category:    CategoryEvent,
		Description: "Session manager ended",
	}

	SelectProcessResponse = RecordID{
		ID:          "SRSP",
		Category:    CategoryEvent,
		Description: "Select process response",
	}

	StatisticsResponse = RecordID{
		ID:          "STRS",
		Category:    CategoryEvent,
		Description: "Select statistics response",
	}

	SessionStarted = RecordID{
		ID:          "SSTR",
		Category:    CategoryEvent,
		Description: "Session started",
	}

	SessionEnded = RecordID{
		ID:          "SEND",
		Category:    CategoryEvent,
		Description: "Session ended",
	}

	SessionError = RecordID{
		ID:          "SERR",
		Category:    CategoryEvent,
		Description: "Session error",
	}

	ShutdownCommand = RecordID{
		ID:          "STOP",
		Category:    CategoryEvent,
		Description: "Shutdown command",
	}

	SubmitProcess = RecordID{
		ID:          "SUBP",
		Category:    CategoryProcess,
		Description: "Submit process",
	}

	ChangeProcess = RecordID{
		ID:          "CHCG",
		Category:    CategoryProcess,
		Description: "Change process",
	}

	DeleteProcess = RecordID{
		ID:          "DELP",
		Category:    CategoryProcess,
		Description: "Delete process",
	}

	TraceOn = RecordID{
		ID:          "TRON",
		Category:    CategoryEvent,
		Description: "Trace on",
	}

	TraceOff = RecordID{
		ID:          "TROFF",
		Category:    CategoryEvent,
		Description: "Trace off",
	}

	UserSecurity = RecordID{
		ID:          "USEC",
		Category:    CategoryProcess,
		Description: "User security",
	}

	// Warnings

	// CheckpointingDisabled
	//
	// 020) CDUA-3557 / APAR IT41867  commit date:  25 Aug 2022
	// --------------------------------------------------------
	// Copy steps to an object store with checkpointing enabled may receive a
	// warning message, XCPK005W, indicating that checkpointing was disabled for
	// the copy step. The message did not indicate why checkpointing was
	// disabled.
	//
	// Source: https://delivery04.dhe.ibm.com/sar/CMA/OSA/0beu4/0/6.2.0.6.iFix015FixList.txt
	CheckpointingDisabled = RecordID{
		ID:          "XCPK",
		Category:    CategoryProcess,
		Description: "Checkpointing was disabled for the copy step",
	}
)

package log

import (
	"github.com/90poe/service-chassis/logging"
	"kafka-topic-analysis/env"
	"os"
)

var (
	EntryFactory logging.LogEntryFactory
)

func init() {

	// Log as JSON instead of the default ASCII formatter.
	logging.Init(env.Settings.LogLevel, os.Stdout, env.Settings.PrettyLogOutput)

	EntryFactory = logging.NewLogEntryFactory("Kafkanalyser")

	// 2000 ERROR

	ParseFileError = EntryFactory.MakeEntry(logging.ERROR, 2000, "FAILED TO PARSE JSON INPUT FILE")
	JSONDecodingError = EntryFactory.MakeEntry(logging.ERROR, 2001, "JSON DECODING ERROR")
	ErrorOpeningFile = EntryFactory.MakeEntry(logging.ERROR, 2002, "ERROR OPENING FILE")

	// 3000 DEBUG

	GenericInfo = EntryFactory.MakeEntry(logging.INFO, 3000, "Generic info")

	// 4000 DEBUG

	GenericDebug = EntryFactory.MakeEntry(logging.DEBUG, 4000, "Generic debug")
}

var (
	// 2000 ERROR

	ParseFileError    logging.EntryFunc
	JSONDecodingError logging.EntryFunc
	ErrorOpeningFile  logging.EntryFunc

	// 3000 INFO

	GenericInfo logging.EntryFunc

	// 4000 DEBUG

	GenericDebug logging.EntryFunc
)

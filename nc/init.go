package nc

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	flag_version  = flag.Bool("V", false, "print version")
	flag_sched    = flag.String("yield", "auto", "CUDA scheduling: auto|spin|yield|sync")
	flag_pagelock = flag.Bool("lock", true, "enable CUDA memeory page-locking")
	flag_maxwarp  = flag.Int("warp", MAX_WARP, "maximum elements per warp")
	flag_maxprocs = flag.Int("threads", 0, "maximum number of CPU threads, 0=auto")
	flag_cpuprof  = flag.String("cpuprof", "", "Write gopprof CPU profile to file")
	//flag_memprof    = flag.String("memprof", "", "Write gopprof memory profile to file")
)

func init() {
	flag.Parse()

	initLog()

	Debug("initializing")

	initGOMAXPROCS()

	initCpuProf()

	if *flag_version {
		PrintInfo(os.Stdout)
	}

	initWarp()

	initCUDA()
}

func initLog() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("#")
}

func initGOMAXPROCS() {
	if *flag_maxprocs == 0 {
		*flag_maxprocs = runtime.NumCPU()
	}
	procs := runtime.GOMAXPROCS(*flag_maxprocs) // sets it
	Log("using up to", procs, "CPU threads")
}

func initWarp() {
	MAX_WARP = *flag_maxwarp
	Log("max WarpLen:", MAX_WARP)
}

func initCpuProf() {
	if *flag_cpuprof != "" {
		f, err := os.Create(*flag_cpuprof)
		PanicErr(err)
		Log("Writing CPU profile to", *flag_cpuprof)
		err = pprof.StartCPUProfile(f)
		PanicErr(err)
		AtExit(pprof.StopCPUProfile)
	}
}

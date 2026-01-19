# Go Program Execution Flow

## User Command
```
go run main.go / ./binary
```

## Go Toolchain
- `go run` → `go build`
- Compile packages
- Link runtime + stdlib `merges the code, the Go runtime, and standard library code into one file so your program can run independently`
- Produce native binary `wont run in other os or cpu architecture`

## Operating System
- Loads ELF/Mach-O/PE `loads the binary to memory`
- Maps .text, .data, .bss `operating system divides the program’s binary into sections and loads them into memory`
- Sets up initial stack
- Jumps to entry point → `runtime.rt0_go` (not `main.main`)

## Runtime Bootstrap
**`runtime.rt0_go`**
- CPU/OS detection 
- TLS setup `thread-local-storage neded for goroutines and os threads`
- argv/env parsing `parses commandline arguments and env variables`
- Signal handlers `nstalls handlers for OS signals (like Ctrl+C, kill, etc.)`
- Create first M (OS thread) `Starts the first “M” (machine), which is an OS thread that will run goroutines.`
- Create g0 (system goroutine) `Creates the initial system goroutine for low level tasks`

## Scheduler Initialization
- Initialize G-M-P model `G-goroutine(light weight thread) M-machine(os thread) P-processor(schedules goroutines)`
- Set GOMAXPROCS → P count `sets P count based on GOMAXPROCS`
- Create idle Ps `creates P objects that are ready for G`
- Bind P to first M `attaches the P to the corresponding M so G can run`
- Start system monitor (sysmon) `network poling garbage collection etc...`

## Memory & GC Initialization
- Heap arena creation `creates a large region of memory for dynamic memory alocation`
- Span & size class setup 
- Stack allocator `alocates stack for each goroutine`
- Write barriers `sets up special checks to track garbage collector`
- GC controller

## Package Initialization Phase
*(single-threaded, deterministic)*

For each package (DFS order):
1. Resolve imports
2. Initialize global variables
3. Run `init()` functions

## Create main Goroutine
- New G for `main.main`
- Place G on run queue
- Scheduler dispatches it

## main.main Executes
- User code starts
- Function calls
- Allocations
- Goroutine creation (`go f()`)
- Channel ops
- Syscalls

## Runtime During Execution

### Scheduler
- Local/global run queues
- Work stealing
- Preemption

### Netpoller
- epoll/kqueue/wepoll
- Async I/O wakeups

### Garbage Collector
- Concurrent marking
- Short STW pauses

## Goroutine Lifecycle
```
new → runnable → running → 
waiting (chan/syscall) → runnable → dead
```

## Program Termination
- `main.main` returns
- Deferred calls executed
- Runtime checks active goroutines
- GC final sweep
- `os.Exit(code)`


# How it Goes

## 1. `go run file.go`
Go looks into `go.mod` for module path, version and dependencies -> dependencies are resolved according to the `go.mod` file and stored in the `$GOPATH/pkg/mod` -> all `.go` files are loaded and build constraints applied according to the OS -> Source code is converted to tokens and Abstract Syntax Tree (AST) is built. Syntax errors are caught here -> type checking happens -> SSA (static single assignment) is generated from the code -> optimization [dead code elimination, inlining and stuff like that] -> machine code generation [native machine code not byte code] -> compiled packages are combined -> redy to run.
###### This is the build process.

## 2. OS loads the program
Loads the program into memory, prepares stack and heap, jumps to the Go runtime entry point (`runtime.rt0_go `)

## 3. Go runtime initializes
Detects OS and CPU architecture, sets memory management, starts garbage collector, prepares sheduler, reads command line arguments.

## 4. G-M-P
G -> goroutines, M -> OS thread, P -> logical processor. `GOMAXPROCS` decides how many P is allocated and now the shceduler is redy.

## 5. Packages are initialised
For each imported package; global variables initialised, `init()` executed. all this happens once and in single thread.

## 6. `main` initialised
main package variables are initialised and main packages `init()` is executed.

## 7. `main()` starts executing
functions are called, memories are allocated on stack or heap, return values are passed back.

## 8. Goroutines are created
goroutines are put in queue and scheduler decides when it is run. scheduler picks a runable goroutine, runs it on the os thread.if a goroutine reads from a channel, sleeps or makes system calls then that goroutine is paused and the os thread is freed so another goroutine can run.

## 9. Garbage Collection runs in the background concurrently (mostly)

## 10. `main()` finishes
`main()` returns, deferred functions are run.

## 11. Program Exits
when `main()` returns and no goroutines are running.

## 12. Controll is released to the OS
Go is now Gone.

```mermaid
flowchart TD
    A[go run file.go] --> B[Read go.mod\nResolve module path and dependencies]
    B --> C[Download dependencies to GOPATH pkg mod]
    C --> D[Load go files\nApply build constraints OS ARCH]
    D --> E[Lexing and Parsing\nBuild AST\nCatch syntax errors]
    E --> F[Type Checking]
    F --> G[Generate SSA]
    G --> H[Optimizations\nInlining DCE etc]
    H --> I[Generate Native Machine Code]
    I --> J[Link compiled packages]
    J --> K[Executable ready to run]

    K --> L[OS loads program]
    L --> M[Prepare stack and heap]
    M --> N[Jump to runtime rt0_go]

    N --> O[Go runtime initializes]
    O --> O1[Detect OS and CPU]
    O --> O2[Initialize memory management]
    O --> O3[Start garbage collector]
    O --> O4[Prepare scheduler]
    O --> O5[Read command line arguments]

    O --> P[GMP Scheduler Model]
    P --> P1[G Goroutines]
    P --> P2[M OS Threads]
    P --> P3[P Logical Processors]
    P --> P4[GOMAXPROCS decides P count]

    P --> Q[Package Initialization]
    Q --> Q1[Initialize global variables]
    Q --> Q2[Execute init functions]
    Q --> Q3[Single threaded once only]

    Q --> R[Main package initialization]
    R --> R1[Main globals initialized]
    R --> R2[Main init executed]

    R --> S[main starts execution]
    S --> S1[Functions called]
    S --> S2[Stack and Heap allocation]
    S --> S3[Return values passed]

    S --> T[Goroutines created]
    T --> T1[Put into runnable queue]
    T --> T2[Scheduler assigns to OS thread]
    T --> T3[Blocking frees OS thread]

    T --> U[Garbage Collector runs concurrently]

    S --> V[main finishes]
    V --> V1[Deferred functions run]

    V --> W[Program exits]
    W --> X[Control returned to OS]
```

# Why it Goes

#### 1. Built for concurrency, goroutines are cheap (KBs for memory), channels make coordination safer. JAVA -> heavy threads, Node.js single threaded (event loop), Python concurrency limited by GIL.
#### 2. High Performance due to being compiled to native binary. Faster than Python and Node safer than C++.
#### 3. Production ready standard library
#### 4. Great for microservices due to light weight threads and super fast startup.
##### The core phyloosophy of go is simplicity, concurrency, productivity and explicitness.
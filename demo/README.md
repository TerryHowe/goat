Functional Demo
===============
This demo compares looping over a set of data rather than manually calling a function for each. I would
guess this is a comparison of functional programming vs structured programming.

In my opinion there are two disadvantages the looping pattern:
1. It is more complex
2. It is harder to debug

Loop
====
The [loop example](https://github.com/TerryHowe/goat/blob/master/demo/main.go#L20:L32) is 11 lines of code, you have an
array and loop. When it fails, you often cannot tell which item failed:

    - Loop Case ----------------------------
    panic: Unexpected boom!
    
    goroutine 1 [running]:
    main.deterministicFailure(0x10cf4a5, 0x5)
            /Users/tlhowe/code/goat/demo/main.go:16 +0x19f
    main.loop_case()
            /Users/tlhowe/code/goat/demo/main.go:30 +0x10d
    main.main()
            /Users/tlhowe/code/goat/demo/main.go:46 +0x38
    
    Process finished with exit code 2

Given that stack trace, you have no idea which item failed.

Functional
==========
The [functional example](https://github.com/TerryHowe/goat/blob/master/demo/main.go#L34:L41) is 6 lines of code, just
function calls. Looking at the stack trace, you can tell which item failed:

    - Functional Case ----------------------
    panic: Unexpected boom!
    
    goroutine 1 [running]:
    main.deterministicFailure(0x10cf4a5, 0x5)
            /Users/tlhowe/code/goat/demo/main.go:16 +0x19f
    main.functional_case()
            /Users/tlhowe/code/goat/demo/main.go:38 +0xc5
    main.main()
            /Users/tlhowe/code/goat/demo/main.go:48 +0x47
    
    Process finished with exit code 2

Looking at the stack trace, you can easily see that it fails on line 38 for "three".
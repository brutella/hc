# Log

This library replaces the standard `log` package to get optional logging behavior based on simple rules. It provides variables to dis-/enable logging for predefined prefixes.

## Example

    import "github.com/brutella/log"
    
    // Disable verbose logging
    log.Verbose = false
    log.Println("[VERB] Cleanup...") // no output

    // Disable error logging
    log.Info = false
    log.Println("[INFO] Cleanup failed") // no output
    
    // Disable error logging
    log.Warn = false
    log.Println("[WARN] Problem ahead") // no output
        
    // Disable error logging
    log.Error = false
    log.Println("[ERRO] Serious problem") // no output

**No code changes required**

Because the library provides the same logging methods as the standard `log` package, you just have to change the import path from `import "log"` to  `import "github.com/brutella/log"`.

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella/](https://github.com/brutella)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)

# License

log is available under the MIT license. See the LICENSE file for more info.
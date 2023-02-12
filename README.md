# Logger

Sane logging in `go`.


# Configuration

*Configuration belongs primarily in the environment*

## Output

It is the job of the consuming application's initiation to redirect `stdout` and `stderr` to the appropriate place(s)

e.g. For local work you desire all messages print to the `stdout` and `stderr` of your shell. For a production workflow, log messages need to be written and appended to a file, in addition to being printed to the terminal.

In the former scenario nothing needs to be done. But for a production workflow, whatever starts the application process may need to run

```
your_consuming_application 2>&1 | tee -a some_log_file.log
```

## Log Level/Severity

Log level for the default logger is set to `INFO`, `WARN`, and `ERROR`.

Log level for consuming client application code can be done by passing the severity for a non-default logger along with where `stdout` and `stderr` should write to, e.g.:

```
	s := logger.New("", logger.Error | logger.Info, os.Stdout, os.Stderr)
```

### Overriding Behavior

The first precedent to set severity is the mask passed during the creation of a new logger (see [Log Level/Severity](#log-levelseverity))

The second precedent are `ENV` variables prefixed with `LOG_LEVEL_`. These variables, if defined, will override behavior set through `New` logger instantiation. If defined and set to something other than `0` that log level will be enabled, vice versa if set to `0`. This allows more granular control at runtime of which logging levels are displayed, as the variables can be set completely independently, i.e.:

- `LOG_LEVEL_ERROR=1 LOG_LEVEL_WARN=0 ...`
- `LOG_LEVEL_ERROR=0 LOG_LEVEL_DEBUG=0 LOG_LEVEL_INFO=1 ...`
- `LOG_LEVEL_ERROR=0 LOG_LEVEL_DEBUG=0 LOG_LEVEL_INFO=0 LOG_LEVEL_WARN=1 ...`

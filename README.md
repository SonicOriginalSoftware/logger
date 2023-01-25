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

Log level for the default logger is set to `WARN` and `ERROR`.

Log level for consuming client application code can be done by passing the severity for a non-default logger, e.g.:

```
	s := logger.New("", logger.Error | logger.Info)
```

### Overriding Behavior

The first precedent to set severity is the mask passed during the creation of a new logger (see [Log Level/Severity](#log-levelseverity))

The second precedent is an `ENV` variable named `LOGLEVEL`. This variable should be a string representation of the number denoting the value of the mask of the logging severity desired, i.e.:

- `ERROR` (`0b0001`) and `WARN` (`0b0010`): `"3"` (`0b0011`)
- `ERROR` (`0b0001`), `WARN` (`0b0010`), and `INFO` (`0b100`): `"7"` (`0b0111`)
- `ERROR` (`0b0001`) and `DEBUG` (`0b1000`): `"9"` (`0b1001`)

The third precedent is for the state of any specific channel of the severity.

There is an `ENV` variable pertaining to each of the severities:

- `LOGERROR`
- `LOGWARN`
- `LOGINFO`
- `LOGDEBUG`

If defined and not equal to `0`, they toggle on that severity. If defined and equal to `0`, they toggle off that specific severity.

If they are not defined then no additional action is taken beyond the previous precedences.

## Multiple Loggers

To configure specific loggers, prefix the `ENV` variable with the same prefix used for that logger.

e.g., if you need a `renderLogger` and a `computationalLogger`, you might create them like this

```
	renderLogger := logger.New("renderer", logger.DefaultSeverity)
	computeLogger := logger.New("computer", logger.DefaultSeverity)
```

The default severity may be fine for most environment but for testing you may want additional `DEBUG` log messages from the `computerLogger`.

In that case, in the `ENV` of whatever is running tests, you can define

```
computer_LOGDEBUG=1
```

and the `computeLogger` will additionally show `DEBUG` messages while the `renderLogger` will only show messages under the default severity (`WARN` and `ERROR`).

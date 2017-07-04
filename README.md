# sql-sentry
A MySQL proxy that programmatically computes possible SQL patterns from a Go application and compares live DB traffic against these patterns.

The analyzer is built on top of Stripe's [safesql static analyzer](https://github.com/stripe/safesql).

## TODO
* Vendoring
* better logging
* Configure alert only
* zerolog

## Status

SQL Sentry now is still in development and should not be used in production.  The project's current goal is proving that SQL whitelisting is practical, performant, and safe.

## Feedback

Email: <chris.church@gmail.com>


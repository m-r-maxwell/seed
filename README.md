<div style="text-align: center">
  <h1>Seed</h1>
 </div>

 <div style="text-align: center">
<img src="seed.png" height="250" width="250">
 </div>

# Seed

A declarative synthetic dataset generation framework.

## Vision

Seed models realistic datasets using configurable generators, relationships,
and business rules instead of simply producing random fake records.

This repository is an architectural scaffold for v0.1.

## Quickstart

Create and activate a Python virtual environment, then install the package in editable mode:

```bash
python -m venv venv
source venv/bin/activate
pip install -e .
```

The project exposes a CLI entrypoint `seed` (configured in `pyproject.toml`). You can also run the CLI directly from the repository root with `PYTHONPATH=src`:

```bash
# run via package entrypoint (after pip install -e .)
seed generate examples/customer.yaml --output customers.csv

# or run directly from the repo (no install required)
PYTHONPATH=src python -m seed.cli.main generate examples/customer.yaml --output customers.csv
```

Other commands:

```bash
seed validate examples/customer.yaml   # validate config and generator params
seed inspect examples/customer.yaml    # print dataset schema summary
python ml/explore_dataset.py dataset.csv  # quick EDA script (requires pandas)
```

## Notes

- `generate` supports a number of built-in generator types: `uuid`, `faker`, `normal`, `categorical`, `const`, `sequence`, and `sample`.
- Use `--quiet` or `-q` with `generate` to run without progress output (useful for CI):

```bash
seed generate examples/customer.yaml --output customers.csv --quiet
```

## Development tasks

This repo includes a `Taskfile.yaml` with common tasks for running and experimenting. See `Taskfile.yaml` for details.

## Next steps / Suggestions

See `suggestions.md` for suggested improvements such as making generators pluggable, adding `inspect --format json`, and unit tests for `validate` and `generate --quiet`.

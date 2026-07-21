import yaml
import typer
from typing import Any, Dict
from rich.console import Console
from faker import Faker

from seed.cli.generate import _build_extended_registry
from seed.config.schema import DatasetConfig

app = typer.Typer()


def _normalize(raw_conf: Dict[str, Any]) -> Dict[str, Any]:
	rows = raw_conf.get("rows") or raw_conf.get("size") or raw_conf.get("count")
	if rows is None:
		raise ValueError("missing rows/size in config")

	fields_raw = raw_conf.get("fields")
	if fields_raw is None:
		raise ValueError("missing fields in config")

	fields: Dict[str, Any] = {}
	if isinstance(fields_raw, list):
		for item in fields_raw:
			name = item.get("name")
			if name in fields:
				raise ValueError(f"duplicate field name: {name}")
			gen = item.get("generator") or {}
			gtype = gen.get("type")
			cfg = {k: v for k, v in gen.items() if k != "type"}
			fields[name] = {"type": gtype, "config": cfg}
	elif isinstance(fields_raw, dict):
		for name, v in fields_raw.items():
			if name in fields:
				raise ValueError(f"duplicate field name: {name}")
			if isinstance(v, dict) and "type" in v:
				fields[name] = {"type": v["type"], "config": v.get("config", {})}
			else:
				# assume mapping to generator spec
				fields[name] = v
	else:
		raise ValueError("unsupported fields format")

	return {"rows": int(rows), "fields": fields}


@app.callback(invoke_without_command=True)
def validate(config: str):
	console = Console()
	console.print("Loading configuration...\n")

	with open(config, "r") as fh:
		raw = yaml.safe_load(fh)

	try:
		norm = _normalize(raw)
	except Exception as e:
		console.print(f"[red]Invalid config:{e}[/red]")
		raise typer.Exit(code=2)

	console.print("[green]✓ Configuration parsed[/green]\n")

	console.print("Validating schema with Pydantic...")
	try:
		cfg = DatasetConfig.model_validate(norm)
	except Exception as e:
		console.print(f"[red]Schema validation failed: {e}[/red]")
		raise typer.Exit(code=2)
	console.print("[green]✓ Schema valid[/green]\n")

	registry = _build_extended_registry()

	# Check generator names and per-type params
	fake = Faker()
	errors = []

	for fname, fcfg in cfg.fields.items():
		gtype = fcfg.type
		gcfg = fcfg.config
		if gtype not in registry:
			errors.append(f"Unknown generator type '{gtype}' for field '{fname}'")
			continue

		# Per-type checks
		if gtype == "faker":
			provider = gcfg.get("provider")
			if not provider:
				errors.append(f"Field '{fname}': faker generator missing 'provider'")
			elif not hasattr(fake, provider):
				errors.append(f"Field '{fname}': faker has no provider '{provider}'")

		if gtype == "normal":
			if "mean" not in gcfg or "stddev" not in gcfg:
				errors.append(f"Field '{fname}': normal generator requires 'mean' and 'stddev'")
			else:
				try:
					std = float(gcfg.get("stddev"))
					if std <= 0:
						errors.append(f"Field '{fname}': normal 'stddev' must be > 0")
				except Exception:
					errors.append(f"Field '{fname}': normal 'stddev' must be numeric")

		if gtype == "categorical":
			values = gcfg.get("values")
			if not values:
				errors.append(f"Field '{fname}': categorical requires 'values'")
			else:
				if isinstance(values, dict):
					total = sum(values.values())
					# accept sums close to 1 or close to 100
					if abs(total - 1.0) > 1e-6 and abs(total - 100.0) > 1e-2:
						errors.append(f"Field '{fname}': categorical probabilities sum to {total} (expected 1.0 or 100)")

	# detect duplicate field names was handled in normalize

	if errors:
		console.print("[red]Validation failed:[/red]")
		for e in errors:
			console.print(f" - {e}")
		raise typer.Exit(code=2)

	console.print("[green]✓ All checks passed[/green]")
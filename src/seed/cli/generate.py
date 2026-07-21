import csv
import random
import typer
import yaml
import uuid
from faker import Faker
from typing import Any, Dict

from rich.console import Console
from rich.progress import Progress, BarColumn, TextColumn, TimeRemainingColumn

from seed.config.schema import DatasetConfig

app = typer.Typer()


def _build_registry():
    def const_gen(cfg, idx):
        return cfg.get("value")

    def sequence_gen(cfg, idx):
        start = cfg.get("start", 0)
        step = cfg.get("step", 1)
        return start + idx * step

    def sample_gen(cfg, idx):
        choices = cfg.get("choices", [])
        return random.choice(choices) if choices else None

    return {
        "const": const_gen,
        "sequence": sequence_gen,
        "sample": sample_gen,
    }


def _build_extended_registry():
    fake = Faker()

    def uuid_gen(cfg: Dict[str, Any], idx: int):
        return str(uuid.uuid4())

    def faker_gen(cfg: Dict[str, Any], idx: int):
        provider = cfg.get("provider")
        if not provider:
            return None
        fn = getattr(fake, provider, None)
        if fn is None:
            return None
        return fn()

    def normal_gen(cfg: Dict[str, Any], idx: int):
        mean = cfg.get("mean", 0)
        stddev = cfg.get("stddev", 1)
        val = random.gauss(mean, stddev)
        if cfg.get("round", False):
            return round(val)
        return val

    def categorical_gen(cfg: Dict[str, Any], idx: int):
        values = cfg.get("values", {})
        if isinstance(values, dict):
            choices = list(values.keys())
            weights = list(values.values())
            return random.choices(choices, weights=weights, k=1)[0]
        if isinstance(values, list):
            return random.choice(values)
        return None

    base = _build_registry()
    base.update({
        "uuid": uuid_gen,
        "faker": faker_gen,
        "normal": normal_gen,
        "categorical": categorical_gen,
    })
    return base


@app.callback(invoke_without_command=True)
def generate(config: str = typer.Argument(...), output: str = "dataset.csv", quiet: bool = typer.Option(False, "--quiet", "-q", help="Suppress console output and progress")):
    """Load YAML config, validate it, generate rows, and export CSV."""
    console = Console()

    if not quiet:
        console.print("Loading configuration...\n")

    with open(config, "r") as fh:
        raw = yaml.safe_load(fh)

    def _normalize(raw_conf: Dict[str, Any]) -> Dict[str, Any]:
        # rows: support 'rows' or 'size'
        rows = raw_conf.get("rows") or raw_conf.get("size") or raw_conf.get("count")
        if rows is None:
            raise typer.BadParameter("missing rows/size in config")

        fields_raw = raw_conf.get("fields")
        if fields_raw is None:
            raise typer.BadParameter("missing fields in config")

        fields: Dict[str, Any] = {}
        # support list of field definitions or mapping
        if isinstance(fields_raw, list):
            for item in fields_raw:
                name = item.get("name")
                gen = item.get("generator") or {}
                gtype = gen.get("type")
                cfg = {k: v for k, v in gen.items() if k != "type"}
                fields[name] = {"type": gtype, "config": cfg}
        elif isinstance(fields_raw, dict):
            for name, v in fields_raw.items():
                if isinstance(v, dict) and "type" in v:
                    fields[name] = {"type": v["type"], "config": v.get("config", {})}
                else:
                    # assume mapping to generator spec
                    fields[name] = v
        else:
            raise typer.BadParameter("unsupported fields format")

        return {"rows": int(rows), "fields": fields}

    norm = _normalize(raw)
    cfg = DatasetConfig.model_validate(norm)

    registry = _build_extended_registry()
    field_items = list(cfg.fields.items())

    console.print("\n[green]✓ Configuration valid[/green]\n")

    used_types = sorted({fcfg.type for _, fcfg in field_items})
    console.print("Registered generators:")
    for t in used_types:
        if t in registry:
            console.print(f"    [green]✓ {t}[/green]")
        else:
            console.print(f"    [red]✗ {t}[/red]")

    console.print(f"\nGenerating {cfg.rows} records...\n")

    # Open output and stream rows; show progress only when not quiet
    with open(output, "w", newline="") as outfh:
        writer = csv.DictWriter(outfh, fieldnames=[n for n, _ in field_items])
        writer.writeheader()

        if quiet:
            for i in range(cfg.rows):
                row = {}
                for name, fcfg in field_items:
                    gen = registry.get(fcfg.type)
                    if gen is None:
                        raise typer.Exit(code=2)
                    val = gen(fcfg.config, i)
                    row[name] = "" if val is None else str(val).replace("\r", " ").replace("\n", " ")
                writer.writerow(row)
        else:
            with Progress(
                TextColumn("{task.description}"),
                BarColumn(),
                TextColumn("{task.percentage:>3.0f}%"),
                TimeRemainingColumn(),
                console=console,
            ) as progress:
                task = progress.add_task("", total=cfg.rows)
                for i in range(cfg.rows):
                    row = {}
                    for name, fcfg in field_items:
                        gen = registry.get(fcfg.type)
                        if gen is None:
                            console.print(f"Unknown generator type: {fcfg.type}", style="red")
                            raise typer.Exit(code=2)
                        val = gen(fcfg.config, i)
                        if val is None:
                            sval = ""
                        else:
                            sval = str(val).replace("\r", " ").replace("\n", " ")
                        row[name] = sval
                    writer.writerow(row)

    if not quiet:
        console.print("\nWriting {}\n".format(output))
        console.print("[bold green]Done![/bold green]\n")

        console.print("Generated:")
        console.print(f"    {output}\n")
        console.print("Rows:")
        console.print(f"    {cfg.rows}")

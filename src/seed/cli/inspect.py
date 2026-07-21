import yaml
import typer
from rich.console import Console


def _normalize_fields_for_inspect(fields_raw):
    # Return list of (name, generator) preserving order
    out = []
    if isinstance(fields_raw, list):
        for item in fields_raw:
            name = item.get("name")
            gen = item.get("generator") or {}
            out.append((name, gen))
    elif isinstance(fields_raw, dict):
        for name, v in fields_raw.items():
            if isinstance(v, dict) and "generator" in v:
                out.append((name, v.get("generator", {})))
            elif isinstance(v, dict) and "type" in v:
                out.append((name, v))
            else:
                out.append((name, v))
    return out


def _format_percentage_map(values: dict):
    total = sum(values.values()) if values else 0
    lines = []
    for k, v in values.items():
        pct = (v / total * 100) if total else 0
        lines.append((k, pct))
    return lines


def inspect(config: str = typer.Argument(...)):
    console = Console()
    with open(config, "r") as fh:
        raw = yaml.safe_load(fh)

    name = raw.get("name") or raw.get("dataset") or "dataset"
    rows = raw.get("rows") or raw.get("size") or raw.get("count") or "?"
    console.print(f"Dataset: {name}\n")
    console.print(f"Rows: {rows}\n")
    console.print("Fields\n")

    fields_raw = raw.get("fields", {})
    fields = _normalize_fields_for_inspect(fields_raw)
    for fname, gen in fields:
        console.print(f"{fname}")
        if not gen:
            console.print("    <unknown>")
            continue
        gtype = gen.get("type")
        if gtype == "faker":
            provider = gen.get("provider")
            console.print(f"    Faker({provider})")
        elif gtype == "uuid":
            console.print("    UUID Generator")
        elif gtype == "normal":
            console.print("    Normal")
            mean = gen.get("mean")
            stddev = gen.get("stddev")
            console.print(f"        mean: {mean}")
            console.print(f"        stddev: {stddev}")
        elif gtype == "categorical":
            console.print("    Categorical")
            values = gen.get("values") or {}
            for k, pct in _format_percentage_map(values):
                # show rounded percentage without trailing zeros (tolerance for float imprecision)
                if abs(pct - round(pct)) < 1e-6:
                    pstr = f"{round(pct):.0f}%"
                else:
                    pstr = f"{pct:.2f}%"
                console.print(f"        {k}: {pstr}")
        else:
            # fallback to showing type and config
            console.print(f"    {gtype}")
            for kk, vv in (gen.items() if isinstance(gen, dict) else []):
                if kk == "type":
                    continue
                console.print(f"        {kk}: {vv}")

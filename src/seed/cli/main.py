import typer
from . import generate,validate,stats,inspect as inspect_cmd

app = typer.Typer(help="Synthetic dataset generation framework")
app.command()(generate.generate)
app.command()(validate.validate)
app.command()(stats.stats)
app.command()(inspect_cmd.inspect)
if __name__=="__main__":
    app()

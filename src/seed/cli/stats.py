import typer
app=typer.Typer()
@app.callback(invoke_without_command=True)
def stats(config:str): print(f'Stats {config}')
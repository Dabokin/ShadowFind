from rich.console import Console
from rich.table import Table

def display_scan_results(host, ports):
    console = Console()
    table = Table(title=f"Scan Results for {host}")
    
    table.add_column("Port", justify="right")
    table.add_column("Service", justify="left")
    table.add_column("Status", justify="center")
    
    for port in ports:
        service = get_service_name(port)  # Нужно реализовать
        table.add_row(str(port), service, "[green]OPEN[/green]")
    
    console.print(table)

def get_service_name(port):
    # Упрощенная версия
    common_ports = {
        80: "HTTP",
        443: "HTTPS",
        22: "SSH",
        # ...
    }
    return common_ports.get(port, "unknown")
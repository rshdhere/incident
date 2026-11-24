"""Check the catalog binary's startup failures, HTTP health and process signals."""

import json
import os
from pathlib import Path
import select
import signal
import socket
import subprocess
import sys
from urllib.request import urlopen


def check_failure(binary, env, setting, value, expected_error):
    result = subprocess.run(
        [binary], env={**env, setting: value}, capture_output=True, timeout=5
    )
    if result.returncode == 0 or expected_error not in json.loads(result.stderr)["error"]:
        raise AssertionError(f"unexpected result for {setting}: {result.stderr!r}")


def check_stop(binary, env, stop_signal):
    with subprocess.Popen([binary], env=env, stderr=subprocess.PIPE) as process:
        try:
            if not select.select([process.stderr], [], [], 5)[0]:
                raise AssertionError("server did not announce startup")
            startup = json.loads(process.stderr.readline())
            port = int(startup["address"].rsplit(":", 1)[1])
            with urlopen(f"http://127.0.0.1:{port}/healthz", timeout=3) as response:
                if json.load(response) != {"status": "ok"}:
                    raise AssertionError("unexpected health response")
            process.send_signal(stop_signal)
            _, logs = process.communicate(timeout=5)
            if process.returncode or 'server stopped' not in logs.decode():
                raise AssertionError(f"server failed to shut down cleanly: {logs!r}")
        finally:
            if process.poll() is None:
                process.kill()
                process.wait()


def main():
    binary = Path(sys.argv[1]).resolve()
    root = Path(__file__).resolve().parents[2]
    env = {
        **os.environ,
        "PORT": "0",
        "INCIDENT_ROOT": str(root),
        "HTTP_READ_TIMEOUT": "2s",
        "HTTP_WRITE_TIMEOUT": "3s",
        "HTTP_IDLE_TIMEOUT": "4s",
        "HTTP_SHUTDOWN_TIMEOUT": "2s",
    }
    check_failure(binary, env, "PORT", "65536", "PORT")
    check_failure(binary, env, "HTTP_READ_TIMEOUT", "0s", "HTTP_READ_TIMEOUT")
    check_failure(binary, env, "HTTP_IDLE_TIMEOUT", "none", "HTTP_IDLE_TIMEOUT")
    with socket.socket() as occupied:
        occupied.bind(("127.0.0.1", 0))
        occupied.listen()
        check_failure(binary, env, "PORT", str(occupied.getsockname()[1]), "listen")
    for stop_signal in (signal.SIGINT, signal.SIGTERM):
        check_stop(binary, env, stop_signal)
    print("Server validated its settings, reported a busy port, and stopped cleanly on SIGINT and SIGTERM.")


if __name__ == "__main__":
    main()

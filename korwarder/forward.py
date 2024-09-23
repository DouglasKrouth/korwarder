from subprocess import check_call

class Forward:
    """
    Shortlived lifecycle of a port-forward subprocess.

    Args:
        keep_alive: Whether we want the port-forward process to attempt to restart on timeout, sys exit.
        max_retries: Number of retries to attempt to restart a port-forward subprocess.
    """

    def __init__(self, keep_alive: bool = True, max_retries: int = 3):
        
        pass

    def spawn_port_forward(self):
        pass

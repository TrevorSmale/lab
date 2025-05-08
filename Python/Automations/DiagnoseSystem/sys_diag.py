# sys_diag.py

import platform
import socket
import os
import subprocess

def get_uptime():
    try:
        output = subprocess.check_output(['uptime', '-p']).decode().strip()
        return output
    except Exception as e:
        return f"Error: {e}"

def get_load():
    try:
        return os.getloadavg()
    except AttributeError:
        return "Not available"

def get_disk_usage(path='/'):
    st = os.statvfs(path)
    total = st.f_blocks * st.f_frsize
    free = st.f_bfree * st.f_frsize
    used = total - free
    return used // (1024**3), total // (1024**3)  # in GB

def main():
    print(f"Hostname: {socket.gethostname()}")
    print(f"OS: {platform.system()} {platform.release()} ({platform.version()})")
    print(f"Architecture: {platform.machine()}")
    print(f"Uptime: {get_uptime()}")
    print(f"Load Average: {get_load()}")
    
    used, total = get_disk_usage("/")
    print(f"Disk Usage (/): {used} GB used / {total} GB total")

if __name__ == "__main__":
    main()
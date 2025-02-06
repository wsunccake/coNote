# ch3 行程排成器

```python
# load.py
N_LOOP=100_000_000

for _ in range(N_LOOP):
    pass
```

```bash
linux:~ $ time python3 load.py
real	0m4.615s
user	0m4.311s
sys	    0m0.033s
# real ~= user + sys

linux:~ $ time sleep 4
real	0m4.058s
user	0m0.000s
sys	    0m0.003s
```

```bash
# multi-load.sh
MULTICPU=0
PROGNAME=$0
SCRIPT_DIR=$(cd $(dirname $0) && pwd)

usage() {
    exec >&2
    echo "$PROGNAME [-m] <process number>
    -m: multi processor"
    exit 1
}

while getopts "m" OPT; do
    case $OPT in
    m)
        MULTICPU=1
	    ;;
	?)
	    usage
	    ;;
    esac
done

shift $((OPTIND - 1))
if [ $# -lt 1 ]; then
    usage
fi

echo "start time: $(date)"

CONCURRENCY=$1
if [ $MULTICPU -eq 0 ]; then
    taskset -p -c 0 $$ > /dev/null
fi

for ((i=0; i<$CONCURRENCY; i++)); do
    time python3 $SCRIPT_DIR/load.py &
done

# wait: wait for background process
for ((i=0; i<$CONCURRENCY; i++)); do
    wait
done

echo "end   time: $(date)"
```

```bash
linux:~ $ bash multi-load.sh 2
linux:~ $ bash multi-load.sh -m 2
```

```python
# sched.py
import sys
import time
import os
import plot_sched

def usage():
    print("""Usage: {} <number of processes>
        * Starts the same number of load processing processes as <number of processes> on logical CPU 0, with each process consuming CPU resources for approximately 100 milliseconds, and waits for all processes to finish.
        * Writes the execution results as a graph to the file "sched-<number of processes>.jpg".
        * The x-axis of the graph represents the elapsed time [milliseconds] from the start of the load processing process, and the y-axis represents the progress [%]""".format(progname, file=sys.stderr))
    sys.exit(1)

NLOOP_FOR_ESTIMATION=100_000_000
nloop_per_msec = None
progname = sys.argv[0]

def estimate_loops_per_msec():
	before = time.perf_counter()
	for _ in  range(NLOOP_FOR_ESTIMATION):
		pass
	after = time.perf_counter()
	return int(NLOOP_FOR_ESTIMATION/(after-before)/1000)

def child_fn(n):
    progress = 100*[None]
    for i in range(100):
        for j in range(nloop_per_msec):
            pass
        progress[i] = time.perf_counter()
    f = open(f"{n}.data","w")
    for i in range(100):
        f.write("{}\t{}\n".format((progress[i]-start)*1000,i))
    f.close()
    exit(0)

if len(sys.argv) < 2:
    usage()

concurrency = int(sys.argv[1])

if concurrency < 1:
    print(f"Please ensure that <concurrency> is an integer greater than or equal to 1: {concurrency}")
    usage()

os.sched_setaffinity(0, {0})

nloop_per_msec = estimate_loops_per_msec()

start = time.perf_counter()

for i in range(concurrency):
    pid = os.fork()
    if (pid < 0):
        exit(1)
    elif pid == 0:
        child_fn(i)

for i in range(concurrency):
    os.wait()
```

```python
# plot_sched.py
import numpy as np
from PIL import Image
import matplotlib
import os

import sys
import getopt

matplotlib.use('Agg')

import matplotlib.pyplot as plt

plt.rcParams['font.family'] = "sans-serif"
plt.rcParams['font.sans-serif'] = "TakaoPGothic"

def png_to_jpg(png_file, jpg_file):
    fig.savefig(png_file)
    Image.open(png_file).convert("RGB").save(jpg_file)
    os.remove(png_file)

def plot_sched(concurrency):
    fig = plt.figure()
    ax = fig.add_subplot(1,1,1)
    for i in range(concurrency):
        x, y = np.loadtxt("{}.data".format(i), unpack=True)
        ax.scatter(x,y,s=1)
    ax.set_title("Visualization of Time Slices (Concurrency={})".format(concurrency))
    ax.set_xlabel("Elapsed Time [ms]")
    ax.set_xlim(0)
    ax.set_ylabel("Progress [%]")
    ax.set_ylim([0,100])
    legend = []
    for i in range(concurrency):
        legend.append("load"+str(i))
    ax.legend(legend)

    pngfilename = "sched-{}.png".format(concurrency)
    jpgfilename = "sched-{}.jpg".format(concurrency)
    png_to_jpg(pngfilename, jpgfilename)

def plot_avg_tat(max_nproc):
    fig = plt.figure()
    ax = fig.add_subplot(1,1,1)
    x, y, _ = np.loadtxt("cpuperf.data", unpack=True)
    ax.scatter(x,y,s=1)
    ax.set_xlim([0, max_nproc+1])
    ax.set_xlabel("Number of Processes")
    ax.set_ylim(0)
    ax.set_ylabel("Average Turnaround Time [seconds]")

    pngfilename = "avg-tat.png"
    jpgfilename = "avg-tat.jpg"
    png_to_jpg(pngfilename, jpgfilename)

def plot_throughput(max_nproc):
    fig = plt.figure()
    ax = fig.add_subplot(1,1,1)
    x, _, y = np.loadtxt("cpuperf.data", unpack=True)
    ax.scatter(x,y,s=1)
    ax.set_xlim([0, max_nproc+1])
    ax.set_xlabel("Number of Processes")
    ax.set_ylim(0)
    ax.set_ylabel("Throughput [Processes/second]")

    pngfilename = "avg-tat.png"
    jpgfilename = "throughput.jpg"
    png_to_jpg(pngfilename, jpgfilename)

def usage():
    print(f'''Usage:{sys.argv[0]} -t sched|avg_tat|throughput -n <n>
    -h help
    -t type of plot
    -n number of process
    ''')

def main(argv):
    try:
        opts, args = getopt.getopt(argv[1:], 'ht:n:', ["help", "type=", "nproc="])
    except getopt.GetoptError:
        print(f'{argv[0]} -t <type> -n <nproc>')
        sys.exit()

    for name, value in opts:
        if name in ('-h', '--help'):
            usage()
            sys.exit()
        elif name in ('-t', '--type'):
            t = value
        elif name in ('-n', '--nproc'):
            n = value

    plot  = {'sched': plot_sched, 'avg_tat': plot_avg_tat, 'throughput': plot_throughput}
    plot[t](int(n))

if __name__ == '__main__':
    main(sys.argv)
```

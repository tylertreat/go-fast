
def munge():
    out = open('data_atomic.txt', 'w')
    data = open('data.txt', 'r')
    atomic = open('atomic.txt', 'r')

    data_lns = [ln for ln in data]
    atomic_lns = [ln for ln in atomic]
    atomic_index = {}
    i = 0
    while i < len(atomic_lns):
        ln = atomic_lns[i]
        if ln.startswith('numactl'):
            ln = ln[0:ln.index('/')]
            if ln not in atomic_index:
                atomic_index[ln] = []
            atomic_index[ln].append(atomic_lns[i+3])
            i += 3
        i += 1

    i = 0
    while i < len(data_lns):
        ln = data_lns[i]
        if ln.startswith('numactl'):
            runs = atomic_index[ln[0:ln.index('/')]]
            run = runs[0]
            atomic_index[ln[0:ln.index('/')]] = runs[1:]
            out.write(ln)
            out.write(data_lns[i+1])
            out.write(data_lns[i+2])
            out.write(data_lns[i+3])
            out.write(run.replace('mx3', 'mx4'))
            i += 2
        i += 1


if __name__ == "__main__":
    munge()

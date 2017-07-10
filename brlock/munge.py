
def munge():
    out = open('data.txt', 'w')
    before = open('before.txt', 'r')
    after = open('after.txt', 'r')

    before_lns = [ln for ln in before]
    after_lns = [ln for ln in after]
    after_index = {}
    i = 0
    while i < len(after_lns):
        ln = after_lns[i]
        if ln.startswith('numactl'):
            ln = ln[0:ln.index('/')]
            if ln not in after_index:
                after_index[ln] = []
            after_index[ln].append(after_lns[i+2])
            i += 2
        i += 1

    i = 0
    while i < len(before_lns):
        ln = before_lns[i]
        if ln.startswith('numactl'):
            runs = after_index[ln[0:ln.index('/')]]
            run = runs[0]
            after_index[ln[0:ln.index('/')]] = runs[1:]
            out.write(ln)
            out.write(before_lns[i+1])
            out.write(before_lns[i+2])
            out.write(run.replace('mx2', 'mx3'))
            i += 2
        i += 1


if __name__ == "__main__":
    munge()

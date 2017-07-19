#!/usr/bin/env Rint
library(grDevices)
library(utils)
#X11(width=12, height=10)

library(ggplot2)
args <- commandArgs(trailingOnly = TRUE)
args <- if (length(args) == 0) Sys.getenv("ARGS") else args
args <- if (args[1] == "") "plot.dat" else args
d <- data.frame(read.table(
			   text=grep('^mx[1234]', readLines(file(args[1])), value=T),
			   col.names=c("mutex", "cores", "readers", "iterations", "wprob", "wwork", "rwork", "refresh", "time", "X"),
			   colClasses=c(rep(NA, 9), rep("NULL"))
			   ))
d$ops = 1/(d$time/d$iterations/d$readers)
d$mutex = sapply(d$mutex, function(x) {
                    if (x == "mx1") "sync.RWMutex"
                    else if (x == "mx2") "DRWMutex"
                    else if (x == "mx3") "DRWMutex (padded)"
                    else if (x == "mx4") "atomic.Value"
})
da <- aggregate(d$ops, by = list(
				 mutex=d$mutex,
				 cores=d$cores,
				 refresh=d$refresh,
				 rwork=d$rwork,
				 wprob=d$wprob,
				 readers=d$readers), FUN = summary)
p <- ggplot(data=da, aes(x = cores, y = x[,c("Mean")], ymin = x[,c("1st Qu.")], ymax = x[,c("3rd Qu.")], color = mutex))
p <- p + geom_line()
p <- p + geom_errorbar()
#p <- p + facet_wrap(~ readers)
#p <- p + facet_grid(refresh ~ rwork, labeller = function(var, val) {paste(var, " = ", val)})
#p <- p + geom_smooth()
p <- p + ggtitle("Multi-Core Scalability of RWMutex on a Read-Heavy Workload")
p <- p + xlab("CPU cores")
p <- p + ylab("Mean ops per second per reader")

p
ggsave("perf.png", plot = p, width = 8, height = 6)

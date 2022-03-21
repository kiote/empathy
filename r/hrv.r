episode02 <- read.csv("pulse.csv", header = FALSE)
plot(episode02$V2, episode02$V1, type = 'l')
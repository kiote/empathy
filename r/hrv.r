# run it from particular experiment data folder
# for example /1645703267
# before, make sure to run ./process.sh script
episode01_pulse <- read.csv("pulseTS1.csv", header = FALSE)
episode02_pulse <- read.csv("pulseTS2.csv", header = FALSE)
plot(episode01_pulse$V2, episode01_pulse$V1, type = 'l')
plot(episode02_pulse$V2, episode02_pulse$V1, type = 'l')
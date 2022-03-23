# run it from particular experiment data folder
# for example /1645703267
# before, make sure to run ./process.sh script
episode01_pulse <- read.csv("pulseTS1.csv", header = FALSE)
episode02_pulse <- read.csv("pulseTS2.csv", header = FALSE)
plot(episode01_pulse$V2, episode01_pulse$V1, type = 'l')
plot(episode02_pulse$V2, episode02_pulse$V1, type = 'l')

# total observations count
length(episode01_pulse$V1)

#install.packages("IDPmisc")
library(IDPmisc)

##
# example for peaks
##
## Analyses of Mass Spectrum between 12000 and 100'000
## without smoothing, without baseline substraction
data(MS)
MS1 <- log10(MS[MS$mz>12000&MS$mz<1e5,])
P <- peaks(MS1, minPH=0.02, minPW=0.001)
plot(MS1, type="l", xlab="log10(mz)", ylab="log10(I)")
points(P,col="blue",cex=1.6)

episode01_pulse1 <- log10(episode01_pulse[episode01_pulse$V1>300,])
pks02 <- peaks(x=episode01_pulse1$V1,y=episode01_pulse1$V2,minPH=0.02)
plot(episode01_pulse1, type="l", ylab="log10(mz)", ylab="log10(I)")
points(pks02,col="blue",cex=1.6)
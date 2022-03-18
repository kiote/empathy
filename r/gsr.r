#GSR handling
episode02 <- read.csv("episode02.csv")
#if needed run lo pass filtering
m <- nrow(episode02)
SampleRate <- 10 #change accordingly samples per second
Filtered <- c() #currently is used 0.5Hz filtering, which means averaging over 2 seconds.
for (i in 1:m){
  if(i<2*SampleRate) Filtered[i]<-0 else Filtered[i]<-mean(episode02$GSR[(i-2*SampleRate):i])
}
episode02<-cbind(episode02,Filtered)
#further, may use episode02$Filtered instead of episode02$GSR in dealing with peaks
#install.packages("IDPmisc")
library(IDPmisc)
#peak extraction
pks02 <- peaks(x=episode02$time,y=episode02$GSR,minPH=0.02)
colnames(pks02) <- c('time','peak','width')
#dealing with peaks' power i.e. relational height of peaks
n <- nrow(pks02)
pks <- c()
for (i in 1:m) {
  pks[i] <- 0
}
#adding peak values as a col in original data set
for (j in 1:n) {
  a <- pks02$time[j]
  for (i in 1:m){
    if (episode02$time[i]==a) pks[i] <- episode02$GSR[i]
  }
}
episode02 <- cbind(episode02,pks)
#adding peak powers as a col in original data set
pwr <- c()
for (i in 1:m) {
  if (episode02$pks[i]==0) pwr[i] <- 0 else {
    b <- i
    while (episode02$GSR[b] >= episode02$GSR[b-1]) {
      pwr[i] <- (episode02$GSR[i]-episode02$GSR[b])
      b <- b-1
    }
  }
}
episode02 <- cbind(episode02,pwr)
#saving data set as csv
write.csv(episode02, file="02.csv", row.names = FALSE)
#if needed plot GSR, peaks and peak's power in time as a graph as pdf
#pdf("02.pdf")
#plot(episode01$time,episode01$GSR,type="l",col="red")
#lines(episode01$time,episode01$pks,type="l",col="blue")
#lines(episode01$time,episode01$pwr,type="l",col="green")
#dev.off()
#if needed, may remove unnecessary "width" column
#pks01<-pks01[,-3]
#plotting pks data set graph as pdf
#pdf("pks02.pdf") #opens graphic device. If necessary, define width = 5, height=10 (in inches)
#plot(pks02$time,pks02$peaks,type="p")
#dev.off()
import pandas as pd
import numpy as np

from scipy.signal import argrelextrema

df = pd.read_csv('rawdata/1645703267/pulseTS1.csv')
argrelextrema(df)
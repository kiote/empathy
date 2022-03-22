## Data processing

Current data format from raw files if not sutable for further data processing.
That's why r/process.sh script should be runned first. It will split gsr_pulse files to two files - separateley with gsr and with pulse. Then it will add timestamp (in milliseconds, starting with 0) to each datapoint.
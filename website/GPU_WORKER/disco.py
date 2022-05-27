#!/bin/python3
import argparse
import time

# 1. take in the commandline arguments
args = argparse.ArgumentParser()

# print the arguments
# "-text_prompts '{\"0\": [\\"]}' --steps 240 --width_height '[1920, 1080]'", prompt)
args.add_argument("--text_prompts", help="the prompt for the model", required=True)
args.add_argument("--steps", help="the number of steps to run the model", required=True)
args.add_argument("--width_height", help="the width and height of the screen", required=True)

# print the arguments
# print(args.parse_args())
print("TESTING")

# pretend to be doing something for 10 seconds
time.sleep(10)

# write a file for testing
fp = open('modelout', 'w')
fp.write('first line')
fp.close()



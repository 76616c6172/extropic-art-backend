#!/bin/python3

import modal
import io
import os
import time
from typing import Optional
import argparse

OUTPUT_DIR = "/home/valar/model/images/pngs/"
MODEL = "runwayml/stable-diffusion-v1-5"
CACHE_PATH = "/root/model_cache"
GPU = modal.gpu.A100()

stub = modal.Stub("serverless-gpu-worker-C")
volume = modal.SharedVolume().persist("serverless-gpu-worker-volume-C")

@stub.function(
	gpu = GPU,
	image =(
		modal.Image.debian_slim()
		.run_commands(["pip3 install torch --extra-index-url https://download.pytorch.org/whl/cu117"])
		.pip_install(["diffusers", "transformers", "scipy", "ftfy", "accelerate"])
	),
	shared_volumes={CACHE_PATH: volume},
	secret=modal.Secret.from_name("my-huggingface-secret"),
)
async def run_large_diffusion_model_C(prompt: str, neg_prompt: str, seed, width, height, steps, scale):
	from diffusers import StableDiffusionPipeline
	from torch import float16
	import torch as torch

	pipe = StableDiffusionPipeline.from_pretrained(
		MODEL,
 		use_auth_token=os.environ["HUGGINGFACE_TOKEN"],
   	torch_dtype=float16,
    cache_dir=CACHE_PATH,
    device_map="auto",
		safety_checker=None,
	)

	generator = torch.Generator("cuda").manual_seed(seed)

	image = pipe(
		prompt,
		negative_prompt= neg_prompt,
		num_inference_steps=steps,
		guidance_scale=scale,
		width=width,
		generator=generator,
		height=height).images[0]

	# Convert PIL Image to PNG byte array.
	buf = io.BytesIO()
	image.save(buf, format="PNG")
	img_bytes = buf.getvalue()

	return img_bytes

if __name__ == "__main__":
	timeAtStartOfRun = time.monotonic()

	parser = argparse.ArgumentParser()
	parser.add_argument("prompt", help="text prompt for the model")
	parser.add_argument("seed", help="seed used by noise generator")
	parser.add_argument("width", help="width of the final image in pixels")
	parser.add_argument("height", help="height of the final image pixels")
	parser.add_argument("steps", help="number of denoising steps")
	parser.add_argument("scale", help="guidance scale")
	parser.add_argument("jobid", help="the id for so output is saved as <jobid>.png")

	args = parser.parse_args()

	prompt = "mdjrny v4 style " + args.prompt
	seed  = int(args.seed)
	width = int(args.width)
	height = int(args.height)
	steps = int(args.steps)
	scale = int(args.scale)
	file_name = args.jobid + ".png"
	neg_prompt = "signature"

# while loop here.. while get latest job

	# Run serverless inference job
	with stub.run():

		img_bytes = run_large_diffusion_model_C(prompt, neg_prompt, seed, width, height, steps, scale)
		output_path = os.path.join(OUTPUT_DIR, file_name)

		with open(output_path, "wb") as f:
			f.write(img_bytes)

	print(f"wrote data to {output_path}")
	timeAtCompletion = time.monotonic() - timeAtStartOfRun
	print(f"finished in: {timeAtCompletion:.2f} seconds")

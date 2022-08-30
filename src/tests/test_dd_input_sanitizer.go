package main

import (
	"project-exia-monorepo/website/exapi"
)

func main() {
	testcases := [...]string{
		"Scenic view of Sharm el Sheikh:1.5 | person:-0.5",
		"Scenic view of:01 Rapture from bioshock, Aquanox, Artstation 4k, Unreal engine",
		"A beautifully powerful, unreal engine, featured on artstation, wide angle:1.5 | purple:-0.5",
		"A beautifully powerful, unreal engine, featured on artstation, wide angle:1.5 | purple:-0.5",
		"A beautifully powerful, unreal engine:9 | featured on artstation, wide angle:1.5 | purple:-0.5"}

	for _, c := range testcases {
		println("Prompt=", c)
		c = exapi.SanitizeUserPrompt(c)

		isTrue := exapi.InputIsValid(c)

		println(isTrue)
		println()
	}

}

// constants used by the API
package exapi

const MAX_PROMPT_LENGTH = 290 // Reject a new job posted by the view if longer than this value

// ENCODING FOT THE MODEL PIPELINES
const stable_diffusion_midjourney_v4 = 1

// const stable_diffusion_512_by_768 = 2
// const stable_diffusion_768_by_512 = 3

// ENCODE RESOLUTIONS
const RES_512_BY_512 = 1  // sm square
const RES_512_BY_768 = 2  // sm portrait
const RES_768_BY_512 = 3  // sm wide
const RES_1024_BY_512 = 4 // sm wider
const RES_512_BY_1024 = 5 // sm higher

const RES_768_BY_768 = "6"  // md square
const RES_768_BY_1024 = "7" // md portrait
const RES_1024_BY_768 = "8" // md wide

// DEFAULT offline
var GPU_STATUS = "online"
var GPU_IS_ONLINE = true

// constants used by the API
package exapi

const MAX_PROMPT_LENGTH = 360 // Reject a new job posted by the view if longer than this value

// ENCODING FOT THE MODEL PIPELINES
const stable_diffusion_midjourney_v4 = 1

// const stable_diffusion_512_by_768 = 2
// const stable_diffusion_768_by_512 = 3

// ENCODE RESOLUTIONS
const RES_512_BY_512 = 1 // square
const RES_512_BY_768 = 2 // portrait
const RES_768_BY_512 = 3 // wide

// DEFAULT offline
var GPU_STATUS = "online"
var GPU_IS_ONLINE = true

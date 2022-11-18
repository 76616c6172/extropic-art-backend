// constants used by the API
package exapi

const MAX_PROMPT_LENGTH = 360 // Reject a new job posted by the view if longer than this value

// ENCODING FOT THE MODEL PIPELINES
const STABLE_DIFFUSION_512_BY_512 = 1
const STABLE_DIFFUSION_512_BY_768 = 2
const STABLE_DIFFUSION_768_BY_512 = 3

// DEFAULT offline
var GPU_STATUS = "offline"

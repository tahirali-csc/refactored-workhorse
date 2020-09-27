package engine

type (
	// Metadata provides execution metadata.
	Metadata struct {
		UID       string            `json:"uid,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
	}

	Spec struct {
		Metadata Metadata `json:"metadata,omitempty"`
		Steps    []*Step  `json:"steps,omitempty"`
	}

	Step struct {
		Metadata Metadata    `json:"metadata,omitempty"`
		Volume   []*VolumeMount `json:"volume,omitempty"`
		Docker   *DockerStep `json:"docker,omitempty"`
	}

	DockerStep struct {
		Args    []string `json:"args,omitempty"`
		Command []string `json:"command,omitempty"`
		Image   string
	}

	VolumeMount struct {
		Source string
		Target string
	}
)

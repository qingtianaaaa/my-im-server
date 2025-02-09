package runtime

func isDocker() bool {
	return false 
}

func isKubernets() bool {
	return false
}

func PrintRuntimeEnv() string {
	var runtimeEnv string 
	if isKubernets() {
		runtimeEnv = "kubernets"
	} else if isDocker() {
		runtimeEnv = "docker"
	} else {
		runtimeEnv = "source"
	}
	return runtimeEnv
}
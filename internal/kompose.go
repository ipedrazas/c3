package internal

func KomposeCMD(composefile string) []string {
	// docker buildx bake -f docker-bake.hcl --push
	command := []string{
		"kompose",
		"-f",
		composefile,
		"-o",
		"/targetdata",
		"--with-kompose-annotation=false",
		"convert",
	}

	return command
}

func KubectlCMD(path string) []string {
	// docker buildx bake -f docker-bake.hcl --push
	command := []string{
		"kubectl",
		"apply",
		"-f",
		path,
	}

	return command
}

func DefaultTargetCMD(compose string) []string {

	command := []string{
		"bash",
		"-c",
		"kompose -f /data/" + compose + " -o ./data --with-kompose-annotation=false convert; kubectl apply -f ./data",
		// "sleep 300",
	}
	return command
}

func SleepCMD() []string {
	command := []string{
		"sleep",
		"300",
	}
	return command
}

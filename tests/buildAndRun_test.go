package tests

import (
	"os"
	"os/exec"
	"testing"
)

func TestLoops(t *testing.T) {
	cmd := exec.Command("go", "build", "-o", "neco", "..")
	err := cmd.Run()

	if err != nil {
		t.Fatalf("Failed to build neco: " + err.Error())
	}

	cmd = exec.Command("./neco", "src/loops.neco")
	output, err := cmd.Output()

	if err != nil {
		t.Fatalf("Failed to build loops.neco: " + err.Error())
	}

	correctOutput := `00 01 02 03 04 05 06 07 08 09 
10 11 12 13 14 15 16 17 18 19 
20 21 22 23 24 25 26 27 28 29 
30 31 32 33 34 35 36 37 38 39 
40 41 42 43 44 45 46 47 48 49 
50 51 52 53 54 55 56 57 58 59 
60 61 62 63 64 65 66 67 68 69 
70 71 72 73 74 75 76 77 78 79 
80 81 82 83 84 85 86 87 88 89 
90 91 92 93 94 95 96 97 98 99 

3... 2... 1... Start!

2 4 8 16 32 64 128 256 512 1024 2048 4096

Cats: 😺 😸 😻 😽 😼 
`
	if string(output) != correctOutput {
		t.Fatalf("Output of loops:\n\"%s\"\nwanted:\n\"%s\"", string(output), correctOutput)
	}
}

func TestMatchStatements(t *testing.T) {
	cmd := exec.Command("go", "build", "-o", "neco", "..")
	err := cmd.Run()

	if err != nil {
		t.Fatalf("Failed to build neco: " + err.Error())
	}

	cmd = exec.Command("./neco", "src/matchStatements.neco")
	output, err := cmd.Output()

	if err != nil {
		t.Fatalf("Failed to build matchStatements.neco: " + err.Error())
	}

	correctOutput := `1
2
<5, 10>
<5, 10>
<5, 10>
11
Not known number: 0
A
B or C
?
B or C
?
D
A
`
	if string(output) != correctOutput {
		t.Fatalf("Output of matchStatements:\n\"%s\"\nwanted:\n\"%s\"", string(output), correctOutput)
	}
}

func TestStructs(t *testing.T) {
	cmd := exec.Command("go", "build", "-o", "neco", "..")
	err := cmd.Run()

	if err != nil {
		t.Fatalf("Failed to build neco: " + err.Error())
	}

	cmd = exec.Command("./neco", "src/structs.neco")
	output, err := cmd.Output()

	if err != nil {
		t.Fatalf("Failed to build structs.neco: " + err.Error())
	}

	correctOutput := `Person{"Daniel", 173}
Daniel
Person{"Peter", 181}
Pet{"Fluffy", Person{"Peter", 181}}
Peter
Pet{"Fluffy", Person{"Dan", 181}}
`
	if string(output) != correctOutput {
		t.Fatalf("Output of structs:\n\"%s\"\nwanted:\n\"%s\"", string(output), correctOutput)
	}

	t.Cleanup(func() {
		os.Remove("neco")
	})
}

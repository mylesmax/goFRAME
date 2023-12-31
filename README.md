<p align="center">
  <img src="https://i.imgur.com/6w0QDh0.png" width="300" height="300">
</p>

# goFRAME

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mylesmax/goFRAME/blob/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/mylesmax/goframe)](https://goreportcard.com/report/github.com/mylesmax/goframe)

a Go-based **F**ramework for **R**apid **A**nalytical **M**odeling in **E**lectrophysiology (**FRAME**). goFRAME is a simple directory that can greatly improve the efficiency and organization of electrophysiological model simulations.

## Authors

- [@mylesmax](https://www.github.com/mylesmax)

## Working Examples

##### 1952 Hodgkin-Huxley Model of a Neuronal Action Potential

[/examples/hodgkinhuxley](https://github.com/mylesmax/goFRAME/tree/master/examples/hodgkinhuxley)

original paper : [https://www.ncbi.nlm.nih.gov/pmc/articles/PMC1392413/
](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC1392413/)

##### 1991 Luo-Rudy Model of a Ventricular Action Potential (LR1)

[/examples/luorudy1991](https://github.com/mylesmax/goFRAME/tree/master/examples/luorudy1991)

original paper : [https://pubmed.ncbi.nlm.nih.gov/1709839/](https://pubmed.ncbi.nlm.nih.gov/1709839/)

## Usage

### Download

Use the following command to download the repository:

```go get -t github.com/mylesmax/goFRAME```

### Model Generation

It is recommended to read the steps below for a general understanding of the framework, but for a more streamlined experience, follow the template tutorial in [/template](https://github.com/mylesmax/goFRAME/tree/master/template/).

##### Step 1: Define Parameters, Gates, and Currents

Parameter, gate, and current formulation is demonstrated in detail in the examples and template. However, here are some key notes:



###### Parameters

Parameters should be defined in the initial ``State``, which is recommended to be generated in a ``func init()``. The general form of a ``State`` is:

```
type State struct {
	V, Cm, RTF, Stim, T float64
	Index                int
	//Nernst Potentials
	E map[string]float64

	//conductances
	GBar map[string]float64

	//currents
	I map[string]float64

	//gates
	Gate map[string]float64

	//ion concentrations
	ConcOut map[string]float64
	ConcIn map[string]float64

	//misc
	Misc map[string]float64
}
```

For the initial ``State``, set ``V=Vrest`` and ``Index=0``.

###### Gates

If a gate is defined by alpha and beta values (``GateAB``), it should be defined as

```
goFRAME.GateAB{Alpha, Beta, ID}
```

where ``Alpha`` and ``Beta`` are ``func (V float64) float64`` and ``ID`` is of type ``string``.

If a gate is defined by an analytical/direct solution (``GateDirect``), it should be defined as

```
goFRAME.GateDirect{Ss, ID}
```

where ``Ss`` is ``func(V float64) float64``.

In both cases, the ``ID`` field should be matched to the gate name (e.g., ``ID`` is ``"m"`` for gate ``m``).

###### Currents

The ``params...interface{}`` is flexible to three parameters: ``params[0]`` is the current ``State``, ``params[1]`` is an ``Out`` (i.e., all saved States, which is ``[]State``), and ``params[2]`` is an ``int`` that represents the index.

##### Step 2: Define Derivative Functions

Most models have similar calculations of dvdt, the change in membrane voltage, which is given in examples. For further customization, the ``params...interface{}`` is flexible to three parameters: ``params[0]`` is the current ``Solver``, ``params[1]`` is an ``int`` that represents the index, and ``params[2]`` is a ``float64`` that represents the sum of all currents.

##### Step 3: Simulate

Define a stimulus array ``Stims`` that contains as many ``Stim`` instances as needed:

```
type Stim struct{
	Start, End, Intensity float64
}
```

Define a ``TSet`` array that contains the simulation start time, simulation end time, and dt:

```
type TSet struct {
	Start, Dt, End float64
}
```

Then, generate an array to capture the evolution of the system, and save its initial value to be the initial ``State``. For example:

```go
out = []goFRAME.State{initial}
```

Now, define the solver:

```
type Solver struct {
	Out         Out
	Tset        TSet
	Stims       Stims
	GatesAB     []GateAB
	GatesDirect []GateDirect
	Currents    []Current
	ToIntegrate map[string]func(params ...interface{}) float64
	Method      string
}
```

For usable methods, see the bottom of the README. After the ``Solver`` has been generated, you may solve the system with

```solver.Solve()```

and write to excel with

```
WriteExcel(O Out, filePath string)
```

### Usable Numerical Methods

"RK1" : Euler Method (Runge-Kutta 1)

## License

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)


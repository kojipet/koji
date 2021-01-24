package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	// create a new koji
	k := newKoji()

	// start the life cycle timers
	go k.kojiLifeCycles()
	go k.kojiHealthMonitor()

	// display the home screen
	k.homeScreen(osCheck())

	// k.menu()

}

func (k *koji) menu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)

		if strings.Compare("exit", text) == 0 {
			os.Exit(0)
		} else if strings.Compare("quit", text) == 0 {
			os.Exit(0)
		} else if strings.Compare("a", text) == 0 {
			k.age++
		} else if strings.Compare("a-", text) == 0 {
			k.age--
		} else if strings.Compare("f-", text) == 0 {
			k.vitals.food--
		} else if strings.Compare("h-", text) == 0 {
			k.vitals.happy--
		} else if strings.Compare("s-", text) == 0 {
			k.vitals.sleep--
		} else if strings.Compare("f", text) == 0 {
			k.vitals.food++
		} else if strings.Compare("h", text) == 0 {
			k.vitals.happy++
		} else if strings.Compare("s", text) == 0 {
			k.vitals.sleep++
		} else if strings.Compare("g", text) == 0 {
			fmt.Printf(brightyellow+"genome:"+nc+" %s", globalGenes)
		} else if strings.Compare("b", text) == 0 {
			k.conjugateGenes()
		} else if strings.Compare("d", text) == 0 {
			k.die()
		} else if strings.Compare("sleep", text) == 0 {
			k.goToSleep()
		}
	}
}

func (k *koji) kojiStatLine() {
	if k.age > 1 {
		fmt.Printf(brightyellow+"age: "+nc+"%v ", k.age)
	} else if k.age <= 1 {

		fmt.Printf(brightyellow + "age: " + nc + "egg ")
	}
	k.statFood()
	k.statSleep()
	k.statHappy()
	fmt.Printf("\n")
}

func (k *koji) statFood() int {

	for {
		score := k.vitals.food
		fmt.Printf(brightblack+"F%v "+nc, score)
		return score
	}

}

func (k *koji) statHappy() int {

	score := k.vitals.happy
	fmt.Printf(brightblack+"H%v "+nc, score)
	return score
}

func (k *koji) statSleep() int {

	score := k.vitals.sleep
	fmt.Printf(brightblack+"S%v "+nc, score)
	return score

}

func (k *koji) homeScreen(isWin bool) {
	go k.menu()

	for {
		cls(isWin)

		k.kojiStatLine()
		k.kojiFace()
		k.kojiChiron(k.kojiHealthMonitor())

		delay(1)
	}

}

func cls(isWin bool) {

	if !isWin {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if isWin {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func (k *koji) kojiFace() {

	if k.alive {
		Debug("koji alive")
		if k.awake {
			Debug("koji awake")
			if k.age <= 1 {
				Debug("koji infant")
				kojiEggFace(k)
			} else if k.age > 1 {
				if k.vitals.happy > 5 {
					kojiHappyFace(k)
				} else if k.vitals.happy <= 5 {
					kojiSadFace(k)
				} else if k.vitals.happy <= 2 {
					kojiStressFace(k)
				}
			}
		} else if !k.awake {
			kojiSleepFace(k)
		}
	} else if !k.alive {
		kojiDeadFace(k)
	}

}

func (k *koji) kojiChiron(kojiMessageQueue []string) {

	for i := 0; i < len(kojiMessageQueue); i++ {
		fmt.Println(kojiMessageQueue[i])
	}
}

func (k *koji) die() {

	k.alive = false
	k.vitals.food = 0
	k.vitals.sleep = 0
	k.vitals.happy = 0

}

func (k *koji) kojiHealthMonitor() []string {

	var kojiMessageQueue []string
	var msgStr string
	if !k.alive {

		k.die()
		kojiMessageQueue = []string{"koji has died"}
		return kojiMessageQueue
	}
	for {
		if k.alive {
			if k.vitals.food > kojiMaxFood+5 {
				k.die()
			}
			if k.vitals.food <= 0 {
				k.die()
			}
			if k.awake {
				if k.vitals.food < kojiMaxFood/2 {
					msgStr = fmt.Sprintf("%s", "koji is hungry")
					kojiMessageQueue = append(kojiMessageQueue, msgStr)
				}
				if k.vitals.happy < kojiMaxHappy/2 {
					msgStr = fmt.Sprintf("%s", "koji is sad")
					kojiMessageQueue = append(kojiMessageQueue, msgStr)
				}
				if k.vitals.sleep < kojiMaxSleep/2 {
					msgStr = fmt.Sprintf("%s", "koji is sleepy")
					kojiMessageQueue = append(kojiMessageQueue, msgStr)
				}
			} else if !k.awake {
				msgStr = fmt.Sprintf("%s", "koji is asleep")
				kojiMessageQueue = append(kojiMessageQueue, msgStr)
			}
		} else if !k.alive {
			msgStr = fmt.Sprintf("%s", "koji has died")
			kojiMessageQueue = append(kojiMessageQueue, msgStr)
		}

		return kojiMessageQueue
	}

}

func newKoji() koji {

	var k koji

	k.vitals = vitals{
		food:  10,
		sleep: 10,
		happy: 10}

	k = koji{
		name:   "koji",
		age:    0,
		awake:  true,
		alive:  true,
		vitals: k.vitals}

	k.getGenome()

	return k
}

func (k *koji) kojiLifeCycles() {

	go k.eatSomething()
	go k.beAwake()
	go k.areYouHappy()
	go k.getOlder()

}

func (k *koji) eatSomething() {
	if !k.alive || !k.awake {
		return
	}
	for {
		if k.awake == true {
			k.vitals.food--
			Debug("koji is eating" + strconv.Itoa(k.vitals.food))
		}
		delay(60) // every minute
	}

}

func (k *koji) beAwake() {
	if !k.alive {
		return
	}

	for {
		if k.vitals.sleep <= 0 {
			k.vitals.sleep = 0
			k.goToSleep()
		}
		if k.awake {
			k.vitals.sleep--
			Debug("koji is awake" + strconv.Itoa(k.vitals.sleep))
		} else if !k.awake {
			k.vitals.sleep++
			Debug("koji is awake" + strconv.Itoa(k.vitals.sleep))
		}
		delay(60 * 60) // every hour
	}

}

func (k *koji) getOlder() {
	if !k.alive {
		return
	}

	for {
		if k.alive {
			k.age++
			delay(60) // every minute
		}
	}

}

func (k *koji) areYouHappy() {
	if !k.alive {
		return
	}

	for {
		if k.awake {
			if k.vitals.sleep < 2 {
				k.vitals.happy--
			} else if k.vitals.food < 2 {
				k.vitals.happy--
			}
		} else if !k.awake {
			k.vitals.happy++
		}
		delay(60)
	}

}

func (k *koji) goToSleep() {
	if !k.alive {
		return
	}
	if k.age > 1 {
		if k.awake == true {
			Debug("koji was awake")
			if k.vitals.sleep <= 2 {
				Debug("koji could sleep")
				k.awake = false
				timeToSleep := (10 - k.vitals.sleep) * 60 * 60
				delay(timeToSleep)
				k.awake = true
				return
			}
		}
	}

}

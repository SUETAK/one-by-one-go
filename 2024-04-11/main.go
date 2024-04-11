package main

func main() {}

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 1 {
		if gas[0] >= cost[0] {
			return 0
		}
		return -1
	}

	gasSum := 0
	costSum := 0
	startIndex := make([]int, 0)
	lastIndex := len(gas) - 1
	for i := 0; i <= lastIndex; i++ {
		if i == lastIndex {
			gasSum += gas[i]
			costSum += cost[i]
			if gas[i] >= cost[i] {
				startIndex = append(startIndex, i)
			}
			continue
		}
		gasSum += gas[i] + gas[lastIndex]
		costSum += cost[i] + cost[lastIndex]
		if gas[i] >= cost[i] {
			startIndex = append(startIndex, i)
		}
		if gas[lastIndex] >= cost[lastIndex] {
			startIndex = append(startIndex, lastIndex)
		}
		lastIndex--
	}
	if gasSum < costSum || gasSum == 0 {
		return -1
	}

	result := -1
breakLabel:
	for _, index := range startIndex {
		nGas := append(gas[index:], gas[:index]...)
		nCost := append(cost[index:], cost[:index]...)

		g := nGas[0]
		count := 1
		for k := 1; k < len(nGas); k++ {
			g = g + nGas[k] - nCost[k-1]
			if g < nCost[k] {
				break
			}
			count++
			if count == len(nGas) {
				result = index
				break breakLabel
			}
		}

	}
	return result
}

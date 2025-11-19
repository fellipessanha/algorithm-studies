package leetcodesgo

func minCost(colors string, neededTime []int) int {
    idx := 1
    answer := 0

    for idx < len(colors){
        if  colors[idx] != colors[idx-1]{
            idx++
            continue
        }
        totalTime := neededTime[idx-1]
        largestTime := neededTime[idx-1]
        repeatedColor := colors[idx]
        for idx < len(colors) && colors[idx] == repeatedColor{
            totalTime += neededTime[idx]
            largestTime = max(neededTime[idx], largestTime)
            idx++
        }
        answer += totalTime - largestTime
    }

    return answer
}


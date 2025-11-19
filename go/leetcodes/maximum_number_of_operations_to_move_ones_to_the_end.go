package leetcodesgo

func maxOperations(s string) int {
    oneCounter:= 0
    operationCounter := 0
    for idx := 0; idx < len(s); idx ++{
        if s[idx] == '1'{
            oneCounter += 1
            continue
        }
        for idx + 1 < len(s) && s[idx+1] == '0' {
            idx++
        }
        operationCounter += oneCounter
    }
    return operationCounter
}


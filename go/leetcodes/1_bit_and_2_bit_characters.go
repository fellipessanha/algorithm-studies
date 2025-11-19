package leetcodesgo

func isOneBitCharacter(bits []int) bool {
    oneCount := 0
    for i := len(bits)-2; i >= 0; i--{
        if bits[i] == 0{
            break
        }
        oneCount++
    }
    return oneCount % 2 == 0
}

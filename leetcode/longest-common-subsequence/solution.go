func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := range dp {
			dp[i] = make([]int, len(text2))
	}
	
	for i, a := range text1 {
			for j, b:= range text2 {
					if a == b {
							if i > 0 && j > 0 {
									dp[i][j] = dp[i - 1][j - 1]
							}
							
							dp[i][j]++
							continue
					}
					
					if i > 0 {
							dp[i][j] = dp[i-1][j]    
					}
					
					if j > 0 && dp[i][j - 1] > dp[i][j] {
							dp[i][j] = dp[i][j - 1]
					}
			}
	}
	
	return dp[len(text1) - 1][len(text2) - 1]
}
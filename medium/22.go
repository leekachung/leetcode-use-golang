func generateParenthesis(n int) []string {
    res := make([]string, 0)
    backtrack(&res, "", 0, 0, n)
    return res
}

func backtrack(res *[]string, cur string, left, right, max int) {
    if len(cur) == max * 2 {
        *res = append(*res, cur)
        return
    }
    if left < max {
        backtrack(res, cur + "(", left + 1, right, max)
    }
    if right < left {
        backtrack(res, cur + ")", left, right + 1, max)
    }
}

func fizzBuzz(n int) []string {
    feedback := []string{}
    for i := 1; i <= n; i++ {
        if i % 15 == 0 {
            feedback = append(feedback, "FizzBuzz")
        } else if i % 3 == 0 {
            feedback = append(feedback, "Fizz")
        } else if i % 5 == 0 {
            feedback = append(feedback, "Buzz")
        } else {
            feedback = append(feedback, strconv.Itoa(i))
        }
    }
    return feedback
}
